package main

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	personController := controller.NewPersonController()
	ginEngine := gin.Default()

	ginEngine.GET("/person", personController.GetAll)
	ginEngine.GET("/person", personController.Create)

	err := ginEngine.Run("localhost:5000")
	if err != nil {
		panic(err)
	}
}
