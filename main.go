package main

import (
	"github.com/gin-gonic/gin"
	"taskung.com/test/m/v2/inits"
	"taskung.com/test/m/v2/routes"
)

func init() {
	inits.LoadEnvVariables()
	inits.ConDB()
}

func main() {
	//call gin function another example => gin.New()
	r := gin.Default()

	routes.Routes(r)

	//*gin.Context is default
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080
}
