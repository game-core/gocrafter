CREATE TABLE master_action_step
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	name VARCHAR(255) NOT NULL COMMENT "アクションステップ名",
	action_step_type INT NOT NULL COMMENT "アクションステップタイプ",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(action_step_type)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
