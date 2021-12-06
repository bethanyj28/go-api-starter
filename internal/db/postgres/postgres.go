package postgres

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // file source
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
)

// Repository is a postgres implementation of repository
type Repository struct {
	conn *pgxpool.Pool
}

// Config is a configuration object for postgres repos
type Config struct {
	Username      string
	Password      string
	Host          string
	Port          string
	DBName        string
	SSL           bool
	RunMigrations bool
}

func buildConnStr(config Config) string {
	sslString := "require"
	if !config.SSL {
		sslString = "disable"
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.Username, config.Password, config.Host, config.Port, config.DBName, sslString)
}

// NewRepository creates a new postgres repository
func NewRepository(config Config) (Repository, error) {
	repo := Repository{}
	fmt.Println(buildConnStr(config))
	if config.RunMigrations {
		if err := runMigrations(config); err != nil {
			return repo, fmt.Errorf("failed to run migrations: %w", err)
		}
	}
	pool, err := pgxpool.Connect(context.Background(), buildConnStr(config))
	if err != nil {
		return repo, fmt.Errorf("failed to connect to database: %w", err)
	}

	repo.conn = pool

	return repo, nil
}

func runMigrations(config Config) error {
	c, err := pgx.ParseConfig(buildConnStr(config))
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	db := stdlib.OpenDB(*c)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to use connection as postgres instance: %w", err)
	}

	file := "file://migrations"
	m, err := migrate.NewWithDatabaseInstance(file, config.DBName, driver)
	if err != nil {
		return fmt.Errorf("failed to create new migration: %w", err)
	}

	return m.Up()
}
