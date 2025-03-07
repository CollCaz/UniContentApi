-- +goose Up
-- +goose StatementBegin
CREATE TABLE hero_images (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    sub_title TEXT NOT NULL,
    image_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (image_id) REFERENCES image(id)
);

CREATE TRIGGER update_hero_images_timestamp
AFTER UPDATE ON hero_images
FOR EACH ROW
BEGIN
    UPDATE hero_images
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS hero_images;
-- +goose StatementEnd

