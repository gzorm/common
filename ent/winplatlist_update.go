// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gzorm/commons/ent/predicate"
	"github.com/gzorm/commons/ent/winplatlist"
)

// WinPlatListUpdate is the builder for updating WinPlatList entities.
type WinPlatListUpdate struct {
	config
	hooks    []Hook
	mutation *WinPlatListMutation
}

// Where appends a list predicates to the WinPlatListUpdate builder.
func (wplu *WinPlatListUpdate) Where(ps ...predicate.WinPlatList) *WinPlatListUpdate {
	wplu.mutation.Where(ps...)
	return wplu
}

// SetCode sets the "code" field.
func (wplu *WinPlatListUpdate) SetCode(s string) *WinPlatListUpdate {
	wplu.mutation.SetCode(s)
	return wplu
}

// SetName sets the "name" field.
func (wplu *WinPlatListUpdate) SetName(s string) *WinPlatListUpdate {
	wplu.mutation.SetName(s)
	return wplu
}

// SetConfig sets the "config" field.
func (wplu *WinPlatListUpdate) SetConfig(s string) *WinPlatListUpdate {
	wplu.mutation.SetConfig(s)
	return wplu
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (wplu *WinPlatListUpdate) SetNillableConfig(s *string) *WinPlatListUpdate {
	if s != nil {
		wplu.SetConfig(*s)
	}
	return wplu
}

// ClearConfig clears the value of the "config" field.
func (wplu *WinPlatListUpdate) ClearConfig() *WinPlatListUpdate {
	wplu.mutation.ClearConfig()
	return wplu
}

// SetRate sets the "rate" field.
func (wplu *WinPlatListUpdate) SetRate(s string) *WinPlatListUpdate {
	wplu.mutation.SetRate(s)
	return wplu
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (wplu *WinPlatListUpdate) SetNillableRate(s *string) *WinPlatListUpdate {
	if s != nil {
		wplu.SetRate(*s)
	}
	return wplu
}

// ClearRate clears the value of the "rate" field.
func (wplu *WinPlatListUpdate) ClearRate() *WinPlatListUpdate {
	wplu.mutation.ClearRate()
	return wplu
}

// SetSort sets the "sort" field.
func (wplu *WinPlatListUpdate) SetSort(i int8) *WinPlatListUpdate {
	wplu.mutation.ResetSort()
	wplu.mutation.SetSort(i)
	return wplu
}

// AddSort adds i to the "sort" field.
func (wplu *WinPlatListUpdate) AddSort(i int8) *WinPlatListUpdate {
	wplu.mutation.AddSort(i)
	return wplu
}

// SetStatus sets the "status" field.
func (wplu *WinPlatListUpdate) SetStatus(i int8) *WinPlatListUpdate {
	wplu.mutation.ResetStatus()
	wplu.mutation.SetStatus(i)
	return wplu
}

// AddStatus adds i to the "status" field.
func (wplu *WinPlatListUpdate) AddStatus(i int8) *WinPlatListUpdate {
	wplu.mutation.AddStatus(i)
	return wplu
}

// SetCreatedAt sets the "created_at" field.
func (wplu *WinPlatListUpdate) SetCreatedAt(i int32) *WinPlatListUpdate {
	wplu.mutation.ResetCreatedAt()
	wplu.mutation.SetCreatedAt(i)
	return wplu
}

// AddCreatedAt adds i to the "created_at" field.
func (wplu *WinPlatListUpdate) AddCreatedAt(i int32) *WinPlatListUpdate {
	wplu.mutation.AddCreatedAt(i)
	return wplu
}

// SetUpdatedAt sets the "updated_at" field.
func (wplu *WinPlatListUpdate) SetUpdatedAt(i int32) *WinPlatListUpdate {
	wplu.mutation.ResetUpdatedAt()
	wplu.mutation.SetUpdatedAt(i)
	return wplu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (wplu *WinPlatListUpdate) AddUpdatedAt(i int32) *WinPlatListUpdate {
	wplu.mutation.AddUpdatedAt(i)
	return wplu
}

// Mutation returns the WinPlatListMutation object of the builder.
func (wplu *WinPlatListUpdate) Mutation() *WinPlatListMutation {
	return wplu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wplu *WinPlatListUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wplu.sqlSave, wplu.mutation, wplu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wplu *WinPlatListUpdate) SaveX(ctx context.Context) int {
	affected, err := wplu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wplu *WinPlatListUpdate) Exec(ctx context.Context) error {
	_, err := wplu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wplu *WinPlatListUpdate) ExecX(ctx context.Context) {
	if err := wplu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wplu *WinPlatListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(winplatlist.Table, winplatlist.Columns, sqlgraph.NewFieldSpec(winplatlist.FieldID, field.TypeInt32))
	if ps := wplu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wplu.mutation.Code(); ok {
		_spec.SetField(winplatlist.FieldCode, field.TypeString, value)
	}
	if value, ok := wplu.mutation.Name(); ok {
		_spec.SetField(winplatlist.FieldName, field.TypeString, value)
	}
	if value, ok := wplu.mutation.Config(); ok {
		_spec.SetField(winplatlist.FieldConfig, field.TypeString, value)
	}
	if wplu.mutation.ConfigCleared() {
		_spec.ClearField(winplatlist.FieldConfig, field.TypeString)
	}
	if value, ok := wplu.mutation.Rate(); ok {
		_spec.SetField(winplatlist.FieldRate, field.TypeString, value)
	}
	if wplu.mutation.RateCleared() {
		_spec.ClearField(winplatlist.FieldRate, field.TypeString)
	}
	if value, ok := wplu.mutation.Sort(); ok {
		_spec.SetField(winplatlist.FieldSort, field.TypeInt8, value)
	}
	if value, ok := wplu.mutation.AddedSort(); ok {
		_spec.AddField(winplatlist.FieldSort, field.TypeInt8, value)
	}
	if value, ok := wplu.mutation.Status(); ok {
		_spec.SetField(winplatlist.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := wplu.mutation.AddedStatus(); ok {
		_spec.AddField(winplatlist.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := wplu.mutation.CreatedAt(); ok {
		_spec.SetField(winplatlist.FieldCreatedAt, field.TypeInt32, value)
	}
	if value, ok := wplu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(winplatlist.FieldCreatedAt, field.TypeInt32, value)
	}
	if value, ok := wplu.mutation.UpdatedAt(); ok {
		_spec.SetField(winplatlist.FieldUpdatedAt, field.TypeInt32, value)
	}
	if value, ok := wplu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(winplatlist.FieldUpdatedAt, field.TypeInt32, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wplu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{winplatlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wplu.mutation.done = true
	return n, nil
}

// WinPlatListUpdateOne is the builder for updating a single WinPlatList entity.
type WinPlatListUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WinPlatListMutation
}

// SetCode sets the "code" field.
func (wpluo *WinPlatListUpdateOne) SetCode(s string) *WinPlatListUpdateOne {
	wpluo.mutation.SetCode(s)
	return wpluo
}

// SetName sets the "name" field.
func (wpluo *WinPlatListUpdateOne) SetName(s string) *WinPlatListUpdateOne {
	wpluo.mutation.SetName(s)
	return wpluo
}

// SetConfig sets the "config" field.
func (wpluo *WinPlatListUpdateOne) SetConfig(s string) *WinPlatListUpdateOne {
	wpluo.mutation.SetConfig(s)
	return wpluo
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (wpluo *WinPlatListUpdateOne) SetNillableConfig(s *string) *WinPlatListUpdateOne {
	if s != nil {
		wpluo.SetConfig(*s)
	}
	return wpluo
}

// ClearConfig clears the value of the "config" field.
func (wpluo *WinPlatListUpdateOne) ClearConfig() *WinPlatListUpdateOne {
	wpluo.mutation.ClearConfig()
	return wpluo
}

// SetRate sets the "rate" field.
func (wpluo *WinPlatListUpdateOne) SetRate(s string) *WinPlatListUpdateOne {
	wpluo.mutation.SetRate(s)
	return wpluo
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (wpluo *WinPlatListUpdateOne) SetNillableRate(s *string) *WinPlatListUpdateOne {
	if s != nil {
		wpluo.SetRate(*s)
	}
	return wpluo
}

// ClearRate clears the value of the "rate" field.
func (wpluo *WinPlatListUpdateOne) ClearRate() *WinPlatListUpdateOne {
	wpluo.mutation.ClearRate()
	return wpluo
}

// SetSort sets the "sort" field.
func (wpluo *WinPlatListUpdateOne) SetSort(i int8) *WinPlatListUpdateOne {
	wpluo.mutation.ResetSort()
	wpluo.mutation.SetSort(i)
	return wpluo
}

// AddSort adds i to the "sort" field.
func (wpluo *WinPlatListUpdateOne) AddSort(i int8) *WinPlatListUpdateOne {
	wpluo.mutation.AddSort(i)
	return wpluo
}

// SetStatus sets the "status" field.
func (wpluo *WinPlatListUpdateOne) SetStatus(i int8) *WinPlatListUpdateOne {
	wpluo.mutation.ResetStatus()
	wpluo.mutation.SetStatus(i)
	return wpluo
}

// AddStatus adds i to the "status" field.
func (wpluo *WinPlatListUpdateOne) AddStatus(i int8) *WinPlatListUpdateOne {
	wpluo.mutation.AddStatus(i)
	return wpluo
}

// SetCreatedAt sets the "created_at" field.
func (wpluo *WinPlatListUpdateOne) SetCreatedAt(i int32) *WinPlatListUpdateOne {
	wpluo.mutation.ResetCreatedAt()
	wpluo.mutation.SetCreatedAt(i)
	return wpluo
}

// AddCreatedAt adds i to the "created_at" field.
func (wpluo *WinPlatListUpdateOne) AddCreatedAt(i int32) *WinPlatListUpdateOne {
	wpluo.mutation.AddCreatedAt(i)
	return wpluo
}

// SetUpdatedAt sets the "updated_at" field.
func (wpluo *WinPlatListUpdateOne) SetUpdatedAt(i int32) *WinPlatListUpdateOne {
	wpluo.mutation.ResetUpdatedAt()
	wpluo.mutation.SetUpdatedAt(i)
	return wpluo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (wpluo *WinPlatListUpdateOne) AddUpdatedAt(i int32) *WinPlatListUpdateOne {
	wpluo.mutation.AddUpdatedAt(i)
	return wpluo
}

// Mutation returns the WinPlatListMutation object of the builder.
func (wpluo *WinPlatListUpdateOne) Mutation() *WinPlatListMutation {
	return wpluo.mutation
}

// Where appends a list predicates to the WinPlatListUpdate builder.
func (wpluo *WinPlatListUpdateOne) Where(ps ...predicate.WinPlatList) *WinPlatListUpdateOne {
	wpluo.mutation.Where(ps...)
	return wpluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wpluo *WinPlatListUpdateOne) Select(field string, fields ...string) *WinPlatListUpdateOne {
	wpluo.fields = append([]string{field}, fields...)
	return wpluo
}

// Save executes the query and returns the updated WinPlatList entity.
func (wpluo *WinPlatListUpdateOne) Save(ctx context.Context) (*WinPlatList, error) {
	return withHooks(ctx, wpluo.sqlSave, wpluo.mutation, wpluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wpluo *WinPlatListUpdateOne) SaveX(ctx context.Context) *WinPlatList {
	node, err := wpluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wpluo *WinPlatListUpdateOne) Exec(ctx context.Context) error {
	_, err := wpluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpluo *WinPlatListUpdateOne) ExecX(ctx context.Context) {
	if err := wpluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wpluo *WinPlatListUpdateOne) sqlSave(ctx context.Context) (_node *WinPlatList, err error) {
	_spec := sqlgraph.NewUpdateSpec(winplatlist.Table, winplatlist.Columns, sqlgraph.NewFieldSpec(winplatlist.FieldID, field.TypeInt32))
	id, ok := wpluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WinPlatList.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wpluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, winplatlist.FieldID)
		for _, f := range fields {
			if !winplatlist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != winplatlist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wpluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wpluo.mutation.Code(); ok {
		_spec.SetField(winplatlist.FieldCode, field.TypeString, value)
	}
	if value, ok := wpluo.mutation.Name(); ok {
		_spec.SetField(winplatlist.FieldName, field.TypeString, value)
	}
	if value, ok := wpluo.mutation.Config(); ok {
		_spec.SetField(winplatlist.FieldConfig, field.TypeString, value)
	}
	if wpluo.mutation.ConfigCleared() {
		_spec.ClearField(winplatlist.FieldConfig, field.TypeString)
	}
	if value, ok := wpluo.mutation.Rate(); ok {
		_spec.SetField(winplatlist.FieldRate, field.TypeString, value)
	}
	if wpluo.mutation.RateCleared() {
		_spec.ClearField(winplatlist.FieldRate, field.TypeString)
	}
	if value, ok := wpluo.mutation.Sort(); ok {
		_spec.SetField(winplatlist.FieldSort, field.TypeInt8, value)
	}
	if value, ok := wpluo.mutation.AddedSort(); ok {
		_spec.AddField(winplatlist.FieldSort, field.TypeInt8, value)
	}
	if value, ok := wpluo.mutation.Status(); ok {
		_spec.SetField(winplatlist.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := wpluo.mutation.AddedStatus(); ok {
		_spec.AddField(winplatlist.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := wpluo.mutation.CreatedAt(); ok {
		_spec.SetField(winplatlist.FieldCreatedAt, field.TypeInt32, value)
	}
	if value, ok := wpluo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(winplatlist.FieldCreatedAt, field.TypeInt32, value)
	}
	if value, ok := wpluo.mutation.UpdatedAt(); ok {
		_spec.SetField(winplatlist.FieldUpdatedAt, field.TypeInt32, value)
	}
	if value, ok := wpluo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(winplatlist.FieldUpdatedAt, field.TypeInt32, value)
	}
	_node = &WinPlatList{config: wpluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wpluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{winplatlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wpluo.mutation.done = true
	return _node, nil
}
