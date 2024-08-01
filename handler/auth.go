package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasmb/getth-cms/service"
	"github.com/lucasmb/getth-cms/view/dashboard"
	"github.com/lucasmb/getth-cms/view/userview"
)

type AuthHandler struct {
	uservice *service.UserService
}

func NewAuthHandler(us *service.UserService) *AuthHandler {
	return &AuthHandler{
		uservice: us,
	}
}

func AuthRegisterHandler(e *echo.Echo, ah *AuthHandler) {
	items := e.Group("/auth")
	items.GET("/login", ah.ShowLogin)
	items.POST("/login", ah.DoLogin)

}


func (h *AuthHandler) ShowLogin(c echo.Context) error {

	return render(c, dashboard.Show())

	//return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *AuthHandler) DoLogin(c echo.Context) error {

	users, err := h.uservice.List(1)
	if err != nil {
		return err
	}
	return Render(c, http.StatusOK, userview.All(users))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}