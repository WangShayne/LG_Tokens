CREATE TABLE `ens` (
  `id` int(64)  NOT NULL AUTO_INCREMENT,
  `domainName` varchar(64) unique NULL COMMENT '域名',
  `pubKey` varchar(64) NOT NULL COMMENT '公钥地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='域名公钥';
