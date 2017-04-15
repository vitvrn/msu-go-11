CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `login` varchar(300) NOT NULL,
  `info` text NULL,
  `balance` int NOT NULL
  `status` int NOT NULL
);

INSERT INTO `users` (`id`, `login`, `info`, `score`) VALUES (1, 'Vasily Romanov', 'company: Mail.ru Group', 100500, 1);
