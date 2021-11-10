package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntsd/alpha-interface-backend/pkg/entities"
)

// CheckLiveness ...
func (h *handlers) CheckLiveness(c *gin.Context) {
	entities.NewResponse(nil).Response(c, http.StatusOK)
}
