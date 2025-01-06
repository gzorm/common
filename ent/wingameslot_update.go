// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gzorm/common/ent/predicate"
	"github.com/gzorm/common/ent/wingameslot"
)

// WinGameSlotUpdate is the builder for updating WinGameSlot entities.
type WinGameSlotUpdate struct {
	config
	hooks    []Hook
	mutation *WinGameSlotMutation
}

// Where appends a list predicates to the WinGameSlotUpdate builder.
func (wgsu *WinGameSlotUpdate) Where(ps ...predicate.WinGameSlot) *WinGameSlotUpdate {
	wgsu.mutation.Where(ps...)
	return wgsu
}

// SetGameID sets the "game_id" field.
func (wgsu *WinGameSlotUpdate) SetGameID(i int32) *WinGameSlotUpdate {
	wgsu.mutation.ResetGameID()
	wgsu.mutation.SetGameID(i)
	return wgsu
}

// AddGameID adds i to the "game_id" field.
func (wgsu *WinGameSlotUpdate) AddGameID(i int32) *WinGameSlotUpdate {
	wgsu.mutation.AddGameID(i)
	return wgsu
}

// SetGameGroupID sets the "game_group_id" field.
func (wgsu *WinGameSlotUpdate) SetGameGroupID(i int8) *WinGameSlotUpdate {
	wgsu.mutation.ResetGameGroupID()
	wgsu.mutation.SetGameGroupID(i)
	return wgsu
}

// AddGameGroupID adds i to the "game_group_id" field.
func (wgsu *WinGameSlotUpdate) AddGameGroupID(i int8) *WinGameSlotUpdate {
	wgsu.mutation.AddGameGroupID(i)
	return wgsu
}

// SetPlatID sets the "plat_id" field.
func (wgsu *WinGameSlotUpdate) SetPlatID(i int32) *WinGameSlotUpdate {
	wgsu.mutation.ResetPlatID()
	wgsu.mutation.SetPlatID(i)
	return wgsu
}

// AddPlatID adds i to the "plat_id" field.
func (wgsu *WinGameSlotUpdate) AddPlatID(i int32) *WinGameSlotUpdate {
	wgsu.mutation.AddPlatID(i)
	return wgsu
}

// SetProvider sets the "provider" field.
func (wgsu *WinGameSlotUpdate) SetProvider(s string) *WinGameSlotUpdate {
	wgsu.mutation.SetProvider(s)
	return wgsu
}

// SetNillableProvider sets the "provider" field if the given value is not nil.
func (wgsu *WinGameSlotUpdate) SetNillableProvider(s *string) *WinGameSlotUpdate {
	if s != nil {
		wgsu.SetProvider(*s)
	}
	return wgsu
}

// ClearProvider clears the value of the "provider" field.
func (wgsu *WinGameSlotUpdate) ClearProvider() *WinGameSlotUpdate {
	wgsu.mutation.ClearProvider()
	return wgsu
}

// SetName sets the "name" field.
func (wgsu *WinGameSlotUpdate) SetName(s string) *WinGameSlotUpdate {
	wgsu.mutation.SetName(s)
	return wgsu
}

// SetNameZh sets the "name_zh" field.
func (wgsu *WinGameSlotUpdate) SetNameZh(s string) *WinGameSlotUpdate {
	wgsu.mutation.SetNameZh(s)
	return wgsu
}

// SetImg sets the "img" field.
func (wgsu *WinGameSlotUpdate) SetImg(s string) *WinGameSlotUpdate {
	wgsu.mutation.SetImg(s)
	return wgsu
}

// SetImgNew sets the "img_new" field.
func (wgsu *WinGameSlotUpdate) SetImgNew(s string) *WinGameSlotUpdate {
	wgsu.mutation.SetImgNew(s)
	return wgsu
}

// SetNillableImgNew sets the "img_new" field if the given value is not nil.
func (wgsu *WinGameSlotUpdate) SetNillableImgNew(s *string) *WinGameSlotUpdate {
	if s != nil {
		wgsu.SetImgNew(*s)
	}
	return wgsu
}

// ClearImgNew clears the value of the "img_new" field.
func (wgsu *WinGameSlotUpdate) ClearImgNew() *WinGameSlotUpdate {
	wgsu.mutation.ClearImgNew()
	return wgsu
}

// SetIsNew sets the "is_new" field.
func (wgsu *WinGameSlotUpdate) SetIsNew(b bool) *WinGameSlotUpdate {
	wgsu.mutation.SetIsNew(b)
	return wgsu
}

// SetIsCasino sets the "is_casino" field.
func (wgsu *WinGameSlotUpdate) SetIsCasino(i int8) *WinGameSlotUpdate {
	wgsu.mutation.ResetIsCasino()
	wgsu.mutation.SetIsCasino(i)
	return wgsu
}

