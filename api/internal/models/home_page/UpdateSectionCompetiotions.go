package home_page

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// JSONB 类型用于存储 PostgreSQL 的 JSONB 数据
type JSONB []string

// Value 实现 driver.Valuer 接口
func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 实现 sql.Scanner 接口
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, j)
}

type UpdatesSectionCompetitions struct {
	BaseModel
	Custom_Id string    `json:"custom_id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	Category  string    `json:"category"`
	Deadline  time.Time `json:"deadline"`
	Tags      JSONB     `json:"tags"`
}

// TableName 方法告诉 GORM 这个结构体对应数据库中的哪一张表
func (UpdatesSectionCompetitions) TableName() string {
	return "updates_section_competitions"
}
