package database

import (
	"fmt"
	"time"

	"github.com/CollCaz/UniSite/database/gen/model"
	t "github.com/CollCaz/UniSite/database/gen/table"
	s "github.com/go-jet/jet/v2/sqlite"
)

type Event struct {
	StarDate  time.Time
	EndDate   time.Time
	Location  string
	PosterUrl string `validate:"url"`
	Name      string
	Content   string
	Language  string `validate:"required,bcp47_language_tag"`
}

type Events []Event

func (d *DataService) scanEvent(stmt s.Statement) (Event, error) {
	dest := model.Event{}
	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		d.logger.Info(stmt.DebugSql())
		return Event{}, err
	}

	event := Event{
		StarDate:  *dest.StartDate,
		EndDate:   *dest.EndDate,
		Location:  dest.Location,
		PosterUrl: dest.PosterURL,
	}

	return event, nil
}

func (d *DataService) scanEvents(eventStmt s.Statement, dataStmt s.Statement) (Events, error) {
	dest := []model.Event{}
	err := eventStmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		d.logger.Info(eventStmt.DebugSql())
		return Events{}, err
	}

	dest2 := []model.EventData{}
	err = dataStmt.Query(d.db, &dest2)
	if err != nil {
		d.logger.Error(err.Error())
		d.logger.Info(eventStmt.DebugSql())
		return Events{}, err
	}

	events := Events{}
	for _, event := range dest {
		events = append(events, Event{
			StarDate:  *event.StartDate,
			EndDate:   *event.EndDate,
			Location:  event.Location,
			PosterUrl: event.PosterURL,
		})

	}
	return events, nil
}

type GetAllEventsArgs struct {
	Language string `validate:"required,bcp47_language_tag"`
}

func (d *DataService) GetAllEvents(args GetAllEventsArgs) (Events, error) {
	stmt := s.
		SELECT(
			t.Event.StartDate,
			t.Event.EndDate,
			t.Event.PosterURL,
			t.EventData.Content,
			t.EventData.Language,
		).FROM(
		t.Event.INNER_JOIN(t.EventData, t.Event.ID.EQ(t.EventData.EventID)),
	).WHERE(t.EventData.Language.EQ(s.String(args.Language)))

	return d.scanEvents(stmt)
}

func (d *DataService) InsertEvent(e Event) (Event, error) {
	insertEventStmt := t.Event.
		INSERT(
			t.Event.StartDate,
			t.Event.EndDate,
			t.Event.PosterURL,
			t.Event.Location,
		).VALUES(
		e.StarDate,
		e.EndDate,
		e.PosterUrl,
		e.Location,
	).RETURNING(
		t.Event.StartDate,
		t.Event.EndDate,
		t.Event.PosterURL,
	)

	event, err := d.scanEvent(insertEventStmt)
	if err != nil {
		return Event{}, err
	}

	insertEventDataStmt := t.EventData.
		INSERT(
			t.EventData.Name,
			t.EventData.Content,
			t.EventData.Language,
		).VALUES(
		e.Name,
		e.Content,
		e.Language,
	).RETURNING(
		t.EventData.Name,
		t.EventData.Content,
		t.EventData.Language,
	)

	eventData, err := d.scanEvent(insertEventDataStmt)

	mergedEvent := Event{
		Location:  event.Location,
		PosterUrl: event.PosterUrl,
		StarDate:  event.StarDate,
		EndDate:   event.EndDate,
		Name:      eventData.Name,
		Content:   eventData.Content,
		Language:  eventData.Content,
	}

	return mergedEvent, nil
}

type SearchAllEventArgs struct {
	Query    string
	Limit    int    `validate:"required,gte=1"`
	Page     int    `validate:"required,gte=0"`
	Language string `validate:"required,bcp47_language_tag"`
}

func (d *DataService) SearchAllEvents(args SearchAllEventArgs) (Events, error) {
	stmt := s.
		SELECT(
			t.Event.StartDate,
			t.Event.EndDate,
			t.Event.PosterURL,
			t.Event.Location,
			t.EventData.Name,
			t.EventData.Content,
			t.EventData.Language,
		).FROM(
		t.Event.INNER_JOIN(t.EventData, t.Event.ID.EQ(t.EventData.EventID)),
	).WHERE(
		s.AND(
			t.EventData.Language.EQ(s.String(args.Language)),
			s.OR(
				t.EventData.Name.LIKE(s.String(fmt.Sprintf("%%%s%%", args.Query))),
				t.Event.Location.LIKE(s.String(fmt.Sprintf("%%%s%%", args.Query))),
			),
		),
	).LIMIT(
		int64(args.Limit),
	).OFFSET(
		int64(args.Page),
	).ORDER_BY(
		s.CASE().
			WHEN(
				t.EventData.Name.EQ(s.String(args.Query)),
			).THEN(s.Int64(1)).
			WHEN(
				t.EventData.Name.LIKE(s.String(fmt.Sprintf("%s%%", args.Query))),
			).THEN(s.Int64(2)).
			WHEN(
				t.EventData.Name.LIKE(s.String(fmt.Sprintf("%%%s%%", args.Query))),
			).THEN(s.Int64(3)).
			WHEN(
				t.Event.Location.EQ(s.String(args.Query)),
			).THEN(s.Int64(1)).
			WHEN(
				t.Event.Location.LIKE(s.String(fmt.Sprintf("%s%%", args.Query))),
			).THEN(s.Int64(2)).
			WHEN(
				t.Event.Location.LIKE(s.String(fmt.Sprintf("%%%s%%", args.Query))),
			).THEN(s.Int64(3)),
		t.EventData.Name,
		t.Event.Location,
		t.Event.ID,
	)

	return d.scanEvents(stmt)
}
