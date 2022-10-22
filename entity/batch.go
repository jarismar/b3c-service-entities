package entity

import "time"

type CompanyBatch struct {
	Id         int64
	Company    *Company
	Qty        int64
	AvgPrice   float64
	TotalPrice float64
	Items      []ItemBatch
	Trades     []Trade
}

type ItemBatch struct {
	Id         int64
	Item       *InvoiceItem
	TaxGroup   *TaxGroup
	Qty        int64
	AvgPrice   float64
	RawPrice   float64
	TotalTaxes float64
}

type Trade struct {
	Id         int64
	TaxGroup   *TaxGroup
	Item       *InvoiceItem
	MarketDate time.Time
	Qty        int64
	AvgPrice   float64
	RawResults float64
	RawPrice   float64
	TotalTax   float64
}

type TradeBatch struct {
	Id             int64
	TaxGroup       *TaxGroup
	StartDate      time.Time
	AccLoss        float64
	CurrentResults float64
	TotalTrade     float64
	TotalTax       float64
}
