use todo;

CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    PRIMARY KEY(`id`),
    UNIQUE(`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO `users` (`id`, `name`, `email`, `password`) VALUES
    (1, 'test1', 'test1@email.test', 'testpass'),
    (2, 'test2', 'test2@email.test', 'testpass'),
    (3, 'test3', 'test3@email.test', 'testpass'),
    (4, 'test4', 'test4@email.test', 'testpass');