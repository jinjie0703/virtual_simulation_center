package our_team

import (
	"database/sql"
	"time"
)

// BaseModel defines the common fields for the our_team package models.
// The ID is a string to match the varchar type in the team_members table.
type BaseModel struct {
	ID        string       `json:"id" gorm:"column:id;primaryKey"`
	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt sql.NullTime `json:"-" gorm:"column:deleted_at"`
}
