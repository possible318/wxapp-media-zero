CREATE TABLE `blog`
(
    `id`        int unsigned                                                  NOT NULL AUTO_INCREMENT,
    `platform`  varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '' COMMENT '来源',
    `item_id`   varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '' COMMENT '组ID',
    `text`      text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci         NOT NULL DEFAULT '' COMMENT '文案',
    `pid`       varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '' COMMENT '图片标识',
    `src`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '地址',
    `index`     int                                                           NOT NULL DEFAULT '0' COMMENT '是否首页',
    `show_type` int                                                           NOT NULL DEFAULT '0' COMMENT '展示类型',
    `deleted`   int                                                           NOT NULL DEFAULT '0' COMMENT '是否删除',
    `add_time`  varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '' COMMENT '发布时间',
    `width`     int                                                           NOT NULL DEFAULT '0' COMMENT '宽度',
    `height`    int                                                           NOT NULL DEFAULT '0' COMMENT '高度',
    `status`    int                                                           NOT NULL DEFAULT '0' COMMENT '状态',
    `like`      int                                                           NOT NULL DEFAULT '0' COMMENT '点赞',
    `dislike`   int                                                           NOT NULL DEFAULT '0' COMMENT '点踩',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_idx_pf_pid` (`platform`, `pid`),
    KEY `idx_like` (`like`),
    KEY `idx_item` (`item_id`),
    KEY `idx_dislike` (`dislike`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;