package taskControllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"taskung.com/test/m/v2/inits"
	"taskung.com/test/m/v2/models"
)

func AddTask(c *gin.Context) {
	var newTask models.TaskModel

	// Manually Bind Json the JSON data into the newTask struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if task_name is empty
	if newTask.TaskName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task_name is required"})
		return
	}

	projectID := c.Param("project_id") // Get the project ID from the URL parameter

	var project models.Project

	// Check if the project with the given ID exists
	if err := inits.DB.Where("id = ?", projectID).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found!"})
		return
	}

	// Set the ProjectID for the new task
	newTask.ProjectTaskID = project.ID

	// Create a new task record in the database
	if err := inits.DB.Create(&newTask).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "created new task!"})
	}
}

// func AddTask(c *gin.Context) {
// 	var project models.Project
// 	var newTask models.TaskModel

// 	// Manually Bind Json the JSON data into the newTask struct
// 	if err := c.ShouldBindJSON(&newTask); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Check if task_name is empty
// 	if newTask.TaskName == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "task_name is required"})
// 		return
// 	}

// 	// Create a new task record in the database
// 	if err := inits.DB.Create(&newTask).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, "created new task!")
// 	}
// 	inits.DB.Where("id = ?", c.Param("id")).First(&project).Model(&project).Association("Task").Append(&newTask)
// }
