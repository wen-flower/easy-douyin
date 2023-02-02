/*!40101 SET NAMES utf8 */;

# CREATE USER 'douyin'@'localhost' IDENTIFIED BY 'douyin';

# CREATE DATABASE douyin;

# 用户服务
CREATE TABLE IF NOT EXISTS `tb_user`
(
    `id`             BIGINT       NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    `uid`            BIGINT       NOT NULL COMMENT '用户ID',
    `username`       VARCHAR(32)  NOT NULL COMMENT '用户名',
    `password`       VARCHAR(128) NOT NULL COMMENT '用户密码',
    `follow_count`   BIGINT       NOT NULL COMMENT '关注总数',
    `follower_count` BIGINT       NOT NULL COMMENT '粉丝总数',
    `created_at`     DATETIME     NOT NULL COMMENT '注册时间',
    `updated_at`     DATETIME     NULL COMMENT '最后一次更新时间',
    `deleted_at`     DATETIME     NULL COMMENT '逻辑删除的时间',

    CONSTRAINT tb_user_uid_uindex UNIQUE KEY (`uid`) COMMENT '用户ID唯一',
    CONSTRAINT tb_user_username_uindex UNIQUE KEY (`username`) COMMENT '用户名唯一'
) comment '用户表';

CREATE TABLE IF NOT EXISTS `tb_follow`
(
    `id`            BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    `uid`           BIGINT   NOT NULL COMMENT '用户ID',
    `followed_user` BIGINT   NOT NULL COMMENT '关注的用户',
    `status`        TINYINT  NOT NULL COMMENT '是否关注（1：关注，0：取消关注）',
    `created_at`    DATETIME NOT NULL COMMENT '第一次关注时间',
    `updated_at`    DATETIME NULL COMMENT '最后一次更新时间',

    CONSTRAINT tb_follow_uid_followed_user_uindex UNIQUE KEY (`uid`, `followed_user`) COMMENT '用户关注同一个人数据唯一'
) comment '用户关注表';

# 当用户关注数据量大到需要根据用户UID水平分表时，反查用户的粉丝时，需要对多个数据多个数据表进行查询
CREATE TABLE IF NOT EXISTS `tb_follower`
(
    `id`            BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    `uid`           BIGINT   NOT NULL COMMENT '用户ID',
    `follower_user` BIGINT   NOT NULL COMMENT '粉丝ID',
    `status`        TINYINT  NOT NULL COMMENT '是否关注（1：关注，0：取消关注）',
    `created_at`    DATETIME NOT NULL COMMENT '第一次关注时间',
    `updated_at`    DATETIME NULL COMMENT '最后一次更新时间',

    CONSTRAINT tb_follower_uid_follower_user_uindex UNIQUE KEY (`uid`, `follower_user`) COMMENT '用户同一个人粉丝数据唯一'
) comment '用户粉丝表';


# 视频服务
CREATE TABLE IF NOT EXISTS `tb_video`
(
    `id`             BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    `vid`            BIGINT      NOT NULL COMMENT '视频ID',
    `uid`            BIGINT      NOT NULL COMMENT '用户ID',
    `title`          VARCHAR(64) NOT NULL COMMENT '视频标题',
    `favorite_count` BIGINT      NOT NULL COMMENT '视频点赞数',
    `comment_count`  BIGINT      NOT NULL COMMENT '视频评论数',
    `created_at`     DATETIME    NOT NULL COMMENT '视频发布时间',
    `updated_at`     DATETIME    NULL COMMENT '最后一次更新时间',
    `deleted_at`     DATETIME    NULL COMMENT '逻辑删除的时间',

    CONSTRAINT tb_video_vid_uindex UNIQUE KEY (`vid`) COMMENT '视频ID唯一'
) comment '视频投稿表';

CREATE TABLE IF NOT EXISTS `tb_favorite`
(
    `id`         BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    `vid`        BIGINT   NOT NULL COMMENT '视频ID',
    `uid`        BIGINT   NOT NULL COMMENT '用户ID',
    `status`     TINYINT  NOT NULL COMMENT '是否点赞（1：已点赞，0：未点赞）',
    `created_at` DATETIME NOT NULL COMMENT '第一次点赞时间',
    `updated_at` DATETIME NULL COMMENT '最后一次更新时间',

    CONSTRAINT tb_favorite_vid_uid_uindex UNIQUE KEY (uid) COMMENT '用户点赞同一个视频数据唯一'
) comment '视频点赞表';

CREATE TABLE IF NOT EXISTS `tb_comment`
(
    `id`         BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    `vid`        BIGINT       NOT NULL COMMENT '视频ID',
    `uid`        BIGINT       NOT NULL COMMENT '用户ID',
    `content`    VARCHAR(512) NOT NULL COMMENT '评论内容',
    `created_at` DATETIME     NOT NULL COMMENT '评论时间',
    `updated_at` DATETIME     NULL COMMENT '最后一次更新时间',
    `deleted_at` DATETIME     NULL COMMENT '逻辑删除的时间'
) comment '用户评论表';


# 聊天服务
# 聊天记录表