// AddIsCasino adds i to the "is_casino" field.
func (wgsu *WinGameSlotUpdate) AddIsCasino(i int8) *WinGameSlotUpdate {
	wgsu.mutation.AddIsCasino(i)
	return wgsu
}

// SetGameTypeID sets the "game_type_id" field.
func (wgsu *WinGameSlotUpdate) SetGameTypeID(s string) *WinGameSlotUpdate {
	wgsu.mutation.SetGameTypeID(s)
	return wgsu
}

// SetNillableGameTypeID sets the "game_type_id" field if the given value is not nil.
func (wgsu *WinGameSlotUpdate) SetNillableGameTypeID(s *string) *WinGameSlotUpdate {
	if s != nil {
		wgsu.SetGameTypeID(*s)
	}
	return wgsu
}

// ClearGameTypeID clears the value of the "game_type_id" field.
func (wgsu *WinGameSlotUpdate) ClearGameTypeID() *WinGameSlotUpdate {
	wgsu.mutation.ClearGameTypeID()
	return wgsu
}

// SetGameTypeName sets the "game_type_name" field.
func (wgsu *WinGameSlotUpdate) SetGameTypeName(s string) *WinGameSlotUpdate {
	wgsu.mutation.SetGameTypeName(s)
	return wgsu
}

// SetNillableGameTypeName sets the "game_type_name" field if the given value is not nil.
func (wgsu *WinGameSlotUpdate) SetNillableGameTypeName(s *string) *WinGameSlotUpdate {
	if s != nil {
		wgsu.SetGameTypeName(*s)
	}
	return wgsu
}

// ClearGameTypeName clears the value of the "game_type_name" field.
func (wgsu *WinGameSlotUpdate) ClearGameTypeName() *WinGameSlotUpdate {
	wgsu.mutation.ClearGameTypeName()
	return wgsu
}

// SetFavoriteStar sets the "favorite_star" field.
func (wgsu *WinGameSlotUpdate) SetFavoriteStar(i int32) *WinGameSlotUpdate {
	wgsu.mutation.ResetFavoriteStar()
	wgsu.mutation.SetFavoriteStar(i)
	return wgsu
}

// AddFavoriteStar adds i to the "favorite_star" field.
func (wgsu *WinGameSlotUpdate) AddFavoriteStar(i int32) *WinGameSlotUpdate {
	wgsu.mutation.AddFavoriteStar(i)
	return wgsu
}

// SetHotStar sets the "hot_star" field.
func (wgsu *WinGameSlotUpdate) SetHotStar(i int32) *WinGameSlotUpdate {
	wgsu.mutation.ResetHotStar()
	wgsu.mutation.SetHotStar(i)
	return wgsu
}

// AddHotStar adds i to the "hot_star" field.
func (wgsu *WinGameSlotUpdate) AddHotStar(i int32) *WinGameSlotUpdate {
	wgsu.mutation.AddHotStar(i)
	return wgsu
}

// SetSort sets the "sort" field.
func (wgsu *WinGameSlotUpdate) SetSort(i int32) *WinGameSlotUpdate {
	wgsu.mutation.ResetSort()
	wgsu.mutation.SetSort(i)
	return wgsu
}

// AddSort adds i to the "sort" field.
func (wgsu *WinGameSlotUpdate) AddSort(i int32) *WinGameSlotUpdate {
	wgsu.mutation.AddSort(i)
	return wgsu
}

// SetStatus sets the "status" field.
func (wgsu *WinGameSlotUpdate) SetStatus(i int8) *WinGameSlotUpdate {
	wgsu.mutation.ResetStatus()
	wgsu.mutation.SetStatus(i)
	return wgsu
}

// AddStatus adds i to the "status" field.
func (wgsu *WinGameSlotUpdate) AddStatus(i int8) *WinGameSlotUpdate {
	wgsu.mutation.AddStatus(i)
	return wgsu
}

