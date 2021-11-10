package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/ntsd/alpha-interface-backend/pkg/config"
)

// Handlers is interface for handlers
type Handlers interface {
	CheckLiveness(c *gin.Context)
	GetMyIotaWallet(c *gin.Context)
	CreateOrder(c *gin.Context)
	GetOrders(c *gin.Context)
	GetCrops(c *gin.Context)
	GetMyWallets(c *gin.Context)
}

// handlers is struct for handlers
type handlers struct {
	Config    config.Config
	Validator *validator.Validate
}

type NewHandlersInput struct {
	Config config.Config
}

// NewHandlers is for create a new Handlers
func NewHandlers(input NewHandlersInput) Handlers {
	return &handlers{
		Config:    input.Config,
		Validator: validator.New(),
	}
}
