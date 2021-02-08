package grpc

import (
	"context"
	"fmt"
	"go-usrv-tmpl/domain"
	"go-usrv-tmpl/protos"
)

func (tr *ThingRequest) ProcessThing(ctx context.Context, ti *protos.ThingIn) (*protos.ThingOut, error) {
	fmt.Print("ProcessThing not implemented")
	return nil, nil
}
