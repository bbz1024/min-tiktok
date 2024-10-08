CREATE TABLE `video`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `userid`     INT UNSIGNED NOT NULL COMMENT '用户id',
    `title`      VARCHAR(255) NOT NULL COMMENT '标题',
    `playurl`    VARCHAR(255) NOT NULL COMMENT '文件名',
    `coverurl`   VARCHAR(255) NOT NULL COMMENT '封面名',
    `content`    text         NULL COMMENT '摘要',
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
