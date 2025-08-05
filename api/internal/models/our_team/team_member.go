package our_team

// TeamMember represents the structure of a team member in the database.
type TeamMember struct {
	BaseModel
	Name              string `json:"name" gorm:"column:name"`
	Title             string `json:"title" gorm:"column:title"`
	Role              string `json:"role" gorm:"column:role"`
	Description       string `json:"description" gorm:"column:description"`
	AvatarURL         string `json:"avatar_url" gorm:"column:avatar_url"`
	Email             string `json:"email" gorm:"column:email"`
	Office            string `json:"office" gorm:"column:office"`
	ResearchInterests string `json:"research_interests" gorm:"column:research_interests"`
	Achievements      string `json:"achievements" gorm:"column:achievements"`
}

// TableName specifies the table name for the TeamMember model
func (TeamMember) TableName() string {
	return "team_members"
}
