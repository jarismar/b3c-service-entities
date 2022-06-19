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

type EarningTax struct {
	Id    int64
	Tax   *Tax
	Value float64
}

type Earning struct {
	Id       int64
	UUID     uuid.UUID
	Type     EarningType
	Company  *Company
	User     *User
	Taxes    []EarningTax
	PayDate  time.Time
	RawValue float64
	NetValue float64
}
