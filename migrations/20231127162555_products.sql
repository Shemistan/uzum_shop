-- +goose Up
-- +goose StatementBegin
CREATE TABLE Products(
    id uuid PRIMARY KEY default gen_random_uuid(),
    name varchar(255) NOT NULL,
    description varchar(255),
    price float4,
    quantity integer NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Products;
-- +goose StatementEnd
