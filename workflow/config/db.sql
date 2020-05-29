CREATE TABLE IF NOT EXISTS `bytearry` (
`id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
`name` varchar(64) collate utf8_bin default NULL COMMENT '名称',
`bytes` longblob COMMENT '资源文件',
PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;