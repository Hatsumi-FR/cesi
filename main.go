package main

import (
	"cesi/adapters/sqlite"
	"cesi/server"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	repo, err := sqlite.NewRepository()
	if err != nil {
		panic(err)
	}
	err = repo.Init()
	if err != nil {
		log.Fatal(err)
	}
	serv := server.NewServer(repo)
	serv.Start()
}
