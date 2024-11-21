package obj

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishDate string `json:"publishDate"`
	UpdateDate  string `json:"updateDate"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	Author      string `json:"author"`
	Icon        string `json:"icon"`
}
