package obj

import "time"

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Uuid           string    `json:"uuid"`
	Username       string    `json:"username"`
	Name           string    `json:"name"`
	Role           string    `json:"role"`
	CreateAt       time.Time `json:"create_at"`
	Status         string    `json:"status"`
	StatusUpdateAt time.Time `json:"status_update_at"`
	StatusReason   string    `json:"status_reason"`
}
