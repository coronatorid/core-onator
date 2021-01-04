CREATE TABLE `reported_cases` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `status` int NOT NULL COMMENT '1 -> confirmed, 0 -> not_confirmed, 2 -> pending',
  `telegram_message_id` varchar(255) NOT NULL,
  `telegram_image_url` varchar(255) NOT NULL,
  `image_path` varchar(255) NOT NULL,
  `image_deleted` boolean,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,

  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
