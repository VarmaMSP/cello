USE `phenopod`;
DROP TABLE IF EXISTS `podcast_keyword`;
DROP TABLE IF EXISTS `episode_keyword`;
DROP TABLE IF EXISTS `keyword`;

CREATE TABLE `keyword` (
    `id`        INT AUTO_INCREMENT,
    `text`      VARCHAR(100),
    PRIMARY KEY (`id`),
    INDEX (`text`)
);

CREATE TABLE `podcast_keyword` (
    `keyword_id` INT,
    `podcast_id` INT,
    FOREIGN KEY (`keyword_id`) REFERENCES `keyword` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`podcast_id`) REFERENCES `podcast` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE `episode_Keyword` (
    `keyword_id` INT,
    `episode_id` INT,
    FOREIGN KEY (`keyword_id`) REFERENCES `keyword` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`episode_id`) REFERENCES `episode` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);


