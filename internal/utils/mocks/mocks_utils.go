package mocks

import "iLikeToKnow.com/internal/model"

func GetEventMocked() (model.Event, bool) {
	stringUUID := "5c0e0e0a-5ff3-4f1a-9f4a-d03f4d8b39e2"
	desc := "Event example"

	return model.Event{
		ID:          stringUUID,
		Title:       "Go Meetup",
		Description: &desc,
	}, true
}
