package obj

import "time"

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Uuid           string    `json:"uuid" gorm:"unique"`
	AuthorID       uint      `json:"author_id"`
	Author         Author    `json:"author" gorm:"foreignKey:AuthorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Username       string    `json:"username"`
	Name           string    `json:"name"`
	Role           string    `json:"role"`
	CreateAt       time.Time `json:"create_at"`
	Status         string    `json:"status"`
	StatusUpdateAt time.Time `json:"status_update_at"`
	StatusReason   string    `json:"status_reason"`
}
