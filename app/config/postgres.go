package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Postgres struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     int    `envconfig:"POSTGRES_PORT" required:"true"`
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Dbname   string `envconfig:"POSTGRES_DATABASE" required:"true"`

	MaxConnectionLifetime time.Duration `envconfig:"DB_MAX_CONN_LIFE_TIME" default:"300s"`
	MaxOpenConnection     int           `envconfig:"DB_MAX_OPEN_CONNECTION" default:"100"`
	MaxIdleConnection     int           `envconfig:"DB_MAX_IDLE_CONNECTION" default:"10"`
}

func (p Postgres) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Dbname)
}

func OpenPostgresDatabaseConnection(p Postgres) (*gorm.DB, error) {
	dbConn, err := gorm.Open(postgres.Open(p.ConnectionString()))
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
