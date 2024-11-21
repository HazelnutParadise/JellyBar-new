package obj

type Article struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Icon        string   `json:"icon"`
	Description string   `json:"description"`
	PublishDate string   `json:"publishDate"`
	UpdateDate  string   `json:"updateDate"`
	Content     string   `json:"content"`
	Media       []string `json:"media"`
	CategoryID  uint
	Category    Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	AuthorID    uint
	Author      Author `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
}
