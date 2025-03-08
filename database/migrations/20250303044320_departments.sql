-- +goose Up
-- +goose StatementBegin
CREATE TABLE department (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    faculty_name TEXT NOT NULL,
    faculty_id INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (faculty_id) REFERENCES faculty(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS department;
-- +goose StatementEnd

