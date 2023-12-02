-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION on_insert()
    RETURNS TRIGGER AS $$
BEGIN
    -- Check if the quantity is greater than the available amount
    IF NEW.quantity > (SELECT quantity FROM products WHERE id = NEW.id) THEN
        RAISE EXCEPTION 'Insufficient quantity for product %', NEW.product_id;
    END IF;

    -- Reduce the amount in the products table based on the basket quantity
    UPDATE products
    SET quantity = quantity - NEW.quantity
    WHERE id = NEW.id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION on_delete()
    RETURNS TRIGGER AS $$
BEGIN
    -- Increase the amount in the products table based on the basket quantity
    UPDATE products
    SET quantity = products.quantity + OLD.quantity
    WHERE id = OLD.id;

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER on_insert_trigger
    BEFORE INSERT ON basket
    FOR EACH ROW
    EXECUTE FUNCTION on_insert();

CREATE TRIGGER on_delete_trigger
    AFTER DELETE ON basket
    FOR EACH ROW
    EXECUTE FUNCTION on_delete();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER on_insert_trigger ON basket;
DROP TRIGGER on_delete_trigger ON basket;
-- +goose StatementEnd
