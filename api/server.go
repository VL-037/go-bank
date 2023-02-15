package api

import (
	db "github.com/VL-037/go-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for the banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer for routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

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
