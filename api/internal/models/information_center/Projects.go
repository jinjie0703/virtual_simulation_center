package information_center

import "api/internal/models"

// Project 定义了项目招募信息的数据结构。
type Project struct {
	models.BaseModel        // 嵌入基础模型
	Title            string `json:"title" db:"title"`
	Summary          string `json:"summary" db:"summary"`
	Deadline         string `json:"deadline" db:"deadline"`
	Tags             string `json:"tags" db:"tags"`
	DetailURL        string `json:"detail_url" db:"detail_url"`
}
