CREATE TABLE users (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       username TEXT NOT NULL UNIQUE,
                       email TEXT NOT NULL UNIQUE,
                       password_hash TEXT NOT NULL,
                       bio TEXT,
                       profile_image TEXT,
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       user_id INTEGER NOT NULL,
                       image_url TEXT NOT NULL,
                       caption TEXT,
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                       FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE comments (
                          id INTEGER PRIMARY KEY AUTOINCREMENT,
                          post_id INTEGER NOT NULL,
                          user_id INTEGER NOT NULL,
                          content TEXT NOT NULL,
                          created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY(post_id) REFERENCES posts(id),
                          FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE likes (
                       user_id INTEGER NOT NULL,
                       post_id INTEGER NOT NULL,
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                       PRIMARY KEY(user_id, post_id),
                       FOREIGN KEY(user_id) REFERENCES users(id),
                       FOREIGN KEY(post_id) REFERENCES posts(id)
);

CREATE TABLE follows (
                         follower_id INTEGER NOT NULL,
                         following_id INTEGER NOT NULL,
                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY(follower_id, following_id),
                         FOREIGN KEY(follower_id) REFERENCES users(id),
                         FOREIGN KEY(following_id) REFERENCES users(id)
);
