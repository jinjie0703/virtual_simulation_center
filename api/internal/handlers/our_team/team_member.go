package our_team

import (
	"virtual_simulation_center/api/internal/models/our_team"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetTeamMembers retrieves all team members from the database.
func GetTeamMembers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var members []our_team.TeamMember
		if err := db.Find(&members).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve team members"})
			return
		}
		c.JSON(http.StatusOK, members)
	}
}
