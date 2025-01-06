// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gzorm/commons/ent/predicate"
	"github.com/gzorm/commons/ent/winuser"
)

// WinUserDelete is the builder for deleting a WinUser entity.
type WinUserDelete struct {
	config
	hooks    []Hook
	mutation *WinUserMutation
}

// Where appends a list predicates to the WinUserDelete builder.
func (wud *WinUserDelete) Where(ps ...predicate.WinUser) *WinUserDelete {
	wud.mutation.Where(ps...)
	return wud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wud *WinUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wud.sqlExec, wud.mutation, wud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wud *WinUserDelete) ExecX(ctx context.Context) int {
	n, err := wud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wud *WinUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(winuser.Table, sqlgraph.NewFieldSpec(winuser.FieldID, field.TypeInt32))
	if ps := wud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wud.mutation.done = true
	return affected, err
}

// WinUserDeleteOne is the builder for deleting a single WinUser entity.
type WinUserDeleteOne struct {
	wud *WinUserDelete
}

// Where appends a list predicates to the WinUserDelete builder.
func (wudo *WinUserDeleteOne) Where(ps ...predicate.WinUser) *WinUserDeleteOne {
	wudo.wud.mutation.Where(ps...)
	return wudo
}

// Exec executes the deletion query.
func (wudo *WinUserDeleteOne) Exec(ctx context.Context) error {
	n, err := wudo.wud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{winuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wudo *WinUserDeleteOne) ExecX(ctx context.Context) {
	if err := wudo.Exec(ctx); err != nil {
		panic(err)
	}
}
