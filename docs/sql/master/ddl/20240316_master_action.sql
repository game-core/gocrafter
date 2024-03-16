CREATE TABLE master_action
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	name VARCHAR(255) NOT NULL COMMENT "アクション名",
	action_type VARCHAR(255) NOT NULL COMMENT "アクションタイプ",
	any_id BIGINT NOT NULL COMMENT "対象のID",
	trigger_action_id BIGINT DEFAULT NULL COMMENT "トリガーになるアクションのID",
	next_action_id BIGINT DEFAULT NULL COMMENT "発火するアクションのID",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(name),
	INDEX(any_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
