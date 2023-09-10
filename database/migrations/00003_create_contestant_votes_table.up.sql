CREATE TABLE contestants_votes ( 
    `contestant_id` bigint unsigned,
    `votes` bigint UNSIGNED DEFAULT 0 COMMENT 'Represents number of votes received to a particular contestant',
    FOREIGN KEY (contestant_id) REFERENCES contestants (`id`) 
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT ='Table to store contestant voting data';