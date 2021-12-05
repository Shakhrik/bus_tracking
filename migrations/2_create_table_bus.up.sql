CREATE TABLE IF NOT EXISTS bus(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    seat_count INT,
    start_time TIME,
    end_time TIME,
    destination_id INT REFERENCES destination(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);