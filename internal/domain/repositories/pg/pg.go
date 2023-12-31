package pg

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gophermart/internal/config"
)

type Pg struct {
	db          *sql.DB
	mDialect    string
	mDir        string
	usersTable  string
	ordersTable string
}

func New() (*Pg, error) {

	var err error
	pg := &Pg{
		mDialect:    "postgres",
		mDir:        "migrations",
		usersTable:  "users",
		ordersTable: "orders",
	}

	pg.db, err = initDB()

	return pg, err
}

func initDB() (*sql.DB, error) {
	conn, err := sql.Open("pgx", config.Instance.DatabaseDsn())

	if err != nil {
		return nil, err
	}

	return conn, nil
}
