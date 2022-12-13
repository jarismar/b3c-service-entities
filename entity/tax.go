package entity

import "time"

const (
	BIV string = "BIV" // broker invoice
	EAR string = "EAR" // earning
	ITB string = "ITB" // item batch
	TRB string = "TRB" // trade batch
	TRD string = "TRD" // trade
)

type Tax struct {
	Id     int64
	Code   string
	Source string
	Rate   float64
}

type TaxGroup struct {
	Id         int64
	Source     string
	ExternalId int64
	Taxes      []TaxInstance
}

type TaxInstance struct {
	Id         int64
	TaxGroupId int64
	Tax        *Tax
	MarketDate time.Time
	TaxValue   float64
	BaseValue  float64
	TaxRate    float64
}
