package models

import (
	"github.com/daremove/go-musthave-diploma-tpl/tree/master/internal/utils"
)

type OrderStatus string

const (
	StatusNew        OrderStatus = "NEW"
	StatusProcessing OrderStatus = "PROCESSING"
	StatusInvalid    OrderStatus = "INVALID"
	StatusProcessed  OrderStatus = "PROCESSED"
)

type Order struct {
	ID         string            `json:"number"`
	Status     OrderStatus       `json:"status"`
	Accrual    *float64          `json:"accrual,omitempty"`
	UploadedAt utils.RFC3339Date `json:"uploaded_at"`
}
