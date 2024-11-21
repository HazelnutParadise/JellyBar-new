package obj

type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Icon        string    `json:"icon"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Articles    []Article `json:"articles" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
