DROP TABLE IF EXISTS `rila`.`article`;
CREATE TABLE `rila`.`article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(64) DEFAULT NULL,
  `title` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
