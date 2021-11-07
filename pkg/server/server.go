package server

import (
	"time"

	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/ntsd/alpha-interface-backend/pkg/config"
	"github.com/ntsd/alpha-interface-backend/pkg/handlers"
	"github.com/ntsd/alpha-interface-backend/pkg/utils/loggers"
	"go.uber.org/zap/zapcore"
)

// Serve is function for the application http server
func Serve(cfg config.Config) error {
	logger := loggers.BuildLogger(zapcore.InfoLevel)

	handlers := handlers.NewHandlers(handlers.NewHandlersInput{
		Config: cfg,
	})

	app := gin.Default()
	app.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	app.Use(ginzap.RecoveryWithZap(logger, true))

	applyRouter(app, handlers)

	if cfg.DevMode {
		pprof.Register(app)
	}

	return app.Run(":" + cfg.Port)
}
