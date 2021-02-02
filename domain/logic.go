package domain

type thingService struct {
	thingRepo ThingRepo
}

func NewThingService(thingRepo ThingRepo) ThingService {
	return &thingService{
		thingRepo,
	}
}

func (t *thingService) Find(field string) (*Thing, error) {
	return t.thingRepo.Find(field)
}

func (t *thingService) Store(thing *Thing) error {
	return t.thingRepo.Store(thing)
}
