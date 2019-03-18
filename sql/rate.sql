CREATE TABLE `rate` (
  `name` varchar(64) unique NULL COMMENT '币种',
  `fBuyPri` varchar(64) NOT NULL COMMENT '买入价格',
  `update_time` varchar(64) NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='汇率';
