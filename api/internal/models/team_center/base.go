package team_center

import (
	"database/sql"
	"time"
)

// BaseModel defines the common fields for the team_center package models.
type BaseModel struct {
	ID        uint         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt sql.NullTime `json:"-" gorm:"column:deleted_at"`
}
