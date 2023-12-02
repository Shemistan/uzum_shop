-- +goose Up
-- +goose StatementBegin
CREATE TABLE Delivery(
                         id uuid PRIMARY KEY default gen_random_uuid(),
                         user_id uuid NOT NULL REFERENCES Users(id),
                         courier_id uuid NOT NULL REFERENCES Users(id),
                         order_id uuid NOT NULL REFERENCES Orders(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Delivery;
-- +goose StatementEnd
