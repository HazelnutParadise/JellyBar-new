package obj

type ListItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Url         string `json:"url"`
	ButtonText  string `json:"buttonText"`
}
