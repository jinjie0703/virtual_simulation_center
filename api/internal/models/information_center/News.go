package information_center

import (
	"virtual_simulation_center/api/internal/database"
)

// News 定义了新闻的数据结构
type News struct {
	BaseModel
	Title       string `json:"title" db:"title"`
	Summary     string `json:"summary" db:"summary"`
	Source      string `json:"source" db:"source"`
	ImageURL    string `json:"image_url" db:"image_url"`
	DetailURL   string `json:"detail_url" db:"detail_url"`
	Category    string `json:"category" db:"category"`
	PublishDate string `json:"publish_date" db:"publish_date"`
}

func GetAllNews() ([]News, error) {
	var news []News
	err := database.DB.Find(&news).Error
	return news, err
}

func GetNewsByID(id int) (News, error) {
	var news News
	err := database.DB.First(&news, id).Error
	return news, err
}
