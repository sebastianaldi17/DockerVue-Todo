CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL DEFAULT '',
    status SMALLINT NOT NULL DEFAULT 0,
    finished BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO todos(content, status)
VALUES ('Sample todo', 1);