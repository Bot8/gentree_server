package database

import (
	"artarn/gentree/infrastructure"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

func GetConnection() (*sqlx.DB, error) {
	databaseConfig := infrastructure.GetConfig().Database
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Name)

	fmt.Println(dsn)

	connection, err := sqlx.Connect("pgx", dsn)

	if err != nil {
		err = errors.New("Cannot add unavailable items to order")
	}

	return connection, err
}
