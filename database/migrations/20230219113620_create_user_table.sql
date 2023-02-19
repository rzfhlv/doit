-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGSERIAL,
    name VARCHAR (255) NOT NULL,
    email VARCHAR (255) NOT NULL,
    username VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
