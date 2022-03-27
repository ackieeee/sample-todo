use todo;

CREATE TABLE IF NOT EXISTS `tasks` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `description` varchar(255),
    `date` datetime DEFAULT CURRENT_TIMESTAMP,
    `status` tinyint(1) DEFAULT 0,
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO `tasks` (`title`, `description`, `date`, `status`) VALUES
    ("test1", "test1 description", "2022-02-01 00:00:00", "0"),
    ("test2", "test2 description", "2022-05-01 00:00:00", "0"),
    ("test3", "test3 description", "2022-03-11 00:00:00", "1");
