package controllers

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"taskung.com/test/m/v2/inits"
	"taskung.com/test/m/v2/models"
)

func DeleteProjectByID(c *gin.Context) {
	projectID := c.Param("project_id")

	// Start a database transaction
	tx := inits.DB.Begin()

	var project models.Project

	// Find the project by ID
	if err := tx.Where("id = ?", projectID).First(&project).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found!"})
		return
	}

	// Delete the project and its associated tasks
	if err := tx.Delete(&project).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}

	// Delete associated tasks
	if err := tx.Where("project_id = ?", projectID).Delete(&models.TaskModel{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tasks"})
		return
	}

	// Commit the transaction
	tx.Commit()

	c.JSON(http.StatusOK, "Deleted project with ID: "+projectID)
}
