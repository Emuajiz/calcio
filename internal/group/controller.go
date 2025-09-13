package group

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	logger  *slog.Logger
	service *Service
}

func NewController(logger *slog.Logger, service *Service) *Controller {
	return &Controller{logger: logger, service: service}
}

func (c *Controller) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/groups")
	g.POST("", c.CreateGroup)
	g.GET("/:id", c.GetGroup)
	g.GET("", c.ListGroups)
	g.PUT("/:id", c.UpdateGroup)
	g.DELETE("/:id", c.DeleteGroup)
}

func (c *Controller) CreateGroup(ctx echo.Context) error {
	var name string
	if err := ctx.Bind(&name); err != nil {
		c.logger.Error("Failed to bind request", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	group, err := c.service.CreateGroup(name)
	if err != nil {
		c.logger.Error("Failed to create group", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, group)
}

func (c *Controller) GetGroup(ctx echo.Context) error {
	id := ctx.Param("id")
	group, err := c.service.GetGroup(id)
	if err != nil {
		c.logger.Error("Failed to get group", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, group)
}

func (c *Controller) ListGroups(ctx echo.Context) error {
	groups, err := c.service.ListGroups()
	if err != nil {
		c.logger.Error("Failed to list groups", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, groups)
}

func (c *Controller) UpdateGroup(ctx echo.Context) error {
	var group Group
	if err := ctx.Bind(&group); err != nil {
		c.logger.Error("Failed to bind request", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id := ctx.Param("id")
	updatedGroup, err := c.service.UpdateGroup(id, group)
	if err != nil {
		c.logger.Error("Failed to update group", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, updatedGroup)
}

func (c *Controller) DeleteGroup(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := c.service.DeleteGroup(id); err != nil {
		c.logger.Error("Failed to delete group", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
