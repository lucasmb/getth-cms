package handler

import (
	"log"
	"strconv"

	"github.com/lucasmb/getth-cms/models"
	"github.com/lucasmb/getth-cms/service"
	"github.com/lucasmb/getth-cms/view/blog"
	"github.com/lucasmb/getth-cms/view/components"

	"github.com/lucasmb/getth-cms/view/userview"

	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	uservice *service.UserService
	pservice *service.PostService
}

func NewDashboardHandler(us *service.UserService, ps *service.PostService) *DashboardHandler {
	return &DashboardHandler{
		uservice: us,
		pservice: ps,
	}
}

func DashboardRegisterHandler(e *echo.Echo, dh *DashboardHandler) {
	dash := e.Group("/dashboard")
	dash.GET("", dh.Home)

	items := dash.Group("/users")
	items.GET("", dh.List)
	items.GET("/:id", dh.Show)

	items.GET("/table", dh.Table)
	items.POST("", dh.Create)


}

func (dh *DashboardHandler) Home(c echo.Context) error {
	posts, err := dh.pservice.List("published")
	if err != nil {
		return err
	}
	//component := pages.ItemListPage(items)
	return render(c, blog.Home(posts))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}



func (dh *DashboardHandler) List(c echo.Context) error {
	users, err := dh.uservice.List(1)
	if err != nil {
		return err
	}
	//component := pages.ItemListPage(items)
	return render(c, userview.All(users))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}

func (dh *DashboardHandler) Table(c echo.Context) error {
	users, err := dh.uservice.List(1)
	if err != nil {
		return err
	}
	//component := pages.ItemListPage(items)
	return render(c, components.UsersTable(users))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}

func (dh *DashboardHandler) Show(c echo.Context) error {

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user, err := dh.uservice.GetById(idParams)
	if err != nil {
		return err
	}
	//component := pages.ItemListPage(items)
	return render(c, userview.Show(user))

	//return component.Render(c.Request().Context(), c.Response().Writer)
}


 func (dh *DashboardHandler) Create(c echo.Context) error {
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

	_, err := dh.uservice.Create(newItem)
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
