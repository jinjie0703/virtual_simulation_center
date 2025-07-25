package home_page

import "gorm.io/gorm"

// PartnerCarousel 结构体映射了 'partner_carousel' 表
type PartnerCarousel struct {
	gorm.Model
	Name string `json:"name"` // name 字段用于图片的 alt 标签
	Logo string `json:"logo"`
	URL  string `json:"url"`
}

// TableName 方法告诉 GORM 这个结构体对应数据库中的哪一张表
func (PartnerCarousel) TableName() string {
	return "partner_carousel"
}
