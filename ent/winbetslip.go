// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gzorm/common/ent/winbetslip"
)

// WinBetslip is the model entity for the WinBetslip schema.
type WinBetslip struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int `json:"id,omitempty"`
	// 回合id
	RoundID string `json:"round_id,omitempty"`
	// 注单号 对应三方拉单transaction_id
	TransactionID string `json:"transaction_id,omitempty"`
	// 游戏平台id
	GamePlatID int32 `json:"game_plat_id,omitempty"`
	// 注单状态 1:待开彩  2:完成  3: 退款
	XbStatus int8 `json:"xb_status,omitempty"`
	// 对应user表id
	XbUID uint32 `json:"xb_uid,omitempty"`
	// 对应user表username
	XbUsername string `json:"xb_username,omitempty"`
	// 盈亏金额
	XbProfit float64 `json:"xb_profit,omitempty"`
	// 投注
	Stake float64 `json:"stake,omitempty"`
	// 有效投注金额
	ValidStake float64 `json:"valid_stake,omitempty"`
	// 派彩
	Payout float64 `json:"payout,omitempty"`
	// 退款金额
	CoinRefund float64 `json:"coin_refund,omitempty"`
	// 投注前金额
	CoinBefore float64 `json:"coin_before,omitempty"`
	// 游戏id 对应game_list表id
	GameID string `json:"game_id,omitempty"`
	// 游戏名称
	GameName string `json:"game_name,omitempty"`
	// 投注方式 1:现金，2:奖金 3:免费旋转 4:活动免费旋转
	AmountType int32 `json:"amount_type,omitempty"`
	// 游戏子类型id 对应game_slot 表id
	GameTypeID string `json:"game_type_id,omitempty"`
	// 游戏大类类型:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能',
	GameGroupID int32 `json:"game_group_id,omitempty"`
	// 体育投注类型
	SportType string `json:"sport_type,omitempty"`
	// 游戏开始时间
	DtStarted int `json:"dt_started,omitempty"`
	// 游戏结束时间
	DtCompleted int `json:"dt_completed,omitempty"`
	// 开奖交易单号
	WinTransactionID string `json:"win_transaction_id,omitempty"`
	// 投注原始json
	BetJSON string `json:"bet_json,omitempty"`
	// 开彩原始json
	RewardJSON string `json:"reward_json,omitempty"`
	// 退款原始json
	RefundJSON string `json:"refund_json,omitempty"`
	// 投注时间
	CreateTimeStr string `json:"create_time_str,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    int32 `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WinBetslip) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case winbetslip.FieldXbProfit, winbetslip.FieldStake, winbetslip.FieldValidStake, winbetslip.FieldPayout, winbetslip.FieldCoinRefund, winbetslip.FieldCoinBefore:
			values[i] = new(sql.NullFloat64)
		case winbetslip.FieldID, winbetslip.FieldGamePlatID, winbetslip.FieldXbStatus, winbetslip.FieldXbUID, winbetslip.FieldAmountType, winbetslip.FieldGameGroupID, winbetslip.FieldDtStarted, winbetslip.FieldDtCompleted, winbetslip.FieldCreatedAt, winbetslip.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case winbetslip.FieldRoundID, winbetslip.FieldTransactionID, winbetslip.FieldXbUsername, winbetslip.FieldGameID, winbetslip.FieldGameName, winbetslip.FieldGameTypeID, winbetslip.FieldSportType, winbetslip.FieldWinTransactionID, winbetslip.FieldBetJSON, winbetslip.FieldRewardJSON, winbetslip.FieldRefundJSON, winbetslip.FieldCreateTimeStr:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WinBetslip fields.
