package domain

type ThingService interface {
	Find(field string) (*Thing, error)
	Store(thing *Thing) error
}
