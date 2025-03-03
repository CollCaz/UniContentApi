-- +goose Up
-- +goose StatementBegin
CREATE TABLE about_section (
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_about_section_timestamp
AFTER UPDATE ON about_section
FOR EACH ROW
BEGIN
    UPDATE about_section
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

CREATE TRIGGER limit_about_section_entries
AFTER INSERT ON about_section
BEGIN
    DELETE FROM about_section
    WHERE id NOT IN (
        SELECT id FROM about_section
        ORDER BY created_at DESC 
        LIMIT 10
    );
END;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS about_section;
-- +goose StatementEnd
