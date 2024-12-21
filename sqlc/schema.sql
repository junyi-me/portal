CREATE TABLE users (
    user_id   BIGSERIAL PRIMARY KEY,
    name text      NOT NULL,
    email text      NOT NULL,
    password text      NOT NULL
);

CREATE TABLE websites (
    website_id   BIGSERIAL PRIMARY KEY,
    name text      NOT NULL,
    url text      NOT NULL
);

CREATE TABLE pages (
    page_id   BIGSERIAL PRIMARY KEY,
    name text      NOT NULL,
    path text      NOT NULL,
    website_id BIGINT REFERENCES websites(website_id) NOT NULL
);

CREATE TABLE comments (
    comment_id   BIGSERIAL PRIMARY KEY,
    text text      NOT NULL,
    user_id BIGINT REFERENCES users(user_id) NOT NULL,
    page_id BIGINT REFERENCES pages(page_id) NOT NULL,
    reply_to BIGINT REFERENCES comments(comment_id),
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

