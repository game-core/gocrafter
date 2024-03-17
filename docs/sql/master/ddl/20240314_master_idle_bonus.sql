CREATE TABLE master_idle_bonus
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	master_idle_bonus_event_id BIGINT NOT NULL COMMENT "放置ボーナスイベントID",
	name VARCHAR(255) NOT NULL COMMENT "放置ボーナス名",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(master_idle_bonus_event_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
