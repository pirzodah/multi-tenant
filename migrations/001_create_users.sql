CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
                                     id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                     name          VARCHAR(100) NOT NULL,
                                     email         VARCHAR(255) UNIQUE NOT NULL,
                                     phone         VARCHAR(20) UNIQUE,
                                     password_hash TEXT NOT NULL,
                                     role          VARCHAR(20) NOT NULL CHECK (role IN ('buyer', 'seller', 'admin')),
                                     created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);