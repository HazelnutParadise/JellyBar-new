package db

import "jellybar/obj"

func GetArticles(categoryId string, onlyPublished bool) (*[]obj.Article, error) {
	var articles []obj.Article
	var err error
	if categoryId != "" {
		if onlyPublished {
			err = database.Where("category_id = ? AND status = ?", categoryId, "publish").Find(&articles).Error
		} else {
			err = database.Where("category_id = ?", categoryId).Find(&articles).Error
		}
	} else {
		if onlyPublished {
			err = database.Where("status = ?", "publish").Preload("Category").Find(&articles).Error
		} else {
			err = database.Preload("Category").Find(&articles).Error
		}
	}
	return &articles, err
}

func AddArticle(article *obj.Article) error {
	err := database.Create(article).Error
	return err
}
