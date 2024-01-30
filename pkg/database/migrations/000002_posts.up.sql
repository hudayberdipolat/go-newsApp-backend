CREATE TYPE postStatus as ENUM ('active', 'passive', 'draft');
CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY,
    post_title VARCHAR(500) NOT NULL,
    post_slug VARCHAR(500) UNIQUE NOT NULL,
    post_desc TEXT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    click_count INTEGER DEFAULT 0,
    post_status postStatus DEFAULT 'draft',
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);