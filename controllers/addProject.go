package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"taskung.com/test/m/v2/inits"
	"taskung.com/test/m/v2/models"
)

func AddProject(c *gin.Context) {
	var newProject models.Project

	// Manually Bind Json the JSON data into the newProject struct
	if err := c.ShouldBindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if task_name is empty
	if newProject.ProjectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "project_name is required"})
		return
	}

	// Create a new task record in the database
	if err := inits.DB.Create(&newProject).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Project"})
		return
	} else {
		c.JSON(http.StatusOK, "created new Project!")
	}

	inits.DB.Create(&newProject)

}
