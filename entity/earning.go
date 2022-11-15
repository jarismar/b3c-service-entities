package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	JCP string = "JCP"
	DIV string = "DIV"
)

type Earning struct {
	Id       int64
	UUID     uuid.UUID
	Type     string
	Company  *Company
	User     *User
	TaxGroup *TaxGroup
	PayDate  time.Time
	RawValue float64
	NetValue float64
}
