CREATE TYPE categoryStatus as ENUM ('active', 'passive');
CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) UNIQUE,
    category_slug VARCHAR(255) UNIQUE,
    category_status categoryStatus DEFAULT 'passive',
    created_at TIMESTAMP
);
CREATE TYPE postStatus as ENUM ('active', 'passive', 'draft');
CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY,
    post_title VARCHAR(500) NOT NULL,
    post_slug VARCHAR(500) UNIQUE NOT NULL,
    post_desc TEXT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    click_count INTEGER DEFAULT 0,
    post_status postStatus DEFAULT 'draft',
    category_id int,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE
);
CREATE TYPE tagStatus as ENUM ('active', 'passive');

CREATE TABLE IF NOT EXISTS tags(
    id SERIAL PRIMARY KEY,
    tag_name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS post_tags(
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    CONSTRAINT fk_post FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

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

CREATE TABLE IF NOT EXISTS user_comment_post(
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    user_comment TEXT NOT NULL,
    created_at TIMESTAMP,
    CONSTRAINT fk_post FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_comment FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TYPE likeType as ENUM ('like','dislike');


CREATE TABLE IF NOT EXISTS user_liked_posts(
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    like_type likeType,
    CONSTRAINT fk_post FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_comment FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE TYPE adminRole as ENUM ('super_admin', 'admin');
CREATE TYPE adminStatus as ENUM ('active', 'passive');

CREATE TABLE IF NOT EXISTS admins(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(100)  NOT NULL,
    phone_number VARCHAR(30)  UNIQUE NOT NULL,
    admin_role adminRole DEFAULT 'admin',
    admin_status adminStatus DEFAULT 'active',
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);  