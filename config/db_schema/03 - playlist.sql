DROP TABLE IF EXISTS `playlist_member`;
DROP TABLE IF EXISTS `playlist`;
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
    `active`                     TINYINT,
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY (`playlist_id`, `episode_id`),
    FOREIGN KEY (`playlist_id`) REFERENCES `playlist` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`episode_id`) REFERENCES `episode` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
