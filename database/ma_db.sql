use ma_db;
DROP TABLE IF EXISTS `bills`;
CREATE TABLE `bills` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type` tinyint NOT NULL COMMENT '1:收入 2:支出',
  `amount` decimal(18,2) NOT NULL COMMENT '金额',
  `description` varchar(300) DEFAULT NULL COMMENT '描述',
  `bill_date` date NOT NULL COMMENT '账单日期',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_by` varchar(50) NOT NULL COMMENT '创建人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int  NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(50)  NOT NULL COMMENT '用户名',
  `password` varchar(100)  NOT NULL COMMENT '密码',
  `gender` int NOT NULL DEFAULT 0 COMMENT '性别：0-男，1-女',
  `nick_name` varchar(50)  NULL DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(1000)  NULL DEFAULT NULL COMMENT '头像',
  `tel` varchar(20)   NULL DEFAULT NULL COMMENT '手机号',
  `qq` varchar(20)  NULL DEFAULT NULL COMMENT 'qq',
  `we_chat` varchar(50)  NULL DEFAULT NULL COMMENT '微信',
  `email` varchar(100)  NULL DEFAULT NULL COMMENT '邮箱',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uq_user_name`(`name`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
