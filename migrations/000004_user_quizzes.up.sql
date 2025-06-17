CREATE TABLE IF NOT EXISTS user_quizzes_points
(
    id                   BIGSERIAL PRIMARY KEY,
    user_id              bigint  NOT NULL REFERENCES users ON DELETE CASCADE,
    quiz_id              bigint  NOT NULL REFERENCES quizzes ON DELETE CASCADE,
    points               INTEGER NOT NULL,
    completed            BOOLEAN NOT NULL,
    currentQuestionOrder INTEGER NOT NULL,
    maxPoints            INTEGER NOT NULL
);
