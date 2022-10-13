package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/trungvdn/simplebank/db/sqlc"
)

type Server struct {
	db     db.Store
	router *gin.Engine
}

func NewServer(db db.Store) *Server {
	server := &Server{
		db: db,
	}
	router := gin.Default()
	server.router = router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
