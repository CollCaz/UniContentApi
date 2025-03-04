-- +goose Up
-- +goose StatementBegin
CREATE TABLE event (
    id INTEGER PRIMARY KEY,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    location TEXT NOT NULL,
    poster_url TEXT NOT NULL,

    content_ar TEXT NOT NULL,
    content_en TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
SELECT 'down SQL query';
-- +goose StatementEnd
