CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS orders (
    order_uid uuid PRIMARY KEY,
    order_data jsonb NOT NULL
);