USE phenopod;

CREATE TABLE `feed` (
    `id`                         INT AUTO_INCREMENT,
    `source`                     VARCHAR(20),
    `source_id`                  VARCHAR(20),
    `url`                        VARCHAR(500),
    `etag`                       VARCHAR(255),
    `last_modified`              VARCHAR(255),
    `refresh_enabled`            TINYINT,
    `refresh_interval`           INT,
    `last_refresh_at`            BIGINT,
    `next_refresh_at`            BIGINT,
    `last_refresh_comment`       VARCHAR(500),
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`url`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `podcast` (
    `id`                         INT,
    `title`                      VARCHAR(500),
    `summary`                    VARCHAR(300),
    `description`                BLOB,
    `image_path`                 VARCHAR(500),
    `language`                   VARCHAR(10),
    `explicit`                   TINYINT DEFAULT 0,
    `author`                     VARCHAR(255),
    `type`                       ENUM('EPISODIC', 'SERIAL') DEFAULT 'EPISODIC',
    `block`                      TINYINT DEFAULT 0,
    `complete`                   TINYINT DEFAULT 0,
    `link`                       VARCHAR(500),
    `owner_name`                 VARCHAR(500),
    `owner_email`                VARCHAR(500),
    `total_episodes`             SMALLINT,
    `total_seasons`              SMALLINT,
    `latest_episode_pub_date`    DATETIME,
    `earliest_episode_pub_date`  DATETIME,    
    `copyright`                  VARCHAR(500),
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`id`) REFERENCES `feed` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `episode` (
    `id`                         INT AUTO_INCREMENT,
    `podcast_id`                 INT,
    `guid`                       VARCHAR(500),
    `title`                      VARCHAR(500),
    `media_url`                  VARCHAR(700),
    `media_type`                 VARCHAR(50),
    `media_size`                 BIGINT,
    `pub_date`                   DATETIME,
    `summary`                    VARCHAR(300),
    `description`                BLOB,
    `duration`                   INT,
    `link`                       VARCHAR(500),
    `image_link`                 VARCHAR(500),
    `explicit`                   TINYINT DEFAULT 0,
    `episode`                    INT,
    `season`                     INT,
    `type`                       ENUM('FULL', 'TRAILER', 'BONUS') DEFAULT 'FULL',
    `block`                      TINYINT DEFAULT 0,
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`podcast_id`) REFERENCES `podcast` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `category` (
    `id`                         INT,
    `parent_id`                  INT, 
    `name`                       VARCHAR(100),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`parent_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE `podcast_category` (
    `podcast_id`                 INT,
    `category_id`                INT,
    PRIMARY KEY (`podcast_id`, `category_id`),
    FOREIGN KEY (`podcast_id`) REFERENCES `podcast` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE `task` (
    `id`                         INT,
    `name`                       VARCHAR(30),
    `type`                       ENUM('PERIODIC', 'IMMEDIATE', 'ONEOFF'),
    `interval_`                  INT,
    `next_run_at`                BIGINT,
    `active`                     TINYINT,
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY (`id`)
);

INSERT INTO `task` (`id`, `name`, `type`, `interval_`, `next_run_at`, `active`, `created_at`, `updated_at`) VALUES (1, 'scrape_trending', 'PERIODIC', 86400, 0, 0, 0, 0);
INSERT INTO `task` (`id`, `name`, `type`, `interval_`, `next_run_at`, `active`, `created_at`, `updated_at`) VALUES (2, 'scrape_categories', 'IMMEDIATE', 0, 0, 0, 0, 0);
INSERT INTO `task` (`id`, `name`, `type`, `interval_`, `next_run_at`, `active`, `created_at`, `updated_at`) VALUES (3, 'scrape_itunes_directory', 'PERIODIC', 21600, 0, 0, 0, 0);
INSERT INTO `task` (`id`, `name`, `type`, `interval_`, `next_run_at`, `active`, `created_at`, `updated_at`) VALUES (4, 'schedule_podcast_refresh', 'PERIODIC', 7200, 0, 0, 0, 0);

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

CREATE TABLE `user` (
    `id`                         INT AUTO_INCREMENT,
    `name`                       VARCHAR(100),
    `email`                      VARCHAR(255),
    `gender`                     VARCHAR(20),
    `sign_in_method`             ENUM('EMAIL', 'GOOGLE', 'FACEBOOK', 'TWITTER'),
    `is_admin`                   TINYINT DEFAULT 0,
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY(`id`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `email_account` (
    `email`                      VARCHAR(100),
    `user_id`                    INT,
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY(`email`),
    FOREIGN KEY(`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE 
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `google_account` (
    `id`                         VARCHAR(50),
    `user_id`                    INT,
    `email`                      VARCHAR(255),
    `family_name`                VARCHAR(100),
    `gender`                     VARCHAR(20),
    `given_name`                 VARCHAR(100),
    `link`                       VARCHAR(500),
    `locale`                     VARCHAR(50), 
    `name`                       VARCHAR(200),
    `picture`                    VARCHAR(500),
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `facebook_account` (
    `id`                         VARCHAR(50),
    `user_id`                    INT,
    `name`                       VARCHAR(200),
    `email`                      VARCHAR(255),
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `twitter_account` (
    `id`                         VARCHAR(50), 
    `user_id`                    INT,
    `name`                       VARCHAR(60),
    `screen_name`                VARCHAR(60),
    `location`                   VARCHAR(100),
    `url`                        VARCHAR(255),
    `description`                MEDIUMBLOB,
    `verified`                   TINYINT,
    `followers_count`            INT,
    `friends_count`              INT,
    `profile_image`              VARCHAR(255),
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `playback` (
    `user_id`                   INT,
    `episode_id`                INT,
    `play_count`                SMALLINT,
    `current_progress`          FLOAT,
    `cumulative_progress`       FLOAT,
    `last_played_at`            DATETIME,
    `created_at`                BIGINT,
    `updated_at`                BIGINT,
    PRIMARY KEY (`user_id`, `episode_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`episode_id`) REFERENCES `episode` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE `subscription` (
    `user_id`                    INT,
    `podcast_id`                 INT,
    `active`                     TINYINT,
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT, 
    PRIMARY KEY (`user_id`, `podcast_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`podcast_id`) REFERENCES `podcast` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE `playlist` (
    `id`                         INT AUTO_INCREMENT,
    `user_id`                    INT,
    `title`                      VARCHAR(170),
    `description`                VARCHAR(170),
    `privacy`                    ENUM('PUBLIC', 'PRIVATE', 'UNLINKED'),
    `episode_count`              INT,
    `preview_image`              VARCHAR(200),
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY (`id`), 
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `playlist_member` (
    `playlist_id`                INT,
    `episode_id`                 INT,
    `position`                   SMALLINT,
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY (`playlist_id`, `episode_id`),
    FOREIGN KEY (`playlist_id`) REFERENCES `playlist` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`episode_id`) REFERENCES `episode` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
