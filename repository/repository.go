package repository

import (
	"cesi/domain"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
)

type Repositorer interface {
	CreatePlayer(p domain.Player) (int, error)
	Init() error
}

type Repository struct {
	Repositorer
	db *sql.DB
}

func NewRepository() (*Repository, error) {
	db, err := connectDatabase()
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func connectDatabase() (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlite3", "./cesi.db")
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func (r *Repository) Init() error {
	sqlFile, err := ioutil.ReadFile("./repository/schema.sql")
	if err != nil {
		return err
	}

	_, err = r.db.Exec(string(sqlFile))
	if err != nil {
		return err
	}

	return nil
}
