package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rajanjaiswal/MyNewContactbook/internals/contact"
	"github.com/rajanjaiswal/MyNewContactbook/internals/user"
	"gorm.io/gorm"

	"github.com/rajanjaiswal/MyNewContactbook/pkg/db/postgres"
)

type server struct {
	C  *gin.Engine
	DB *gorm.DB
}

func GetServer() *server {

	return &server{
		C:  gin.Default(),
		DB: postgres.ConnectDatabase(),
	}

}

func (s *server) Run() {
	s.initRoutes()
	log.Fatal(s.C.Run())

}

func (s *server) initRoutes() {
	// routes or Endpoints is giving direction
	r := s.C
	// r.GET("/", controllers.Welcome)
	// r.GET("/ping", controllers.Ping)

	// CONTACTS ENDPOINTS
	// r.GET("/contacts", controllers.GetAllContacts)

	// r.POST("/contacts", controllers.CreateContact)

	// r.PUT("/contacts/:id", controllers.UpdateContacts)
	// r.GET("/contacts/:id", controllers.GetContactById)
	// r.DELETE("/contacts/:id", controllers.DeleteContactsById)
	contact.RegisterRoutes(r, contact.NewService(contact.NewRepository(*s.DB)))
	user.RegisterRoutes(r, user.NewService(user.NewRepository(*s.DB)))
}
