package server

import (
	"fmt"
	"github.com/J4stEu/solib/internal/app/config"
	"github.com/J4stEu/solib/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// Server - server_errors structure
type Server struct {
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New - new server_errors instance
func New(config *config.Config, logger *logrus.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
		router: mux.NewRouter(),
	}
}

// Start - start server_errors instance
func (srv *Server) Start() error {
	if err := srv.ConfigureLogger(); err != nil {
		return err
	}
	srv.ConfigureRouter()
	if err := srv.ConfigureStore(); err != nil {
		return err
	}
	instance := fmt.Sprintf("%s:%s",
		srv.config.Server.ServerAddr,
		strconv.Itoa(int(srv.config.Server.ServerPort)),
	)
	srv.logger.Info(fmt.Sprintf("Starting application (%s)...", instance))
	return http.ListenAndServe(instance, srv.router)
}
