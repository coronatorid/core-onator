CREATE TABLE `exposed_users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `confirmed_cases_id` int NOT NULL,
  `lat` double NOT NULL,
  `long` double NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,

  KEY `user_id` (`user_id`)
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
