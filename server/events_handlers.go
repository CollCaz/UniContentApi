package server

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

func (s *Server) registerEventsSubrouteOn(parentRoute *fuego.Server) *fuego.Server {
	eventsRoute := fuego.Group(parentRoute, "/events")
	fuego.Get(eventsRoute, "", s.GetEvents)
	fuego.Get(eventsRoute, "/search", s.SearchAllEvents)
	fuego.Post(eventsRoute, "", s.PostEvent)
	fuego.Put(eventsRoute, "", s.PutEvent)

	return eventsRoute
}

func (s *Server) GetEvents(c fuego.ContextNoBody) (d.Events, error) {
	return s.db.GetAllEvents(d.GetAllEventsArgs{Language: "ar"})
}

func (s *Server) SearchAllEvents(c fuego.ContextWithBody[d.SearchAllEventArgs]) (d.Events, error) {
	body, err := c.Body()
	if err != nil {
		return d.Events{}, err
	}

	return s.db.SearchAllEvents(body)
}

func (s *Server) PostEvent(c fuego.ContextWithBody[d.Event]) (d.Event, error) {
	body, err := c.Body()
	if err != nil {
		return d.Event{}, err
	}

	return s.db.InsertEvent(body)
}

func (s *Server) PutEvent(c fuego.ContextWithBody[d.UpdateEventArgs]) (d.Event, error) {
	body, err := c.Body()
	if err != nil {
		return d.Event{}, err
	}

	return s.db.UpdateEvent(body)
}
