package entity

import "time"

type InvoiceTax struct {
	Id    int64
	Tax   *Tax
	Value float64
}

type InvoiceItem struct {
	Id      int64
	Company *Company
	Qty     int64
	Balance int64
	Price   float64
	Debit   bool
	Order   int64
}

type Invoice struct {
	Id          int64
	Number      int64
	User        *User
	FileName    string
	MarketDate  time.Time
	BillingDate time.Time
	RawValue    float64
	NetValue    float64
	Items       []InvoiceItem
	Taxes       []InvoiceTax
}
