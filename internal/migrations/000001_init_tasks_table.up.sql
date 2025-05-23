CREATE TABLE IF NOT EXISTS tasks (
                                     id SERIAL PRIMARY KEY,
                                     title TEXT NOT NULL,
                                     content TEXT,
                                     status BOOLEAN DEFAULT false,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);