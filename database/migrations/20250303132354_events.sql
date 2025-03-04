-- +goose Up
-- +goose StatementBegin
CREATE TABLE event (
    id INTEGER PRIMARY KEY,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    location TEXT NOT NULL,
    poster_url TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE event_data (
    id INTEGER PRIMARY KEY,
    event_id INTEGER NOT NULL,
    language TEXT NOT NULL,
    name TEXT NOT NULL,
    content TEXT NOT NULL,

    FOREIGN KEY (event_id) REFERENCES event(id)
);

CREATE TRIGGER update_event_timestamp
AFTER UPDATE ON event
FOR EACH ROW
BEGIN
    UPDATE department
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event;
DROP TABLE event_data;
-- +goose StatementEnd
