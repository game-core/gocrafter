CREATE TABLE master_login_bonus_event
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	master_login_bonus_id BIGINT NOT NULL COMMENT "ログボーナスID",
	name VARCHAR(255) NOT NULL COMMENT "イベント名",
	reset_hour INT NOT NULL COMMENT "リセット時間",
	interval_hour INT NOT NULL COMMENT "間隔時間",
	repeat_setting INT NOT NULL COMMENT "繰り返し設定",
	start_at VARCHAR(255) NOT NULL COMMENT "開始日時",
	end_at VARCHAR(255) DEFAULT NULL COMMENT "終了日時",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(master_login_bonus_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
