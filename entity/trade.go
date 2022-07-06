package entity

type Trade struct {
	Invoice        *Invoice
	InvoiceItemsIn []InvoiceItem
	InvoiceItemOut *InvoiceItem
}
