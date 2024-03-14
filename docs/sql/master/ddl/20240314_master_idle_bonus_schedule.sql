CREATE TABLE master_idle_bonus_schedule
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	master_idle_bonus_id BIGINT NOT NULL COMMENT "放置ボーナスID",
	step INT NOT NULL COMMENT "ステップ",
	name VARCHAR(255) NOT NULL COMMENT "放置ボーナススケジュール名",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(master_idle_bonus_id),
	INDEX(step),
	INDEX(master_idle_bonus_id,step)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
