-- +goose Up
-- +goose StatementBegin
INSERT INTO products (name, price, description, quantity) VALUES ('Apple', 100, 'generic fruit', 100), ('Orange', 200, 'generic fruit', 100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM products WHERE name IN ('Apple', 'Orange');
-- +goose StatementEnd
