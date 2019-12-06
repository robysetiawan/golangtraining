CREATE DATABASE db_twitter;
USE db_twitter;
CREATE TABLE users (
	id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    fullname VARCHAR(255),
    status TINYINT,
    gender ENUM('L','P'),
    email VARCHAR(255) UNIQUE,
    password TEXT,
    location VARCHAR(255),
    created_at DATETIME, -- timestamp default current_timestamp
    updated_at DATETIME -- timestamp default current_timestamp on update current_timestamp
);

ALTER TABLE users
MODIFY COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
MODIFY COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;

CREATE TABLE user_avatars(
	user_id INT PRIMARY KEY,
    url TEXT,
	created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
ALTER TABLE user_avatars
MODIFY COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
MODIFY COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;

CREATE TABLE favourites(
	tweet_id INT,
	user_id INT,
	created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY(tweet_id,user_id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(tweet_id) REFERENCES tweets(id)
);
ALTER TABLE favourites
MODIFY COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
MODIFY COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;

CREATE TABLE user_followers(
	id INT AUTO_INCREMENT PRIMARY KEY ,
	user_id INT,
    following_id INT,
	created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
ALTER TABLE user_followers
MODIFY COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
MODIFY COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;

CREATE TABLE tweets(
	id INT AUTO_INCREMENT PRIMARY KEY ,
	user_id INT,
    type ENUM('text','image','video'),
    message TEXT,
    parent_id INT,
    favourite_count INT,
	created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
ALTER TABLE tweets
MODIFY COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
MODIFY COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
