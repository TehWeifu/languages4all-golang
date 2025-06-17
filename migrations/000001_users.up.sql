CREATE TABLE IF NOT EXISTS users
(
    id          BIGSERIAL PRIMARY KEY,
    android_uid TEXT NOT NULL,
    name        text NOT NULL
);
