package obj

import "time"

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Uuid           string    `json:"uuid" gorm:"unique"`
	Username       string    `json:"username"`
	Name           string    `json:"name"`
	Role           UserRole  `json:"-"`
	CreateAt       time.Time `json:"create_at"`
	Status         string    `json:"status"`
	StatusUpdateAt time.Time `json:"status_update_at"`
	StatusReason   string    `json:"status_reason"`
	Author         Author    `json:"author" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserRole uint8

const (
	UserRoleUser UserRole = iota + 1
	UserRoleAuthor
	UserRoleEditor
	UserRoleAdmin
)

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusBanned    UserStatus = "banned"
)
