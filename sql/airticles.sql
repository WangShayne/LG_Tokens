CREATE TABLE `articles` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL,
  `keyName` varchar(64) DEFAULT NULL,
  `type` varchar(64) DEFAULT NULL,
  `codes` text,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='文章html代码';