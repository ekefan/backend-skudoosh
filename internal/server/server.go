package server

import (
	"fmt"

	"time"

	db "github.com/ekefan/backend-skudoosh/internal/db/sqlc"
	"github.com/ekefan/backend-skudoosh/internal/token"
	"github.com/ekefan/backend-skudoosh/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	config     utils.Config
	tokenMaker token.Maker
}

// NewServer creates a new http server, sets up api routes
// returns the server instance, or an error on error
func NewServer(store db.Store, config utils.Config) (*Server, error) {
	makeToken, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: makeToken,
		config:     config,
	}
	server.setUpRouter()
	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Change to your frontend's origin
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/")
	authRoutes.Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/users/itineraries", server.createTrip)

	// authRoutes.DELETE("/users/itinerary/delete", server.deleteTrip)

	server.router = router
}

// Start: runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
