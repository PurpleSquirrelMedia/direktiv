package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/vorteil/direktiv/pkg/dlog"
)

type Handler struct {
	db            *sql.DB
	args          *HandlerArgs
	queueMutex    sync.Mutex
	logQueue      chan *log15.Record
	queuedLogs    []log15.Record
	closed        chan bool
	brokerManager *dlog.BrokerManager
}

type HandlerArgs struct {
	Driver                      *sql.DB
	InsertFrequencyMilliSeconds int
	Namespace                   string
	InstanceID                  string
	BrokerManager               *dlog.BrokerManager
}

func NewHandler(args *HandlerArgs) (*Handler, error) {

	out := new(Handler)

	out.args = args
	out.db = args.Driver
	out.brokerManager = args.BrokerManager

	return out.init()
}

func (l *Logger) initDB() error {

	tx, err := l.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`create table if not exists logs (
		id serial primary key,
		namespace text,
		instance text,
		time bigint,
		lvl int,
		msg text,
		ctx bytea
	)`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`create index if not exists "idx_log_namespace_instance" on logs (namespace, instance)`)
	if err != nil {
		return err
	}

	return tx.Commit()

}

func (h *Handler) onboarder() {

	for {

		// ensure logs are inserted in order
		r, more := <-h.logQueue

		h.queueMutex.Lock()

		if !more {
			close(h.closed)
			h.queueMutex.Unlock()
			return
		}

		h.queuedLogs = append(h.queuedLogs, *r)

		h.queueMutex.Unlock()

	}

}

func (h *Handler) dispatcher() {

	var closed bool

	ticker := time.NewTicker(time.Millisecond * time.Duration(h.args.InsertFrequencyMilliSeconds))
	defer ticker.Stop()

	for !closed {

		select {
		case <-h.closed:
			closed = true
		case <-ticker.C:
		}

		h.queueMutex.Lock()

		var err error
		rowValues := make([]string, 0)
		vals := make([]interface{}, 0)

		if len(h.queuedLogs) == 0 {
			goto nextIter
		}

		for i, msg := range h.queuedLogs {

			ctxMap := make(map[string]string, 0)
			for i, c := range msg.Ctx {
				if i%2 == 1 {
					ctxMap[fmt.Sprintf("%s", msg.Ctx[i-1])] = fmt.Sprintf("%v", c)
				}
			}

			b, err := json.Marshal(ctxMap)
			if err != nil {
				fmt.Printf("(todo: improve this log!) %s", err.Error())
			}

			idx := i * 6
			if h.args.InstanceID != "" {
				rowValues = append(rowValues, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)\n", idx+1, idx+2, idx+3, idx+4, idx+5, idx+6))
				vals = append(vals, h.args.Namespace, h.args.InstanceID, msg.Time.UnixNano(), msg.Lvl, msg.Msg, fmt.Sprintf("%s", b))
				err = h.brokerManager.Publish(h.args.InstanceID, dlog.LogEntry{
					Level:     msg.Lvl.String(),
					Timestamp: msg.Time.UnixNano(),
					Message:   msg.Msg,
					Context:   ctxMap,
				})
				if err != nil {
					fmt.Printf("(todo: improve this log!!!) %s\n", err.Error())
				}
			} else {
				rowValues = append(rowValues, fmt.Sprintf("($%d, $%d, $%d,  $%d, $%d)\n", idx+1, idx+2, idx+3, idx+4, idx+5))
				vals = append(vals, h.args.Namespace, msg.Time.UnixNano(), msg.Lvl, msg.Msg, fmt.Sprintf("%s", b))
			}

		}

		if h.args.InstanceID != "" {
			_, err = h.db.Exec(fmt.Sprintf("insert into logs (namespace, instance, time, lvl, msg, ctx) values %s", strings.Join(rowValues, ", ")), vals...)
			if err != nil {
				fmt.Printf("(todo: improve this log!) %s", err.Error())
			}
		} else {
			_, err = h.db.Exec(fmt.Sprintf("insert into logs (namespace, time, lvl, msg, ctx) values %s", strings.Join(rowValues, ", ")), vals...)
			if err != nil {
				fmt.Printf("(todo: improve this log!) %s", err.Error())
			}
		}

	nextIter:
		h.queuedLogs = h.queuedLogs[:0]
		h.queueMutex.Unlock()

	}

}

func (h *Handler) init() (*Handler, error) {

	h.queuedLogs = make([]log15.Record, 0)
	h.logQueue = make(chan *log15.Record, 100)
	h.closed = make(chan bool)

	go h.onboarder()
	go h.dispatcher()

	return h, nil

}

func (h *Handler) Log(r *log15.Record) error {
	h.logQueue <- r
	return nil
}

func (h *Handler) Close() error {

	defer func() {
		_ = recover()
	}()

	close(h.logQueue)
	return nil

}
