package team_center

import (
	"net/http"
	"strconv"
	"virtual_simulation_center/api/internal/database"
	"virtual_simulation_center/api/internal/models/team_center"

	"github.com/gin-gonic/gin"
)

// GetAllCompetitionTeams godoc
// @Summary Get all competition teams
// @Description Get all competition teams
// @Tags team_center
// @Accept json
// @Produce json
// @Success 200 {array} team_center.CompetitionTeam
// @Router /team_center/competition_teams [get]
func GetAllCompetitionTeams(c *gin.Context) {
	var competitionTeams []team_center.CompetitionTeam
	if err := database.DB.Find(&competitionTeams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, competitionTeams)
}

// GetCompetitionTeamByID godoc
// @Summary Get a competition team by ID
// @Description Get a competition team by ID
// @Tags team_center
// @Accept json
// @Produce json
// @Param id path int true "Competition Team ID"
// @Success 200 {object} team_center.CompetitionTeam
// @Router /team_center/competition_teams/{id} [get]
func GetCompetitionTeamByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var competitionTeam team_center.CompetitionTeam
	if err := database.DB.First(&competitionTeam, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Competition team not found"})
		return
	}
	c.JSON(http.StatusOK, competitionTeam)
}

// CreateCompetitionTeam godoc
// @Summary Create a new competition team
// @Description Create a new competition team
// @Tags team_center
// @Accept json
// @Produce json
// @Param competition_team body team_center.CompetitionTeam true "Competition Team"
// @Success 201 {object} team_center.CompetitionTeam
// @Router /team_center/competition_teams [post]
func CreateCompetitionTeam(c *gin.Context) {
	var competitionTeam team_center.CompetitionTeam
	if err := c.ShouldBindJSON(&competitionTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if competitionTeam.ContactWechat == "" && competitionTeam.ContactQq == "" && competitionTeam.ContactEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one contact method is required"})
		return
	}

	if err := database.DB.Create(&competitionTeam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, competitionTeam)
}
