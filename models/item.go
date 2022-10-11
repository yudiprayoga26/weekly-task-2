package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID          string `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Qty         int    `json:"qty" form:"qty"`
	CategoryID  int    `json:"category_id" form:"category_id"`
}
