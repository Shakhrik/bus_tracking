CREATE TABLE IF NOT EXISTS user_type(
    id SERIAL PRIMARY KEY,
    alias VARCHAR(255)
);


CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    user_type_id INT REFERENCES user_type(id) ON DELETE SET NULL,
    name VARCHAR(255),
    login VARCHAR(255),
    password VARCHAR(255),
    bus_id INT REFERENCES bus(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);