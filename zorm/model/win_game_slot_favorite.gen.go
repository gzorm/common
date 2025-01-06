// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinGameSlotFavorite = "win_game_slot_favorite"

// WinGameSlotFavorite 游戏收藏列表
type WinGameSlotFavorite struct {
	ID         int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	UID        int64  `gorm:"column:uid;type:int;not null;comment:UID" json:"uid"`                                // UID
	GameID     int64  `gorm:"column:game_id;type:int;not null;comment:游戏类型ID" json:"gameId"`                      // 游戏类型ID
	GameSlotID string `gorm:"column:game_slot_id;type:varchar(48);not null;comment:游戏(子老虎机)ID" json:"gameSlotId"` // 游戏(子老虎机)ID
	CreatedAt  int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt  int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinGameSlotFavorite's table name
func (*WinGameSlotFavorite) TableName() string {
	return TableNameWinGameSlotFavorite
}
