CREATE TABLE contestants ( 
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Represents name of the contestant',
    `description` text COLLATE utf8mb4_unicode_ci NULL COMMENT 'Represents description for the contestant',
    `image` varchar(255) NULL COMMENT 'Represents image link of the contestant',
    `is_nominated` boolean DEFAULT 0 COMMENT 'Represents whether the contestant is nominated or not',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT ='Table to store device id registrations';

-- Added index to is_nominated and name columns
ALTER TABLE contestants
ADD KEY `is_nominated`(`is_nominated`),
ADD KEY `name`(`name`);