package gapi

import (
	"fmt"

	db "github.com/KYLS/simplebank/db/sqlc"
	"github.com/KYLS/simplebank/pb"
	"github.com/KYLS/simplebank/token"
	"github.com/KYLS/simplebank/util"
	"github.com/gin-gonic/gin"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new gRPC server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
