package controllers

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"taskung.com/test/m/v2/inits"
	"taskung.com/test/m/v2/models"
)

func DeleteTaskByID(c *gin.Context) {
	// get the posts
	// models call one element
	var modelTask models.TaskModel
	taskId := c.Param("id")
	if err := inits.DB.Where("id = ?", c.Param("id")).First(&modelTask).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maybe the task has already been deleted!"})
		return
	}
	inits.DB.Delete(&modelTask, taskId)
	c.JSON(http.StatusOK, "deleted: task id "+taskId)
}
