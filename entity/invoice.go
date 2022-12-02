package entity

import "time"

type InvoiceItem struct {
	Id         int64
	Company    *Company
	ItemBatch  *ItemBatch
	MarketDate time.Time
	Qty        int64
	Price      float64
	Debit      bool
	Order      int64
}

type Invoice struct {
	Id            int64
	Number        int64
	User          *User
	TaxGroup      *TaxGroup
	FileName      string
	MarketDate    time.Time
	BillingDate   time.Time
	TotalSold     float64
	TotalAcquired float64
	RawValue      float64
	NetValue      float64
	Items         []InvoiceItem
}
