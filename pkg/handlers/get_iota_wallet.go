package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntsd/alpha-interface-backend/pkg/entities"
)

func (h *handlers) GetMyIotaWallet(c *gin.Context) {
	entities.NewResponse(entities.IOTAWallet{
		Address: "fn0nPelxak2LGyr50r6byg",
		Balance: 1000000,
	}).Response(c, http.StatusOK)
}
