-- 確認用
CREATE TABLE `example` (
    `id`           BIGINT       NOT NULL AUTO_INCREMENT COMMENT "ID",
    `example_key`  VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "KEY",
    `example_name` VARCHAR(50)  NOT NULL                COMMENT "名前",
    `deleted_at`   TIMESTAMP    DEFAULT NULL            COMMENT "削除日時",
    `created_at`   TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`   TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
