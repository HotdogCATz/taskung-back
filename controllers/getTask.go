package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"taskung.com/test/m/v2/inits"
	"taskung.com/test/m/v2/models"
)

// "fmt"

//	func GetTask(c *gin.Context) {
//		// []models call array (all elements)
//		// var project models.Project
//		var modelTask []models.TaskModel
//		inits.DB.Find(&modelTask)
//		c.JSON(http.StatusOK, gin.H{"task_data": modelTask})
//	}
func GetTask(c *gin.Context) {
	// Get the project ID from the URL parameter
	projectID := c.Param("project_id")

	var tasks []models.TaskModel

	// Find all tasks associated with the project
	if err := inits.DB.Where("project_id = ?", projectID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tasks not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task_data": tasks})
}

// func GetTaskByID(c *gin.Context) {
// 	// get the posts
// 	// models call one element
// 	var modelTask models.TaskModel

// if err := inits.DB.Where("id = ?", c.Param("id")).First(&modelTask).Error; err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 	return
// }

//		c.JSON(http.StatusOK, gin.H{"task_data": modelTask})
//	}
func GetTaskByID(c *gin.Context) {
	// Get the task ID from the URL parameter
	taskID := c.Param("id")
	projectID := c.Param("project_id")

	var task models.TaskModel

	// Find the task by ID
	if err := inits.DB.Where("project_id = ?", projectID).Where("id = ?", taskID).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task_data": task})
}
