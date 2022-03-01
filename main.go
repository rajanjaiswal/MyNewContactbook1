package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rajanjaiswal/MyNewContactbook/controllers"
	"github.com/rajanjaiswal/MyNewContactbook/models"
)

// routes or endpoints
func main() {
	fmt.Print("Starting ContactBook Api")
	//r := .Default()
	r := gin.Default()

	//database connection
	models.ConnectDatabase()

	initRoutes(r)

	r.Run() // starting server,
}

func initRoutes(r *gin.Engine) {
	// routes or Endpoints is giving direction
	r.GET("/", controllers.Welcome)
	r.GET("/ping", controllers.Ping)

	// CONTACTS ENDPOINTS
	r.GET("/contacts", controllers.GetAllContacts)

	r.POST("/contacts", controllers.CreateContact)
	r.DELETE("/contacts", controllers.DeleteContacts)
	r.PUT("/contacts", controllers.UpdateContacts)
	r.GET("/contacts/:id", controllers.GetContactById)

}
