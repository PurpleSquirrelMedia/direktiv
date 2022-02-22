// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/secrets/ent/namespacesecret"
	"github.com/direktiv/direktiv/pkg/secrets/ent/predicate"
)

// NamespaceSecretUpdate is the builder for updating NamespaceSecret entities.
type NamespaceSecretUpdate struct {
	config
	hooks    []Hook
	mutation *NamespaceSecretMutation
}

// Where appends a list predicates to the NamespaceSecretUpdate builder.
func (nsu *NamespaceSecretUpdate) Where(ps ...predicate.NamespaceSecret) *NamespaceSecretUpdate {
	nsu.mutation.Where(ps...)
	return nsu
}

// SetNs sets the "ns" field.
func (nsu *NamespaceSecretUpdate) SetNs(s string) *NamespaceSecretUpdate {
	nsu.mutation.SetNs(s)
	return nsu
}

// SetName sets the "name" field.
func (nsu *NamespaceSecretUpdate) SetName(s string) *NamespaceSecretUpdate {
	nsu.mutation.SetName(s)
	return nsu
}

// SetSecret sets the "secret" field.
func (nsu *NamespaceSecretUpdate) SetSecret(b []byte) *NamespaceSecretUpdate {
	nsu.mutation.SetSecret(b)
	return nsu
}

// Mutation returns the NamespaceSecretMutation object of the builder.
func (nsu *NamespaceSecretUpdate) Mutation() *NamespaceSecretMutation {
	return nsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nsu *NamespaceSecretUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nsu.hooks) == 0 {
		if err = nsu.check(); err != nil {
			return 0, err
		}
		affected, err = nsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NamespaceSecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nsu.check(); err != nil {
				return 0, err
			}
			nsu.mutation = mutation
			affected, err = nsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nsu.hooks) - 1; i >= 0; i-- {
			if nsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nsu *NamespaceSecretUpdate) SaveX(ctx context.Context) int {
	affected, err := nsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nsu *NamespaceSecretUpdate) Exec(ctx context.Context) error {
	_, err := nsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsu *NamespaceSecretUpdate) ExecX(ctx context.Context) {
	if err := nsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nsu *NamespaceSecretUpdate) check() error {
	if v, ok := nsu.mutation.Secret(); ok {
		if err := namespacesecret.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`ent: validator failed for field "NamespaceSecret.secret": %w`, err)}
		}
	}
	return nil
}

func (nsu *NamespaceSecretUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   namespacesecret.Table,
			Columns: namespacesecret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: namespacesecret.FieldID,
			},
		},
	}
	if ps := nsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nsu.mutation.Ns(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: namespacesecret.FieldNs,
		})
	}
	if value, ok := nsu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: namespacesecret.FieldName,
		})
	}
	if value, ok := nsu.mutation.Secret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: namespacesecret.FieldSecret,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{namespacesecret.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NamespaceSecretUpdateOne is the builder for updating a single NamespaceSecret entity.
type NamespaceSecretUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NamespaceSecretMutation
}

// SetNs sets the "ns" field.
func (nsuo *NamespaceSecretUpdateOne) SetNs(s string) *NamespaceSecretUpdateOne {
	nsuo.mutation.SetNs(s)
	return nsuo
}

// SetName sets the "name" field.
func (nsuo *NamespaceSecretUpdateOne) SetName(s string) *NamespaceSecretUpdateOne {
	nsuo.mutation.SetName(s)
	return nsuo
}

// SetSecret sets the "secret" field.
func (nsuo *NamespaceSecretUpdateOne) SetSecret(b []byte) *NamespaceSecretUpdateOne {
	nsuo.mutation.SetSecret(b)
	return nsuo
}

// Mutation returns the NamespaceSecretMutation object of the builder.
func (nsuo *NamespaceSecretUpdateOne) Mutation() *NamespaceSecretMutation {
	return nsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nsuo *NamespaceSecretUpdateOne) Select(field string, fields ...string) *NamespaceSecretUpdateOne {
	nsuo.fields = append([]string{field}, fields...)
	return nsuo
}

// Save executes the query and returns the updated NamespaceSecret entity.
func (nsuo *NamespaceSecretUpdateOne) Save(ctx context.Context) (*NamespaceSecret, error) {
	var (
		err  error
		node *NamespaceSecret
	)
	if len(nsuo.hooks) == 0 {
		if err = nsuo.check(); err != nil {
			return nil, err
		}
		node, err = nsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NamespaceSecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nsuo.check(); err != nil {
				return nil, err
			}
			nsuo.mutation = mutation
			node, err = nsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nsuo.hooks) - 1; i >= 0; i-- {
			if nsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nsuo *NamespaceSecretUpdateOne) SaveX(ctx context.Context) *NamespaceSecret {
	node, err := nsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nsuo *NamespaceSecretUpdateOne) Exec(ctx context.Context) error {
	_, err := nsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsuo *NamespaceSecretUpdateOne) ExecX(ctx context.Context) {
	if err := nsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nsuo *NamespaceSecretUpdateOne) check() error {
	if v, ok := nsuo.mutation.Secret(); ok {
		if err := namespacesecret.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`ent: validator failed for field "NamespaceSecret.secret": %w`, err)}
		}
	}
	return nil
}

func (nsuo *NamespaceSecretUpdateOne) sqlSave(ctx context.Context) (_node *NamespaceSecret, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   namespacesecret.Table,
			Columns: namespacesecret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: namespacesecret.FieldID,
			},
		},
	}
	id, ok := nsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NamespaceSecret.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, namespacesecret.FieldID)
		for _, f := range fields {
			if !namespacesecret.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != namespacesecret.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nsuo.mutation.Ns(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: namespacesecret.FieldNs,
		})
	}
	if value, ok := nsuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: namespacesecret.FieldName,
		})
	}
	if value, ok := nsuo.mutation.Secret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: namespacesecret.FieldSecret,
		})
	}
	_node = &NamespaceSecret{config: nsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{namespacesecret.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
