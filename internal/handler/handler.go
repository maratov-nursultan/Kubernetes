package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/maratov-nursultan/Kubernetes/internal/manager/user"
	"github.com/maratov-nursultan/Kubernetes/internal/model"
)

type Handler struct {
	userManager user.UserSDK
}

func NewHandler(userManager user.UserSDK) *Handler {
	return &Handler{
		userManager: userManager,
	}
}

func (h *Handler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(model.GetUserRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := h.userManager.GetUser(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func (h *Handler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(model.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	id, err := h.userManager.CreateUser(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(200, id)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(model.DeleteUserRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	err := h.userManager.Delete(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(200, "Success")
}

func (h *Handler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(model.UpdateUserRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	err := h.userManager.Update(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(200, "Success")
}
