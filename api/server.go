package api

import (
	"fmt"

	db "github.com/atercode/SimplySacco/db/sqlc"
	"github.com/atercode/SimplySacco/token"
	"github.com/atercode/SimplySacco/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Cannot create a token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/login", server.loginMember)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/members", server.listMembers)
	authRoutes.POST("/members", server.createMember)
	authRoutes.GET("/members/:id", server.getMember)

	server.router = router
}

//Http server start and graceful shutdown logic
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
