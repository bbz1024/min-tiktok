create table `messages`
(
    `id`             int unsigned not null auto_increment,
    `touserid`       int unsigned not null,
    `fromuserid`     int unsigned not null,
    `conversationid` varchar(255) not null,
    `content`        text         not null,
    `createdat`      timestamp    NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (`id`),
    index `idx_conversationid` (`conversationid`)
) engine = innodb
  default charset = utf8mb4;
