package persist

import (
	"github.com/Eric-GreenComb/one-recharge/bean"
)

// CreateRechargeRate CreateRechargeRate Persist
func (persist *Persist) CreateRechargeRate(rate bean.RechargeRate) error {
	err := persist.db.Create(&rate).Error
	return err
}

// RechargeRateInfo RechargeRateInfo Persist
func (persist *Persist) RechargeRateInfo() (bean.RechargeRate, error) {

	var rate bean.RechargeRate
	err := persist.db.Table("recharge_rates").Last(&rate).Error

	return rate, err
}

// ListRechargeRate ListRechargeRate Persist
func (persist *Persist) ListRechargeRate(page, limit int) ([]bean.RechargeRate, error) {

	var rates []bean.RechargeRate

	page--
	if page < 0 {
		page = 0
	}

	_0ffset := page * limit

	var err error

	err = persist.db.Table("recharge_rates").Order("id desc").Select("*").Limit(limit).Offset(_0ffset).Find(&rates).Error
	return rates, err
}
