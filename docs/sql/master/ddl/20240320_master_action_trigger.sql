CREATE TABLE master_action_trigger
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	name VARCHAR(255) NOT NULL COMMENT "アクショントリガー名",
	action_trigger_type INT NOT NULL COMMENT "アクショントリガータイプ",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(action_trigger_type)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
