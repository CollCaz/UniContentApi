-- +goose Up
-- +goose StatementBegin
CREATE TABLE hero_images (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    sub_title TEXT NOT NULL,
    image_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (image_id) REFERENCES image (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS hero_images;
-- +goose StatementEnd
