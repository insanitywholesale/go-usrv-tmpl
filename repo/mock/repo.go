package mock

import (
	"errors"
	"go-usrv-tmpl/domain"
)

type mockRepo []*domain.Thing

var thingList mockRepo = []*domain.Thing{}

func mkThingList() mockRepo {
	var thingList mockRepo = []*domain.Thing{
		&domain.Thing{
			Field:  "value",
			URL:    "http://here.local",
			Number: 6,
		},
	}
	return thingList
}

func NewMockRepo() (domain.ThingRepo, error) {
	thingList = mkThingList()
	repo := &mockRepo{}
	return repo, nil
}

func (r *mockRepo) Find(field string) (*domain.Thing, error) {
	for _, thing := range thingList {
		if thing.Field == field {
			return thing, nil
		}
	}
	return nil, errors.New("Thing with Field" + field + "not found")
}

func (r *mockRepo) Store(thing *domain.Thing) error {
	thingList = append(thingList, thing)
	return nil
}
