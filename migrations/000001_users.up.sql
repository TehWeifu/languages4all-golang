CREATE TABLE IF NOT EXISTS users
(
    id          BIGSERIAL PRIMARY KEY,
    android_uid TEXT                        NOT NULL,
    name        text                        NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT now(),
    version     integer                     NOT NULL DEFAULT 1
);
