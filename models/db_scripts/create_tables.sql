-- 用户表
CREATE TABLE `t_users`
(
    `id`              INT(20)                                       NOT NULL AUTO_INCREMENT,
    `username`        VARCHAR(64) COLLATE utf8mb4_general_ci unique NOT NULL,
    `password`        VARCHAR(64) COLLATE utf8mb4_general_ci        NOT NULL,
    `last_login_time` TIMESTAMP                                     NULL DEFAULT NULL,
    `create_time`     TIMESTAMP                                          DEFAULT CURRENT_TIMESTAMP,
    `update_time`     TIMESTAMP                                          DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
    comment '用户表';


-- 分类表
CREATE TABLE `t_categories`
(
    `id`          INT(10)      NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(128) NOT NULL,
    `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
    comment '支出分类表';

-- 记账记录表
CREATE TABLE `t_transactions`
(
    `id`               INT(20)        NOT NULL AUTO_INCREMENT,
    `user_id`          INT(20)        NOT NULL,
    `category_id`      INT(10)        NOT NULL,
    `title`            VARCHAR(255)   NOT NULL COMMENT '记账标题',
    `desc`             TEXT           NULL COMMENT '记账描述',
    `amount`           DECIMAL(10, 2) NOT NULL,
    `transaction_date` DATETIME       NOT NULL,
    `create_time`      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_time`      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `t_users` (`id`),
    FOREIGN KEY (`category_id`) REFERENCES `t_categories` (`id`),
    INDEX `idx_user_category` (`user_id`, `category_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
    comment '记账记录表';


-- 用户与分类关系表
CREATE TABLE `t_user2cate`
(
    `id`          INT(20) NOT NULL AUTO_INCREMENT,
    `user_id`     INT(20) NOT NULL,
    `category_id` INT(10) NOT NULL,
    `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `t_users` (`id`),
    FOREIGN KEY (`category_id`) REFERENCES `t_categories` (`id`),
    INDEX `idx_user_category` (`user_id`, `category_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
    comment '用户与分类关系表';

