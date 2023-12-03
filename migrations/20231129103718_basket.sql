-- +goose Up
-- +goose StatementBegin
CREATE TABLE Basket(
    id uuid PRIMARY KEY default gen_random_uuid(),
    product_id uuid NOT NULL REFERENCES Products(id),
    quantity int NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Basket;
-- +goose StatementEnd
