-- +goose Up
-- +goose StatementBegin
CREATE TABLE faculty (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    abbreviation TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO faculty (name, abbreviation) VALUES ("Science", "Sci");
INSERT INTO faculty (name, abbreviation) VALUES ("Engineering", "Eng");

CREATE TRIGGER update_faculty_timestamp
AFTER UPDATE ON faculty
FOR EACH ROW
BEGIN
    UPDATE faculty
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS faculty;
-- +goose StatementEnd

