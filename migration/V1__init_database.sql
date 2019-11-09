/*
  Table Name : accounts
 */
CREATE TABLE IF NOT EXISTS `accounts`
(
    `id`           int(11)      NOT NULL,
    `name`         varchar(255) NOT NULL,
    `create_time`  datetime,
    `update_time`  datetime,
    `deleted_time` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

/*
  Table Name : books
 */
CREATE TABLE IF NOT EXISTS `books`
(
    `id`           int(11)      NOT NULL,
    `title`        varchar(255) NOT NULL,
    `isbn_13`      varchar(13) DEFAULT NULL,
    `account_id`   int(11)      NOT NULL,
    `create_time`  datetime,
    `update_time`  datetime,
    `deleted_time` datetime    DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;