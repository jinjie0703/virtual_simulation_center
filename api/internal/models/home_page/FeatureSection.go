package home_page

import (
	"encoding/json"

	"gorm.io/gorm"
)

// FeatureSection 结构体映射了 'achievement_showcase' 表
type FeatureSection struct {
	gorm.Model
	Title        string          `json:"Title" gorm:"column:title"`
	Description  string          `json:"Description" gorm:"column:description"`
	Image        string          `json:"Image" gorm:"column:image"`
	Tags         json.RawMessage `json:"Tags" gorm:"column:tags"`
	Author       string          `json:"Author" gorm:"column:author"`
	AuthorTitle  string          `json:"AuthorTitle" gorm:"column:authorTitle"`
	AuthorAvatar string          `json:"AuthorAvatar" gorm:"column:authorAvatar"`
	Contact1     string          `json:"Contact1" gorm:"column:contact1"`
	Contact2     string          `json:"Contact2" gorm:"column:contact2"`
	ProjectUrl   string          `json:"ProjectUrl" gorm:"column:projectUrl"`
	VideoUrl     string          `json:"VideoUrl" gorm:"column:videoUrl"`
	Gallery      json.RawMessage `json:"Gallery" gorm:"column:gallery"`
}

// TableName 方法告诉 GORM 这个结构体对应数据库中的哪一张表
func (FeatureSection) TableName() string {
	return "achievement_showcase"
}
