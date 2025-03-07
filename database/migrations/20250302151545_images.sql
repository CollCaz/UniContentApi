-- +goose Up
-- +goose StatementBegin
CREATE TABLE image (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    image_url TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_image_timestamp
AFTER UPDATE ON image
FOR EACH ROW
BEGIN
    UPDATE image
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE image;
-- +goose StatementEnd
