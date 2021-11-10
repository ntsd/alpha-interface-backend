package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntsd/alpha-interface-backend/pkg/entities"
)

// GetCrops ...
func (h *handlers) GetCrops(c *gin.Context) {
	crops := []entities.Crop{
		{
			Idx:     0,
			Country: "germany",
			Name:    "rice",
		},
		{
			Idx:     1,
			Country: "germany",
			Name:    "potato",
		},
		{
			Idx:     2,
			Country: "italy",
			Name:    "rice",
		},
		{
			Idx:     3,
			Country: "italy",
			Name:    "potato",
		},
	}
	entities.NewResponse(crops).Response(c, http.StatusOK)
}
