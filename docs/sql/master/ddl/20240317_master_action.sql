CREATE TABLE master_action
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	name VARCHAR(255) NOT NULL COMMENT "アクション名",
	action_step_type INT NOT NULL COMMENT "アクションステップタイプ",
	action_trigger_type INT NOT NULL COMMENT "アクショントリガータイプ",
	any_id BIGINT DEFAULT NULL COMMENT "対象のID",
	trigger_action_id BIGINT DEFAULT NULL COMMENT "トリガーになるアクションのID",
	expiration INT DEFAULT NULL COMMENT "有効期限",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(name),
	INDEX(action_step_type)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
