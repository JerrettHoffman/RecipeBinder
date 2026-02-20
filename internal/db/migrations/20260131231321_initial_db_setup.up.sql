BEGIN;

CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    hashed_password VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS recipes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    author_id INTEGER NOT NULL,
    uploader_id INTEGER NOT NULL,
    prep_time INTEGER,
    total_time INTEGER,
    --image goes here TODO: chat with clove about what this means
    steps TEXT,
    ingredient_text TEXT,
    yeild VARCHAR(255),
    ingredient_vector tsvector
	GENERATED ALWAYS AS (to_tsvector('english', ingredient_text)) STORED,
    FOREIGN KEY(author_id) REFERENCES authors(id),
    FOREIGN KEY(uploader_id) REFERENCES users(id)
);

CREATE INDEX ingredient_index ON recipes USING GIN (ingredient_vector);

COMMIT;
