package json

import (
	"encoding/json"
	"go-usrv-tmpl/domain"
)

type Thing struct{}

func (t *Thing) Decode(input []byte) (*domain.Thing, error) {
	thing := &domain.Thing{}
	err := json.Unmarshal(input, thing)
	if err != nil {
		return nil, err
	}
	return thing, nil
}

func (t *Thing) Encode(input *domain.Thing) ([]byte, error) {
	jsonThing, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	return jsonThing, nil
}
