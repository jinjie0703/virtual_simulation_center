package team_center

import (
	"time"
)

type CompetitionTeam struct {
	BaseModel
	Name              string    `json:"name" gorm:"column:name"`
	Leader            string    `json:"leader" gorm:"column:leader"`
	Description       string    `json:"description" gorm:"column:description"`
	Status            string    `json:"status" gorm:"column:status"`
	RecruitmentNumber int       `json:"recruitmentNumber" gorm:"column:recruitment_number"`
	Deadline          string    `json:"deadline" gorm:"column:deadline"`
	Tags              string    `json:"tags" gorm:"column:tags"`
	CreatedAt         time.Time `json:"createdAt" gorm:"column:created_at"`
	ContactWechat     string    `json:"contact_wechat" gorm:"column:contact_wechat"`
	ContactQq         string    `json:"contact_qq" gorm:"column:contact_qq"`
	ContactEmail      string    `json:"contact_email" gorm:"column:contact_email"`
}

func (CompetitionTeam) TableName() string {
	return "competition_teams"
}
