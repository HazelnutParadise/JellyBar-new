package obj

type Article struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PublishDate string   `json:"publishDate"`
	UpdateDate  string   `json:"updateDate"`
	Content     string   `json:"content"`
	Media       []string `json:"media"`
	Category    string   `json:"category"`
	Author      string   `json:"author"`
	Icon        string   `json:"icon"`
}
