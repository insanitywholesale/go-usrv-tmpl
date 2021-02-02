package msgpack

import (
	"github.com/vmihailenco/msgpack"
	"go-usrv-tmpl/domain"
)

type Thing struct{}

func (t *Thing) Decode(input []byte) (*domain.Thing, error) {
	thing := &domain.Thing{}
	err := msgpack.Unmarshal(input, thing)
	if err != nil {
		return nil, err
	}
	return thing, nil
}

func (t *Thing) Encode(input *domain.Thing) ([]byte, error) {
	msgpackThing, err := msgpack.Marshal(input)
	if err != nil {
		return nil, err
	}
	return msgpackThing, nil
}
