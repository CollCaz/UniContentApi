-- +goose Up
-- +goose StatementBegin
CREATE TABLE department (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    faculty_name TEXT NOT NULL,
    faculty_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (faculty_id) REFERENCES faculty(id)
);

INSERT INTO department (name, faculty_name) VALUES ("ComputerScience", "Science");
INSERT INTO department (name, faculty_name) VALUES ("ElectricalEngineering" ,"Engineering");

CREATE TRIGGER update_department_timestamp
AFTER UPDATE ON department
FOR EACH ROW
BEGIN
    UPDATE department
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

CREATE TRIGGER update_department_faculty_id
AFTER INSERT ON department
FOR EACH ROW
BEGIN
    UPDATE orders
    SET faculty_id = (SELECT id FROM faculty WHERE name = NEW.faculty_name)
    WHERE id = NEW.id;
END;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS department;
-- +goose StatementEnd

