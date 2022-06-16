package entity

import "time"

type EarningType string

const (
	JCP EarningType = "JCP"
	DIV EarningType = "DIV"
)

type Earning struct {
	Id       int64
	Type     EarningType
	Company  *Company
	User     *User
	Taxes    []InvoiceTax
	Paydate  time.Time
	RawValue float64
	NetValue float64
}
