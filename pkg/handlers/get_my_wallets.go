package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntsd/alpha-interface-backend/pkg/entities"
)

func (h *handlers) GetMyWallets(c *gin.Context) {
	wallets := []entities.Wallet{
		{
			Amount:  100,
			CropIdx: 0,
		},
		{
			Amount:  50,
			CropIdx: 1,
		},
		{
			Amount:  10,
			CropIdx: 2,
		},
	}
	entities.NewResponse(wallets).Response(c, http.StatusOK)
}
