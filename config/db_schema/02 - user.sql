USE `phenopod`;

CREATE TABLE `user` (
    `id`                         INT AUTO_INCREMENT,
    `uid`                        VARCHAR(12),
    `name`                       VARCHAR(100),
    `email`                      VARCHAR(255),
    `gender`                     VARCHAR(20),
    `sign_in_method`             enum('EMAIL','GOOGLE','FACEBOOK','TWITTER', 'GUEST'),
    `is_admin`                   TINYINT DEFAULT 0,
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY(`id`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE INDEX `user_uid` ON `user` (`uid`);

CREATE TABLE `guest_account` (
    `id`                         VARCHAR(255),
    `user_id`                    INT,
    `device_uuid`                VARCHAR(255),
    `device_os`                  VARCHAR(100),
    `device_model`               VARCHAR(100),
    `created_at`                 BIGINT,
    `updated_at`                 BIGINT,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE 
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
