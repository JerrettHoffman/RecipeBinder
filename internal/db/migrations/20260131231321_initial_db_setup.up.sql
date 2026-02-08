BEGIN;

CREATE TABLE IF NOT EXISTS ingredients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    display_name VARCHAR(255) NOT NULL
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
    FOREIGN KEY(author_id) REFERENCES authors(author_id),
    FOREIGN KEY(uploader_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS recipe_ingredients (
    recipe_id INTEGER NOT NULL,
    ingredient_id INTEGER NOT NULL,
    FOREIGN KEY(recipe_id) REFERENCES recipes(recipe_id),
    FOREIGN KEY(ingredient_id) REFERENCES ingredients(ingredient_id)
);

COMMIT;
