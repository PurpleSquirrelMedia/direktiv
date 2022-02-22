// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/vardata"
	"github.com/direktiv/direktiv/pkg/flow/ent/varref"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// VarRefCreate is the builder for creating a VarRef entity.
type VarRefCreate struct {
	config
	mutation *VarRefMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (vrc *VarRefCreate) SetName(s string) *VarRefCreate {
	vrc.mutation.SetName(s)
	return vrc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (vrc *VarRefCreate) SetNillableName(s *string) *VarRefCreate {
	if s != nil {
		vrc.SetName(*s)
	}
	return vrc
}

// SetBehaviour sets the "behaviour" field.
func (vrc *VarRefCreate) SetBehaviour(s string) *VarRefCreate {
	vrc.mutation.SetBehaviour(s)
	return vrc
}

// SetNillableBehaviour sets the "behaviour" field if the given value is not nil.
func (vrc *VarRefCreate) SetNillableBehaviour(s *string) *VarRefCreate {
	if s != nil {
		vrc.SetBehaviour(*s)
	}
	return vrc
}

// SetID sets the "id" field.
func (vrc *VarRefCreate) SetID(u uuid.UUID) *VarRefCreate {
	vrc.mutation.SetID(u)
	return vrc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vrc *VarRefCreate) SetNillableID(u *uuid.UUID) *VarRefCreate {
	if u != nil {
		vrc.SetID(*u)
	}
	return vrc
}

// SetVardataID sets the "vardata" edge to the VarData entity by ID.
func (vrc *VarRefCreate) SetVardataID(id uuid.UUID) *VarRefCreate {
	vrc.mutation.SetVardataID(id)
	return vrc
}

// SetVardata sets the "vardata" edge to the VarData entity.
func (vrc *VarRefCreate) SetVardata(v *VarData) *VarRefCreate {
	return vrc.SetVardataID(v.ID)
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (vrc *VarRefCreate) SetNamespaceID(id uuid.UUID) *VarRefCreate {
	vrc.mutation.SetNamespaceID(id)
	return vrc
}

// SetNillableNamespaceID sets the "namespace" edge to the Namespace entity by ID if the given value is not nil.
func (vrc *VarRefCreate) SetNillableNamespaceID(id *uuid.UUID) *VarRefCreate {
	if id != nil {
		vrc = vrc.SetNamespaceID(*id)
	}
	return vrc
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (vrc *VarRefCreate) SetNamespace(n *Namespace) *VarRefCreate {
	return vrc.SetNamespaceID(n.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (vrc *VarRefCreate) SetWorkflowID(id uuid.UUID) *VarRefCreate {
	vrc.mutation.SetWorkflowID(id)
	return vrc
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (vrc *VarRefCreate) SetNillableWorkflowID(id *uuid.UUID) *VarRefCreate {
	if id != nil {
		vrc = vrc.SetWorkflowID(*id)
	}
	return vrc
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (vrc *VarRefCreate) SetWorkflow(w *Workflow) *VarRefCreate {
	return vrc.SetWorkflowID(w.ID)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (vrc *VarRefCreate) SetInstanceID(id uuid.UUID) *VarRefCreate {
	vrc.mutation.SetInstanceID(id)
	return vrc
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (vrc *VarRefCreate) SetNillableInstanceID(id *uuid.UUID) *VarRefCreate {
	if id != nil {
		vrc = vrc.SetInstanceID(*id)
	}
	return vrc
}

// SetInstance sets the "instance" edge to the Instance entity.
func (vrc *VarRefCreate) SetInstance(i *Instance) *VarRefCreate {
	return vrc.SetInstanceID(i.ID)
}

// Mutation returns the VarRefMutation object of the builder.
func (vrc *VarRefCreate) Mutation() *VarRefMutation {
	return vrc.mutation
}

// Save creates the VarRef in the database.
func (vrc *VarRefCreate) Save(ctx context.Context) (*VarRef, error) {
	var (
		err  error
		node *VarRef
	)
	vrc.defaults()
	if len(vrc.hooks) == 0 {
		if err = vrc.check(); err != nil {
			return nil, err
		}
		node, err = vrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VarRefMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vrc.check(); err != nil {
				return nil, err
			}
			vrc.mutation = mutation
			if node, err = vrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vrc.hooks) - 1; i >= 0; i-- {
			if vrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vrc *VarRefCreate) SaveX(ctx context.Context) *VarRef {
	v, err := vrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vrc *VarRefCreate) Exec(ctx context.Context) error {
	_, err := vrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vrc *VarRefCreate) ExecX(ctx context.Context) {
	if err := vrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vrc *VarRefCreate) defaults() {
	if _, ok := vrc.mutation.ID(); !ok {
		v := varref.DefaultID()
		vrc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vrc *VarRefCreate) check() error {
	if v, ok := vrc.mutation.Name(); ok {
		if err := varref.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "VarRef.name": %w`, err)}
		}
	}
	if _, ok := vrc.mutation.VardataID(); !ok {
		return &ValidationError{Name: "vardata", err: errors.New(`ent: missing required edge "VarRef.vardata"`)}
	}
	return nil
}

func (vrc *VarRefCreate) sqlSave(ctx context.Context) (*VarRef, error) {
	_node, _spec := vrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (vrc *VarRefCreate) createSpec() (*VarRef, *sqlgraph.CreateSpec) {
	var (
		_node = &VarRef{config: vrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: varref.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: varref.FieldID,
			},
		}
	)
	if id, ok := vrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := vrc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: varref.FieldName,
		})
		_node.Name = value
	}
	if value, ok := vrc.mutation.Behaviour(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: varref.FieldBehaviour,
		})
		_node.Behaviour = value
	}
	if nodes := vrc.mutation.VardataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   varref.VardataTable,
			Columns: []string{varref.VardataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vardata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.var_data_varrefs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vrc.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   varref.NamespaceTable,
			Columns: []string{varref.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.namespace_vars = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vrc.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   varref.WorkflowTable,
			Columns: []string{varref.WorkflowColumn},
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
		_node.workflow_vars = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vrc.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   varref.InstanceTable,
			Columns: []string{varref.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.instance_vars = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VarRefCreateBulk is the builder for creating many VarRef entities in bulk.
type VarRefCreateBulk struct {
	config
	builders []*VarRefCreate
}

// Save creates the VarRef entities in the database.
func (vrcb *VarRefCreateBulk) Save(ctx context.Context) ([]*VarRef, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vrcb.builders))
	nodes := make([]*VarRef, len(vrcb.builders))
	mutators := make([]Mutator, len(vrcb.builders))
	for i := range vrcb.builders {
		func(i int, root context.Context) {
			builder := vrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VarRefMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vrcb *VarRefCreateBulk) SaveX(ctx context.Context) []*VarRef {
	v, err := vrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vrcb *VarRefCreateBulk) Exec(ctx context.Context) error {
	_, err := vrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vrcb *VarRefCreateBulk) ExecX(ctx context.Context) {
	if err := vrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