func (wb *WinBetslip) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case winbetslip.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wb.ID = int(value.Int64)
		case winbetslip.FieldRoundID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field round_id", values[i])
			} else if value.Valid {
				wb.RoundID = value.String
			}
		case winbetslip.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				wb.TransactionID = value.String
			}
		case winbetslip.FieldGamePlatID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field game_plat_id", values[i])
			} else if value.Valid {
				wb.GamePlatID = int32(value.Int64)
			}
		case winbetslip.FieldXbStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field xb_status", values[i])
			} else if value.Valid {
				wb.XbStatus = int8(value.Int64)
			}
		case winbetslip.FieldXbUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field xb_uid", values[i])
			} else if value.Valid {
				wb.XbUID = uint32(value.Int64)
			}
		case winbetslip.FieldXbUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field xb_username", values[i])
			} else if value.Valid {
				wb.XbUsername = value.String
			}
		case winbetslip.FieldXbProfit:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field xb_profit", values[i])
			} else if value.Valid {
				wb.XbProfit = value.Float64
			}
		case winbetslip.FieldStake:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field stake", values[i])
			} else if value.Valid {
				wb.Stake = value.Float64
			}
		case winbetslip.FieldValidStake:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field valid_stake", values[i])
			} else if value.Valid {
				wb.ValidStake = value.Float64
			}
		case winbetslip.FieldPayout:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field payout", values[i])
			} else if value.Valid {
				wb.Payout = value.Float64
			}
		case winbetslip.FieldCoinRefund:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field coin_refund", values[i])
			} else if value.Valid {
				wb.CoinRefund = value.Float64
			}
		case winbetslip.FieldCoinBefore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field coin_before", values[i])
			} else if value.Valid {
				wb.CoinBefore = value.Float64
			}
		case winbetslip.FieldGameID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field game_id", values[i])
			} else if value.Valid {
				wb.GameID = value.String
			}
		case winbetslip.FieldGameName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field game_name", values[i])
			} else if value.Valid {
				wb.GameName = value.String
			}
		case winbetslip.FieldAmountType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount_type", values[i])
			} else if value.Valid {
				wb.AmountType = int32(value.Int64)
			}
		case winbetslip.FieldGameTypeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field game_type_id", values[i])
			} else if value.Valid {
				wb.GameTypeID = value.String
			}
		case winbetslip.FieldGameGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field game_group_id", values[i])
			} else if value.Valid {
				wb.GameGroupID = int32(value.Int64)
			}
		case winbetslip.FieldSportType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sport_type", values[i])
			} else if value.Valid {
				wb.SportType = value.String
			}
		case winbetslip.FieldDtStarted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dt_started", values[i])
			} else if value.Valid {
				wb.DtStarted = int(value.Int64)
			}
		case winbetslip.FieldDtCompleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dt_completed", values[i])
			} else if value.Valid {
				wb.DtCompleted = int(value.Int64)
			}
		case winbetslip.FieldWinTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field win_transaction_id", values[i])
			} else if value.Valid {
				wb.WinTransactionID = value.String
			}
		case winbetslip.FieldBetJSON:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bet_json", values[i])
			} else if value.Valid {
				wb.BetJSON = value.String
			}
		case winbetslip.FieldRewardJSON:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reward_json", values[i])
			} else if value.Valid {
				wb.RewardJSON = value.String
			}
		case winbetslip.FieldRefundJSON:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refund_json", values[i])
			} else if value.Valid {
				wb.RefundJSON = value.String
			}
		case winbetslip.FieldCreateTimeStr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field create_time_str", values[i])
			} else if value.Valid {
				wb.CreateTimeStr = value.String
			}
		case winbetslip.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wb.CreatedAt = int32(value.Int64)
			}
		case winbetslip.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				wb.UpdatedAt = int32(value.Int64)
			}
		default:
			wb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WinBetslip.
// This includes values selected through modifiers, order, etc.
func (wb *WinBetslip) Value(name string) (ent.Value, error) {
	return wb.selectValues.Get(name)
}

// Update returns a builder for updating this WinBetslip.
// Note that you need to call WinBetslip.Unwrap() before calling this method if this WinBetslip
// was returned from a transaction, and the transaction was committed or rolled back.
func (wb *WinBetslip) Update() *WinBetslipUpdateOne {
	return NewWinBetslipClient(wb.config).UpdateOne(wb)
}

