package database

import (
	"fmt"
	"time"

	"github.com/CollCaz/UniSite/database/gen/model"
	t "github.com/CollCaz/UniSite/database/gen/table"
	s "github.com/go-jet/jet/v2/sqlite"
)

type EventData struct {
	Id       int32 `json:",omitempty"`
	Name     string
	Content  string
	Language string `validate:"required,bcp47_language_tag"`
}

type Event struct {
	Id        int32 `json:",omitempty"`
	StarDate  time.Time
	EndDate   time.Time
	Location  string
	PosterUrl string `validate:"url"`
	EventData EventData
}

type Events []Event

type joinedEventModel struct {
	model.Event
	model.EventData
}

func (d *DataService) scanEvent(stmt s.Statement) (Event, error) {
	dest := joinedEventModel{}
	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		d.logger.Info(stmt.DebugSql())
		return Event{}, err
	}

	event := Event{
		Id:        *dest.Event.ID,
		StarDate:  *dest.StartDate,
		EndDate:   *dest.EndDate,
		Location:  dest.Location,
		PosterUrl: dest.PosterURL,
		EventData: EventData{
			Id:       *dest.EventData.ID,
			Name:     dest.Name,
			Content:  dest.Content,
			Language: dest.Language,
		},
	}

	return event, nil
}

func (d *DataService) scanEvents(stmt s.Statement) (Events, error) {
	dest := []joinedEventModel{}
	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		d.logger.Info(stmt.DebugSql())
		return Events{}, err
	}

	events := Events{}
	for _, event := range dest {
		events = append(events, Event{
			StarDate:  *event.StartDate,
			EndDate:   *event.EndDate,
			Location:  event.Location,
			PosterUrl: event.PosterURL,
			EventData: EventData{
				Name:     event.Name,
				Content:  event.Content,
				Language: event.Language,
			},
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
			t.Event.ID,
			t.Event.StartDate,
			t.Event.EndDate,
			t.Event.PosterURL,
			t.EventData.ID,
			t.EventData.Name,
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
		t.Event.ID,
		t.Event.StartDate,
		t.Event.EndDate,
		t.Event.PosterURL,
	)

	event, err := d.scanEvent(insertEventStmt)
	if err != nil {
		return Event{}, err
	}

	insertDataStmt := t.EventData.
		INSERT(
			t.EventData.Name,
			t.EventData.Content,
			t.EventData.Language,
			t.EventData.EventID,
		).VALUES(
		e.EventData.Name,
		e.EventData.Content,
		e.EventData.Language,
		event.Id,
	).RETURNING(
		t.EventData.Name,
		t.EventData.Content,
		t.EventData.Language,
		t.EventData.EventID,
	)

	data, err := d.scanEvent(insertDataStmt)
	if err != nil {
		return Event{}, err
	}

	event.EventData = data.EventData
	return event, nil
}

type UpdateEventArgs struct {
	Id  int32
	New Event
}

func (d *DataService) UpdateEvent(args UpdateEventArgs) (Event, error) {
	insertEventStmt := t.Event.
		UPDATE(
			t.Event.StartDate,
			t.Event.EndDate,
			t.Event.PosterURL,
			t.Event.Location,
		).
		SET(
			args.New.StarDate,
			args.New.EndDate,
			args.New.PosterUrl,
			args.New.Location,
		).
		WHERE(t.Event.ID.EQ(s.Int32(args.Id))).
		RETURNING(
			t.Event.ID,
			t.Event.StartDate,
			t.Event.EndDate,
			t.Event.PosterURL,
		)

	event, err := d.scanEvent(insertEventStmt)
	if err != nil {
		return Event{}, err
	}

	insertDataStmt := t.EventData.
		UPDATE(
			t.EventData.Name,
			t.EventData.Content,
			t.EventData.Language,
			t.EventData.EventID,
		).
		SET(
			args.New.EventData.Name,
			args.New.EventData.Content,
			args.New.EventData.Language,
			event.Id,
		).
		WHERE(t.EventData.ID.EQ(s.Int32(args.Id))).
		RETURNING(
			t.EventData.Name,
			t.EventData.Content,
			t.EventData.Language,
			t.EventData.EventID,
		)

	data, err := d.scanEvent(insertDataStmt)
	if err != nil {
		return Event{}, err
	}

	event.EventData = data.EventData
	return event, nil
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
		).
		FROM(
			t.Event.INNER_JOIN(t.EventData, t.Event.ID.EQ(t.EventData.EventID)),
		).
		WHERE(
			s.AND(
				t.EventData.Language.EQ(s.String(args.Language)),
				s.OR(
					t.EventData.Name.LIKE(s.String(fmt.Sprintf("%%%s%%", args.Query))),
					t.Event.Location.LIKE(s.String(fmt.Sprintf("%%%s%%", args.Query))),
				),
			),
		).
		LIMIT(
			int64(args.Limit),
		).
		OFFSET(
			int64(args.Page),
		).
		ORDER_BY(
			s.
				CASE().
				WHEN(
					t.EventData.Name.EQ(s.String(args.Query)),
				).
				THEN(s.Int64(1)).
				WHEN(
					t.EventData.Name.LIKE(s.String(fmt.Sprintf("%s%%", args.Query))),
				).
				THEN(s.Int64(2)).
				WHEN(
					t.EventData.Name.LIKE(s.String(fmt.Sprintf("%%%s%%", args.Query))),
				).
				THEN(s.Int64(3)).
				WHEN(
					t.Event.Location.EQ(s.String(args.Query)),
				).
				THEN(s.Int64(1)).
				WHEN(
					t.Event.Location.LIKE(s.String(fmt.Sprintf("%s%%", args.Query))),
				).
				THEN(s.Int64(2)).
				WHEN(
					t.Event.Location.LIKE(s.String(fmt.Sprintf("%%%s%%", args.Query))),
				).
				THEN(s.Int64(3)),
			t.EventData.Name,
			t.Event.Location,
			t.Event.ID,
		)

	return d.scanEvents(stmt)
}
