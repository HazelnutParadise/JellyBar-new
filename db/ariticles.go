package db

import (
	"fmt"
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

	tx := database.Begin()
	var err error

	if article.CategoryID == 0 {
		newCate := obj.Category{
			Name: article.Category.Name,
		}
		// 新增類別
		if err = tx.Create(&newCate).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("create category error: %v", err)
		}

		// 取得新類別 ID
		var createdCategory obj.Category
		if err = tx.Where("name = ?", newCate.Name).First(&createdCategory).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("find category error: %v", err)
		}

		article.CategoryID = createdCategory.ID
		fmt.Printf("設置文章類別 ID: %d\n", article.CategoryID)
	}
	newArticle := obj.Article{
		Title:      article.Title,
		Content:    article.Content,
		CategoryID: article.CategoryID,
		AuthorID:   article.AuthorID,
		Status:     article.Status,
		PublishAt:  article.PublishAt,
		UpdateAt:   article.UpdateAt,
	}
	// 新增文章
	if err = tx.Create(&newArticle).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("create article error: %v", err)
	}

	return tx.Commit().Error
}

func UpdateArticle(id string, article *obj.Article) error {
	article.UpdateAt = time.Now()
	err := database.Model(&obj.Article{}).Where("id = ?", id).Updates(article).Error
	return err
}

func DeleteArticle(id string) error {
	err := database.Where("id = ?", id).Delete(&obj.Article{}).Error
	return err
}
