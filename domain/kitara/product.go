package kitara

import (
	"errors"
	"fmt"
)

type Product struct {
	ID    int
	Name  string
	Stock []*ProductVariant
}

type ProductVariant struct {
	ID    int    `json:"id"`
	Sku   string `json:"sku"`
	Size  string `json:"size"`
	Color string `json:"color"`
	Qty   int    `json:"qty"`
}

func NewProductVariant(id int, sku string, size string, color string, qty int) *ProductVariant {
	return &ProductVariant{
		ID:    id,
		Sku:   sku,
		Size:  size,
		Color: color,
		Qty:   qty,
	}
}

func (pv *ProductVariant) ReduceQty(qty int) error {
	if pv.Qty < 0 {
		return errors.New("Product is out of stock")
	}

	if pv.Qty-qty < 0 {
		return errors.New(fmt.Sprintf("Only %d qty left", pv.Qty))
	}

	pv.Qty = pv.Qty - qty
	return nil
}

type ReserveProductReq struct {
	CustomerID       string `json:"customerId"`
	ProductVariantID int    `json:"productVariantId"`
	Qty              int    `json:"qty"`
}

type ReserveProductResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    *ReserveProductReq `json:"data"`
}

func NewReserveProductResponse(success bool, msg string, req *ReserveProductReq) *ReserveProductResponse {
	return &ReserveProductResponse{
		Success: success,
		Message: msg,
		Data:    req,
	}
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
