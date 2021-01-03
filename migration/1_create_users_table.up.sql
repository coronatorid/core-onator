CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `phone` varchar(255) NOT NULL,
  `state` tinyint NOT NULL COMMENT '1 -> active, 0 -> inactive',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,

  UNIQUE KEY `users_index_phone` (`phone`),
  KEY `phone_state` (`phone`, `state`)
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
