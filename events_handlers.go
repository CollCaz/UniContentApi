package main

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

func (a *app) SearchAllEvents(c fuego.ContextWithBody[d.SearchAllEventArgs]) (d.Events, error) {
	body, err := c.Body()
	if err != nil {
		return d.Events{}, err
	}

	return a.db.SearchAllEvents(body)
}

func (a *app) PostEvent(c fuego.ContextWithBody[d.Event]) (d.Event, error) {
	body, err := c.Body()
	if err != nil {
		return d.Event{}, err
	}

	return a.db.InsertEvent(body)
}
