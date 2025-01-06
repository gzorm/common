// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gzorm/commons/ent/wincoinlog"
)

// WinCoinLog is the model entity for the WinCoinLog schema.
type WinCoinLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UID
	UID int32 `json:"uid,omitempty"`
	// 用户名
	Username string `json:"username,omitempty"`
	// 类型:1-存款 2-提款 3-投注 4-派彩 5-返水 6-佣金 7-活动(奖励) 8-系统调账 9-退款 10-佣金钱包转主账户余额 11-小费
	Category int8 `json:"category,omitempty"`
	// 关联ID
	ReferID int `json:"refer_id,omitempty"`
	// 金额
	Coin float64 `json:"coin,omitempty"`
	// 实际金额
	CoinReal float64 `json:"coin_real,omitempty"`
	// 游戏平台ID
	PlatID int32 `json:"plat_id,omitempty"`
	// 收支类型:0-支出 1-收入
	OutIn int8 `json:"out_in,omitempty"`
	// 三方游戏ID
	GameID int32 `json:"game_id,omitempty"`
	// 前金额
	CoinBefore float64 `json:"coin_before,omitempty"`
	// 帐变后金额
	CoinAfter float64 `json:"coin_after,omitempty"`
	// 状态:0-处理中 1-成功 2-失败
	Status int8 `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    int32 `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WinCoinLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wincoinlog.FieldCoin, wincoinlog.FieldCoinReal, wincoinlog.FieldCoinBefore, wincoinlog.FieldCoinAfter:
			values[i] = new(sql.NullFloat64)
		case wincoinlog.FieldID, wincoinlog.FieldUID, wincoinlog.FieldCategory, wincoinlog.FieldReferID, wincoinlog.FieldPlatID, wincoinlog.FieldOutIn, wincoinlog.FieldGameID, wincoinlog.FieldStatus, wincoinlog.FieldCreatedAt, wincoinlog.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case wincoinlog.FieldUsername:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WinCoinLog fields.
func (wcl *WinCoinLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wincoinlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wcl.ID = int(value.Int64)
		case wincoinlog.FieldUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				wcl.UID = int32(value.Int64)
			}
		case wincoinlog.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				wcl.Username = value.String
			}
		case wincoinlog.FieldCategory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				wcl.Category = int8(value.Int64)
			}
		case wincoinlog.FieldReferID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field refer_id", values[i])
			} else if value.Valid {
				wcl.ReferID = int(value.Int64)
			}
		case wincoinlog.FieldCoin:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field coin", values[i])
			} else if value.Valid {
				wcl.Coin = value.Float64
			}
		case wincoinlog.FieldCoinReal:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field coin_real", values[i])
			} else if value.Valid {
				wcl.CoinReal = value.Float64
			}
		case wincoinlog.FieldPlatID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field plat_id", values[i])
			} else if value.Valid {
				wcl.PlatID = int32(value.Int64)
			}
		case wincoinlog.FieldOutIn:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field out_in", values[i])
			} else if value.Valid {
				wcl.OutIn = int8(value.Int64)
			}
		case wincoinlog.FieldGameID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field game_id", values[i])
			} else if value.Valid {
				wcl.GameID = int32(value.Int64)
			}
		case wincoinlog.FieldCoinBefore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field coin_before", values[i])
			} else if value.Valid {
				wcl.CoinBefore = value.Float64
			}
		case wincoinlog.FieldCoinAfter:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field coin_after", values[i])
			} else if value.Valid {
				wcl.CoinAfter = value.Float64
			}
		case wincoinlog.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				wcl.Status = int8(value.Int64)
			}
		case wincoinlog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wcl.CreatedAt = int32(value.Int64)
			}
		case wincoinlog.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				wcl.UpdatedAt = int32(value.Int64)
			}
		default:
			wcl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WinCoinLog.
// This includes values selected through modifiers, order, etc.
func (wcl *WinCoinLog) Value(name string) (ent.Value, error) {
	return wcl.selectValues.Get(name)
}

// Update returns a builder for updating this WinCoinLog.
// Note that you need to call WinCoinLog.Unwrap() before calling this method if this WinCoinLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (wcl *WinCoinLog) Update() *WinCoinLogUpdateOne {
	return NewWinCoinLogClient(wcl.config).UpdateOne(wcl)
}

// Unwrap unwraps the WinCoinLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wcl *WinCoinLog) Unwrap() *WinCoinLog {
	_tx, ok := wcl.config.driver.(*txDriver)
	if !ok {
		panic("ent: WinCoinLog is not a transactional entity")
	}
	wcl.config.driver = _tx.drv
	return wcl
}

// String implements the fmt.Stringer.
func (wcl *WinCoinLog) String() string {
	var builder strings.Builder
	builder.WriteString("WinCoinLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wcl.ID))
	builder.WriteString("uid=")
	builder.WriteString(fmt.Sprintf("%v", wcl.UID))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(wcl.Username)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", wcl.Category))
	builder.WriteString(", ")
	builder.WriteString("refer_id=")
	builder.WriteString(fmt.Sprintf("%v", wcl.ReferID))
	builder.WriteString(", ")
	builder.WriteString("coin=")
	builder.WriteString(fmt.Sprintf("%v", wcl.Coin))
	builder.WriteString(", ")
	builder.WriteString("coin_real=")
	builder.WriteString(fmt.Sprintf("%v", wcl.CoinReal))
	builder.WriteString(", ")
	builder.WriteString("plat_id=")
	builder.WriteString(fmt.Sprintf("%v", wcl.PlatID))
	builder.WriteString(", ")
	builder.WriteString("out_in=")
	builder.WriteString(fmt.Sprintf("%v", wcl.OutIn))
	builder.WriteString(", ")
	builder.WriteString("game_id=")
	builder.WriteString(fmt.Sprintf("%v", wcl.GameID))
	builder.WriteString(", ")
	builder.WriteString("coin_before=")
	builder.WriteString(fmt.Sprintf("%v", wcl.CoinBefore))
	builder.WriteString(", ")
	builder.WriteString("coin_after=")
	builder.WriteString(fmt.Sprintf("%v", wcl.CoinAfter))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", wcl.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", wcl.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", wcl.UpdatedAt))
	builder.WriteByte(')')
	return builder.String()
}

// WinCoinLogs is a parsable slice of WinCoinLog.
type WinCoinLogs []*WinCoinLog
