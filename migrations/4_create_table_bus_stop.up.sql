CREATE TABLE IF NOT EXISTS bus_stop(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    distance BIGINT,
    destination_id INT REFERENCES destination(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);