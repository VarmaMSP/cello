ALTER TABLE episode_playback ADD last_played_at DATETIME AFTER current_time_;
UPDATE episode_playback SET last_played_at = FROM_UNIXTIME(updated_at);
ALTER TABLE episode_playback MODIFY last_played_at DATETIME NOT NULL;

CREATE TABLE `playlist` (
    `id` VARCHAR(20),
    `title` VARCHAR(500),
    `created_by` VARCHAR(20),
    `privacy` ENUM('PUBLIC', 'PRIVATE', 'ANONYMOUS'),
    `created_at` BIGINT,
    `updated_at` BIGINT,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`created_by`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `playlist_item` (
    `playlist_id` VARCHAR(20),
    `episode_id` VARCHAR(20),
    `active` TINYINT,
    `created_at` BIGINT,
    `updated_at` BIGINT,
    UNIQUE KEY (`playlist_id`, `episode_id`),
    FOREIGN KEY (`playlist_id`) REFERENCES `playlist` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`episode_id`) REFERENCES `episode` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
