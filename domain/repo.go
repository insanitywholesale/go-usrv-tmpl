package domain

type ThingRepo interface {
	Find(field string) (*Thing, error)
	Store(thing *Thing) error
	//Delete(thing *Thing) error
}
