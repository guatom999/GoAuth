package repositories

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB
}

func (r customerRepositoryDB) 