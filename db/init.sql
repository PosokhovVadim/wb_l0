CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS orders (
    order_uid uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    track_number VARCHAR(60) NOT NULL,
    entry VARCHAR(60) NOT NULL,
    delivery_id uuid,
    payment_id uuid,
    locale VARCHAR(10),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey VARCHAR(10),
    sm_id INT,
    date_created TIMESTAMP,
    oof_shard VARCHAR(10),
    FOREIGN KEY (delivery_id) REFERENCES delivery(delivery_id),
    FOREIGN KEY (payment_id) REFERENCES payment(payment_id)
);


CREATE TABLE IF NOT EXISTS delivery (
    delivery_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    zip VARCHAR(20) NOT NULL,
    city VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS payment (
    payment_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    transaction VARCHAR(255) NOT NULL,
    request_id VARCHAR(255),
    currency VARCHAR(10) NOT NULL,
    provider VARCHAR(255) NOT NULL,
    amount INT NOT NULL,
    payment_dt BIGINT NOT NULL,
    bank VARCHAR(255) NOT NULL,
    delivery_cost INT NOT NULL,
    goods_total INT NOT NULL,
    custom_fee INT NOT NULL
);


CREATE TABLE IF NOT EXISTS order_items (
    item_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    order_uid VARCHAR(255),
    chrt_id INT NOT NULL,
    track_number VARCHAR(60) NOT NULL,
    price INT NOT NULL,
    rid VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    sale INT NOT NULL,
    size VARCHAR(10) NOT NULL,
    total_price INT NOT NULL,
    nm_id INT NOT NULL,
    brand VARCHAR(255) NOT NULL,
    status INT NOT NULL,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
);
