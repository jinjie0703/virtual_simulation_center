package home_page

import (
	"time"

	"gorm.io/gorm"
)

// UpdatesSection 结构体映射了 'updates_section' 表
type UpdatesSection struct {
	gorm.Model
	Type        string    `json:"type"` // 'news' or 'competition'
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// TableName 方法告诉 GORM 这个结构体对应数据库中的哪一张表
func (UpdatesSection) TableName() string {
	return "updates_section"
}
