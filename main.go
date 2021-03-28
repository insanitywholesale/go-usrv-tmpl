package main

import (
	"go-usrv-tmpl/domain"
	mockrepo "go-usrv-tmpl/repo/mock"
	"log"
	"os"
)

func chooseRepo() (domain.ThingRepo, error) {
	repo, err := mockrepo.NewMockRepo()
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func main() {
	repo, err := chooseRepo()
	if err != nil {
		log.Fatalln("error with repository:", err)
	}
	log.Println("repo", repo, "selected")
	port := os.Getenv("PORT")
	if port == "" {
		port = "53011"
	}
	log.Println("port", port, "selected")
}
