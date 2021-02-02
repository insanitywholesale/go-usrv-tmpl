package xml

import (
	"encoding/xml"
	"go-usrv-tmpl/domain"
)

type Thing struct{}

func (t *Thing) Decode(input []byte) (*domain.Thing, error) {
	thing := &domain.Thing{}
	err := xml.Unmarshal(input, thing)
	if err != nil {
		return nil, err
	}
	return thing, nil
}

func (t *Thing) Encode(input *domain.Thing) ([]byte, error) {
	xmlThing, err := xml.Marshal(input)
	if err != nil {
		return nil, err
	}
	return xmlThing, nil
}
