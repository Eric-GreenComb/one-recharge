package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/one-recharge/bean"
	"github.com/Eric-GreenComb/one-recharge/persist"
)

// CreateRechargeRate CreateRechargeRate
func CreateRechargeRate(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_token := _formParams.Token
	_chip := _formParams.Chip

	if _token == "" || _chip == "" {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "There are some empty fields."})
		return
	}

	_int64Token, err := strconv.ParseInt(_token, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_int64Chip, err := strconv.ParseInt(_chip, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	var _rate bean.RechargeRate
	_rate.Token = _int64Token
	_rate.Chip = _int64Chip

	err = persist.GetPersist().CreateRechargeRate(_rate)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "ok"})

}

// RechargeRateInfo Get RechargeRate
func RechargeRateInfo(c *gin.Context) {

	_rate, err := persist.GetPersist().RechargeRateInfo()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "token": _rate.Token, "chip": _rate.Chip})
}

// ListRechargeRate ListRechargeRate
func ListRechargeRate(c *gin.Context) {

	_page := c.Params.ByName("page")
	_limit := c.Params.ByName("limit")

	_nPage, _ := strconv.Atoi(_page)
	_nLimit, _ := strconv.Atoi(_limit)

	_rates, err := persist.GetPersist().ListRechargeRate(_nPage, _nLimit)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get rates error"})
		return
	}

	c.JSON(http.StatusOK, _rates)
}

// Recharge Recharge
func Recharge(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	fmt.Println(_formParams)

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _formParams})
}
