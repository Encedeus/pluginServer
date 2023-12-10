// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Encedeus/pluginServer/ent/plugin"
	"github.com/Encedeus/pluginServer/ent/publication"
	"github.com/Encedeus/pluginServer/ent/source"
	"github.com/Encedeus/pluginServer/ent/user"
	"github.com/google/uuid"
)

// PluginCreate is the builder for creating a Plugin entity.
type PluginCreate struct {
	config
	mutation *PluginMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PluginCreate) SetName(s string) *PluginCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetOwnerID sets the "owner_id" field.
func (pc *PluginCreate) SetOwnerID(u uuid.UUID) *PluginCreate {
	pc.mutation.SetOwnerID(u)
	return pc
}

// SetSourceID sets the "source_id" field.
func (pc *PluginCreate) SetSourceID(i int) *PluginCreate {
	pc.mutation.SetSourceID(i)
	return pc
}

// SetID sets the "id" field.
func (pc *PluginCreate) SetID(u uuid.UUID) *PluginCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PluginCreate) SetNillableID(u *uuid.UUID) *PluginCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PluginCreate) SetOwner(u *User) *PluginCreate {
	return pc.SetOwnerID(u.ID)
}

// SetSource sets the "source" edge to the Source entity.
func (pc *PluginCreate) SetSource(s *Source) *PluginCreate {
	return pc.SetSourceID(s.ID)
}

// AddPublicationIDs adds the "publications" edge to the Publication entity by IDs.
func (pc *PluginCreate) AddPublicationIDs(ids ...int) *PluginCreate {
	pc.mutation.AddPublicationIDs(ids...)
	return pc
}

// AddPublications adds the "publications" edges to the Publication entity.
func (pc *PluginCreate) AddPublications(p ...*Publication) *PluginCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPublicationIDs(ids...)
}

// Mutation returns the PluginMutation object of the builder.
func (pc *PluginCreate) Mutation() *PluginMutation {
	return pc.mutation
}

// Save creates the Plugin in the database.
func (pc *PluginCreate) Save(ctx context.Context) (*Plugin, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PluginCreate) SaveX(ctx context.Context) *Plugin {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PluginCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PluginCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PluginCreate) defaults() {
	if _, ok := pc.mutation.ID(); !ok {
		v := plugin.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PluginCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Plugin.name"`)}
	}
	if _, ok := pc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Plugin.owner_id"`)}
	}
	if _, ok := pc.mutation.SourceID(); !ok {
		return &ValidationError{Name: "source_id", err: errors.New(`ent: missing required field "Plugin.source_id"`)}
	}
	if _, ok := pc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Plugin.owner"`)}
	}
	if _, ok := pc.mutation.SourceID(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`ent: missing required edge "Plugin.source"`)}
	}
	return nil
}

func (pc *PluginCreate) sqlSave(ctx context.Context) (*Plugin, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
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
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PluginCreate) createSpec() (*Plugin, *sqlgraph.CreateSpec) {
	var (
		_node = &Plugin{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(plugin.Table, sqlgraph.NewFieldSpec(plugin.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(plugin.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   plugin.OwnerTable,
			Columns: []string{plugin.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   plugin.SourceTable,
			Columns: []string{plugin.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(source.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SourceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PublicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   plugin.PublicationsTable,
			Columns: []string{plugin.PublicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(publication.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PluginCreateBulk is the builder for creating many Plugin entities in bulk.
type PluginCreateBulk struct {
	config
	err      error
	builders []*PluginCreate
}

// Save creates the Plugin entities in the database.
func (pcb *PluginCreateBulk) Save(ctx context.Context) ([]*Plugin, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Plugin, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PluginMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PluginCreateBulk) SaveX(ctx context.Context) []*Plugin {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PluginCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PluginCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
