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
create table t_transactions
(
    id               int auto_increment
        primary key,
    user_id          int                                  not null,
    category_id      int                                  not null,
    type             tinyint(1) default 1                 null comment '1支出｜2收入',
    title            varchar(255)                         not null comment '记账标题',
    description      text                                 null comment '记账描述',
    amount           decimal(10, 2)                       not null,
    transaction_date datetime                             not null,
    create_time      timestamp  default CURRENT_TIMESTAMP null,
    update_time      timestamp  default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint t_transactions_ibfk_1
        foreign key (user_id) references t_users (id),
    constraint t_transactions_ibfk_2
        foreign key (category_id) references t_categories (id)
)
    comment '记账记录表' collate = utf8mb4_general_ci;

create index idx_user_category_type
    on t_transactions (user_id, category_id, type);


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

