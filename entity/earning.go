package entity

import "time"

type Earning struct {
	Id       int64
	Company  *Company
	User     *User
	Paydate  time.Time
	NetValue float64
	Tax      float64
}
