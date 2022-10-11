package main

import (
	"weekly-task-2/controllers"
	"weekly-task-2/database"
	"weekly-task-2/middlewares"

	"github.com/labstack/echo/v4"
)

func init() {
	database.InitDB()
	database.InitialMigration()
}

func main() {
	// create a new echo instance
	e := echo.New()

	middlewares.LogMiddleware(e)

	// Route / to handler function
	e.GET("/items", controllers.GetItemsController)
	e.GET("/items/:id", controllers.GetItemController)
	e.POST("/items", controllers.CreateItemController)
	e.DELETE("/items/:id", controllers.DeleteItemController)
	e.PUT("/items/:id", controllers.UpdateItemController)

	e.GET("/categories", controllers.GetCategoriesController)
	e.GET("/categories/:id", controllers.GetCategoryController)
	e.POST("/categories", controllers.CreateCategoryController)
	e.DELETE("/categories/:id", controllers.DeleteCategoryController)
	e.PUT("/categories/:id", controllers.UpdateCategoryController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
