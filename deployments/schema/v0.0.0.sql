DROP DATABASE IF EXISTS `phenopod`;
CREATE DATABASE `phenopod`;
USE `phenopod`;

CREATE TABLE `podcast` (
    `id` VARCHAR(20),
    `title` VARCHAR(500) NOT NULL,
    `description` BLOB NOT NULL,
    `image_path` VARCHAR(500) NOT NULL,
    `language` VARCHAR(10) NOT NULL,
    `explicit` TINYINT DEFAULT 0,
    `author` VARCHAR(255) NOT NULL,
    `type` ENUM('EPISODIC', 'SERIAL') DEFAULT 'EPISODIC',
	`block` TINYINT DEFAULT 0,
    `complete` TINYINT DEFAULT 0,
    `link` VARCHAR(500) NOT NULL,
    `owner_name` VARCHAR(500) NOT NULL,
    `owner_email` VARCHAR(500) NOT NULL,
    `copyright` VARCHAR(500) NOT NULL,
    `feed_url` VARCHAR(500) NOT NULL,
    `feed_etag` VARCHAR(255) NOT NULL,
    `feed_last_modified` VARCHAR(255) NOT NULL,
    `new_feed_url` VARCHAR(500) NOT NULL,
    `created_at` INT NOT NULL,
    `updated_at` INT NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`title`),
    UNIQUE KEY (`feed_url`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `episode` (
    `id` VARCHAR(20),
    `podcast_id` VARCHAR(20) NOT NULL,
    `guid` VARCHAR(500) NOT NULL,
    `title` VARCHAR(500) NOT NULL,
    `media_url` VARCHAR(700) NOT NULL,
    `media_type` VARCHAR(50) NOT NULL,
    `media_size` BIGINT NOT NULL,
    `pub_date` DATETIME NOT NULL,
    `description` BLOB NOT NULL,
    `duration` INT NOT NULL,
    `link` VARCHAR(500) NOT NULL,
    `image_link` VARCHAR(500) NOT NULL,
    `explicit` TINYINT DEFAULT 0,
    `episode` INT NOT NULL,
    `season` INT NOT NULL,
    `type` ENUM('FULL', 'TRAILER', 'BONUS') DEFAULT 'FULL',
    `block` TINYINT DEFAULT 0,
    `created_at` INT NOT NULL,
    `updated_at` INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`podcast_id`) REFERENCES `podcast` (`id`)
        ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `itunes_meta` (
    `itunes_id` VARCHAR(15),
    `feed_url` VARCHAR(500),
    `scrapped_at` DATETIME NOT NULL,
    `added_to_db` ENUM('SUCCESS', 'FAILURE', 'PENDING') DEFAULT 'PENDING',
    `updated_at` INT NOT NULL,
    PRIMARY KEY (`itunes_id`),
    UNIQUE KEY (`feed_url`)
);

CREATE TABLE `category` (
    `id` INT,
    `name` VARCHAR(100) NOT NULL,
    `parent_id` INT, 
    PRIMARY KEY (`id`),
    FOREIGN KEY (`parent_id`) REFERENCES `category` (`id`)
        ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE `podcast_category` (
    `podcast_id` VARCHAR(20) NOT NULL,
    `category_id` INT NOT NULL,
    FOREIGN KEY (`podcast_id`) REFERENCES `podcast` (`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
        ON UPDATE CASCADE ON DELETE NO ACTION
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `job_schedule` (
    `job_name` VARCHAR(100), 
    `type` ENUM('PERIODIC', 'ONEOFF', 'IMMEDIATE'),
    `run_at` INT NOT NULL, 
    `run_after` INT NOT NULL,
    `is_active` TINYINT DEFAULT 1,
    `created_at` INT NOT NULL,
    `updated_at` INT NOT NULL,
    PRIMARY KEY(`job_name`)
);

INSERT INTO `category` (`id`, `name`) VALUES (1, 'Arts');
INSERT INTO `category` (`id`, `name`) VALUES (2, 'Business');
INSERT INTO `category` (`id`, `name`) VALUES (3, 'Comedy');
INSERT INTO `category` (`id`, `name`) VALUES (4, 'Education');
INSERT INTO `category` (`id`, `name`) VALUES (5, 'Games & Hobbies');
INSERT INTO `category` (`id`, `name`) VALUES (6, 'Government & Organizations');
INSERT INTO `category` (`id`, `name`) VALUES (7, 'Health');
INSERT INTO `category` (`id`, `name`) VALUES (8, 'Music');
INSERT INTO `category` (`id`, `name`) VALUES (9, 'News & Politics');
INSERT INTO `category` (`id`, `name`) VALUES (10, 'Religion & Spirituality');
INSERT INTO `category` (`id`, `name`) VALUES (11, 'Science & Medicine');
INSERT INTO `category` (`id`, `name`) VALUES (12, 'Society & Culture');
INSERT INTO `category` (`id`, `name`) VALUES (13, 'Sports & Recreation');
INSERT INTO `category` (`id`, `name`) VALUES (14, 'Technology');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (15, 1, 'Design');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (16, 1, 'Fashion & Beauty');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (17, 1, 'Food');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (18, 1, 'Literature');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (19, 1, 'Performing Arts');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (20, 1, 'Visual Arts');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (21, 2, 'Business News');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (22, 2, 'Careers');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (23, 2, 'Investing');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (24, 2, 'Management & Marketing');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (25, 2, 'Shopping');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (26, 4, 'Educational Technology');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (27, 4, 'Higher Education');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (28, 4, 'K-12');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (29, 4, 'Training');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (30, 5, 'Automotive');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (31, 5, 'Aviation');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (32, 5, 'Hobbies');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (33, 5, 'Other Games');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (34, 5, 'Video Games');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (35, 6, 'Local');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (36, 6, 'National');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (37, 6, 'Non-Profit');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (38, 7, 'Alternative Health');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (39, 7, 'Fitness & Nutrition');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (40, 7, 'Self-Help');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (41, 7, 'Sexuality');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (42, 7, 'Kids & Family');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (43, 10, 'Buddhism');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (44, 10, 'Christianity');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (45, 10, 'Hinduism');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (46, 10, 'Islam');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (47, 10, 'Judaism');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (48, 10, 'Other');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (49, 10, 'Spirituality');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (50, 11, 'Medicine');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (51, 11, 'Natural Sciences');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (52, 11, 'Social Sciences');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (53, 12, 'History');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (54, 12, 'Personal Journals');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (55, 12, 'Philosophy');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (56, 12, 'Places & Travel');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (57, 13, 'Amateur');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (58, 13, 'College & High School');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (59, 13, 'Outdoor');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (60, 13, 'Professional');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (61, 13, 'TV & Film');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (62, 14, 'Gadgets');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (63, 14, 'Podcasting');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (64, 14, 'Software How-To');
INSERT INTO `category` (`id`, `parent_id`, `name`) VALUES (65, 14, 'Tech News');
