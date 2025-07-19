package home_page

import "gorm.io/gorm"

// HomePageCarousel 结构体精确地映射了 'home_page_carousel' 表
type HomeCarousel struct {
	gorm.Model        // 这包含了 ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	// 你的数据库列名是 imageUrl (驼峰)，而不是 image_url (蛇形)
	// 所以我们必须用 gorm 标签来明确指定列名，这一点非常重要！
	ImageUrl string `json:"imageUrl" gorm:"column:imageUrl"`
}

// TableName 方法告诉 GORM 这个结构体对应数据库中的哪一张表。
// 如果没有这个方法，GORM 会默认去寻找 'home_page_carousels' (复数形式) 这张表。
func (HomeCarousel) TableName() string {
	return "HomeCarousel"
}
