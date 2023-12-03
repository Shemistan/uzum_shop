-- +goose Up
-- +goose StatementBegin
CREATE TABLE Orders(
                       id uuid PRIMARY KEY default gen_random_uuid(),
                       user_id uuid NOT NULL REFERENCES Users(id),
                       baskets_id uuid[] NOT NULL,

                       address varchar(255) NOT NULL,
                       address_location_x float NOT NULL,
                       address_location_y float NOT NULL,
                       pick_up_location_x float NOT NULL,
                       pick_up_location_y float NOT NULL,
                       metadata varchar(255) NOT NULL,
                       total_price float NOT NULL,

                       created_at timestamp NOT NULL default now(),
                       started_at timestamp,
                       delivered_at timestamp,
                       status varchar(255) NOT NULL default 'pending'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Orders;
-- +goose StatementEnd
