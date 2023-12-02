-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users(
                      id uuid PRIMARY KEY default gen_random_uuid(),
                      name varchar(255) NOT NULL,
                      surname varchar(255) NOT NULL,
                      phone varchar(255) NOT NULL,
                      login varchar(255) NOT NULL,
                      password_hash varchar(255) NOT NULL,
                      role varchar(255) NOT NULL default 'user',
                      address varchar(255) NOT NULL,
                      location_x float NOT NULL,
                      location_y float NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users
-- +goose StatementEnd
