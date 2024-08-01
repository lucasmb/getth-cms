package handler

import (
	"log"
	"strconv"

	"github.com/lucasmb/getth-cms/models"
	"github.com/lucasmb/getth-cms/service"
	"github.com/lucasmb/getth-cms/view/components"
	"github.com/lucasmb/getth-cms/view/userview"

	"github.com/labstack/echo/v4"
)

// type UserHandler struct {}

// func (h UserHandler ) HandleUserShow(c echo.Context) error {
// 	fmt.Printf("User Handler")
// 	//return user.Show().Render(c.Request().Context(), )
// 	return render(c, user.Show())
// }

// func (h UserHandler ) HandleUserAll(c echo.Context) error {

// 	return render(c, user.All())
// }


type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{
		service: us,
	}
}

func UserRegisterHandler(e *echo.Echo, h *UserHandler) {
	items := e.Group("/users")
	items.GET("", h.List)
	items.GET("/:id", h.Show)

	items.GET("/table", h.Table)
	items.POST("", h.Create)
	items.GET("/resume", h.Resumee)

	//items.DELETE("/:itemId", h.Delete)
}

func (h *UserHandler) Table(c echo.Context) error {
	users, err := h.service.List(1)
	if err != nil {
		return err
	}
	//component := pages.ItemListPage(items)
	return render(c, components.UsersTable(users))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *UserHandler) Resumee(c echo.Context) error {

	return render(c, userview.Resumee())

}

func (h *UserHandler) List(c echo.Context) error {
	users, err := h.service.List(1)
	if err != nil {
		return err
	}
	//component := pages.ItemListPage(items)
	return render(c, userview.All(users))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *UserHandler) Show(c echo.Context) error {

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user, err := h.service.GetById(idParams)
	if err != nil {
		return err
	}
	//component := pages.ItemListPage(items)
	return render(c, userview.Show(user))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}


 func (h *UserHandler) Create(c echo.Context) error {
	// isError := false
	log.Printf("POST Create user")

	if c.Request().Method == "POST" {
	newItem := models.User{
		FirstName: c.FormValue("first_name"),
		LastName: c.FormValue("last_name"),
		Username: c.FormValue("username"),
		Email: c.FormValue("email"),
		Password: "Test123",
		Role: "user",
		TenantId: 1,


	}
	log.Printf("User: +v", newItem)

	_, err := h.service.Create(newItem)
	if err != nil {
		log.Print(err) 
		//return err
	}

	setFlashmessages(c, "success", "User created successfully!!")


	c.Response().Header().Set("HX-Trigger", "newUser")
	c.Response().WriteHeader(201)
	//return render(c, components.UserForm())
	// return false
	// return c.Redirect(http.StatusSeeOther, "/users") 
	}

	return c.JSON(201, true)
	// return render(c, userview.Show())

	//component := partials.ItemRow(item)
	//return component.Render(c.Request().Context(), c.Response().Writer)
}
