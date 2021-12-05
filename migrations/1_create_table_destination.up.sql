CREATE TABLE IF NOT EXISTS destination(
    id SERIAL PRIMARY KEY,
    from_place VARCHAR(255),
    to_place VARCHAR(255),
    price INT,
    distance BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);