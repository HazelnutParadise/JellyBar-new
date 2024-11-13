package obj

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishDate string `json:"publishDate"`
	UpdateDate  string `json:"updateDate"`
	Url         string `json:"url"`
	Icon        string `json:"icon"`
	ButtonText  string `json:"buttonText"`
	Category    string `json:"category"`
	Author      string `json:"author"`
}
