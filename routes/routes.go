package routes

import (
	"log"

	controllers "employe/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(c *gin.Engine) {
	log.Println("Setting Up Routes")

	c.GET("/employe", controllers.GetDataEmploye)
	c.GET("/employe/:id", controllers.GetDataEmployeById)
	c.POST("/employe", controllers.AddEmploye)
	c.GET("/employe/delete/:id", controllers.DeleteEmploye)

}
