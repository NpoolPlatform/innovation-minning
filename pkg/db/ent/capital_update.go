// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/innovation-mining/pkg/db/ent/capital"
	"github.com/NpoolPlatform/innovation-mining/pkg/db/ent/predicate"
)

// CapitalUpdate is the builder for updating Capital entities.
type CapitalUpdate struct {
	config
	hooks    []Hook
	mutation *CapitalMutation
}

// Where appends a list predicates to the CapitalUpdate builder.
func (cu *CapitalUpdate) Where(ps ...predicate.Capital) *CapitalUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CapitalUpdate) SetName(s string) *CapitalUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetIntroduction sets the "introduction" field.
func (cu *CapitalUpdate) SetIntroduction(s string) *CapitalUpdate {
	cu.mutation.SetIntroduction(s)
	return cu
}

// SetLogo sets the "logo" field.
func (cu *CapitalUpdate) SetLogo(s string) *CapitalUpdate {
	cu.mutation.SetLogo(s)
	return cu
}

// SetCreateAt sets the "create_at" field.
func (cu *CapitalUpdate) SetCreateAt(u uint32) *CapitalUpdate {
	cu.mutation.ResetCreateAt()
	cu.mutation.SetCreateAt(u)
	return cu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cu *CapitalUpdate) SetNillableCreateAt(u *uint32) *CapitalUpdate {
	if u != nil {
		cu.SetCreateAt(*u)
	}
	return cu
}

// AddCreateAt adds u to the "create_at" field.
func (cu *CapitalUpdate) AddCreateAt(u uint32) *CapitalUpdate {
	cu.mutation.AddCreateAt(u)
	return cu
}

// SetUpdateAt sets the "update_at" field.
func (cu *CapitalUpdate) SetUpdateAt(u uint32) *CapitalUpdate {
	cu.mutation.ResetUpdateAt()
	cu.mutation.SetUpdateAt(u)
	return cu
}

// AddUpdateAt adds u to the "update_at" field.
func (cu *CapitalUpdate) AddUpdateAt(u uint32) *CapitalUpdate {
	cu.mutation.AddUpdateAt(u)
	return cu
}

// SetDeleteAt sets the "delete_at" field.
func (cu *CapitalUpdate) SetDeleteAt(u uint32) *CapitalUpdate {
	cu.mutation.ResetDeleteAt()
	cu.mutation.SetDeleteAt(u)
	return cu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cu *CapitalUpdate) SetNillableDeleteAt(u *uint32) *CapitalUpdate {
	if u != nil {
		cu.SetDeleteAt(*u)
	}
	return cu
}

// AddDeleteAt adds u to the "delete_at" field.
func (cu *CapitalUpdate) AddDeleteAt(u uint32) *CapitalUpdate {
	cu.mutation.AddDeleteAt(u)
	return cu
}

// Mutation returns the CapitalMutation object of the builder.
func (cu *CapitalUpdate) Mutation() *CapitalMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CapitalUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CapitalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CapitalUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CapitalUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CapitalUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CapitalUpdate) defaults() {
	if _, ok := cu.mutation.UpdateAt(); !ok {
		v := capital.UpdateDefaultUpdateAt()
		cu.mutation.SetUpdateAt(v)
	}
}

func (cu *CapitalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   capital.Table,
			Columns: capital.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: capital.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capital.FieldName,
		})
	}
	if value, ok := cu.mutation.Introduction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capital.FieldIntroduction,
		})
	}
	if value, ok := cu.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capital.FieldLogo,
		})
	}
	if value, ok := cu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldCreateAt,
		})
	}
	if value, ok := cu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldCreateAt,
		})
	}
	if value, ok := cu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldUpdateAt,
		})
	}
	if value, ok := cu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldUpdateAt,
		})
	}
	if value, ok := cu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldDeleteAt,
		})
	}
	if value, ok := cu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{capital.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CapitalUpdateOne is the builder for updating a single Capital entity.
type CapitalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CapitalMutation
}

