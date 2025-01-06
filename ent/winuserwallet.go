// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gzorm/commons/ent/winuserwallet"
)

// WinUserWallet is the model entity for the WinUserWallet schema.
type WinUserWallet struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// 用户名
	Username string `json:"username,omitempty"`
	// 账户余额
	Coin float64 `json:"coin,omitempty"`
	// 版本号
	Version int `json:"version,omitempty"`
	// 13位时间戳
	ModifyAt int `json:"modify_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    int32 `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WinUserWallet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case winuserwallet.FieldCoin:
			values[i] = new(sql.NullFloat64)
		case winuserwallet.FieldID, winuserwallet.FieldVersion, winuserwallet.FieldModifyAt, winuserwallet.FieldCreatedAt, winuserwallet.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case winuserwallet.FieldUsername:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WinUserWallet fields.
func (wuw *WinUserWallet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case winuserwallet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wuw.ID = int32(value.Int64)
		case winuserwallet.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				wuw.Username = value.String
			}
		case winuserwallet.FieldCoin:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field coin", values[i])
			} else if value.Valid {
				wuw.Coin = value.Float64
			}
		case winuserwallet.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				wuw.Version = int(value.Int64)
			}
		case winuserwallet.FieldModifyAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field modify_at", values[i])
			} else if value.Valid {
				wuw.ModifyAt = int(value.Int64)
			}
		case winuserwallet.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wuw.CreatedAt = int32(value.Int64)
			}
		case winuserwallet.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				wuw.UpdatedAt = int32(value.Int64)
			}
		default:
			wuw.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WinUserWallet.
// This includes values selected through modifiers, order, etc.
func (wuw *WinUserWallet) Value(name string) (ent.Value, error) {
	return wuw.selectValues.Get(name)
}

// Update returns a builder for updating this WinUserWallet.
// Note that you need to call WinUserWallet.Unwrap() before calling this method if this WinUserWallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (wuw *WinUserWallet) Update() *WinUserWalletUpdateOne {
	return NewWinUserWalletClient(wuw.config).UpdateOne(wuw)
}

// Unwrap unwraps the WinUserWallet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wuw *WinUserWallet) Unwrap() *WinUserWallet {
	_tx, ok := wuw.config.driver.(*txDriver)
	if !ok {
		panic("ent: WinUserWallet is not a transactional entity")
	}
	wuw.config.driver = _tx.drv
	return wuw
}

// String implements the fmt.Stringer.
func (wuw *WinUserWallet) String() string {
	var builder strings.Builder
	builder.WriteString("WinUserWallet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wuw.ID))
	builder.WriteString("username=")
	builder.WriteString(wuw.Username)
	builder.WriteString(", ")
	builder.WriteString("coin=")
	builder.WriteString(fmt.Sprintf("%v", wuw.Coin))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", wuw.Version))
	builder.WriteString(", ")
	builder.WriteString("modify_at=")
	builder.WriteString(fmt.Sprintf("%v", wuw.ModifyAt))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", wuw.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", wuw.UpdatedAt))
	builder.WriteByte(')')
	return builder.String()
}

// WinUserWallets is a parsable slice of WinUserWallet.
type WinUserWallets []*WinUserWallet
