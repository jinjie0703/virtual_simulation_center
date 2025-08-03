package information_center

import (
	"api/internal/database"
	"api/internal/models"
)

// Competition 定义了竞赛信息的数据结构。
type Competition struct {
	models.BaseModel         // 嵌入基础模型
	Title            string `json:"title" db:"title"`
	Summary          string `json:"summary" db:"summary"`
	Deadline         string `json:"deadline" db:"deadline"`
	Level            string `json:"level" db:"level"`
	Tags             string `json:"tags" db:"tags"`
	DetailURL        string `json:"detail_url" db:"detail_url"`
}

func GetAllCompetitions() ([]Competition, error) {
	var competitions []Competition
	err := database.DB.Find(&competitions).Error
	return competitions, err
}
