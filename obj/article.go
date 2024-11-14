package obj

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishDate string `json:"publishDate"`
	UpdateDate  string `json:"updateDate"`
	Content     string `json:"content"`
	Url         string `json:"url"`
	ButtonText  string `json:"buttonText"`
	Category    string `json:"category"`
	Author      string `json:"author"`
}
