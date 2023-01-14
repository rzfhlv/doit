-- +goose Up
-- +goose StatementBegin
CREATE TABLE investors (
    id BIGINT NOT NULL,
    name VARCHAR (255) NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE investors;
-- +goose StatementEnd
