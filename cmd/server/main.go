package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bethanyj28/go-api-starter/internal/db/postgres"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Config represents environment variables
type Config struct {
	Server struct {
		Address string        `default:"0.0.0.0:8080"`
		Timeout time.Duration `default:"5s"`
	}
	Postgres struct {
		Username string `default:"postgres"`
		Password string `default:"postgres"`
		Host     string `default:"postgres"`
		Port     string `default:"5432"`
		DBName   string `default:"database"`
		UseSSL   bool   `default:"false"`
	}
}

func main() {
	logger := logrus.New()
	var c Config
	if err := envconfig.Process("api", &c); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to read envconfig").Error())
	}

	db, err := postgres.NewRepository(postgres.Config{
		Username:      c.Postgres.Username,
		Password:      c.Postgres.Password,
		Host:          c.Postgres.Host,
		Port:          c.Postgres.Port,
		DBName:        c.Postgres.DBName,
		SSL:           c.Postgres.UseSSL,
		RunMigrations: true,
	})
	if err != nil {
		logger.Fatal(fmt.Errorf("failed to create new postgres repository: %w", err).Error())
	}

	svr := newServer(ServerConfig{
		Addr:    c.Server.Address,
		Timeout: c.Server.Timeout,
	}, &db, logger)

	log.Fatal(svr.httpServer.ListenAndServe())
}
