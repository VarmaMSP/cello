USE `phenopod`;

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

ALTER TABLE `user` MODIFY COLUMN `sign_in_method` enum('EMAIL','GOOGLE','FACEBOOK','TWITTER', 'GUEST');