package manager

import (
	"fmt"
	"myGoSkeleton/config"
	"os"

	"github.com/jmoiron/sqlx"
)

type InfraManager interface {
	PostgreConn() *sqlx.DB
}

type infraManager struct {
	postgreConn *sqlx.DB
}

func (i *infraManager) PostgreConn() *sqlx.DB {
	return i.postgreConn
}

func NewInfraManager(configDatabase *config.ConfigDatabase) InfraManager {
	urlPostgresql := configDatabase.PostgreConn()
	conn, err := sqlx.Connect("pgx", urlPostgresql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &infraManager{
		postgreConn: conn,
	}
}
