package home_page

import "gorm.io/gorm"

type HomeCarousel struct {
	gorm.Model
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	ImageUrl string `json:"imageUrl" gorm:"column:imageUrl"`
}

// TableName 方法告诉 GORM 这个结构体对应数据库中的哪一张表。
func (HomeCarousel) TableName() string {
	return "home_carousel"
}
