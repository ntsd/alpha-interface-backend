package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntsd/alpha-interface-backend/pkg/constants"
	"github.com/ntsd/alpha-interface-backend/pkg/entities"
	"github.com/ntsd/alpha-interface-backend/pkg/handlers"
)

// applyRouter is function for create http multiplexer router
func applyRouter(app *gin.Engine, handlers handlers.Handlers) {

	apiRoute := app.Group("/orders")
	{
		apiRoute.GET("/", handlers.GetOrders)
		apiRoute.POST("/", handlers.CreateOrder)
	}

	apiRoute = app.Group("/crops")
	{
		apiRoute.GET("/", handlers.GetCrops)
	}

	apiRoute = app.Group("/wallets")
	{
		apiRoute.GET("/", handlers.GetMyWallets)
	}

	apiRoute = app.Group("/iota-wallets")
	{
		apiRoute.GET("/", handlers.GetMyIotaWallet)
	}

	app.NoRoute(func(c *gin.Context) {
		entities.NewErrorResponse(&entities.ErrorStatus{
			Code:    constants.NotFoundErrorCode,
			Message: constants.NotFoundErrorMessage,
		}).Response(c, http.StatusNotFound)
	})
}
