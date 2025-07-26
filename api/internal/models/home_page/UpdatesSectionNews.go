package home_page

import (
	"time"
)

type UpdatesSectionNews struct {
	BaseModel
	Custom_Id string    `json:"custom_id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	Category  string    `json:"category"`
	Date      time.Time `json:"date"`
}

// TableName 方法告诉 GORM 这个结构体对应数据库中的哪一张表
func (UpdatesSectionNews) TableName() string {
	return "updates_section_news"
}
