package obj

import "time"

type Article struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	PublishDate time.Time `json:"publishDate"`
	UpdateDate  time.Time `json:"updateDate"`
	Content     string    `json:"content"`
	Media       []string  `json:"media" gorm:"type:json"`
	CategoryID  uint
	Category    Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	AuthorID    uint
	Author      Author `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
}
