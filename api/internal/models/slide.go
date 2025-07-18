// internal/models/slide.go
package models

import "gorm.io/gorm"

// Slide 对应数据库中的 'slides' 表
type Slide struct {
	gorm.Model
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ImageUrl string `json:"imageUrl" gorm:"column:image_url"` // 明确指定列名
}