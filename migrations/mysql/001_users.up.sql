CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `utype` enum('client','executor') NOT NULL,
  `balance` int(11) NOT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=utf8;