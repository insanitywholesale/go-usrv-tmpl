package grpc

import (
	"go-usrv-tmpl/domain"
	"log"
)

func hey() {
	smth := domain.NewThingService(nil)
	log.Println("thingservice", smth)
}
