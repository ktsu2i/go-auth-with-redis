use `auth_db`;

CREATE TABLE IF NOT EXISTS `user` (
    `id` varchar(255) NOT NULL,
    `username` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `hashed_password` varchar(255) NOT NULL,
    `created_at` datetime(6) NOT NULL,
    `updated_at` datetime(6) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
