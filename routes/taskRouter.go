package routes

import (
	"github.com/gin-gonic/gin"
	"taskung.com/test/m/v2/controllers"
)

func Routes(router *gin.Engine) {
	// project routes
	router.GET("/project", controllers.GetProject)
	router.GET("/project/:id", controllers.GetProjectByID)

	// task routes
	router.GET("project/:id/task", controllers.GetTask)
	router.GET("/project/:project_id/task/:id", controllers.GetTaskByID)
	router.POST("project/:id/task", controllers.AddTask)
	router.PUT("project/:id/task", controllers.UpdateTask)

}
