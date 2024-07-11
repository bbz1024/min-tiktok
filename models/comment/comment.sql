create table comment
(
    `id`              int unsigned not null auto_increment comment '评论id',
    `videoid`    int unsigned not null comment '视频id',
    `userid`     int unsigned not null comment '用户id',
    `content`    varchar(255) null comment '评论内容',
    `createdat`       timestamp    NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updatedat`       timestamp    NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (`id`),
    index comment_video (`videoid`)
) engine = innodb
  default charset = utf8mb4;
