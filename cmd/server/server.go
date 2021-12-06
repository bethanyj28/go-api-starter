package main

import (
	"net/http"
	"time"

	"github.com/bethanyj28/go-api-starter/internal/db"
	"github.com/sirupsen/logrus"
)

type server struct {
	httpServer *http.Server
	logger     *logrus.Logger
	store      db.Store
}

// ServerConfig contains connection values for the server
type ServerConfig struct {
	Addr    string
	Timeout time.Duration
}

func newServer(svrConfig ServerConfig, store db.Store, logger *logrus.Logger) *server {
	svr := &server{logger: logger, store: store}
	svr.buildHTTPServer(svrConfig.Addr, svrConfig.Timeout)

	return svr
}
