package entity

import (
	"time"

	"github.com/google/uuid"
)

type EarningType string

const (
	JCP EarningType = "JCP"
	DIV EarningType = "DIV"
)

type Earning struct {
	Id       int64
	UUID     uuid.UUID
	Type     EarningType
	Company  *Company
	User     *User
	Taxes    []InvoiceTax
	PayDate  time.Time
	RawValue float64
	NetValue float64
}
