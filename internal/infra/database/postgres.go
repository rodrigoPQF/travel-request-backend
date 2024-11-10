package database

import (
	"fmt"
	"log"

	"github.com/rodrigoPQF/travel-request-backend/internal/config"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/models"
	"github.com/rodrigoPQF/travel-request-backend/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type RepositoryTransaction func(interface{}) error


type PostgresDB struct {
	Session *gorm.DB
}

func NewPostgresDb()(*PostgresDB, error) {
	envs := config.GetEnvs()

	connectionUrl := utils.GenerateConnectionURL(utils.ConnectionURL{
		Username: envs.POSTGRES_USER,
		Password: envs.POSTGRES_PASSWORD,
		Host: envs.POSTGRES_HOST,
		DatabaseName: envs.POSTGRES_DB,
	})


	session, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{Logger: gormLogger.Default})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %W", err)
	}

	return &PostgresDB{Session: session}, nil
}

func (s *PostgresDB) Close() error {
	log.Println("close database connection")

	db, err := s.Session.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (s *PostgresDB) Tx(exec RepositoryTransaction) error {
	log.Println("initializing transaction database")

	return s.Session.Transaction(func(tx *gorm.DB) error {
		return exec(&PostgresDB{
			Session: tx,
		})})
}

func (s *PostgresDB) Migrations() error {
	s.Session.Exec(`
    DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
            CREATE TYPE status AS ENUM ('REQUESTED', 'APPROVED', 'CANCELED');
        END IF;
    END $$;
`)

if err := s.Session.AutoMigrate(&models.TravelRequest{}); err != nil {
	return err
}
	log.Println("initializing migrations database")

	return nil
}