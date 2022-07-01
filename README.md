<!--
 * @Author: liziwei01
 * @Date: 2022-03-03 15:20:51
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-02 05:33:27
 * @Description: README
-->
# github.com/liziwei01/gin-file-download

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

This is a course project written by graduate student Ziwei Li from JHU

## Interface

## Mysql require

```bash

# In database db_idiary_user:

DROP table if exists `tb_user_private_info`;
create table `tb_user_private_info`(
    `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'int',
    `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'clear text',
    `password` varchar(2000) NOT NULL DEFAULT '' COMMENT 'hashed by bcrypt',
    `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT 'clear text',
    `address` varchar(255) NOT NULL DEFAULT '' COMMENT 'client ip address',
    `profile` varchar(255) NOT NULL DEFAULT '' COMMENT 'generated file name uploading to aliyun OSS',
    `db_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'update time',
    primary key (`user_id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户隐私信息表';
```
