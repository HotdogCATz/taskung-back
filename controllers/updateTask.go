package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"taskung.com/test/m/v2/inits"
	"taskung.com/test/m/v2/models"
)

func UpdateTask(c *gin.Context) {
	var modelTask models.TaskModel
	var updateTask models.TaskModel

	// Manually Bind Json the JSON data into the newTask struct
	if err := c.ShouldBindJSON(&modelTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create a new task record in the database
	if err := inits.DB.First(&updateTask, modelTask.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	} else {
		c.JSON(http.StatusOK, "update task successfully!")
	}

	inits.DB.First(&updateTask, modelTask.ID)

	if modelTask.TaskName != "" {
		updateTask.TaskName = modelTask.TaskName
	}
	if modelTask.Status != "" {
		updateTask.Status = modelTask.Status
	}
	if modelTask.Description != "" {
		updateTask.Description = modelTask.Description
	}
	inits.DB.Save(updateTask)
}
