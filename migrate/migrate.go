package main

import (
	"taskung.com/test/m/v2/inits"
	"taskung.com/test/m/v2/models"
)

func init() {
	// inits.LoadEnvVariables()
	inits.ConDB()
}

func main() {
	inits.DB.AutoMigrate(&models.TaskModel{})
	inits.DB.AutoMigrate(&models.Project{})
}
