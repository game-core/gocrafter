CREATE TABLE master_login_bonus
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	master_login_bonus_event_id BIGINT NOT NULL COMMENT "ログボーナスイベントID",
	name VARCHAR(255) NOT NULL COMMENT "ログインボーナス名",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(master_login_bonus_event_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
