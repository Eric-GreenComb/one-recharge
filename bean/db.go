package bean

import (
	"github.com/jinzhu/gorm"
)

// Recharge 充值记录
type Recharge struct {
	gorm.Model
	TxID        string `gorm:"not null" form:"txid" json:"txid"`       // 用户交易txid
	Address     string `gorm:"not null" form:"address" json:"address"` // 用户账户
	TokenAmount int64  `gorm:"not null" form:"token" json:"token"`     // 充值Token数量
	ChipAmount  string `gorm:"not null" form:"chip" json:"chip"`       // 兑换的筹码
	Rate        string `gorm:"not null" form:"rate" json:"rate"`       // 筹码兑换比例:1000:100
}

// RechargeRate token-筹码兑换比例
type RechargeRate struct {
	gorm.Model
	Token int64 `gorm:"not null" form:"token" json:"token"` // token数量
	Chip  int64 `gorm:"not null" form:"chip" json:"chip"`   // 筹码数量
}
