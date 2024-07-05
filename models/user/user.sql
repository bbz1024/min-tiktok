create table `users`
(
    `id`              int unsigned not null auto_increment comment '用户id',
    `username`        varchar(32)  not null unique comment '用户名',
    `password`        varchar(255) not null comment '密码',
    `avatar`          varchar(255)      default null comment '头像',
    `backgroundimage` varchar(255)      default null comment '背景图片',
    `signature`       text              default null comment '签名',
    `createdat`       timestamp    NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updatedat`       timestamp    NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    primary key (`id`),
    index `idx_username` (`username`)
) engine = innodb
  default charset = utf8mb4;
