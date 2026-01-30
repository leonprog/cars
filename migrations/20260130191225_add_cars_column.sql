-- +goose Up
-- +goose StatementBegin
CREATE TABLE cars
(
    id           SERIAL PRIMARY KEY,
    mark         TEXT NOT NULL,
    model_car    TEXT NOT NULL,
    owner_count  INTEGER,
    price        INTEGER,
    currency     VARCHAR(3),
    options      JSONB NOT NULL DEFAULT '{}'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cars;
-- +goose StatementEnd