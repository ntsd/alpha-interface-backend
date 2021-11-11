package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"
	"github.com/ntsd/alpha-interface-backend/pkg/entities"
)

// GetOrders ...
func (h *handlers) GetOrders(c *gin.Context) {
	orders := []entities.Order{
		{
			Idx:     0,
			CropIdx: 0,
			Price:   5,
			Iota:    50,
			Amount:  10,
			Owner:   wasmlib.ScAgentID{},
			Status:  "opening",
			Type:    "sell",
		},
		{
			Idx:     1,
			CropIdx: 1,
			Price:   10,
			Iota:    150,
			Amount:  15,
			Owner:   wasmlib.ScAgentID{},
			Status:  "opening",
			Type:    "buy",
		},
		{
			Idx:     2,
			CropIdx: 1,
			Price:   5,
			Iota:    10,
			Amount:  2,
			Owner:   wasmlib.ScAgentID{},
			Status:  "opening",
			Type:    "buy",
		},
	}
	entities.NewResponse(orders).Response(c, http.StatusOK)
}
