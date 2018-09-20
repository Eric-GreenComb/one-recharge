package bean

import ()

// FormParams FormParams
type FormParams struct {
	Params string `form:"params" json:"params"` // params
	Key    string `form:"key" json:"key"`       // key
	Value  string `form:"value" json:"value"`   // value
	Token  string `form:"token" json:"token"`   // token
	Chip   string `form:"chip" json:"chip"`     // chip

	Nums       string `form:"nums" json:"nums"`               // nums
	Address    string `form:"address" json:"address"`         // address
	CreateTime int64  `form:"create_time" json:"create_time"` // create_time
	IsAuth     string `form:"isauth" json:"isauth"`           // isauth
}
