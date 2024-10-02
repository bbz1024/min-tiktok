-- 创建一个单独的用户
CREATE USER IF NOT EXISTS 'lzb200244'@'%' IDENTIFIED BY 'lzb200244';
GRANT ALL PRIVILEGES ON *.* TO 'lzb200244'@'%';

create database if not exists tiktok character set utf8mb4;

create database if not exists gorse character set utf8mb4;