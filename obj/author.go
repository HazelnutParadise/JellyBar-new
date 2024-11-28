package obj

type Author struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id"`
	Name         string    `json:"name"`
	ProfileImage string    `json:"profile_image"`
	Description  string    `json:"description"`
	Articles     []Article `json:"articles" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
