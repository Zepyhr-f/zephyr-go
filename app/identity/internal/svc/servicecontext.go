package svc

import (
	"log"

	"zephyr-go/app/identity/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(postgres.Open(c.Postgres.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Init PostgreSQL failed: %v", err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
