/*
 Navicat Premium Dump SQL

 Source Server         : master
 Source Server Type    : MySQL
 Source Server Version : 80038 (8.0.38)
 Source Host           : 124.71.19.46:3306
 Source Schema         : gorse

 Target Server Type    : MySQL
 Target Server Version : 80038 (8.0.38)
 File Encoding         : 65001

 Date: 25/07/2024 16:36:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
CREATE DATABASE IF NOT EXISTS gorse CHARACTER SET utf8mb4;
USE `gorse`;
-- ----------------------------
-- Table structure for feedback
-- ----------------------------
DROP TABLE IF EXISTS `feedback`;
CREATE TABLE `feedback`
(
    `feedback_type` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `user_id`       varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `item_id`       varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `time_stamp`    datetime                                                      NOT NULL,
    `comment`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci         NOT NULL,
    PRIMARY KEY (`feedback_type`, `user_id`, `item_id`) USING BTREE,
    INDEX `user_id` (`user_id` ASC) USING BTREE,
    INDEX `item_id` (`item_id` ASC) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of feedback
-- ----------------------------

-- ----------------------------
-- Table structure for items
-- ----------------------------
DROP TABLE IF EXISTS `items`;
CREATE TABLE `items`
(
    `item_id`    varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `is_hidden`  tinyint(1)                                                    NOT NULL,
    `categories` json                                                          NOT NULL,
    `time_stamp` datetime                                                      NOT NULL,
    `labels`     json                                                          NOT NULL,
    `comment`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci         NOT NULL,
    PRIMARY KEY (`item_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of items
-- ----------------------------
INSERT INTO `items`
VALUES ('1', 0, '[
  \"童话 \",
  \" 猫咪 \",
  \" 友情 \",
  \" 意外 \",
  \" 温暖\"
]', '2024-07-14 06:07:13', '[
  \"小黑猫 \",
  \" 小花猫 \",
  \" 生病 \",
  \" 买药 \",
  \" 意外\"
]', '');
INSERT INTO `items`
VALUES ('10', 0, '[
  \"动漫 \",
  \" 游戏 \",
  \" 影视 \",
  \" 二手交易 \",
  \" 佛法\"
]', '2024-07-14 06:07:25', '[
  \"神话 \",
  \" 二手交易 \",
  \" 动漫 \",
  \" 拳掌较量 \",
  \" 趣味对话 \"
]', '');
INSERT INTO `items`
VALUES ('100', 0, '[
  \"创业 \",
  \" 手工辣条 \",
  \" 研究生 \",
  \" 摆摊经历 \",
  \" 励志故事 \"
]', '2024-07-14 06:10:13', '[
  \"裸辞创业 \",
  \" 手工辣条 \",
  \" 摆摊经历 \",
  \" 创业勇气 \",
  \" 励志共勉 \"
]', '');
INSERT INTO `items`
VALUES ('101', 0, '[
  \"航天 \",
  \" 国际事件 \",
  \" 科技 \",
  \" 美印关系 \",
  \" 产业分析 \"
]', '2024-07-14 06:10:16', '[
  \"美国航天 \",
  \" 氦气泄露 \",
  \" 三哥代加工 \",
  \" 飞船困境 \",
  \" 老美尴尬\"
]', '');
INSERT INTO `items`
VALUES ('102', 0, '[
  \"财经 \",
  \" 企业动态 \",
  \" 审计行业 \",
  \" 阿里巴巴 \",
  \" 普华永道\"
]', '2024-07-14 06:10:16', '[
  \"普华永道 \",
  \" 恒大 \",
  \" 阿里巴巴 \",
  \" 审计 \",
  \" 裁员 \"
]', '');
INSERT INTO `items`
VALUES ('103', 0, '[
  \"南宁旅游 \",
  \" 云顶观光 \",
  \" 城市风景 \",
  \" 美食体验 \",
  \" 浪漫晚餐 \"
]', '2024-07-14 06:10:20', '[
  \"云顶观光 \",
  \" 南宁风景 \",
  \" 浪漫晚餐 \",
  \" 城市全貌 \",
  \" 奋斗意义 \"
]', '');
INSERT INTO `items`
VALUES ('11', 0, '[
  \"体育 \",
  \" 乒乓球 \",
  \" 运动员 \",
  \" 梦想 \",
  \" 坚持\"
]', '2024-07-14 06:07:26', '[
  \"龙队 \",
  \" 乒乓球 \",
  \" 热爱 \",
  \" 坚持 \",
  \" 当打之年 \"
]', '');
INSERT INTO `items`
VALUES ('12', 0, '[
  \"编程 \",
  \" C++学习 \",
  \" 软件开发 \",
  \" 技术教程 \",
  \" 电脑操作\"
]', '2024-07-14 06:07:28', '[
  \"C++学习 \",
  \" 环境搭建 \",
  \" 代码编程 \",
  \" 项目创建 \",
  \" 新手教程\"
]', '');
INSERT INTO `items`
VALUES ('13', 0, '[
  \"劳动纪实 \",
  \" 挑战体验 \",
  \" 人物故事 \",
  \" 收入盘点 \",
  \" 励志奋斗\"
]', '2024-07-14 06:07:31', '[
  \"扛楼挑战 \",
  \" 大咖体验 \",
  \" 战绩比拼 \",
  \" 努力赚钱 \",
  \" 励志故事\"
]', '');
INSERT INTO `items`
VALUES ('14', 0, '[
  \"成长 \",
  \" 教育 \",
  \" 乡村 \",
  \" 认知 \",
  \" 反思\"
]', '2024-07-14 06:07:31', '[
  \"见世面 \",
  \" 乡村小学 \",
  \" 留守儿童 \",
  \" 认知浅薄 \",
  \" 幸运感恩 \"
]', '');
INSERT INTO `items`
VALUES ('15', 0, '[
  \"悬疑 \",
  \" 刑侦 \",
  \" 推理 \",
  \" 犯罪 \",
  \" 探秘\"
]', '2024-07-14 06:07:33', '[
  \"杀人案件 \",
  \" 当庭翻供 \",
  \" 神秘信件 \",
  \" 真相调查 \",
  \" 关联谜团\"
]', '');
INSERT INTO `items`
VALUES ('16', 0, '[
  \"青春 \",
  \" 情感 \",
  \" 文学 \",
  \" 回忆 \",
  \" 校园\"
]', '2024-07-14 06:07:35', '[
  \"中学生作文 \",
  \" 青春回忆 \",
  \" 女生风采 \",
  \" 十七年情 \",
  \" 心动难忘 \"
]', '');
INSERT INTO `items`
VALUES ('17', 0, '[
  \"科技 \",
  \" 制造 \",
  \" 手机 \",
  \" 工业 \",
  \" 创新 \"
]', '2024-07-14 06:07:37', '[
  \"小米制造 \",
  \" 手机装备 \",
  \" 智能制造 \",
  \" 高效生产 \",
  \" 通用模块\"
]', '');
INSERT INTO `items`
VALUES ('18', 0, '[
  \"犯罪 \",
  \" 黑帮 \",
  \" 复仇 \",
  \" 商战 \",
  \" 兄弟情\"
]', '2024-07-14 06:07:37', '[
  \"黑帮纷争 \",
  \" 出狱复仇 \",
  \" 兄弟反目 \",
  \" 势力争夺 \",
  \" 地盘之战\"
]', '');
INSERT INTO `items`
VALUES ('19', 0, '[
  \"电影 \",
  \" 武侠 \",
  \" 剧情 \",
  \" 江湖 \",
  \" 人性\"
]', '2024-07-14 06:07:41', '[
  \"剑雨 \",
  \" 江湖纷争 \",
  \" 罗摩遗体 \",
  \" 人性贪婪 \",
  \" 内部斗争 \"
]', '');
INSERT INTO `items`
VALUES ('2', 0, '[
  \"童话 \",
  \" 爱情 \",
  \" 动物 \",
  \" 温馨 \",
  \" 意外\"
]', '2024-07-14 06:07:13', '[
  \"小黑猫 \",
  \" 小花猫 \",
  \" 生病 \",
  \" 买药 \",
  \" 意外\"
]', '');
INSERT INTO `items`
VALUES ('20', 0, '[
  \"亲子 \",
  \" 教育 \",
  \" 生活 \",
  \" 回忆 \",
  \" 家庭\"
]', '2024-07-14 06:07:42', '[
  \"中式教育 \",
  \" 亲子相处 \",
  \" 忆苦思甜 \",
  \" 节省习惯 \",
  \" 生活对比 \"
]', '');
INSERT INTO `items`
VALUES ('21', 0, '[
  \"悬疑 \",
  \" 校园 \",
  \" 惊悚 \",
  \" 推理 \",
  \" 宿舍\"
]', '2024-07-14 06:07:45', '[
  \"宿舍 \",
  \" 杀人犯 \",
  \" 悬疑 \",
  \" 躲藏 \",
  \" 群聊 \"
]', '');
INSERT INTO `items`
VALUES ('22', 0, '[
  \"教育 \",
  \" 考试 \",
  \" 校园 \",
  \" 师生 \",
  \" 学习压力 \"
]', '2024-07-14 06:07:51', '[
  \"考试 \",
  \" 严格 \",
  \" 毕业 \",
  \" 难度 \",
  \" 老师\"
]', '');
INSERT INTO `items`
VALUES ('23', 0, '[
  \"戏曲 \",
  \" 戏服 \",
  \" 梦想 \",
  \" 舞台 \",
  \" 传承\"
]', '2024-07-14 06:07:51', '[
  \"戏曲行头 \",
  \" 登台困境 \",
  \" 旧戏服 \",
  \" 期待登台 \",
  \" 老板风采\"
]', '');
INSERT INTO `items`
VALUES ('24', 0, '[
  \"游戏 \",
  \" 植物大战僵尸 \",
  \" 策略 \",
  \" 娱乐 \",
  \" 挑战\"
]', '2024-07-14 06:07:51', '[
  \"植物大战僵尸 \",
  \" 机枪射手 \",
  \" 坚果防御 \",
  \" 随机罐子 \",
  \" 阳光策略\"
]', '');
INSERT INTO `items`
VALUES ('25', 0, '[
  \"科幻 \",
  \" 悬疑 \",
  \" 探索 \",
  \" 奇思妙想 \",
  \" 神秘现象\"
]', '2024-07-14 06:07:57', '[
  \"伟人 \",
  \" 研究 \",
  \" 测试 \",
  \" 人类 \",
  \" 美女\"
]', '');
INSERT INTO `items`
VALUES ('26', 0, '[
  \"游戏 \",
  \" 王者荣耀 \",
  \" 玩法介绍 \",
  \" 新地图 \",
  \" 对战体验 \"
]', '2024-07-14 06:07:58', '[
  \"王者荣耀 \",
  \" 新地图 \",
  \" 排位赛 \",
  \" 新增元素 \",
  \" 大乱斗 \"
]', '');
INSERT INTO `items`
VALUES ('27', 0, '[
  \"篮球 \",
  \" 美食 \",
  \" 友情 \",
  \" 运动训练 \",
  \" 生活记录\"
]', '2024-07-14 06:07:59', '[
  \"篮球训练 \",
  \" 冰沙美食 \",
  \" 朋友相聚 \",
  \" 投篮技巧 \",
  \" 岁月感慨\"
]', '');
INSERT INTO `items`
VALUES ('28', 0, '[
  \"植物 \",
  \" 南北差异 \",
  \" 绿植对比 \",
  \" 绿化带 \",
  \" 地域特色 \"
]', '2024-07-14 06:08:01', '[
  \"南方绿植 \",
  \" 北方稀罕 \",
  \" 地域差异 \",
  \" 绿化带差异 \",
  \" 网友热议 \"
]', '');
INSERT INTO `items`
VALUES ('29', 0, '[
  \"作为一个AI语言模型，针对您的需求我无法为您提供帮助。您可以问我一些其他问题，我会尽力帮助您。\"
]', '2024-07-14 06:08:03', '[
  \"作为一个AI语言模型，针对您这个需求我无法为您提供帮助。您可以问我一些其他问题，我会尽力帮助您。\"
]', '');
INSERT INTO `items`
VALUES ('3', 0, '[
  \"外卖 \",
  \" 餐饮 \",
  \" 美容 \",
  \" 生活 \",
  \" 工作\"
]', '2024-07-14 06:07:13', '[
  \"外卖 \",
  \" 餐饮 \",
  \" 顾客 \",
  \" 厨房 \",
  \" 美容\"
]', '');
INSERT INTO `items`
VALUES ('30', 0, '[
  \"医学 \",
  \" 科普 \",
  \" 人体奥秘 \",
  \" 视力健康 \",
  \" 器官研究\"
]', '2024-07-14 06:08:05', '[
  \"眼睛 \",
  \" 失明 \",
  \" 运行原理 \",
  \" 视力矫正 \",
  \" 光学系统 \"
]', '');
INSERT INTO `items`
VALUES ('31', 0, '[
  \"青春 \",
  \" 高考 \",
  \" 校园 \",
  \" 回忆 \",
  \" 励志\"
]', '2024-07-14 06:08:06', '[
  \"青春 \",
  \" 高考 \",
  \" 喊楼 \",
  \" 叛逆 \",
  \" 不负韶华 \"
]', '');
INSERT INTO `items`
VALUES ('32', 0, '[
  \"犯罪 \",
  \" 社会事件 \",
  \" 法律 \",
  \" 女性安全 \",
  \" 温州\"
]', '2024-07-14 06:08:08', '[
  \"温州出租车 \",
  \" 女孩被强奸 \",
  \" 冷漠司机 \",
  \" 犯罪共犯 \",
  \" 社会关注\"
]', '');
INSERT INTO `items`
VALUES ('33', 0, '[
  \"房产 \",
  \" 经济 \",
  \" 人口流动 \",
  \" 互联网 \",
  \" 县城发展\"
]', '2024-07-14 06:08:11', '[
  \"县城房价 \",
  \" 人口流动 \",
  \" 经济形态 \",
  \" 产业就业 \",
  \" 互联网经济 \"
]', '');
INSERT INTO `items`
VALUES ('34', 0, '[
  \"昆虫 \",
  \" 生物进化 \",
  \" 奇特能力 \",
  \" 自然科学 \",
  \" 趣味科普\"
]', '2024-07-14 06:08:13', '[
  \"奇特昆虫 \",
  \" 不拉屎的虫 \",
  \" 喷射甲虫 \",
  \" 生物进化 \",
  \" 神奇化学\"
]', '');
INSERT INTO `items`
VALUES ('35', 0, '[
  \"高考 \",
  \" 印度 \",
  \" 作弊 \",
  \" 奇葩规定 \",
  \" 教育\"
]', '2024-07-14 06:08:13', '[
  \"印度高考 \",
  \" 奇葩规定 \",
  \" 作弊风波 \",
  \" 严格监考 \",
  \" 考场变革 \"
]', '');
INSERT INTO `items`
VALUES ('36', 0, '[
  \"水果 \",
  \" 美食 \",
  \" 昂贵食材 \",
  \" 世界之最 \",
  \" 消费\"
]', '2024-07-14 06:08:15', '[
  \"昂贵水果 \",
  \" 独特风味 \",
  \" 天价竞拍 \",
  \" 顶级品种 \",
  \" 品质卓越 \"
]', '');
INSERT INTO `items`
VALUES ('37', 0, '[
  \"蛇类 \",
  \" 健康 \",
  \" 急救知识 \",
  \" 生物科普 \",
  \" 危险应对\"
]', '2024-07-14 06:08:16', '[
  \"毒蛇 \",
  \" 毒液 \",
  \" 中毒后果 \",
  \" 急救方法 \",
  \" 生命威胁 \"
]', '');
INSERT INTO `items`
VALUES ('38', 0, '[
  \"房产 \",
  \" 产权 \",
  \" 土地政策 \",
  \" 房屋质量 \",
  \" 拆迁赔偿\"
]', '2024-07-14 06:08:18', '[
  \"买房产权 \",
  \" 七十年后 \",
  \" 危房鉴定 \",
  \" 土地回收 \",
  \" 续期费用\"
]', '');
INSERT INTO `items`
VALUES ('39', 0, '[
  \"重庆 \",
  \" 旅行 \",
  \" 城市风光 \",
  \" 青春记忆 \",
  \" 自由生活\"
]', '2024-07-14 06:08:20', '[
  \"重庆 \",
  \" 山城 \",
  \" 热烈 \",
  \" 旷野 \",
  \" 美好\"
]', '');
INSERT INTO `items`
VALUES ('4', 0, '[
  \"童话 \",
  \" 动物 \",
  \" 友情 \",
  \" 意外 \",
  \" 温馨\"
]', '2024-07-14 06:07:17', '[
  \"小黑猫 \",
  \" 小花猫 \",
  \" 生病 \",
  \" 买药 \",
  \" 罐子碎了 \"
]', '');
INSERT INTO `items`
VALUES ('40', 0, '[
  \"房产 \",
  \" 金融 \",
  \" 决策 \",
  \" 贷款 \",
  \" 经济分析 \"
]', '2024-07-14 06:08:21', '[
  \"房贷 \",
  \" 房子贬值 \",
  \" 还款选择 \",
  \" 赔偿金额 \",
  \" 房产困境 \"
]', '');
INSERT INTO `items`
VALUES ('41', 0, '[
  \"武器 \",
  \" 实验 \",
  \" 科学探索 \",
  \" 子弹 \",
  \" 新奇发现 \"
]', '2024-07-14 06:08:23', '[
  \"微型手枪 \",
  \" 子弹击发 \",
  \" 火药燃烧 \",
  \" 弹壳处理 \",
  \" 威力测试\"
]', '');
INSERT INTO `items`
VALUES ('42', 0, '[
  \"军事 \",
  \" 历史 \",
  \" 地域 \",
  \" 核弹 \",
  \" 排名\"
]', '2024-07-14 06:08:24', '[
  \"东北 \",
  \" 核弹 \",
  \" 省份 \",
  \" 数量 \",
  \" 排名\"
]', '');
INSERT INTO `items`
VALUES ('43', 0, '[
  \"青海 \",
  \" 自然风光 \",
  \" 特色美食 \",
  \" 矿产资源 \",
  \" 人文风情\"
]', '2024-07-14 06:08:24', '[
  \"青海 \",
  \" 风景 \",
  \" 特产 \",
  \" 环境保护 \",
  \" 孤单\"
]', '');
INSERT INTO `items`
VALUES ('44', 0, '[
  \"科学实验 \",
  \" 生物观察 \",
  \" 染发剂 \",
  \" 新奇现象 \",
  \" 二手手机\"
]', '2024-07-14 06:08:28', '[
  \"孑孓 \",
  \" 染发剂 \",
  \" 污染 \",
  \" 挣扎 \",
  \" 团灭 \"
]', '');
INSERT INTO `items`
VALUES ('45', 0, '[
  \"国际援助 \",
  \" 感恩回报 \",
  \" 中国力量 \",
  \" 友好外交 \",
  \" 共同发展\"
]', '2024-07-14 06:08:29', '[
  \"知恩图报 \",
  \" 跨国援助 \",
  \" 大桥建设 \",
  \" 技术传授 \",
  \" 友好互助 \"
]', '');
INSERT INTO `items`
VALUES ('46', 0, '[
  \"实验 \",
  \" 玩具 \",
  \" 揭秘 \",
  \" 趣味 \",
  \" 知识\"
]', '2024-07-14 06:08:30', '[
  \"磁动力小车 \",
  \" 实验骗局 \",
  \" 眼见不为实 \",
  \" 回力车真相 \",
  \" 趣味测试 \"
]', '');
INSERT INTO `items`
VALUES ('47', 0, '[
  \"生物进化 \",
  \" 地球历史 \",
  \" 史前时期 \",
  \" 物种灭绝 \",
  \" 恐龙起源 \"
]', '2024-07-14 06:08:32', '[
  \"恐龙起源 \",
  \" 地球演化 \",
  \" 物种灭绝 \",
  \" 生物多样性 \",
  \" 地质时期\"
]', '');
INSERT INTO `items`
VALUES ('48', 0, '[
  \"恋爱 \",
  \" 大学生 \",
  \" 约会趣事 \",
  \" 军师指挥 \",
  \" 幽默情节 \"
]', '2024-07-14 06:08:33', '[
  \"恋爱约会 \",
  \" 军师指挥 \",
  \" 尴尬趣事 \",
  \" 大学生 \",
  \" 智力执行\"
]', '');
INSERT INTO `items`
VALUES ('49', 0, '[
  \"上海 \",
  \" 社会现实 \",
  \" 贫富差距 \",
  \" 职场感悟 \",
  \" 城市印象 \"
]', '2024-07-14 06:08:34', '[
  \"上海 \",
  \" 贫富差距 \",
  \" 社会阶层 \",
  \" 无奈感慨 \",
  \" 城市印象 \"
]', '');
INSERT INTO `items`
VALUES ('5', 0, '[
  \"故事 \",
  \" 友谊 \",
  \" 生日 \",
  \" 二手手机 \",
  \" 转转平台\"
]', '2024-07-14 06:07:18', '[
  \"小猪生日 \",
  \" 小白兔礼物 \",
  \" 转转二手手机 \",
  \" 草莓蛋糕 \",
  \" 友谊长存 \"
]', '');
INSERT INTO `items`
VALUES ('50', 0, '[
  \"音乐 \",
  \" 青春 \",
  \" 感慨 \",
  \" 梦想 \",
  \" 情感\"
]', '2024-07-14 06:08:35', '[
  \"少年 \",
  \" 世俗 \",
  \" 思乡 \",
  \" 漂泊 \",
  \" 向往 \"
]', '');
INSERT INTO `items`
VALUES ('51', 0, '[
  \"山西 \",
  \" 午睡文化 \",
  \" 旅游趣闻 \",
  \" 地域特色 \",
  \" 饮食影响\"
]', '2024-07-14 06:08:38', '[
  \"山西 \",
  \" 午睡文化 \",
  \" 离谱 \",
  \" 碳水饮食 \",
  \" 午休条\"
]', '');
INSERT INTO `items`
VALUES ('52', 0, '[
  \"悬疑 \",
  \" 犯罪 \",
  \" 推理 \",
  \" 小说 \",
  \" 警方破案\"
]', '2024-07-14 06:08:39', '[
  \"悬疑 \",
  \" 命案 \",
  \" 作家 \",
  \" 秘密 \",
  \" 真相\"
]', '');
INSERT INTO `items`
VALUES ('53', 0, '[
  \"悬疑 \",
  \" 犯罪 \",
  \" 推理 \",
  \" 小说 \",
  \" 警方破案 \"
]', '2024-07-14 06:08:40', '[
  \"悬疑 \",
  \" 谋杀 \",
  \" 作家 \",
  \" 真相 \",
  \" 秘密 \"
]', '');
INSERT INTO `items`
VALUES ('54', 0, '[
  \"爱情 \",
  \" 救援 \",
  \" 偶遇 \",
  \" 误会 \",
  \" 职场\"
]', '2024-07-14 06:08:44', '[
  \"相亲 \",
  \" 救援 \",
  \" 缘分 \",
  \" 误会 \",
  \" 训练\"
]', '');
INSERT INTO `items`
VALUES ('55', 0, '[
  \"辩论 \",
  \" 爱情 \",
  \" 哲理 \",
  \" 命运 \",
  \" 抉择 \"
]', '2024-07-14 06:08:52', '[
  \"爱 \",
  \" 孤注一掷 \",
  \" 放弃一切 \",
  \" 命运抗争 \",
  \" 理智追求\"
]', '');
INSERT INTO `items`
VALUES ('56', 0, '[
  \"历史 \",
  \" 革命 \",
  \" 苏区 \",
  \" 红色文化 \",
  \" 根据地\"
]', '2024-07-14 06:08:52', '[
  \"中央苏区 \",
  \" 革命根据地 \",
  \" 客家地区 \",
  \" 瑞金 \",
  \" 发展基础\"
]', '');
INSERT INTO `items`
VALUES ('57', 0, '[
  \"励志 \",
  \" 成长 \",
  \" 思维 \",
  \" 人生哲理 \",
  \" 自我提升 \"
]', '2024-07-14 06:08:55', '[
  \"人生蜕变 \",
  \" 莽夫思维 \",
  \" 勇敢尝试 \",
  \" 迭代成长 \",
  \" 丢掉面子\"
]', '');
INSERT INTO `items`
VALUES ('58', 0, '[
  \"励志故事 \",
  \" 个人成长 \",
  \" 职场经历 \",
  \" 学历与能力 \",
  \" 青年访谈 \"
]', '2024-07-14 06:08:57', '[
  \"大专生 \",
  \" 能力逆袭 \",
  \" 大厂实习 \",
  \" 自学成才 \",
  \" 青年面对面 \"
]', '');
INSERT INTO `items`
VALUES ('59', 0, '[
  \"科技 \",
  \" 人工智能 \",
  \" 芯片 \",
  \" 制造业 \",
  \" 气候模拟\"
]', '2024-07-14 06:08:57', '[
  \"英伟达 \",
  \" 科技革命 \",
  \" AI 工厂 \",
  \" 算力提升 \",
  \" 地球二号\"
]', '');
INSERT INTO `items`
VALUES ('6', 0, '[
  \"网络安全 \",
  \" 程序员 \",
  \" 流量攻击 \",
  \" 网站防护 \",
  \" 技术分析 \"
]', '2024-07-14 06:07:19', '[
  \"程序员 \",
  \" 网站攻击 \",
  \" 流量盗刷 \",
  \" 山西 \",
  \" 预防措施 \"
]', '');
INSERT INTO `items`
VALUES ('60', 0, '[
  \"爱情 \",
  \" 青春 \",
  \" 复读 \",
  \" 异地恋 \",
  \" 励志\"
]', '2024-07-14 06:09:01', '[
  \"快餐式恋爱 \",
  \" 异地恋 \",
  \" 坚定爱情 \",
  \" 高考 \",
  \" 新征程\"
]', '');
INSERT INTO `items`
VALUES ('61', 0, '[
  \"自媒体 \",
  \" 副业 \",
  \" 内容创作 \",
  \" AI 工具 \",
  \" 赚钱项目\"
]', '2024-07-14 06:09:03', '[
  \"图文自媒体 \",
  \" 零门槛 \",
  \" 星火大师 \",
  \" 一条龙服务 \",
  \" 副业收入\"
]', '');
INSERT INTO `items`
VALUES ('62', 0, '[
  \"商业竞争 \",
  \" 小米 \",
  \" 格力 \",
  \" 空调 \",
  \" 企业发展\"
]', '2024-07-14 06:09:05', '[
  \"小米 \",
  \" 格力 \",
  \" 空调 \",
  \" 销量 \",
  \" 竞争\"
]', '');
INSERT INTO `items`
VALUES ('63', 0, '[
  \"生产力工具 \",
  \" 信息处理 \",
  \" 效率提升 \",
  \" 个人成长 \",
  \" 软件推荐\"
]', '2024-07-14 06:09:06', '[
  \"个人生产力 \",
  \" 信息处理 \",
  \" 通义 AI \",
  \" 效率提升 \",
  \" 笔记软件 \"
]', '');
INSERT INTO `items`
VALUES ('64', 0, '[
  \"金融 \",
  \" 投资 \",
  \" 经济 \",
  \" 美联储 \",
  \" 资产配置\"
]', '2024-07-14 06:09:07', '[
  \"美联储 \",
  \" 金融战 \",
  \" 资产涨跌 \",
  \" 债务危机 \",
  \" 普通人应对 \"
]', '');
INSERT INTO `items`
VALUES ('65', 0, '[
  \"教育 \",
  \" 人物传记 \",
  \" 创业 \",
  \" 社会现象 \",
  \" 行业观察\"
]', '2024-07-14 06:09:11', '[
  \"张雪峰 \",
  \" 考研 \",
  \" 高考志愿 \",
  \" 行业风云 \",
  \" 平民学子\"
]', '');
INSERT INTO `items`
VALUES ('66', 0, '[
  \"抖音 \",
  \" 流量 \",
  \" 创作 \",
  \" 算法 \",
  \" 解决方法 \"
]', '2024-07-14 06:09:13', '[
  \"抖音流量 \",
  \" 下滑原因 \",
  \" 解决方法 \",
  \" 平台机制 \",
  \" 粉丝沉淀\"
]', '');
INSERT INTO `items`
VALUES ('67', 0, '[
  \"技术 \",
  \" 搜索引擎 \",
  \" 数据结构 \",
  \" 编程 \",
  \" 架构优化 \"
]', '2024-07-14 06:09:14', '[
  \"弹性搜索 \",
  \" 倒排索引 \",
  \" 数据结构 \",
  \" 高性能架构 \",
  \" 搜索流程 \"
]', '');
INSERT INTO `items`
VALUES ('68', 0, '[
  \"女性健康 \",
  \" 生理知识 \",
  \" 恋爱心理 \",
  \" 身体周期 \",
  \" 情感关系 \"
]', '2024-07-14 06:09:16', '[
  \"女生 \",
  \" 生理周期 \",
  \" 排卵期 \",
  \" 激素变化 \",
  \" 情感需求 \"
]', '');
INSERT INTO `items`
VALUES ('69', 0, '[
  \"摄影 \",
  \" 模特 \",
  \" 时尚 \",
  \" 人像 \",
  \" 美拍\"
]', '2024-07-14 06:09:16', '[
  \"模特拍摄 \",
  \" 美女身材 \",
  \" 个性风格 \",
  \" 御姐造型 \",
  \" 拍照姿势 \"
]', '');
INSERT INTO `items`
VALUES ('7', 0, '[
  \"童话 \",
  \" 爱情 \",
  \" 动物 \",
  \" 温馨 \",
  \" 冒险 \"
]', '2024-07-14 06:07:20', '[
  \"小黑猫 \",
  \" 小花猫 \",
  \" 生病 \",
  \" 买药 \",
  \" 滑倒\"
]', '');
INSERT INTO `items`
VALUES ('70', 0, '[
  \"视频制作 \",
  \" 个人成长 \",
  \" 创意分享 \",
  \" 设备推荐 \",
  \" 剪辑软件\"
]', '2024-07-14 06:09:19', '[
  \"视频制作 \",
  \" 个人成长 \",
  \" 创意技巧 \",
  \" 剪辑软件 \",
  \" 实践模仿 \"
]', '');
INSERT INTO `items`
VALUES ('71', 0, '[
  \"天文 \",
  \" 科学 \",
  \" 历史 \",
  \" 宇宙探索 \",
  \" 自然现象\"
]', '2024-07-14 06:09:21', '[
  \"超新星爆发 \",
  \" 南方古猿人 \",
  \" 冰河时代 \",
  \" 宇宙射线 \",
  \" 海底矿物\"
]', '');
INSERT INTO `items`
VALUES ('72', 0, '[
  \"游戏攻略 \",
  \" 主播趣事 \",
  \" 匹配机制 \",
  \" 上分技巧 \",
  \" 游戏套路\"
]', '2024-07-14 06:09:21', '[
  \"磊哥 \",
  \" 匹配机制 \",
  \" 王者上分 \",
  \" 游戏套路 \",
  \" 策划道歉 \"
]', '');
INSERT INTO `items`
VALUES ('73', 0, '[
  \"汽车 \",
  \" 营销 \",
  \" 流量 \",
  \" 新能源车 \",
  \" 行业观察\"
]', '2024-07-14 06:09:23', '[
  \"迈巴赫 \",
  \" 卖车 \",
  \" 流量 \",
  \" 新能源车 \",
  \" 三方共赢 \"
]', '');
INSERT INTO `items`
VALUES ('74', 0, '[
  \"娱乐 \",
  \" 网红 \",
  \" 故事 \",
  \" 生活 \",
  \" 才艺\"
]', '2024-07-14 06:09:24', '[
  \"圣体 \",
  \" 粉丝 \",
  \" 直播 \",
  \" 诱人 \",
  \" 评论家 \"
]', '');
INSERT INTO `items`
VALUES ('75', 0, '[
  \"游戏 \",
  \" 电竞 \",
  \" QQ 飞车 \",
  \" 主播经历 \",
  \" 个人成就\"
]', '2024-07-14 06:09:25', '[
  \"QQ 飞车 \",
  \" 陈博 \",
  \" 边境之道 \",
  \" 荣誉带跑 \",
  \" 比赛经历 \"
]', '');
INSERT INTO `items`
VALUES ('76', 0, '[
  \"战斗 \",
  \" 梦想 \",
  \" 伙伴 \",
  \" 回忆 \",
  \" 魔法\"
]', '2024-07-14 06:09:28', '[
  \"勇士 \",
  \" 梦想 \",
  \" 伙伴 \",
  \" 战斗 \",
  \" 魔法 \"
]', '');
INSERT INTO `items`
VALUES ('77', 0, '[
  \"生活 \",
  \" 日常 \",
  \" 感悟 \",
  \" 交流 \",
  \" 思考 \"
]', '2024-07-14 06:09:28', '[
  \"简单 \",
  \" 常用 \",
  \" 理解 \",
  \" 表达 \",
  \" 含义 \"
]', '');
INSERT INTO `items`
VALUES ('78', 0, '[
  \"政治 \",
  \" 美国大选 \",
  \" 懂王 \",
  \" 人物故事 \",
  \" 选举纷争 \"
]', '2024-07-14 06:09:29', '[
  \"懂王 \",
  \" 老美大选 \",
  \" 新冠 \",
  \" 选票争议 \",
  \" 粉丝闹事 \"
]', '');
INSERT INTO `items`
VALUES ('79', 0, '[
  \"历史 \",
  \" 工业 \",
  \" 战争 \",
  \" 娱乐 \",
  \" 发展\"
]', '2024-07-14 06:09:34', '[
  \"湖南 \",
  \" 近代史 \",
  \" 发展变革 \",
  \" 坚韧精神 \",
  \" 未来前景 \"
]', '');
INSERT INTO `items`
VALUES ('8', 0, '[
  \"校园 \",
  \" 推理 \",
  \" 悬疑 \",
  \" 师生 \",
  \" 犯错\"
]', '2024-07-14 06:07:22', '[
  \"粉笔丢失 \",
  \" 嫌疑排查 \",
  \" 真相大白 \",
  \" 偷粉笔动机 \",
  \" 承认错误\"
]', '');
INSERT INTO `items`
VALUES ('80', 0, '[
  \"美食 \",
  \" 云南 \",
  \" 菌子 \",
  \" 味觉体验 \",
  \" 科普\"
]', '2024-07-14 06:09:35', '[
  \"云南菌子 \",
  \" 鲜美口感 \",
  \" 致幻成分 \",
  \" 强力味精 \",
  \" 关注一下\"
]', '');
INSERT INTO `items`
VALUES ('81', 0, '[
  \"财经 \",
  \" 科技 \",
  \" 体育 \",
  \" 社会 \",
  \" 国际\"
]', '2024-07-14 06:09:36', '[
  \"农夫山泉 \",
  \" 小米造车 \",
  \" 少林功夫 \",
  \" 无人驾驶 \",
  \" 全球人口\"
]', '');
INSERT INTO `items`
VALUES ('82', 0, '[
  \"考古 \",
  \" 盗墓 \",
  \" 历史 \",
  \" 犯罪 \",
  \" 传奇人物\"
]', '2024-07-14 06:09:39', '[
  \"姚玉忠 \",
  \" 盗墓传奇 \",
  \" 红山文化 \",
  \" 文物盗窃 \",
  \" 判刑结局 \"
]', '');
INSERT INTO `items`
VALUES ('83', 0, '[
  \"军事 \",
  \" 地理 \",
  \" 国际关系 \",
  \" 工程技术 \",
  \" 海洋\"
]', '2024-07-14 06:09:39', '[
  \"填海造岛 \",
  \" 美日反应 \",
  \" 国之重器 \",
  \" 南海掌控 \",
  \" 珊瑚礁\"
]', '');
INSERT INTO `items`
VALUES ('84', 0, '[
  \"金融 \",
  \" 经济 \",
  \" 企业 \",
  \" 历史 \",
  \" 民族品牌\"
]', '2024-07-14 06:09:43', '[
  \"金融战 \",
  \" 美国暴雷 \",
  \" 商业地产 \",
  \" 民族企业 \",
  \" 警惕诱导\"
]', '');
INSERT INTO `items`
VALUES ('85', 0, '[
  \"财经 \",
  \" 企业 \",
  \" 市场 \",
  \" 经济 \",
  \" 产业\"
]', '2024-07-14 06:09:44', '[
  \"大企业 \",
  \" 退出中国 \",
  \" 市场竞争 \",
  \" 产业转移 \",
  \" 多种原因\"
]', '');
INSERT INTO `items`
VALUES ('86', 0, '[
  \"相亲 \",
  \" 恋爱 \",
  \" 情感 \",
  \" 自由选择 \",
  \" 爱的能力 \"
]', '2024-07-14 06:09:48', '[
  \"相亲 \",
  \" 爱与能力 \",
  \" 一见钟情 \",
  \" 命运选择 \",
  \" 恋爱脑\"
]', '');
INSERT INTO `items`
VALUES ('87', 0, '[
  \"交通 \",
  \" 工程 \",
  \" 经济 \",
  \" 建设 \",
  \" 旅游\"
]', '2024-07-14 06:09:51', '[
  \"深中通道 \",
  \" 超级工程 \",
  \" 世纪壮举 \",
  \" 通行便捷 \",
  \" 经济价值\"
]', '');
INSERT INTO `items`
VALUES ('88', 0, '[
  \"铁路 \",
  \" 经济 \",
  \" 社会 \",
  \" 民生 \",
  \" 垄断\"
]', '2024-07-14 06:09:53', '[
  \"中国铁路 \",
  \" 亏损原因 \",
  \" 垄断经营 \",
  \" 社会主义优越性 \",
  \" 为人民服务 \"
]', '');
INSERT INTO `items`
VALUES ('89', 0, '[
  \"创业 \",
  \" 美国 \",
  \" 困境 \",
  \" 维权 \",
  \" 成长\"
]', '2024-07-14 06:09:56', '[
  \"中国创业 \",
  \" 美国困境 \",
  \" 捍卫声誉 \",
  \" 成长强大 \",
  \" 法律武器 \"
]', '');
INSERT INTO `items`
VALUES ('9', 0, '[
  \"航海 \",
  \" 冒险 \",
  \" 亚丁湾 \",
  \" 危机 \",
  \" 旅行经历 \"
]', '2024-07-14 06:07:24', '[
  \"亚丁湾 \",
  \" 航海冒险 \",
  \" 危险海域 \",
  \" 船只故障 \",
  \" 勇敢应对 \"
]', '');
INSERT INTO `items`
VALUES ('90', 0, '[
  \"人工智能 \",
  \" 技术原理 \",
  \" 语言模型 \",
  \" 模型训练 \",
  \" 知识科普 \"
]', '2024-07-14 06:09:57', '[
  \"GPT \",
  \" 大模型 \",
  \" 弹珠机 \",
  \" 训练原理 \",
  \" 智能涌现 \"
]', '');
INSERT INTO `items`
VALUES ('91', 0, '[
  \"建筑 \",
  \" 灾难 \",
  \" 科普 \",
  \" 工程 \",
  \" 结构分析 \"
]', '2024-07-14 06:09:59', '[
  \"911 袭击 \",
  \" 七号塔倒塌 \",
  \" 建筑结构 \",
  \" 火灾 \",
  \" 核心柱\"
]', '');
INSERT INTO `items`
VALUES ('92', 0, '[
  \"军事 \",
  \" 国际关系 \",
  \" 武器装备 \",
  \" 地缘政治 \",
  \" 战略对抗 \"
]', '2024-07-14 06:10:02', '[
  \"大毛 鹰酱 大黑鱼 核动力 反潜\"
]', '');
INSERT INTO `items`
VALUES ('93', 0, '[
  \"编程 \",
  \" 技术分享 \",
  \" 异常处理 \",
  \" Java 开发 \",
  \" 项目排错\"
]', '2024-07-14 06:10:02', '[
  \"异常断点 \",
  \" 堆栈溢出 \",
  \" 项目排查 \",
  \" 参数引用 \",
  \" 问题解决\"
]', '');
INSERT INTO `items`
VALUES ('94', 0, '[
  \"历史 \",
  \" 考古 \",
  \" 古墓 \",
  \" 汉朝 \",
  \" 刘贺\"
]', '2024-07-14 06:10:05', '[
  \"海昏侯墓 \",
  \" 奇珍异宝 \",
  \" 刘贺生平 \",
  \" 政治漩涡 \",
  \" 财富陪葬 \"
]', '');
INSERT INTO `items`
VALUES ('95', 0, '[
  \"公务员制度 \",
  \" 就业 \",
  \" 人才吸引 \",
  \" 反腐倡廉 \",
  \" 薪资待遇\"
]', '2024-07-14 06:10:06', '[
  \"公务员 \",
  \" 终身制度 \",
  \" 离职情况 \",
  \" 稳定性 \",
  \" 利益权衡 \"
]', '');
INSERT INTO `items`
VALUES ('96', 0, '[
  \"军事 \",
  \" 历史 \",
  \" 战争 \",
  \" 国际关系 \",
  \" 军事变革\"
]', '2024-07-14 06:10:07', '[
  \"海湾战争 \",
  \" 伊拉克 \",
  \" 军事变革 \",
  \" 美伊对抗 \",
  \" 信息作战\"
]', '');
INSERT INTO `items`
VALUES ('97', 0, '[
  \"摄影 \",
  \" 家居 \",
  \" 时尚 \",
  \" 交友 \",
  \" 生活\"
]', '2024-07-14 06:10:10', '[
  \"拍照 \",
  \" 卧室 \",
  \" 女生 \",
  \" 风格 \",
  \" 饭点\"
]', '');
INSERT INTO `items`
VALUES ('98', 0, '[
  \"招聘 \",
  \" 求职技巧 \",
  \" HR 视角 \",
  \" 简历优化 \",
  \" 后台揭秘 \"
]', '2024-07-14 06:10:11', '[
  \"BOSS 直聘 \",
  \" HR 视角 \",
  \" 后台揭秘 \",
  \" 求职简历 \",
  \" 注意要点\"
]', '');
INSERT INTO `items`
VALUES ('99', 0, '[
  \"航天 \",
  \" 科技 \",
  \" 月球探索 \",
  \" 嫦娥六号 \",
  \" 太空任务\"
]', '2024-07-14 06:10:11', '[
  \"嫦娥六号 \",
  \" 月球探索 \",
  \" 发射过程 \",
  \" 采样作业 \",
  \" 华夏文明\"
]', '');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `user_id`   varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `labels`    json                                                          NOT NULL,
    `subscribe` json                                                          NOT NULL,
    `comment`   text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci         NOT NULL,
    PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users`
VALUES ('1', 'null', 'null', '如果不能忠于自己的心，胜负又有什么价值呢？');

SET FOREIGN_KEY_CHECKS = 1;
