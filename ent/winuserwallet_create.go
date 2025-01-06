// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gzorm/commons/ent/winuserwallet"
)

// WinUserWalletCreate is the builder for creating a WinUserWallet entity.
type WinUserWalletCreate struct {
	config
	mutation *WinUserWalletMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (wuwc *WinUserWalletCreate) SetUsername(s string) *WinUserWalletCreate {
	wuwc.mutation.SetUsername(s)
	return wuwc
}

// SetCoin sets the "coin" field.
func (wuwc *WinUserWalletCreate) SetCoin(f float64) *WinUserWalletCreate {
	wuwc.mutation.SetCoin(f)
	return wuwc
}

// SetVersion sets the "version" field.
func (wuwc *WinUserWalletCreate) SetVersion(i int) *WinUserWalletCreate {
	wuwc.mutation.SetVersion(i)
	return wuwc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (wuwc *WinUserWalletCreate) SetNillableVersion(i *int) *WinUserWalletCreate {
	if i != nil {
		wuwc.SetVersion(*i)
	}
	return wuwc
}

// SetModifyAt sets the "modify_at" field.
func (wuwc *WinUserWalletCreate) SetModifyAt(i int) *WinUserWalletCreate {
	wuwc.mutation.SetModifyAt(i)
	return wuwc
}

// SetCreatedAt sets the "created_at" field.
func (wuwc *WinUserWalletCreate) SetCreatedAt(i int32) *WinUserWalletCreate {
	wuwc.mutation.SetCreatedAt(i)
	return wuwc
}

// SetUpdatedAt sets the "updated_at" field.
func (wuwc *WinUserWalletCreate) SetUpdatedAt(i int32) *WinUserWalletCreate {
	wuwc.mutation.SetUpdatedAt(i)
	return wuwc
}

// SetID sets the "id" field.
func (wuwc *WinUserWalletCreate) SetID(i int32) *WinUserWalletCreate {
	wuwc.mutation.SetID(i)
	return wuwc
}

// Mutation returns the WinUserWalletMutation object of the builder.
func (wuwc *WinUserWalletCreate) Mutation() *WinUserWalletMutation {
	return wuwc.mutation
}

// Save creates the WinUserWallet in the database.
func (wuwc *WinUserWalletCreate) Save(ctx context.Context) (*WinUserWallet, error) {
	return withHooks(ctx, wuwc.sqlSave, wuwc.mutation, wuwc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wuwc *WinUserWalletCreate) SaveX(ctx context.Context) *WinUserWallet {
	v, err := wuwc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wuwc *WinUserWalletCreate) Exec(ctx context.Context) error {
	_, err := wuwc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuwc *WinUserWalletCreate) ExecX(ctx context.Context) {
	if err := wuwc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuwc *WinUserWalletCreate) check() error {
	if _, ok := wuwc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "WinUserWallet.username"`)}
	}
	if _, ok := wuwc.mutation.Coin(); !ok {
		return &ValidationError{Name: "coin", err: errors.New(`ent: missing required field "WinUserWallet.coin"`)}
	}
	if _, ok := wuwc.mutation.ModifyAt(); !ok {
		return &ValidationError{Name: "modify_at", err: errors.New(`ent: missing required field "WinUserWallet.modify_at"`)}
	}
	if _, ok := wuwc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WinUserWallet.created_at"`)}
	}
	if _, ok := wuwc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WinUserWallet.updated_at"`)}
	}
	return nil
}

func (wuwc *WinUserWalletCreate) sqlSave(ctx context.Context) (*WinUserWallet, error) {
	if err := wuwc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wuwc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wuwc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	wuwc.mutation.id = &_node.ID
	wuwc.mutation.done = true
	return _node, nil
}

func (wuwc *WinUserWalletCreate) createSpec() (*WinUserWallet, *sqlgraph.CreateSpec) {
	var (
		_node = &WinUserWallet{config: wuwc.config}
		_spec = sqlgraph.NewCreateSpec(winuserwallet.Table, sqlgraph.NewFieldSpec(winuserwallet.FieldID, field.TypeInt32))
	)
	if id, ok := wuwc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wuwc.mutation.Username(); ok {
		_spec.SetField(winuserwallet.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := wuwc.mutation.Coin(); ok {
		_spec.SetField(winuserwallet.FieldCoin, field.TypeFloat64, value)
		_node.Coin = value
	}
	if value, ok := wuwc.mutation.Version(); ok {
		_spec.SetField(winuserwallet.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := wuwc.mutation.ModifyAt(); ok {
		_spec.SetField(winuserwallet.FieldModifyAt, field.TypeInt, value)
		_node.ModifyAt = value
	}
	if value, ok := wuwc.mutation.CreatedAt(); ok {
		_spec.SetField(winuserwallet.FieldCreatedAt, field.TypeInt32, value)
		_node.CreatedAt = value
	}
	if value, ok := wuwc.mutation.UpdatedAt(); ok {
		_spec.SetField(winuserwallet.FieldUpdatedAt, field.TypeInt32, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// WinUserWalletCreateBulk is the builder for creating many WinUserWallet entities in bulk.
type WinUserWalletCreateBulk struct {
	config
	builders []*WinUserWalletCreate
}

// Save creates the WinUserWallet entities in the database.
func (wuwcb *WinUserWalletCreateBulk) Save(ctx context.Context) ([]*WinUserWallet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wuwcb.builders))
	nodes := make([]*WinUserWallet, len(wuwcb.builders))
	mutators := make([]Mutator, len(wuwcb.builders))
	for i := range wuwcb.builders {
		func(i int, root context.Context) {
			builder := wuwcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WinUserWalletMutation)
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
					_, err = mutators[i+1].Mutate(root, wuwcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wuwcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, wuwcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wuwcb *WinUserWalletCreateBulk) SaveX(ctx context.Context) []*WinUserWallet {
	v, err := wuwcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wuwcb *WinUserWalletCreateBulk) Exec(ctx context.Context) error {
	_, err := wuwcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuwcb *WinUserWalletCreateBulk) ExecX(ctx context.Context) {
	if err := wuwcb.Exec(ctx); err != nil {
		panic(err)
	}
}
