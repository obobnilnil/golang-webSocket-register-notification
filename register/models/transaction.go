package models

type TransactionInputRegister struct { // not use yet
	ID        string        `json:"id"`
	ActionLog []interface{} `json:"actionLog"`
}

type UserInfo struct {
	Role      string
	CompanyID string
}
