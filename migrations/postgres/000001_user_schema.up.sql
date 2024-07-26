CREATE TABLE IF NOT EXISTS users (
       id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
       full_name VARCHAR NOT NULL,
       email VARCHAR NOT NULL UNIQUE,
       registration_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       role VARCHAR NOT NULL
);


