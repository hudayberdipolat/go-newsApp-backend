CREATE TYPE categoryStatus as ENUM ('active', 'passive');
CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) UNIQUE,
    category_slug VARCHAR(255) UNIQUE,
    category_status categoryStatus DEFAULT 'passive',
    created_at TIMESTAMP
);