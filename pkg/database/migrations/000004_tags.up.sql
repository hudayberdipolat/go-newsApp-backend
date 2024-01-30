CREATE TYPE tagStatus as ENUM ('active', 'passive');

CREATE TABLE IF NOT EXISTS tags(
    id SERIAL PRIMARY KEY,
    tag_name VARCHAR(255) NOT NULL UNIQUE
);