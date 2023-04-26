package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"himymo/models"
	"himymo/repository"
	"himymo/request"
	"himymo/response"
	"himymo/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(db *gorm.DB) *Handler {
	repo := &repository.Repository{DB: db}
	return &Handler{
		service: service.NewService(repo),
	}
}

func (h *Handler) CreateTransaction(c *gin.Context) {
	var request requests.TransactionCreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response := response.NewErrorResponse(http.StatusBadRequest, err)
		c.JSON(response.Code, response)
		return
	}

	transaction := models.Transaction{
		Name:  request.Name,
		Type:  request.Type,
		Value: request.Value,
		Date:  request.Date,
	}

	if err := h.service.CreateTransaction(&transaction); err != nil {
		response := response.NewErrorResponse(http.StatusInternalServerError, err)
		c.JSON(response.Code, response)
		return
	}

	response := response.NewSuccessResponse(transaction)
	c.JSON(response.Code, response)
}

func (h *Handler) GetTransactionById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := response.NewErrorResponse(http.StatusBadRequest, err)
		c.JSON(response.Code, response)
		return
	}

	transaction, err := h.service.GetTransactionById(id)
	if err != nil {
		response := response.NewErrorResponse(http.StatusNotFound, err)
		c.JSON(response.Code, response)
		return
	}

	response := response.NewSuccessResponse(transaction)
	c.JSON(response.Code, response)
}

func (h *Handler) GetTransactions(c *gin.Context) {
	transactions, err := h.service.GetTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(transactions))
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	var request requests.TransactionCreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid ID parameter")))
		return
	}

	// Cek apakah data yang diminta terdapat di database
	transaction, err := h.service.GetTransactionById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}
	if transaction == nil {
		c.JSON(http.StatusNotFound, response.NewErrorResponse(http.StatusNotFound, fmt.Errorf("transaction not found")))
		return
	}

	transaction.Name = request.Name
	transaction.Type = request.Type
	transaction.Value = request.Value
	transaction.Date = request.Date

	if err := h.service.UpdateTransaction(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(transaction))
}

func (h *Handler) DeleteTransactionById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := response.NewErrorResponse(http.StatusBadRequest, err)
		c.JSON(response.Code, response)
		return
	}

	if err := h.service.DeleteTransactionById(id); err != nil {
		response := response.NewErrorResponse(http.StatusInternalServerError, err)
		c.JSON(response.Code, response)
		return
	}

	response := response.NewSuccessResponse(gin.H{"message": "transaction deleted"})
	c.JSON(response.Code, response)
}