// SetDevice sets the "device" field.
func (wgsu *WinGameSlotUpdate) SetDevice(i int8) *WinGameSlotUpdate {
	wgsu.mutation.ResetDevice()
	wgsu.mutation.SetDevice(i)
	return wgsu
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (wgsu *WinGameSlotUpdate) SetNillableDevice(i *int8) *WinGameSlotUpdate {
	if i != nil {
		wgsu.SetDevice(*i)
	}
	return wgsu
}

// AddDevice adds i to the "device" field.
func (wgsu *WinGameSlotUpdate) AddDevice(i int8) *WinGameSlotUpdate {
	wgsu.mutation.AddDevice(i)
	return wgsu
}

// ClearDevice clears the value of the "device" field.
func (wgsu *WinGameSlotUpdate) ClearDevice() *WinGameSlotUpdate {
	wgsu.mutation.ClearDevice()
	return wgsu
}

// SetCreatedAt sets the "created_at" field.
func (wgsu *WinGameSlotUpdate) SetCreatedAt(i int32) *WinGameSlotUpdate {
	wgsu.mutation.ResetCreatedAt()
	wgsu.mutation.SetCreatedAt(i)
	return wgsu
}

// AddCreatedAt adds i to the "created_at" field.
func (wgsu *WinGameSlotUpdate) AddCreatedAt(i int32) *WinGameSlotUpdate {
	wgsu.mutation.AddCreatedAt(i)
	return wgsu
}

// SetUpdatedAt sets the "updated_at" field.
func (wgsu *WinGameSlotUpdate) SetUpdatedAt(i int32) *WinGameSlotUpdate {
	wgsu.mutation.ResetUpdatedAt()
	wgsu.mutation.SetUpdatedAt(i)
	return wgsu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (wgsu *WinGameSlotUpdate) AddUpdatedAt(i int32) *WinGameSlotUpdate {
	wgsu.mutation.AddUpdatedAt(i)
	return wgsu
}

// SetUpdatedUser sets the "updated_user" field.
func (wgsu *WinGameSlotUpdate) SetUpdatedUser(s string) *WinGameSlotUpdate {
	wgsu.mutation.SetUpdatedUser(s)
	return wgsu
}

// SetNillableUpdatedUser sets the "updated_user" field if the given value is not nil.
func (wgsu *WinGameSlotUpdate) SetNillableUpdatedUser(s *string) *WinGameSlotUpdate {
	if s != nil {
		wgsu.SetUpdatedUser(*s)
	}
	return wgsu
}

// ClearUpdatedUser clears the value of the "updated_user" field.
func (wgsu *WinGameSlotUpdate) ClearUpdatedUser() *WinGameSlotUpdate {
	wgsu.mutation.ClearUpdatedUser()
	return wgsu
}

// SetMaintenance sets the "maintenance" field.
func (wgsu *WinGameSlotUpdate) SetMaintenance(s string) *WinGameSlotUpdate {
	wgsu.mutation.SetMaintenance(s)
	return wgsu
}

// SetNillableMaintenance sets the "maintenance" field if the given value is not nil.
func (wgsu *WinGameSlotUpdate) SetNillableMaintenance(s *string) *WinGameSlotUpdate {
	if s != nil {
		wgsu.SetMaintenance(*s)
	}
	return wgsu
}

// ClearMaintenance clears the value of the "maintenance" field.
func (wgsu *WinGameSlotUpdate) ClearMaintenance() *WinGameSlotUpdate {
	wgsu.mutation.ClearMaintenance()
	return wgsu
}

// SetOperatorName sets the "operator_name" field.
func (wgsu *WinGameSlotUpdate) SetOperatorName(s string) *WinGameSlotUpdate {
	wgsu.mutation.SetOperatorName(s)
	return wgsu
}

// SetNillableOperatorName sets the "operator_name" field if the given value is not nil.
func (wgsu *WinGameSlotUpdate) SetNillableOperatorName(s *string) *WinGameSlotUpdate {
	if s != nil {
		wgsu.SetOperatorName(*s)
	}
	return wgsu
}

// ClearOperatorName clears the value of the "operator_name" field.
func (wgsu *WinGameSlotUpdate) ClearOperatorName() *WinGameSlotUpdate {
	wgsu.mutation.ClearOperatorName()
	return wgsu
}

// Mutation returns the WinGameSlotMutation object of the builder.
func (wgsu *WinGameSlotUpdate) Mutation() *WinGameSlotMutation {
	return wgsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wgsu *WinGameSlotUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wgsu.sqlSave, wgsu.mutation, wgsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wgsu *WinGameSlotUpdate) SaveX(ctx context.Context) int {
	affected, err := wgsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wgsu *WinGameSlotUpdate) Exec(ctx context.Context) error {
	_, err := wgsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wgsu *WinGameSlotUpdate) ExecX(ctx context.Context) {
	if err := wgsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wgsu *WinGameSlotUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(wingameslot.Table, wingameslot.Columns, sqlgraph.NewFieldSpec(wingameslot.FieldID, field.TypeString))
	if ps := wgsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wgsu.mutation.GameID(); ok {
		_spec.SetField(wingameslot.FieldGameID, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.AddedGameID(); ok {
		_spec.AddField(wingameslot.FieldGameID, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.GameGroupID(); ok {
		_spec.SetField(wingameslot.FieldGameGroupID, field.TypeInt8, value)
	}
	if value, ok := wgsu.mutation.AddedGameGroupID(); ok {
		_spec.AddField(wingameslot.FieldGameGroupID, field.TypeInt8, value)
	}
	if value, ok := wgsu.mutation.PlatID(); ok {
		_spec.SetField(wingameslot.FieldPlatID, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.AddedPlatID(); ok {
		_spec.AddField(wingameslot.FieldPlatID, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.Provider(); ok {
		_spec.SetField(wingameslot.FieldProvider, field.TypeString, value)
	}
	if wgsu.mutation.ProviderCleared() {
		_spec.ClearField(wingameslot.FieldProvider, field.TypeString)
	}
	if value, ok := wgsu.mutation.Name(); ok {
		_spec.SetField(wingameslot.FieldName, field.TypeString, value)
	}
	if value, ok := wgsu.mutation.NameZh(); ok {
		_spec.SetField(wingameslot.FieldNameZh, field.TypeString, value)
	}
	if value, ok := wgsu.mutation.Img(); ok {
		_spec.SetField(wingameslot.FieldImg, field.TypeString, value)
	}
	if value, ok := wgsu.mutation.ImgNew(); ok {
		_spec.SetField(wingameslot.FieldImgNew, field.TypeString, value)
	}
	if wgsu.mutation.ImgNewCleared() {
		_spec.ClearField(wingameslot.FieldImgNew, field.TypeString)
	}
	if value, ok := wgsu.mutation.IsNew(); ok {
		_spec.SetField(wingameslot.FieldIsNew, field.TypeBool, value)
	}
	if value, ok := wgsu.mutation.IsCasino(); ok {
		_spec.SetField(wingameslot.FieldIsCasino, field.TypeInt8, value)
	}
	if value, ok := wgsu.mutation.AddedIsCasino(); ok {
		_spec.AddField(wingameslot.FieldIsCasino, field.TypeInt8, value)
	}
	if value, ok := wgsu.mutation.GameTypeID(); ok {
		_spec.SetField(wingameslot.FieldGameTypeID, field.TypeString, value)
	}
	if wgsu.mutation.GameTypeIDCleared() {
		_spec.ClearField(wingameslot.FieldGameTypeID, field.TypeString)
	}
	if value, ok := wgsu.mutation.GameTypeName(); ok {
		_spec.SetField(wingameslot.FieldGameTypeName, field.TypeString, value)
	}
	if wgsu.mutation.GameTypeNameCleared() {
		_spec.ClearField(wingameslot.FieldGameTypeName, field.TypeString)
	}
	if value, ok := wgsu.mutation.FavoriteStar(); ok {
		_spec.SetField(wingameslot.FieldFavoriteStar, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.AddedFavoriteStar(); ok {
		_spec.AddField(wingameslot.FieldFavoriteStar, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.HotStar(); ok {
		_spec.SetField(wingameslot.FieldHotStar, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.AddedHotStar(); ok {
		_spec.AddField(wingameslot.FieldHotStar, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.Sort(); ok {
		_spec.SetField(wingameslot.FieldSort, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.AddedSort(); ok {
		_spec.AddField(wingameslot.FieldSort, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.Status(); ok {
		_spec.SetField(wingameslot.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := wgsu.mutation.AddedStatus(); ok {
		_spec.AddField(wingameslot.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := wgsu.mutation.Device(); ok {
		_spec.SetField(wingameslot.FieldDevice, field.TypeInt8, value)
	}
	if value, ok := wgsu.mutation.AddedDevice(); ok {
		_spec.AddField(wingameslot.FieldDevice, field.TypeInt8, value)
	}
	if wgsu.mutation.DeviceCleared() {
		_spec.ClearField(wingameslot.FieldDevice, field.TypeInt8)
	}
	if value, ok := wgsu.mutation.CreatedAt(); ok {
		_spec.SetField(wingameslot.FieldCreatedAt, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(wingameslot.FieldCreatedAt, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.UpdatedAt(); ok {
		_spec.SetField(wingameslot.FieldUpdatedAt, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(wingameslot.FieldUpdatedAt, field.TypeInt32, value)
	}
	if value, ok := wgsu.mutation.UpdatedUser(); ok {
		_spec.SetField(wingameslot.FieldUpdatedUser, field.TypeString, value)
	}
	if wgsu.mutation.UpdatedUserCleared() {
		_spec.ClearField(wingameslot.FieldUpdatedUser, field.TypeString)
	}
	if value, ok := wgsu.mutation.Maintenance(); ok {
		_spec.SetField(wingameslot.FieldMaintenance, field.TypeString, value)
	}
	if wgsu.mutation.MaintenanceCleared() {
		_spec.ClearField(wingameslot.FieldMaintenance, field.TypeString)
	}
	if value, ok := wgsu.mutation.OperatorName(); ok {
		_spec.SetField(wingameslot.FieldOperatorName, field.TypeString, value)
	}
	if wgsu.mutation.OperatorNameCleared() {
		_spec.ClearField(wingameslot.FieldOperatorName, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wgsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wingameslot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wgsu.mutation.done = true
	return n, nil
}

// WinGameSlotUpdateOne is the builder for updating a single WinGameSlot entity.
type WinGameSlotUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WinGameSlotMutation
}

// SetGameID sets the "game_id" field.
func (wgsuo *WinGameSlotUpdateOne) SetGameID(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetGameID()
	wgsuo.mutation.SetGameID(i)
	return wgsuo
}

// AddGameID adds i to the "game_id" field.
func (wgsuo *WinGameSlotUpdateOne) AddGameID(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddGameID(i)
	return wgsuo
}

// SetGameGroupID sets the "game_group_id" field.
func (wgsuo *WinGameSlotUpdateOne) SetGameGroupID(i int8) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetGameGroupID()
	wgsuo.mutation.SetGameGroupID(i)
	return wgsuo
}

// AddGameGroupID adds i to the "game_group_id" field.
func (wgsuo *WinGameSlotUpdateOne) AddGameGroupID(i int8) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddGameGroupID(i)
	return wgsuo
}

// SetPlatID sets the "plat_id" field.
func (wgsuo *WinGameSlotUpdateOne) SetPlatID(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetPlatID()
	wgsuo.mutation.SetPlatID(i)
	return wgsuo
}

// AddPlatID adds i to the "plat_id" field.
func (wgsuo *WinGameSlotUpdateOne) AddPlatID(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddPlatID(i)
	return wgsuo
}

// SetProvider sets the "provider" field.
func (wgsuo *WinGameSlotUpdateOne) SetProvider(s string) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetProvider(s)
	return wgsuo
}

// SetNillableProvider sets the "provider" field if the given value is not nil.
func (wgsuo *WinGameSlotUpdateOne) SetNillableProvider(s *string) *WinGameSlotUpdateOne {
	if s != nil {
		wgsuo.SetProvider(*s)
	}
	return wgsuo
}

// ClearProvider clears the value of the "provider" field.
func (wgsuo *WinGameSlotUpdateOne) ClearProvider() *WinGameSlotUpdateOne {
	wgsuo.mutation.ClearProvider()
	return wgsuo
}

// SetName sets the "name" field.
func (wgsuo *WinGameSlotUpdateOne) SetName(s string) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetName(s)
	return wgsuo
}

// SetNameZh sets the "name_zh" field.
func (wgsuo *WinGameSlotUpdateOne) SetNameZh(s string) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetNameZh(s)
	return wgsuo
}

// SetImg sets the "img" field.
func (wgsuo *WinGameSlotUpdateOne) SetImg(s string) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetImg(s)
	return wgsuo
}

// SetImgNew sets the "img_new" field.
func (wgsuo *WinGameSlotUpdateOne) SetImgNew(s string) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetImgNew(s)
	return wgsuo
}

// SetNillableImgNew sets the "img_new" field if the given value is not nil.
func (wgsuo *WinGameSlotUpdateOne) SetNillableImgNew(s *string) *WinGameSlotUpdateOne {
	if s != nil {
		wgsuo.SetImgNew(*s)
	}
	return wgsuo
}

// ClearImgNew clears the value of the "img_new" field.
func (wgsuo *WinGameSlotUpdateOne) ClearImgNew() *WinGameSlotUpdateOne {
	wgsuo.mutation.ClearImgNew()
	return wgsuo
}

// SetIsNew sets the "is_new" field.
func (wgsuo *WinGameSlotUpdateOne) SetIsNew(b bool) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetIsNew(b)
	return wgsuo
}

// SetIsCasino sets the "is_casino" field.
func (wgsuo *WinGameSlotUpdateOne) SetIsCasino(i int8) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetIsCasino()
	wgsuo.mutation.SetIsCasino(i)
	return wgsuo
}

// AddIsCasino adds i to the "is_casino" field.
func (wgsuo *WinGameSlotUpdateOne) AddIsCasino(i int8) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddIsCasino(i)
	return wgsuo
}

// SetGameTypeID sets the "game_type_id" field.
func (wgsuo *WinGameSlotUpdateOne) SetGameTypeID(s string) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetGameTypeID(s)
	return wgsuo
}

// SetNillableGameTypeID sets the "game_type_id" field if the given value is not nil.
func (wgsuo *WinGameSlotUpdateOne) SetNillableGameTypeID(s *string) *WinGameSlotUpdateOne {
	if s != nil {
		wgsuo.SetGameTypeID(*s)
	}
	return wgsuo
}

// ClearGameTypeID clears the value of the "game_type_id" field.
func (wgsuo *WinGameSlotUpdateOne) ClearGameTypeID() *WinGameSlotUpdateOne {
	wgsuo.mutation.ClearGameTypeID()
	return wgsuo
}

// SetGameTypeName sets the "game_type_name" field.
func (wgsuo *WinGameSlotUpdateOne) SetGameTypeName(s string) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetGameTypeName(s)
	return wgsuo
}

// SetNillableGameTypeName sets the "game_type_name" field if the given value is not nil.
func (wgsuo *WinGameSlotUpdateOne) SetNillableGameTypeName(s *string) *WinGameSlotUpdateOne {
	if s != nil {
		wgsuo.SetGameTypeName(*s)
	}
	return wgsuo
}

// ClearGameTypeName clears the value of the "game_type_name" field.
func (wgsuo *WinGameSlotUpdateOne) ClearGameTypeName() *WinGameSlotUpdateOne {
	wgsuo.mutation.ClearGameTypeName()
	return wgsuo
}

// SetFavoriteStar sets the "favorite_star" field.
func (wgsuo *WinGameSlotUpdateOne) SetFavoriteStar(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetFavoriteStar()
	wgsuo.mutation.SetFavoriteStar(i)
	return wgsuo
}

// AddFavoriteStar adds i to the "favorite_star" field.
func (wgsuo *WinGameSlotUpdateOne) AddFavoriteStar(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddFavoriteStar(i)
	return wgsuo
}

// SetHotStar sets the "hot_star" field.
func (wgsuo *WinGameSlotUpdateOne) SetHotStar(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetHotStar()
	wgsuo.mutation.SetHotStar(i)
	return wgsuo
}

// AddHotStar adds i to the "hot_star" field.
func (wgsuo *WinGameSlotUpdateOne) AddHotStar(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddHotStar(i)
	return wgsuo
}

// SetSort sets the "sort" field.
func (wgsuo *WinGameSlotUpdateOne) SetSort(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetSort()
	wgsuo.mutation.SetSort(i)
	return wgsuo
}

// AddSort adds i to the "sort" field.
func (wgsuo *WinGameSlotUpdateOne) AddSort(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddSort(i)
	return wgsuo
}

// SetStatus sets the "status" field.
func (wgsuo *WinGameSlotUpdateOne) SetStatus(i int8) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetStatus()
	wgsuo.mutation.SetStatus(i)
	return wgsuo
}

// AddStatus adds i to the "status" field.
func (wgsuo *WinGameSlotUpdateOne) AddStatus(i int8) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddStatus(i)
	return wgsuo
}

// SetDevice sets the "device" field.
func (wgsuo *WinGameSlotUpdateOne) SetDevice(i int8) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetDevice()
	wgsuo.mutation.SetDevice(i)
	return wgsuo
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (wgsuo *WinGameSlotUpdateOne) SetNillableDevice(i *int8) *WinGameSlotUpdateOne {
	if i != nil {
		wgsuo.SetDevice(*i)
	}
	return wgsuo
}

// AddDevice adds i to the "device" field.
func (wgsuo *WinGameSlotUpdateOne) AddDevice(i int8) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddDevice(i)
	return wgsuo
}

// ClearDevice clears the value of the "device" field.
func (wgsuo *WinGameSlotUpdateOne) ClearDevice() *WinGameSlotUpdateOne {
	wgsuo.mutation.ClearDevice()
	return wgsuo
}

// SetCreatedAt sets the "created_at" field.
func (wgsuo *WinGameSlotUpdateOne) SetCreatedAt(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetCreatedAt()
	wgsuo.mutation.SetCreatedAt(i)
	return wgsuo
}

// AddCreatedAt adds i to the "created_at" field.
func (wgsuo *WinGameSlotUpdateOne) AddCreatedAt(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddCreatedAt(i)
	return wgsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (wgsuo *WinGameSlotUpdateOne) SetUpdatedAt(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.ResetUpdatedAt()
	wgsuo.mutation.SetUpdatedAt(i)
	return wgsuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (wgsuo *WinGameSlotUpdateOne) AddUpdatedAt(i int32) *WinGameSlotUpdateOne {
	wgsuo.mutation.AddUpdatedAt(i)
	return wgsuo
}

// SetUpdatedUser sets the "updated_user" field.
func (wgsuo *WinGameSlotUpdateOne) SetUpdatedUser(s string) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetUpdatedUser(s)
	return wgsuo
}

// SetNillableUpdatedUser sets the "updated_user" field if the given value is not nil.
func (wgsuo *WinGameSlotUpdateOne) SetNillableUpdatedUser(s *string) *WinGameSlotUpdateOne {
	if s != nil {
		wgsuo.SetUpdatedUser(*s)
	}
	return wgsuo
}

// ClearUpdatedUser clears the value of the "updated_user" field.
func (wgsuo *WinGameSlotUpdateOne) ClearUpdatedUser() *WinGameSlotUpdateOne {
	wgsuo.mutation.ClearUpdatedUser()
	return wgsuo
}

// SetMaintenance sets the "maintenance" field.
func (wgsuo *WinGameSlotUpdateOne) SetMaintenance(s string) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetMaintenance(s)
	return wgsuo
}

// SetNillableMaintenance sets the "maintenance" field if the given value is not nil.
func (wgsuo *WinGameSlotUpdateOne) SetNillableMaintenance(s *string) *WinGameSlotUpdateOne {
	if s != nil {
		wgsuo.SetMaintenance(*s)
	}
	return wgsuo
}

// ClearMaintenance clears the value of the "maintenance" field.
func (wgsuo *WinGameSlotUpdateOne) ClearMaintenance() *WinGameSlotUpdateOne {
	wgsuo.mutation.ClearMaintenance()
	return wgsuo
}

// SetOperatorName sets the "operator_name" field.
func (wgsuo *WinGameSlotUpdateOne) SetOperatorName(s string) *WinGameSlotUpdateOne {
	wgsuo.mutation.SetOperatorName(s)
	return wgsuo
}

// SetNillableOperatorName sets the "operator_name" field if the given value is not nil.
func (wgsuo *WinGameSlotUpdateOne) SetNillableOperatorName(s *string) *WinGameSlotUpdateOne {
	if s != nil {
		wgsuo.SetOperatorName(*s)
	}
	return wgsuo
}

// ClearOperatorName clears the value of the "operator_name" field.
func (wgsuo *WinGameSlotUpdateOne) ClearOperatorName() *WinGameSlotUpdateOne {
	wgsuo.mutation.ClearOperatorName()
	return wgsuo
}

// Mutation returns the WinGameSlotMutation object of the builder.
func (wgsuo *WinGameSlotUpdateOne) Mutation() *WinGameSlotMutation {
	return wgsuo.mutation
}

// Where appends a list predicates to the WinGameSlotUpdate builder.
func (wgsuo *WinGameSlotUpdateOne) Where(ps ...predicate.WinGameSlot) *WinGameSlotUpdateOne {
	wgsuo.mutation.Where(ps...)
	return wgsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wgsuo *WinGameSlotUpdateOne) Select(field string, fields ...string) *WinGameSlotUpdateOne {
	wgsuo.fields = append([]string{field}, fields...)
	return wgsuo
}

// Save executes the query and returns the updated WinGameSlot entity.
func (wgsuo *WinGameSlotUpdateOne) Save(ctx context.Context) (*WinGameSlot, error) {
	return withHooks(ctx, wgsuo.sqlSave, wgsuo.mutation, wgsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wgsuo *WinGameSlotUpdateOne) SaveX(ctx context.Context) *WinGameSlot {
	node, err := wgsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wgsuo *WinGameSlotUpdateOne) Exec(ctx context.Context) error {
	_, err := wgsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wgsuo *WinGameSlotUpdateOne) ExecX(ctx context.Context) {
	if err := wgsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wgsuo *WinGameSlotUpdateOne) sqlSave(ctx context.Context) (_node *WinGameSlot, err error) {
	_spec := sqlgraph.NewUpdateSpec(wingameslot.Table, wingameslot.Columns, sqlgraph.NewFieldSpec(wingameslot.FieldID, field.TypeString))
	id, ok := wgsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WinGameSlot.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wgsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wingameslot.FieldID)
		for _, f := range fields {
			if !wingameslot.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wingameslot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wgsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wgsuo.mutation.GameID(); ok {
		_spec.SetField(wingameslot.FieldGameID, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.AddedGameID(); ok {
		_spec.AddField(wingameslot.FieldGameID, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.GameGroupID(); ok {
		_spec.SetField(wingameslot.FieldGameGroupID, field.TypeInt8, value)
	}
	if value, ok := wgsuo.mutation.AddedGameGroupID(); ok {
		_spec.AddField(wingameslot.FieldGameGroupID, field.TypeInt8, value)
	}
	if value, ok := wgsuo.mutation.PlatID(); ok {
		_spec.SetField(wingameslot.FieldPlatID, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.AddedPlatID(); ok {
		_spec.AddField(wingameslot.FieldPlatID, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.Provider(); ok {
		_spec.SetField(wingameslot.FieldProvider, field.TypeString, value)
	}
	if wgsuo.mutation.ProviderCleared() {
		_spec.ClearField(wingameslot.FieldProvider, field.TypeString)
	}
	if value, ok := wgsuo.mutation.Name(); ok {
		_spec.SetField(wingameslot.FieldName, field.TypeString, value)
	}
	if value, ok := wgsuo.mutation.NameZh(); ok {
		_spec.SetField(wingameslot.FieldNameZh, field.TypeString, value)
	}
	if value, ok := wgsuo.mutation.Img(); ok {
		_spec.SetField(wingameslot.FieldImg, field.TypeString, value)
	}
	if value, ok := wgsuo.mutation.ImgNew(); ok {
		_spec.SetField(wingameslot.FieldImgNew, field.TypeString, value)
	}
	if wgsuo.mutation.ImgNewCleared() {
		_spec.ClearField(wingameslot.FieldImgNew, field.TypeString)
	}
	if value, ok := wgsuo.mutation.IsNew(); ok {
		_spec.SetField(wingameslot.FieldIsNew, field.TypeBool, value)
	}
	if value, ok := wgsuo.mutation.IsCasino(); ok {
		_spec.SetField(wingameslot.FieldIsCasino, field.TypeInt8, value)
	}
	if value, ok := wgsuo.mutation.AddedIsCasino(); ok {
		_spec.AddField(wingameslot.FieldIsCasino, field.TypeInt8, value)
	}
	if value, ok := wgsuo.mutation.GameTypeID(); ok {
		_spec.SetField(wingameslot.FieldGameTypeID, field.TypeString, value)
	}
	if wgsuo.mutation.GameTypeIDCleared() {
		_spec.ClearField(wingameslot.FieldGameTypeID, field.TypeString)
	}
	if value, ok := wgsuo.mutation.GameTypeName(); ok {
		_spec.SetField(wingameslot.FieldGameTypeName, field.TypeString, value)
	}
	if wgsuo.mutation.GameTypeNameCleared() {
		_spec.ClearField(wingameslot.FieldGameTypeName, field.TypeString)
	}
	if value, ok := wgsuo.mutation.FavoriteStar(); ok {
		_spec.SetField(wingameslot.FieldFavoriteStar, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.AddedFavoriteStar(); ok {
		_spec.AddField(wingameslot.FieldFavoriteStar, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.HotStar(); ok {
		_spec.SetField(wingameslot.FieldHotStar, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.AddedHotStar(); ok {
		_spec.AddField(wingameslot.FieldHotStar, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.Sort(); ok {
		_spec.SetField(wingameslot.FieldSort, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.AddedSort(); ok {
		_spec.AddField(wingameslot.FieldSort, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.Status(); ok {
		_spec.SetField(wingameslot.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := wgsuo.mutation.AddedStatus(); ok {
		_spec.AddField(wingameslot.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := wgsuo.mutation.Device(); ok {
		_spec.SetField(wingameslot.FieldDevice, field.TypeInt8, value)
	}
	if value, ok := wgsuo.mutation.AddedDevice(); ok {
		_spec.AddField(wingameslot.FieldDevice, field.TypeInt8, value)
	}
	if wgsuo.mutation.DeviceCleared() {
		_spec.ClearField(wingameslot.FieldDevice, field.TypeInt8)
	}
	if value, ok := wgsuo.mutation.CreatedAt(); ok {
		_spec.SetField(wingameslot.FieldCreatedAt, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(wingameslot.FieldCreatedAt, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(wingameslot.FieldUpdatedAt, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(wingameslot.FieldUpdatedAt, field.TypeInt32, value)
	}
	if value, ok := wgsuo.mutation.UpdatedUser(); ok {
		_spec.SetField(wingameslot.FieldUpdatedUser, field.TypeString, value)
	}
	if wgsuo.mutation.UpdatedUserCleared() {
		_spec.ClearField(wingameslot.FieldUpdatedUser, field.TypeString)
	}
	if value, ok := wgsuo.mutation.Maintenance(); ok {
		_spec.SetField(wingameslot.FieldMaintenance, field.TypeString, value)
	}
	if wgsuo.mutation.MaintenanceCleared() {
		_spec.ClearField(wingameslot.FieldMaintenance, field.TypeString)
	}
	if value, ok := wgsuo.mutation.OperatorName(); ok {
		_spec.SetField(wingameslot.FieldOperatorName, field.TypeString, value)
	}
	if wgsuo.mutation.OperatorNameCleared() {
		_spec.ClearField(wingameslot.FieldOperatorName, field.TypeString)
	}
	_node = &WinGameSlot{config: wgsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wgsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wingameslot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wgsuo.mutation.done = true
	return _node, nil
}
