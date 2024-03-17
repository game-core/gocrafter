CREATE TABLE master_action_run
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	name VARCHAR(255) NOT NULL COMMENT "実行されるアクション名",
	action_id BIGINT NOT NULL COMMENT "アクションID",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(name),
	INDEX(action_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
