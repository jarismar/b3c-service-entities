package entity

import "time"

const (
	BIV string = "BIV"
	EAR string = "EAR"
	ITB string = "ITB"
	TRB string = "TRB"
	TRD string = "TRD"
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
	Tax        *Tax
	MarketDate time.Time
	TaxValue   float64
	BaseValue  float64
	TaxRate    float64
}
