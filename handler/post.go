package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lucasmb/getth-cms/service"
	"github.com/lucasmb/getth-cms/view/blog"
)


type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(ps *service.PostService) *PostHandler {
	return &PostHandler{
		service: ps,
	}
}

func PostRegisterHandler(e *echo.Echo, h *PostHandler) {
	items := e.Group("/blog")
	items.GET("", h.Home)
	items.GET("/posts", h.List)
	items.GET("/posts/:id", h.Show)
	//items.POST("", h.Create)

	//items.DELETE("/:itemId", h.Delete)
}

func (h *PostHandler) Home(c echo.Context) error {
	posts, err := h.service.List("published")
	if err != nil {
		return err
	}
	//component := pages.ItemListPage(items)
	return render(c, blog.Home(posts))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *PostHandler) List(c echo.Context) error {
	posts, err := h.service.List("published")
	if err != nil {
		return err
	}
	//component := pages.ItemListPage(items)
	return render(c, blog.Posts(posts))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *PostHandler) Show(c echo.Context) error {
	posts, err := h.service.List("published")
	if err != nil {
		return err
	}
	//component := pages.ItemListPage(items)
	return render(c, blog.Posts(posts))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}