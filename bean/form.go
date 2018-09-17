package bean

import ()

// FormParams FormParams
type FormParams struct {
	Params string `form:"params" json:"params"` // params
	Key    string `form:"key" json:"key"`       // key
	Value  string `form:"value" json:"value"`   // value
	Token  string `form:"token" json:"token"`   // token
	Chip   string `form:"chip" json:"chip"`     // chip
}
