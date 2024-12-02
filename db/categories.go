package db

import "jellybar/obj"

func GetCategories(preloadArticles bool) (*[]obj.Category, error) {
	var categories []obj.Category
	var err error
	if preloadArticles {
		err = database.Preload("Articles").Preload("Articles.Category").Find(&categories).Error
	} else {
		err = database.Find(&categories).Error
	}
	return &categories, err
}

func AddCategory(category *obj.Category) error {
	err := database.Create(category).Error
	return err
}

func UpdateCategory(category *obj.Category) error {
	err := database.Model(category).Updates(category).Error
	return err
}

func DeleteCategory(category *obj.Category) error {
	err := database.Delete(category).Error
	return err
}
