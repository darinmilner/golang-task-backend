package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Db struct {
	*GoalsRepo
}

func NewDB(dataSourceName string) (*Db, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Could not open the DB %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Could not open the DB %w", err)
	}
	return &Db{
		GoalsRepo: &GoalsRepo{DB: db},
	}, nil
}
