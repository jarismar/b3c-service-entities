package entity

import "time"

type Trade struct {
	Id             int64
	Qty            int64
	MarketDate     time.Time
	InvoiceItemIn  *InvoiceItem
	InvoiceItemOut *InvoiceItem
}
