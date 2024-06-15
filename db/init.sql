CREATE TABLE IF NOT EXISTS orders (
    order_uid VARCHAR(36) PRIMARY KEY,
    order_data jsonb NOT NULL
);


CREATE INDEX IF NOT EXISTS orders_order_uid_idx ON orders (order_uid);