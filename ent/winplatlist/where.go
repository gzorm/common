// Code generated by ent, DO NOT EDIT.

package winplatlist

import (
	"entgo.io/ent/dialect/sql"
	"github.com/gzorm/common/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldCode, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldName, v))
}

// Config applies equality check predicate on the "config" field. It's identical to ConfigEQ.
func Config(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldConfig, v))
}

// Rate applies equality check predicate on the "rate" field. It's identical to RateEQ.
func Rate(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldRate, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldSort, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldUpdatedAt, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldContainsFold(FieldCode, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldContainsFold(FieldName, v))
}

// ConfigEQ applies the EQ predicate on the "config" field.
func ConfigEQ(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldConfig, v))
}

// ConfigNEQ applies the NEQ predicate on the "config" field.
func ConfigNEQ(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNEQ(FieldConfig, v))
}

// ConfigIn applies the In predicate on the "config" field.
func ConfigIn(vs ...string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIn(FieldConfig, vs...))
}

// ConfigNotIn applies the NotIn predicate on the "config" field.
func ConfigNotIn(vs ...string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotIn(FieldConfig, vs...))
}

// ConfigGT applies the GT predicate on the "config" field.
func ConfigGT(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGT(FieldConfig, v))
}

// ConfigGTE applies the GTE predicate on the "config" field.
func ConfigGTE(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGTE(FieldConfig, v))
}

// ConfigLT applies the LT predicate on the "config" field.
func ConfigLT(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLT(FieldConfig, v))
}

// ConfigLTE applies the LTE predicate on the "config" field.
func ConfigLTE(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLTE(FieldConfig, v))
}

// ConfigContains applies the Contains predicate on the "config" field.
func ConfigContains(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldContains(FieldConfig, v))
}

// ConfigHasPrefix applies the HasPrefix predicate on the "config" field.
func ConfigHasPrefix(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldHasPrefix(FieldConfig, v))
}

// ConfigHasSuffix applies the HasSuffix predicate on the "config" field.
func ConfigHasSuffix(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldHasSuffix(FieldConfig, v))
}

// ConfigIsNil applies the IsNil predicate on the "config" field.
func ConfigIsNil() predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIsNull(FieldConfig))
}

// ConfigNotNil applies the NotNil predicate on the "config" field.
func ConfigNotNil() predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotNull(FieldConfig))
}

// ConfigEqualFold applies the EqualFold predicate on the "config" field.
func ConfigEqualFold(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEqualFold(FieldConfig, v))
}

// ConfigContainsFold applies the ContainsFold predicate on the "config" field.
func ConfigContainsFold(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldContainsFold(FieldConfig, v))
}

// RateEQ applies the EQ predicate on the "rate" field.
func RateEQ(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldRate, v))
}

// RateNEQ applies the NEQ predicate on the "rate" field.
func RateNEQ(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNEQ(FieldRate, v))
}

// RateIn applies the In predicate on the "rate" field.
func RateIn(vs ...string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIn(FieldRate, vs...))
}

// RateNotIn applies the NotIn predicate on the "rate" field.
func RateNotIn(vs ...string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotIn(FieldRate, vs...))
}

// RateGT applies the GT predicate on the "rate" field.
func RateGT(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGT(FieldRate, v))
}

// RateGTE applies the GTE predicate on the "rate" field.
func RateGTE(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGTE(FieldRate, v))
}

// RateLT applies the LT predicate on the "rate" field.
func RateLT(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLT(FieldRate, v))
}

// RateLTE applies the LTE predicate on the "rate" field.
func RateLTE(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLTE(FieldRate, v))
}

// RateContains applies the Contains predicate on the "rate" field.
func RateContains(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldContains(FieldRate, v))
}

// RateHasPrefix applies the HasPrefix predicate on the "rate" field.
func RateHasPrefix(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldHasPrefix(FieldRate, v))
}

// RateHasSuffix applies the HasSuffix predicate on the "rate" field.
func RateHasSuffix(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldHasSuffix(FieldRate, v))
}

// RateIsNil applies the IsNil predicate on the "rate" field.
func RateIsNil() predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIsNull(FieldRate))
}

// RateNotNil applies the NotNil predicate on the "rate" field.
func RateNotNil() predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotNull(FieldRate))
}

// RateEqualFold applies the EqualFold predicate on the "rate" field.
func RateEqualFold(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEqualFold(FieldRate, v))
}

// RateContainsFold applies the ContainsFold predicate on the "rate" field.
func RateContainsFold(v string) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldContainsFold(FieldRate, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLTE(FieldSort, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLTE(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int32) predicate.WinPlatList {
	return predicate.WinPlatList(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WinPlatList) predicate.WinPlatList {
	return predicate.WinPlatList(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WinPlatList) predicate.WinPlatList {
	return predicate.WinPlatList(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WinPlatList) predicate.WinPlatList {
	return predicate.WinPlatList(func(s *sql.Selector) {
		p(s.Not())
	})
}
