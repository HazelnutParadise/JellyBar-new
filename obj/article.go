package obj

type Article struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PublishDate string   `json:"publishDate"`
	UpdateDate  string   `json:"updateDate"`
	Content     string   `json:"content"`
	Media       []string `json:"media"`
	CategoryID  uint     `json:"categoryId"`
	AuthorID    uint     `json:"authorId"`
	Icon        string   `json:"icon"`
}
