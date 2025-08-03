package information_center

import (
	"api/internal/database"
	"api/internal/models"
)

// News 定义了新闻文章的数据结构。
type News struct {
	models.BaseModel        // 嵌入基础模型，自动获得 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	Title            string `json:"title" db:"title"`
	Summary          string `json:"summary" db:"summary"`
	PublishDate      string `json:"publish_date" db:"publish_date"`
	Category         string `json:"category" db:"category"`
	DetailURL        string `json:"detail_url" db:"detail_url"`
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
