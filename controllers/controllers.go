package controllers

import (
	"net/http"
	"weekly-task-2/models"
	"weekly-task-2/services"

	"github.com/labstack/echo/v4"
)

func GetItemsController(c echo.Context) error {
	items, err := services.GetAllItems()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all items",
		"items":   items,
	})
}

func GetItemController(c echo.Context) error {
	var itemId string = c.Param("id")

	item, err := services.GetItemByID(itemId)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "item not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "item found",
		"item":    item,
	})
}

func CreateItemController(c echo.Context) error {
	itemInput := models.Item{}
	c.Bind(&itemInput)

	item, err := services.CreateItem(itemInput)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new item",
		"item":    item,
	})
}

func DeleteItemController(c echo.Context) error {
	var itemId string = c.Param("id")

	isDeleted := services.DeleteItem(itemId)

	if !isDeleted {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "item not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "item deleted",
	})
}

func UpdateItemController(c echo.Context) error {
	var itemId string = c.Param("id")

	var itemInput models.Item = models.Item{}

	c.Bind(&itemInput)

	item, err := services.UpadteItem(itemInput, itemId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "item updated",
		"item":    item,
	})
}

// create function for category
func GetCategoriesController(c echo.Context) error {
	categories, err := services.GetAllCategories()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success get all categories",
		"categories": categories,
	})
}

func GetCategoryController(c echo.Context) error {
	var categoryId string = c.Param("id")

	category, err := services.GetCategoryByID(categoryId)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "category not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "category found",
		"category": category,
	})
}

func CreateCategoryController(c echo.Context) error {
	categoryInput := models.Category{}
	c.Bind(&categoryInput)

	category, err := services.CreateCategory(categoryInput)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new category",
		"category": category,
	})
}

func DeleteCategoryController(c echo.Context) error {
	var categoryId string = c.Param("id")

	isDeleted := services.DeleteCategory(categoryId)

	if !isDeleted {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "category not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "category deleted",
	})
}

func UpdateCategoryController(c echo.Context) error {
	var categoryId string = c.Param("id")

	var categoryInput models.Item = models.Item{}

	c.Bind(&categoryInput)

	category, err := services.UpadteItem(categoryInput, categoryId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "category updated",
		"category": category,
	})
}
