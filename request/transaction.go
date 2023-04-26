package requests

import (
	"time"
)

type TransactionCreateRequest struct {
	Name  string    `json:"name" binding:"required"`
	Type  int       `json:"type" binding:"gte=0,lte=1"`
	Value int       `json:"value" binding:"required"`
	Date  time.Time `json:"date"`
}
