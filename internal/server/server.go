package server

import (

	db "github.com/ekefan/backend-skudoosh/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)


type Server struct {
	store db.Store
	router *gin.Engine
	//config utils.Config
	//tokenMaker token.Maker
}

// NewServer creates a new http server, sets up api routes
// returns the server instance, or an error on error
func NewServer(store db.Store) (*Server, error) {
	server := &Server{
		store: store,
	}

	server.setUpRouter()
	return server, nil
}

func (server *Server) setUpRouter(){
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)


	server.router = router
}


// Start: runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
