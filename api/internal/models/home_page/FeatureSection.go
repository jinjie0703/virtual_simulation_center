package home_page

import "gorm.io/gorm"

// FeatureSection 结构体映射了 'feature_section' 表
type FeatureSection struct {
	gorm.Model
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TableName 方法告诉 GORM 这个结构体对应数据库中的哪一张表
func (FeatureSection) TableName() string {
	return "feature_section"
}
