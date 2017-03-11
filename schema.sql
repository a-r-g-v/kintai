CREATE TABLE users (id integer primary key autoincrement, name varchar(64), created_at TIMESTAMP DEFAULT (DATETIME('now','localtime')));
