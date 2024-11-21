package obj

type Category struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Articles    []Article `json:"articles" gorm:"constraint:OnUpdate:CASCADE;"`
}
