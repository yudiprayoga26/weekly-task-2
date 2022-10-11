package services

import (
	"errors"
	"weekly-task-2/database"
	"weekly-task-2/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category = []models.Category{}

	if err := database.DB.Find(&categories).Error; err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}

func GetCategoryByID(id string) (models.Category, error) {
	var category models.Category = models.Category{}

	if err := database.DB.Find(&category, "id = ?", id).Error; err != nil {
		return models.Category{}, err
	}

	if category.ID == 0 {
		return models.Category{}, errors.New("category not found")
	}

	return category, nil
}

func CreateCategory(input models.Category) (models.Category, error) {
	var category models.Category = models.Category{
		Name: input.Name,
	}

	if err := database.DB.Save(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func UpdateCategory(input models.Category, id string) (models.Category, error) {
	category, err := GetCategoryByID(id)

	if err != nil {
		return models.Category{}, nil
	}

	category.Name = input.Name

	if err := database.DB.Save(&category).Error; err != nil {
		return models.Category{}, nil
	}

	return category, nil
}

func DeleteCategory(id string) bool {
	category, err := GetCategoryByID(id)

	if err != nil {
		return false
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		return false
	}

	return true
}
