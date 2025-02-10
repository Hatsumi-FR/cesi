package main

import (
	"cesi/repository"
	"cesi/server"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	repo, err := repository.NewRepository()
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
