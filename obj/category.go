package obj

type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description"`
	Articles    []Article `json:"articles" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
