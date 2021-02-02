package domain

type ThingSerializer interface {
	Decode(input []byte) (*Thing, error)
	Encode(input *Thing) ([]byte, error)
}
