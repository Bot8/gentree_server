package database

import (
	"artarn/gentree/infrastructure"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/jackc/pgx/stdlib"
)

func GetConnection() *sqlx.DB {
	databaseConfig := infrastructure.GetConfig().Database
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Name)

	connection, err := sqlx.Connect("pgx", dsn)

	if err != nil {
		err = errors.New("cannot establish connection")
		log.Fatal(err)
	}

	err = connection.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return connection
}