// Unwrap unwraps the WinBetslip entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wb *WinBetslip) Unwrap() *WinBetslip {
	_tx, ok := wb.config.driver.(*txDriver)
	if !ok {
		panic("ent: WinBetslip is not a transactional entity")
	}
	wb.config.driver = _tx.drv
	return wb
}

// String implements the fmt.Stringer.
func (wb *WinBetslip) String() string {
	var builder strings.Builder
	builder.WriteString("WinBetslip(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wb.ID))
	builder.WriteString("round_id=")
	builder.WriteString(wb.RoundID)
	builder.WriteString(", ")
	builder.WriteString("transaction_id=")
	builder.WriteString(wb.TransactionID)
	builder.WriteString(", ")
	builder.WriteString("game_plat_id=")
	builder.WriteString(fmt.Sprintf("%v", wb.GamePlatID))
	builder.WriteString(", ")
	builder.WriteString("xb_status=")
	builder.WriteString(fmt.Sprintf("%v", wb.XbStatus))
	builder.WriteString(", ")
	builder.WriteString("xb_uid=")
	builder.WriteString(fmt.Sprintf("%v", wb.XbUID))
	builder.WriteString(", ")
	builder.WriteString("xb_username=")
	builder.WriteString(wb.XbUsername)
	builder.WriteString(", ")
	builder.WriteString("xb_profit=")
	builder.WriteString(fmt.Sprintf("%v", wb.XbProfit))
	builder.WriteString(", ")
	builder.WriteString("stake=")
	builder.WriteString(fmt.Sprintf("%v", wb.Stake))
	builder.WriteString(", ")
	builder.WriteString("valid_stake=")
	builder.WriteString(fmt.Sprintf("%v", wb.ValidStake))
	builder.WriteString(", ")
	builder.WriteString("payout=")
	builder.WriteString(fmt.Sprintf("%v", wb.Payout))
	builder.WriteString(", ")
	builder.WriteString("coin_refund=")
	builder.WriteString(fmt.Sprintf("%v", wb.CoinRefund))
	builder.WriteString(", ")
	builder.WriteString("coin_before=")
	builder.WriteString(fmt.Sprintf("%v", wb.CoinBefore))
	builder.WriteString(", ")
	builder.WriteString("game_id=")
	builder.WriteString(wb.GameID)
	builder.WriteString(", ")
	builder.WriteString("game_name=")
	builder.WriteString(wb.GameName)
	builder.WriteString(", ")
	builder.WriteString("amount_type=")
	builder.WriteString(fmt.Sprintf("%v", wb.AmountType))
	builder.WriteString(", ")
	builder.WriteString("game_type_id=")
	builder.WriteString(wb.GameTypeID)
	builder.WriteString(", ")
	builder.WriteString("game_group_id=")
	builder.WriteString(fmt.Sprintf("%v", wb.GameGroupID))
	builder.WriteString(", ")
	builder.WriteString("sport_type=")
	builder.WriteString(wb.SportType)
	builder.WriteString(", ")
	builder.WriteString("dt_started=")
	builder.WriteString(fmt.Sprintf("%v", wb.DtStarted))
	builder.WriteString(", ")
	builder.WriteString("dt_completed=")
	builder.WriteString(fmt.Sprintf("%v", wb.DtCompleted))
	builder.WriteString(", ")
	builder.WriteString("win_transaction_id=")
	builder.WriteString(wb.WinTransactionID)
	builder.WriteString(", ")
	builder.WriteString("bet_json=")
	builder.WriteString(wb.BetJSON)
	builder.WriteString(", ")
	builder.WriteString("reward_json=")
	builder.WriteString(wb.RewardJSON)
	builder.WriteString(", ")
	builder.WriteString("refund_json=")
	builder.WriteString(wb.RefundJSON)
	builder.WriteString(", ")
	builder.WriteString("create_time_str=")
	builder.WriteString(wb.CreateTimeStr)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", wb.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", wb.UpdatedAt))
	builder.WriteByte(')')
	return builder.String()
}

// WinBetslips is a parsable slice of WinBetslip.
type WinBetslips []*WinBetslip