// SetName sets the "name" field.
func (cuo *CapitalUpdateOne) SetName(s string) *CapitalUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetIntroduction sets the "introduction" field.
func (cuo *CapitalUpdateOne) SetIntroduction(s string) *CapitalUpdateOne {
	cuo.mutation.SetIntroduction(s)
	return cuo
}

// SetLogo sets the "logo" field.
func (cuo *CapitalUpdateOne) SetLogo(s string) *CapitalUpdateOne {
	cuo.mutation.SetLogo(s)
	return cuo
}

// SetCreateAt sets the "create_at" field.
func (cuo *CapitalUpdateOne) SetCreateAt(u uint32) *CapitalUpdateOne {
	cuo.mutation.ResetCreateAt()
	cuo.mutation.SetCreateAt(u)
	return cuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cuo *CapitalUpdateOne) SetNillableCreateAt(u *uint32) *CapitalUpdateOne {
	if u != nil {
		cuo.SetCreateAt(*u)
	}
	return cuo
}

// AddCreateAt adds u to the "create_at" field.
func (cuo *CapitalUpdateOne) AddCreateAt(u uint32) *CapitalUpdateOne {
	cuo.mutation.AddCreateAt(u)
	return cuo
}

// SetUpdateAt sets the "update_at" field.
func (cuo *CapitalUpdateOne) SetUpdateAt(u uint32) *CapitalUpdateOne {
	cuo.mutation.ResetUpdateAt()
	cuo.mutation.SetUpdateAt(u)
	return cuo
}

// AddUpdateAt adds u to the "update_at" field.
func (cuo *CapitalUpdateOne) AddUpdateAt(u uint32) *CapitalUpdateOne {
	cuo.mutation.AddUpdateAt(u)
	return cuo
}

// SetDeleteAt sets the "delete_at" field.
func (cuo *CapitalUpdateOne) SetDeleteAt(u uint32) *CapitalUpdateOne {
	cuo.mutation.ResetDeleteAt()
	cuo.mutation.SetDeleteAt(u)
	return cuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cuo *CapitalUpdateOne) SetNillableDeleteAt(u *uint32) *CapitalUpdateOne {
	if u != nil {
		cuo.SetDeleteAt(*u)
	}
	return cuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (cuo *CapitalUpdateOne) AddDeleteAt(u uint32) *CapitalUpdateOne {
	cuo.mutation.AddDeleteAt(u)
	return cuo
}

// Mutation returns the CapitalMutation object of the builder.
func (cuo *CapitalUpdateOne) Mutation() *CapitalMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CapitalUpdateOne) Select(field string, fields ...string) *CapitalUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Capital entity.
func (cuo *CapitalUpdateOne) Save(ctx context.Context) (*Capital, error) {
	var (
		err  error
		node *Capital
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CapitalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CapitalUpdateOne) SaveX(ctx context.Context) *Capital {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CapitalUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CapitalUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CapitalUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateAt(); !ok {
		v := capital.UpdateDefaultUpdateAt()
		cuo.mutation.SetUpdateAt(v)
	}
}

func (cuo *CapitalUpdateOne) sqlSave(ctx context.Context) (_node *Capital, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   capital.Table,
			Columns: capital.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: capital.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Capital.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, capital.FieldID)
		for _, f := range fields {
			if !capital.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != capital.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capital.FieldName,
		})
	}
	if value, ok := cuo.mutation.Introduction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capital.FieldIntroduction,
		})
	}
	if value, ok := cuo.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capital.FieldLogo,
		})
	}
	if value, ok := cuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldCreateAt,
		})
	}
	if value, ok := cuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldCreateAt,
		})
	}
	if value, ok := cuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldUpdateAt,
		})
	}
	if value, ok := cuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldUpdateAt,
		})
	}
	if value, ok := cuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldDeleteAt,
		})
	}
	if value, ok := cuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldDeleteAt,
		})
	}
	_node = &Capital{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{capital.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}