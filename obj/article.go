package obj

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Icon        string `json:"icon"`
	ButtonText  string `json:"buttonText"`
	Category    string `json:"category"`
}
