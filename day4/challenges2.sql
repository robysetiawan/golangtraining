USE db_twitter;
-- problem 1.1
INSERT INTO users(username,fullname,status,gender,email,password,location) VALUES
('aldo','nurhadi aldo',1,'L','nurhadialdo@gmail.com','123456','jambi'),
('bedul','bedul',1,'L','bedul@gmail.com','123456','malang');
SELECT * FROM users;

-- problem 1.2
INSERT INTO tweets(user_id,type,message,parent_id,favourite_count) VALUES
(1,'text','halo perkenalkan nama saya nurhadi',0,0),
(1,'text','ini adalah tweet kedua saya',0,0),
(1,'text','ini adalah tweet ketiga saya',0,0),
(2,'text','halo perkenalkan nama saya bedul',0,0),
(2,'text','saya suka makan gado-gado',0,0),
(2,'text','hari ini saya di training of trainer altera academy',0,0);
SELECT * FROM tweets;

-- problem 1.3
INSERT INTO favourites(tweet_id ,user_id) VALUES(1,2);
SELECT * FROM favourites;

-- problem 2.1
SELECT username,fullname FROM users WHERE gender = 'L';

-- problem 2.2
SELECT message FROM tweets WHERE id = 4;

-- problem 2.3
SELECT * FROM users WHERE fullname LIKE '%a%' AND created_at >= CURRENT_DATE() - INTERVAL 7 DAY;

-- problem 2.4
SELECT COUNT(gender) AS jumlah FROM users WHERE gender = 'L';

-- problem 2.5
SELECT * FROM users ORDER BY fullname ASC;

-- problem 3
UPDATE tweets SET message = 'Hello ini updated tweet' WHERE id = 4;

-- problem 4
DELETE FROM tweets WHERE id = 4;