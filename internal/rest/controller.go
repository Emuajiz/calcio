package rest

import (
	"calcio/internal/group"
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	// Add fields as necessary
	echo   *echo.Echo
	logger *slog.Logger

	groupController *group.Controller
}

func NewController(e *echo.Echo, l *slog.Logger, gs *group.Controller) *Controller {
	return &Controller{
		echo:            e,
		logger:          l,
		groupController: gs,
	}
}

func (c *Controller) RegisterRoutes() {
	// Define your routes here
	c.echo.GET("/health", func(ctx echo.Context) error {
		return ctx.String(200, "OK")
	})

	c.groupController.RegisterRoutes(c.echo)
}

func (c *Controller) Start(address string) {
	go func() {
		if err := c.echo.Start(address); err != nil {
			c.logger.Error("Failed to start server", "error", err)
		}
	}()
}

func (c *Controller) Stop(ctx context.Context) error {
	return c.echo.Shutdown(ctx)
}
