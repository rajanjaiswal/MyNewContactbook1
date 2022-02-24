package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rajanjaiswal/MyNewContactbook/controller"
)

// routes or endpoints
func main() {
	fmt.Print("Starting ContactBook Api")
	//r := .Default()
	r := gin.Default()
	Routes(r)

	r.Run() // starting server,
}

func Routes(r *gin.Engine) {
	r.GET("/", controller.Welcome)
	r.GET("/ping", controller.Ping)

	// CONTACTS ENDPOINTS
	r.GET("/contacts", controllers.GetAllContacts)

	r.POST("/contacts", controllers.GetContacts)
	r.DELETE("/contacts", controllers.DelectContacts)
	r.PUT("/contacts", controllers.UpdateContacts)

}
