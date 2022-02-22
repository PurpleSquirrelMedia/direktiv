// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/direktiv/direktiv/pkg/flow/ent/ref"
	"github.com/direktiv/direktiv/pkg/flow/ent/route"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// RouteUpdate is the builder for updating Route entities.
type RouteUpdate struct {
	config
	hooks    []Hook
	mutation *RouteMutation
}

// Where appends a list predicates to the RouteUpdate builder.
func (ru *RouteUpdate) Where(ps ...predicate.Route) *RouteUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (ru *RouteUpdate) SetWorkflowID(id uuid.UUID) *RouteUpdate {
	ru.mutation.SetWorkflowID(id)
	return ru
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (ru *RouteUpdate) SetWorkflow(w *Workflow) *RouteUpdate {
	return ru.SetWorkflowID(w.ID)
}

// SetRefID sets the "ref" edge to the Ref entity by ID.
func (ru *RouteUpdate) SetRefID(id uuid.UUID) *RouteUpdate {
	ru.mutation.SetRefID(id)
	return ru
}

// SetRef sets the "ref" edge to the Ref entity.
func (ru *RouteUpdate) SetRef(r *Ref) *RouteUpdate {
	return ru.SetRefID(r.ID)
}

// Mutation returns the RouteMutation object of the builder.
func (ru *RouteUpdate) Mutation() *RouteMutation {
	return ru.mutation
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (ru *RouteUpdate) ClearWorkflow() *RouteUpdate {
	ru.mutation.ClearWorkflow()
	return ru
}

// ClearRef clears the "ref" edge to the Ref entity.
func (ru *RouteUpdate) ClearRef() *RouteUpdate {
	ru.mutation.ClearRef()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RouteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RouteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RouteUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RouteUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RouteUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RouteUpdate) check() error {
	if _, ok := ru.mutation.WorkflowID(); ru.mutation.WorkflowCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Route.workflow"`)
	}
	if _, ok := ru.mutation.RefID(); ru.mutation.RefCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Route.ref"`)
	}
	return nil
}

func (ru *RouteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   route.Table,
			Columns: route.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: route.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ru.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.WorkflowTable,
			Columns: []string{route.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.WorkflowTable,
			Columns: []string{route.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.RefTable,
			Columns: []string{route.RefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ref.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.RefTable,
			Columns: []string{route.RefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ref.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{route.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RouteUpdateOne is the builder for updating a single Route entity.
type RouteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RouteMutation
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (ruo *RouteUpdateOne) SetWorkflowID(id uuid.UUID) *RouteUpdateOne {
	ruo.mutation.SetWorkflowID(id)
	return ruo
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (ruo *RouteUpdateOne) SetWorkflow(w *Workflow) *RouteUpdateOne {
	return ruo.SetWorkflowID(w.ID)
}

// SetRefID sets the "ref" edge to the Ref entity by ID.
func (ruo *RouteUpdateOne) SetRefID(id uuid.UUID) *RouteUpdateOne {
	ruo.mutation.SetRefID(id)
	return ruo
}

// SetRef sets the "ref" edge to the Ref entity.
func (ruo *RouteUpdateOne) SetRef(r *Ref) *RouteUpdateOne {
	return ruo.SetRefID(r.ID)
}

// Mutation returns the RouteMutation object of the builder.
func (ruo *RouteUpdateOne) Mutation() *RouteMutation {
	return ruo.mutation
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (ruo *RouteUpdateOne) ClearWorkflow() *RouteUpdateOne {
	ruo.mutation.ClearWorkflow()
	return ruo
}

// ClearRef clears the "ref" edge to the Ref entity.
func (ruo *RouteUpdateOne) ClearRef() *RouteUpdateOne {
	ruo.mutation.ClearRef()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RouteUpdateOne) Select(field string, fields ...string) *RouteUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Route entity.
func (ruo *RouteUpdateOne) Save(ctx context.Context) (*Route, error) {
	var (
		err  error
		node *Route
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RouteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RouteUpdateOne) SaveX(ctx context.Context) *Route {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RouteUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RouteUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RouteUpdateOne) check() error {
	if _, ok := ruo.mutation.WorkflowID(); ruo.mutation.WorkflowCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Route.workflow"`)
	}
	if _, ok := ruo.mutation.RefID(); ruo.mutation.RefCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Route.ref"`)
	}
	return nil
}

func (ruo *RouteUpdateOne) sqlSave(ctx context.Context) (_node *Route, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   route.Table,
			Columns: route.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: route.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Route.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, route.FieldID)
		for _, f := range fields {
			if !route.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != route.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ruo.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.WorkflowTable,
			Columns: []string{route.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.WorkflowTable,
			Columns: []string{route.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.RefTable,
			Columns: []string{route.RefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ref.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.RefTable,
			Columns: []string{route.RefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ref.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Route{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{route.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
