package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntsd/alpha-interface-backend/pkg/entities"
)

type CreateOrderInput struct {
	CropIdx int32  `json:"cropIdx"`
	Type    string `json:"type"`
	Price   int64  `json:"price"`
	Amount  int64  `json:"amount"`
}

func (h *handlers) CreateOrder(c *gin.Context) {
	createOrderInput := &CreateOrderInput{}
	err := h.decodeAndValidate(c.Request.Body, createOrderInput)
	if err != nil {
		entities.NewErrorResponse(err).Response(c, http.StatusBadRequest)
		return
	}
	entities.NewResponse(nil).Response(c, http.StatusOK)
}
