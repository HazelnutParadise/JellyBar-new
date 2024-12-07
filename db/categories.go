package db

import "jellybar/obj"

type GetCategoriesOption struct {
	PreloadArticles bool
	OnlyPublished   bool
}

func GetCategories(option GetCategoriesOption) (*[]obj.Category, error) {
	var categories []obj.Category
	var err error
	query := database

	if option.PreloadArticles {
		if option.OnlyPublished {
			query = query.Preload("Articles", "status = ?", "publish")
		} else {
			query = query.Preload("Articles")
		}
		query = query.Preload("Articles.Category")
	}

	err = query.Find(&categories).Error
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
