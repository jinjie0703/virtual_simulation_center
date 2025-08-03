package models

import (
	"database/sql"
	"time"
)

// BaseModel 包含所有数据表共有的字段。
// 它通过 Go 的“结构体嵌入”被其他模型复用。
type BaseModel struct {
	ID        int          `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"-" db:"deleted_at"` // 在JSON响应中隐藏此字段
}
