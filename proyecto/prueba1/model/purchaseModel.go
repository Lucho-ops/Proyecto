package model

type Purchase struct {
	ClientID int64     `json:"clientId"`
	Name     string    `json:"nombre"`
	Purchase bool      `json:"compro"`
	Tdc      string    `json:"tdc"`
	Amount   float64   `json:"monto"`
	Date     string `json:"date"`
}

type ConsolidatedPurchase struct {
	Total       float64  `json:"total"`
	PurchaseTDC CardType `json:"comprasPorTDC"`
	NotPurchase int    `json:"nocompraron"`
	BuyHigher   float64    `json:"compraMasAlta"`
}

type CardType struct {
	Gold float64 `json:"oro"`
	Amex float64 `json:"amex"`
}
