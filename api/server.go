package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/joekingsleyMukundi/Gatekeeper/db/sqlc"
	"github.com/joekingsleyMukundi/Gatekeeper/tokens"
	"github.com/joekingsleyMukundi/Gatekeeper/utils"
	"github.com/joekingsleyMukundi/Gatekeeper/workers"
)

type Server struct {
	TokenMaker      tokens.Maker
	config          utils.Config
	store           db.Store
	Router          *gin.Engine
	taskDistributor workers.TaskDistributor
}

func NewSever(config utils.Config, store db.Store, taskDistributor workers.TaskDistributor) (*Server, error) {
	tokenMaker, err := tokens.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("ERROR: cannot create token: %s", err)
	}
	server := &Server{
		TokenMaker:      tokenMaker,
		config:          config,
		store:           store,
		taskDistributor: taskDistributor,
	}
	server.routerSetup()
	return server, nil
}

func (server *Server) routerSetup() {
	router := gin.Default()
	// TO DO : Create user suth apis
	router.POST("/users/register", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/auth/renew_access", server.renewAccessToken)

	server.Router = router
}
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
