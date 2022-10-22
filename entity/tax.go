package entity

import "time"

type TaxSourceType string

const (
	BIV TaxSourceType = "BIV"
	EAR TaxSourceType = "EAR"
	ITB TaxSourceType = "ITB"
	TRB TaxSourceType = "TRB"
	TRD TaxSourceType = "TRD"
)

type Tax struct {
	Id     int64
	Code   string
	Source string
	Rate   float64
}

type TaxGroup struct {
	Id         int64
	Source     TaxSourceType
	ExternalId string
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
