PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    nickname TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    age     INTEGER,
    first_name TEXT,
    last_name TEXT,
    gender TEXT,
    password TEXT NOT NULL,
    created_at TEXT DEFAULT (datetime('now', '+1 hour')) NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    post_id TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    category_name TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TEXT DEFAULT (datetime('now', '+1 hour')) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE IF NOT EXISTS categories (
    category_id INTEGER PRIMARY KEY AUTOINCREMENT,
    category_name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    post_id TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TEXT DEFAULT (datetime('now', '+1 hour')) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (post_id) REFERENCES posts (post_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Reactions (
    user_id INTEGER NOT NULL,
    post_id TEXT,
    comment_id TEXT,
    reaction_type TEXT NOT NULL CHECK (reaction_type IN ('like', 'dislike')),
    created_at TEXT DEFAULT (datetime('now', '+1 hour')) NOT NULL,
    PRIMARY KEY (user_id, post_id, comment_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (post_id) REFERENCES posts (post_id) ON DELETE CASCADE,
    FOREIGN KEY (comment_id) REFERENCES comments (comment_id)
);

CREATE TABLE IF NOT EXISTS sessions (
    session_id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    created_at TEXT DEFAULT (datetime('now', '+1 hour')) NOT NULL,
    expired_at TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS chats (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user1 VARCHAR(255) NOT NULL,
    user2 VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sender VARCHAR(64),
    receiver VARCHAR(64),
    message TEXT,
    time TEXT NOT NULL,
    status VARCHAR(64)
);
INSERT
    OR IGNORE INTO categories (category_name)
VALUES
    ('Technology'),
    ('Sport'),
    ('Health'),
    ('Lifestyle'),
    ('Education');