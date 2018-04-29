CREATE TABLE `tasks` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `description` varchar(255) NOT NULL,
  `creator` int(11) NOT NULL,
  `executor` int(11) NOT NULL,
  `price` int(11) NOT NULL,
  `status` enum('done','not done') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
