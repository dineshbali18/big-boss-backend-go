CREATE TABLE users ( 
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `device_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Represents device id of the user',
    `votes` tinyint(3) UNSIGNED DEFAULT NULL COMMENT 'Represents number of votes left to the particular deviceID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp at which the row was created',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT ='Table to store device id registrations';
  
-- Add unique index to deviceID
CREATE UNIQUE INDEX device_id_idx
ON users (id,device_id);