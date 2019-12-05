DROP TABLE IF EXISTS `subscription`;
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

DROP TABLE IF EXISTS `playlist`;
CREATE TABLE `playlist` (
    `id`                         INT AUTO_INCREMENT,
    `user_id`                    INT,
    `title`                      VARCHAR(170),
    `description`                VARCHAR(250),
    `privacy`                    ENUM('PUBLIC', 'PRIVATE', 'UNLINKED'),
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `playlist_member`;
CREATE TABLE `playlist_member` (
    `episode_id`                 INT,
    `playlist_id`                INT,
    `active`                     TINYINT,
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY (`playlist_id`, `episode_id`),
    FOREIGN KEY (`playlist_id`) REFERENCES `playlist` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`episode_id`) REFERENCES `episode` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
