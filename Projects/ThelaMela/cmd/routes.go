package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (app *Application) getRoutes() *echo.Echo {
	e := echo.New()
	e.GET("/", HomeHandler)
	e.GET("/users", app.GetUsers)
	e.GET("/users/:user_id", app.GetUser)
	e.POST("/users/signup", app.signup)
	e.POST("/users/login", app.login)

	e.GET("/foods", app.GetFoods)
	e.GET("/foods/:food_id", app.GetFoodItem)
	e.POST("/foods", app.Createfood)
	e.PATCH("/foods/:food_id", app.UpdateFood)

	e.GET("/invoices", app.GetInvoice)
	e.GET("/invoices/:invoice_id", app.GetInvoiceItem)
	e.POST("/invoices", app.CreateInvoice)
	e.PATCH("/invoices/:invoice_id", app.UpdateInvoice)

	e.GET("/menus", app.GetMenus)
	e.GET("/menus/:menu_id", app.GetMenu)
	e.POST("/menus", app.CreateMenu)
	e.PATCH("/menus/:menu_id", app.UpdateMenu)

	e.GET("/orders", app.GetOrders)
	e.GET("/orders/:order_id", app.GetOrder)
	e.POST("/orders", app.CreateOrder)
	e.PATCH("/orders/:order_id", app.UpdateOrder)

	e.GET("/tables", app.GetTables)
	e.GET("/tables/:table_id", app.GetTable)
	e.POST("/tables", app.CreateTable)
	e.PATCH("/tables/:table_id", app.UpdateTable)

	e.GET("/orderItems", app.GetOrderItems)
	e.GET("/orderItems/:orderItem_id", app.GetOrderItem)
	e.GET("/orderItems-order/:order_id", app.GetOrderItemsByOrder)
	e.POST("/orderItems", app.CreateOrderItem)
	e.PATCH("/orderItems/:orderItem_id", app.UpdateOrderItem)

	return e
}

func HomeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
