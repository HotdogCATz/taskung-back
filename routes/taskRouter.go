package routes

import (
	"github.com/gin-gonic/gin"
	"taskung.com/test/m/v2/controllers"
)

func Routes(router *gin.Engine) {
	// login - register routes
	router.POST("/register", controllers.Register)
	// project routes
	router.GET("/project", controllers.GetProject)
	router.GET("/project/:project_id", controllers.GetProjectByID)
	router.POST("/project", controllers.AddProject)
	router.PUT("/project", controllers.UpdateProject)
	router.PUT("/project/:project_id", controllers.UpdateProjectByID)
	router.DELETE("/project/:project_id", controllers.DeleteProjectByID)

	// task routes
	router.GET("/project/:project_id/task", controllers.GetTask)
	router.GET("/project/:project_id/task/:id", controllers.GetTaskByID)
	router.POST("project/:project_id/task", controllers.AddTask)
	router.PUT("project/:project_id/task", controllers.UpdateTask)
	router.PUT("/project/:project_id/task/:id", controllers.UpdateTaskByID)
	router.DELETE("/project/:project_id/task/:id", controllers.DeleteTaskByID)

}
