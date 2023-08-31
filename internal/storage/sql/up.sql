CREATE TABLE IF NOT EXISTS segment (
    slug VARCHAR(255) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS user_segment (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "user"(id) ON DELETE CASCADE,
    segment_slug VARCHAR(255) REFERENCES segment(slug) ON DELETE CASCADE,
    CONSTRAINT unique_user_segment UNIQUE (user_id, segment_slug)
);
