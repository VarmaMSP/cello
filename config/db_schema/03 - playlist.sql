USE `phenopod`;

CREATE TABLE `playback_position` (
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
    `uid`                        VARCHAR(12),
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

CREATE INDEX `user_uid` ON `user` (`uid`);

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
