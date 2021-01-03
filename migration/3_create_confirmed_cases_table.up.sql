CREATE TABLE `confirmed_cases` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `status` int NOT NULL COMMENT '1 -> positive, 2 -> suspek, 3 -> probable, 4 -> kontak erat',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,

  KEY `user_id` (`user_id`)
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
