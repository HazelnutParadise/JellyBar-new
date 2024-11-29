package db

import "jellybar/obj"

func GetCategories() ([]obj.Category, error) {
	var categories []obj.Category
	err := database.Find(&categories).Error
	return categories, err
}
