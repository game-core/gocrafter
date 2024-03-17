CREATE TABLE master_idle_bonus_item
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	master_idle_bonus_schedule_id BIGINT NOT NULL COMMENT "放置ボーナススケジュールID",
	master_item_id BIGINT NOT NULL COMMENT "アイテムID",
	name VARCHAR(255) NOT NULL COMMENT "放置ボーナスアイテム名",
	count INT NOT NULL COMMENT "個数",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(master_idle_bonus_schedule_id),
	INDEX(master_item_id),
	INDEX(master_idle_bonus_schedule_id,master_item_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
