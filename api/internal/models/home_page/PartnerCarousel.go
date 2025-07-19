package home_page

// Partner 定义了合作伙伴的数据结构
type Partner struct {
	ID   int    `json:"id"`
	Name string `json:"name"` // name 字段用于图片的 alt 标签，符合您的要求
	Logo string `json:"logo"`
	URL  string `json:"url"`
}
