package entity

import "time"

type InvoiceItem struct {
	Id         int64
	Company    *Company
	MarketDate time.Time
	Qty        int64
	Balance    int64
	Price      float64
	Debit      bool
	Order      int64
}

type Invoice struct {
	Id          int64
	Number      int64
	User        *User
	TaxGroup    *TaxGroup
	FileName    string
	MarketDate  time.Time
	BillingDate time.Time
	RawValue    float64
	NetValue    float64
	Items       []InvoiceItem
}
