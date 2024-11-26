package obj

import "time"

type Article struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishAt   time.Time `json:"publishDate"`
	UpdateAt    time.Time `json:"updateDate"`
	Status      string    `json:"status"`
	Content     string    `json:"content"`
	Media       []string  `json:"media" gorm:"type:json"`
	CategoryID  uint
	Category    Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	AuthorID    uint
	Author      Author `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
}
