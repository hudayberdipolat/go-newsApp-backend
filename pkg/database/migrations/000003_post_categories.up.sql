CREATE TABLE IF NOT EXISTS post_categories(
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    CONSTRAINT fk_post FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE
);