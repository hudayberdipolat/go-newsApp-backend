CREATE TYPE adminRole as ENUM ('super_admin', 'admin');

CREATE TABLE IF NOT EXISTS admins(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(100)  NOT NULL,
    phone_number VARCHAR(30)  NOT NULL,
    admin_role adminRole DEFAULT 'admin',
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);