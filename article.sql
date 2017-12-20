create database sample encoding 'UTF-8';
\c sample;
--
-- Table structure for table `article`
--

DROP TABLE IF EXISTS article;
CREATE TABLE article (
    id              serial,        -- primary key.
    title           varchar(45) NOT NULL,   -- title.
    content         text NOT NULL,         -- content.
    created_at      date,          -- create date.
    updated_at      date,          -- update date.
    deleted_at      date           -- delete date.
);

-- test data
INSERT INTO article (title, content, created_at, updated_at) VALUES ('title1', 'content1', NOW(), NOW());
INSERT INTO article (title, content, created_at, updated_at) VALUES ('title2', 'content2', NOW(), NOW());
INSERT INTO article (title, content, created_at, updated_at) VALUES ('title3', 'content3', NOW(), NOW());
INSERT INTO article (title, content, created_at, updated_at) VALUES ('title4', 'content4', NOW(), NOW());
INSERT INTO article (title, content, created_at, updated_at) VALUES ('title5', 'content5', NOW(), NOW());
