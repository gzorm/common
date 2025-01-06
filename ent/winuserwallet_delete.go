// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gzorm/commons/ent/predicate"
	"github.com/gzorm/commons/ent/winuserwallet"
)

// WinUserWalletDelete is the builder for deleting a WinUserWallet entity.
type WinUserWalletDelete struct {
	config
	hooks    []Hook
	mutation *WinUserWalletMutation
}

// Where appends a list predicates to the WinUserWalletDelete builder.
func (wuwd *WinUserWalletDelete) Where(ps ...predicate.WinUserWallet) *WinUserWalletDelete {
	wuwd.mutation.Where(ps...)
	return wuwd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wuwd *WinUserWalletDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wuwd.sqlExec, wuwd.mutation, wuwd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wuwd *WinUserWalletDelete) ExecX(ctx context.Context) int {
	n, err := wuwd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wuwd *WinUserWalletDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(winuserwallet.Table, sqlgraph.NewFieldSpec(winuserwallet.FieldID, field.TypeInt32))
	if ps := wuwd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wuwd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wuwd.mutation.done = true
	return affected, err
}

// WinUserWalletDeleteOne is the builder for deleting a single WinUserWallet entity.
type WinUserWalletDeleteOne struct {
	wuwd *WinUserWalletDelete
}

// Where appends a list predicates to the WinUserWalletDelete builder.
func (wuwdo *WinUserWalletDeleteOne) Where(ps ...predicate.WinUserWallet) *WinUserWalletDeleteOne {
	wuwdo.wuwd.mutation.Where(ps...)
	return wuwdo
}

// Exec executes the deletion query.
func (wuwdo *WinUserWalletDeleteOne) Exec(ctx context.Context) error {
	n, err := wuwdo.wuwd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{winuserwallet.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wuwdo *WinUserWalletDeleteOne) ExecX(ctx context.Context) {
	if err := wuwdo.Exec(ctx); err != nil {
		panic(err)
	}
}
