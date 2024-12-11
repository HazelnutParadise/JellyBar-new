package db

import (
	"jellybar/obj"
	"time"
)

func GetArticleByID(articleID uint, onlyPublished bool) (*obj.Article, error) {
	var articles []obj.Article
	var err error
	if onlyPublished {
		err = database.Preload("Category").Preload("Author").Where("id = ? AND status = ?", articleID, "publish").Find(&articles).Error
	} else {
		err = database.Preload("Category").Preload("Author").Where("id = ?", articleID).Find(&articles).Error
	}
	if err != nil || len(articles) == 0 {
		return nil, err
	}
	return &articles[0], nil
}

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
	article.PublishAt = time.Now()
	article.UpdateAt = time.Now()
	err := database.Create(article).Error
	return err
}
