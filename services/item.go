package services

import (
	"errors"
	"weekly-task-2/database"
	"weekly-task-2/models"
)

func GetAllItems() ([]models.Item, error) {
	var items []models.Item = []models.Item{}

	if err := database.DB.Find(&items).Error; err != nil {
		return []models.Item{}, err
	}

	return items, nil
}

func GetItemByID(id string) (models.Item, error) {
	var item models.Item = models.Item{}

	if err := database.DB.Find(&item, "id = ?", id).Error; err != nil {
		return models.Item{}, err
	}

	if item.ID == "" {
		return models.Item{}, errors.New("item not found")
	}

	return item, nil
}

func CreateItem(input models.Item) (models.Item, error) {
	var user models.Item = models.Item{
		Name:        input.Name,
		Description: input.Description,
		Qty:         input.Qty,
		CategoryID:  input.CategoryID,
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return models.Item{}, err
	}

	return user, nil
}

func UpadteItem(input models.Item, id string) (models.Item, error) {
	item, err := GetItemByID(id)

	if err != nil {
		return models.Item{}, nil
	}

	item.Name = input.Name
	item.Description = input.Description
	item.Qty = input.Qty

	if err := database.DB.Save(&item).Error; err != nil {
		return models.Item{}, nil
	}

	return item, nil
}

func DeleteItem(id string) bool {
	item, err := GetItemByID(id)

	if err != nil {
		return false
	}

	if err := database.DB.Delete(&item).Error; err != nil {
		return false
	}

	return true
}
