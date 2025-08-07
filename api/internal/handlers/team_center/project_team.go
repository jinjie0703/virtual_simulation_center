package team_center

import (
	"net/http"
	"strconv"
	"virtual_simulation_center/api/internal/database"
	"virtual_simulation_center/api/internal/models/team_center"

	"github.com/gin-gonic/gin"
)

// GetAllProjectTeams godoc
// @Summary Get all project teams
// @Description Get all project teams
// @Tags team_center
// @Accept json
// @Produce json
// @Success 200 {array} team_center.ProjectTeam
// @Router /team_center/project_teams [get]
func GetAllProjectTeams(c *gin.Context) {
	var projectTeams []team_center.ProjectTeam
	if err := database.DB.Find(&projectTeams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projectTeams)
}

// GetProjectTeamByID godoc
// @Summary Get a project team by ID
// @Description Get a project team by ID
// @Tags team_center
// @Accept json
// @Produce json
// @Param id path int true "Project Team ID"
// @Success 200 {object} team_center.ProjectTeam
// @Router /team_center/project_teams/{id} [get]
func GetProjectTeamByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var projectTeam team_center.ProjectTeam
	if err := database.DB.First(&projectTeam, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project team not found"})
		return
	}
	c.JSON(http.StatusOK, projectTeam)
}

// CreateProjectTeam godoc
// @Summary Create a new project team
// @Description Create a new project team
// @Tags team_center
// @Accept json
// @Produce json
// @Param project_team body team_center.ProjectTeam true "Project Team"
// @Success 201 {object} team_center.ProjectTeam
// @Router /team_center/project_teams [post]
func CreateProjectTeam(c *gin.Context) {
	var projectTeam team_center.ProjectTeam
	if err := c.ShouldBindJSON(&projectTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if projectTeam.ContactWechat == "" && projectTeam.ContactQq == "" && projectTeam.ContactEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one contact method is required"})
		return
	}

	if err := database.DB.Create(&projectTeam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, projectTeam)
}
