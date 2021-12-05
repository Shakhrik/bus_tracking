CREATE TABLE IF NOT EXISTS bus(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    seat_count INT,
    start_time TIME,
    end_time TIME,
    is_full boolean DEFAULT false,
    destination_id INT REFERENCES destination(id) ON DELETE SET NULL,
    bus_stop_id INT REFERENCES bus_stop(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);