-- +goose Up
-- +goose StatementBegin
CREATE TABLE event (
    id INTEGER PRIMARY KEY,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    location TEXT NOT NULL,
    poster_id INTEGER NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (poster_id) REFERENCES image (id)
);


CREATE TABLE event_data (
    id INTEGER PRIMARY KEY,
    event_id INTEGER NOT NULL,
    language TEXT NOT NULL,
    name TEXT NOT NULL,
    content TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (event_id) REFERENCES event (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event;
DROP TABLE event_data;
-- +goose StatementEnd
