CREATE TABLE `billings` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `sender` int(11) NOT NULL,
  `reciever` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `time_bill` timestamp NOT NULL,
  `task_id` int(11) NOT NULL,
  `btype` enum('pay','hold') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
