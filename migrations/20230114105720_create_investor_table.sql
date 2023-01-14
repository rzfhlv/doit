-- +goose Up
-- +goose StatementBegin
CREATE TABLE investors (
    id BIGSERIAL,
    name VARCHAR (255) NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE investors;
-- +goose StatementEnd
