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
		apiRoute.GET("/", handlers.CheckLiveness)
		apiRoute.POST("/", handlers.CheckLiveness)
	}

	apiRoute = app.Group("/positions")
	{
		apiRoute.GET("/", handlers.CheckLiveness)
	}

	app.NoRoute(func(c *gin.Context) {
		entities.NewErrorResponse(&entities.ErrorStatus{
			Code:    constants.NotFoundErrorCode,
			Message: constants.NotFoundErrorMessage,
		}).Response(c, http.StatusNotFound)
	})
}
