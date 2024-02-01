CREATE TYPE userStatus as ENUM ('active' ,'passive' );

CREATE TABLE IF NOT EXISTS users(
        id SERIAL PRIMARY KEY,
        full_name VARCHAR(100) NOT NULL,
        phone_number VARCHAR(30) UNIQUE NOT NULL,
        password VARCHAR(255)  NOT NULL,
        status userStatus DEFAULT 'active',
        created_at TIMESTAMP,
        updated_at TIMESTAMP
);