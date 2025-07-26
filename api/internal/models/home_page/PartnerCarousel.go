package home_page

// PartnerCarousel 结构体定义了公司和学校轮播数据的通用结构
type PartnerCarousel struct {
	BaseModel
	Name string `json:"name"`
	Logo string `json:"logo"`
	URL  string `json:"url"`
}
