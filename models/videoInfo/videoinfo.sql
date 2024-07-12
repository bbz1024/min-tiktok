CREATE TABLE `videoinfo`
(
    `id`           INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `videoid`      INT UNSIGNED NOT NULL Comment '视频id',
    `videosummery` text         NULL COMMENT '视频摘要',
    `keyword`      varchar(255) NULL COMMENT '关键词', # 运动 | 游戏
    `category`     varchar(255) NULL COMMENT '分类',
    `created_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
#      FOREIGN KEY (`videoid`) REFERENCES `video` (`id`) avoid  #constraint
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
