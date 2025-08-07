package team_center

import (
	"time"
	"virtual_simulation_center/api/internal/models/our_team"
)

type ProjectTeam struct {
	our_team.BaseModel
	Name              string    `json:"name" gorm:"column:name"`
	Leader            string    `json:"leader" gorm:"column:leader"`
	Description       string    `json:"description" gorm:"column:description"`
	Difficulty        string    `json:"difficulty" gorm:"column:difficulty"`
	RecruitmentNumber int       `json:"recruitmentNumber" gorm:"column:recruitment_number"`
	Duration          string    `json:"duration" gorm:"column:duration"`
	Tags              string    `json:"tags" gorm:"column:tags"`
	CreatedAt         time.Time `json:"createdAt" gorm:"column:created_at"`
	ContactWechat     string    `json:"contact_wechat" gorm:"column:contact_wechat"`
	ContactQq         string    `json:"contact_qq" gorm:"column:contact_qq"`
	ContactEmail      string    `json:"contact_email" gorm:"column:contact_email"`
}

func (ProjectTeam) TableName() string {
	return "project_teams"
}
