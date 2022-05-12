package api

import (
	db "github.com/atercode/SimplySacco/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/members", server.listMembers)
	router.POST("/members", server.createMember)
	router.GET("/members/:id", server.getMember)

	server.router = router

	return server
}

//Http server start and graceful shutdown logic
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
