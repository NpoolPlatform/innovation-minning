// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/launchtime"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/predicate"
)

// LaunchTimeDelete is the builder for deleting a LaunchTime entity.
type LaunchTimeDelete struct {
	config
	hooks    []Hook
	mutation *LaunchTimeMutation
}

// Where appends a list predicates to the LaunchTimeDelete builder.
func (ltd *LaunchTimeDelete) Where(ps ...predicate.LaunchTime) *LaunchTimeDelete {
	ltd.mutation.Where(ps...)
	return ltd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ltd *LaunchTimeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ltd.hooks) == 0 {
		affected, err = ltd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LaunchTimeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ltd.mutation = mutation
			affected, err = ltd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ltd.hooks) - 1; i >= 0; i-- {
			if ltd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ltd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ltd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltd *LaunchTimeDelete) ExecX(ctx context.Context) int {
	n, err := ltd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ltd *LaunchTimeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: launchtime.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: launchtime.FieldID,
			},
		},
	}
	if ps := ltd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ltd.driver, _spec)
}

// LaunchTimeDeleteOne is the builder for deleting a single LaunchTime entity.
type LaunchTimeDeleteOne struct {
	ltd *LaunchTimeDelete
}

// Exec executes the deletion query.
func (ltdo *LaunchTimeDeleteOne) Exec(ctx context.Context) error {
	n, err := ltdo.ltd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{launchtime.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ltdo *LaunchTimeDeleteOne) ExecX(ctx context.Context) {
	ltdo.ltd.ExecX(ctx)
}
