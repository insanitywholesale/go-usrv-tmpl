package main

import (
	"go-usrv-tmpl/domain"
	mockrepo "go-usrv-tmpl/repo/mock"
	"log"
)

func chooseRepo() domain.ThingRepo {
	repo, err := mockrepo.NewMockRepo()
	if err != nil {
		log.Fatal(err)
	}
	return repo
}

func main() {
	repo := chooseRepo()
	log.Println("main() still in primitive state but here's a repo", repo)
}
