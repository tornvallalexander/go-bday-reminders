package api

import (
	"github.com/gin-gonic/gin"
	db "go-bday-reminders/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new server instance
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/birthdays", server.createBirthday)
	router.GET("/birthdays/:id", server.getBirthday)
	router.GET("/birthdays", server.listBirthday)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
