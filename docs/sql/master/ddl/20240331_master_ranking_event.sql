CREATE TABLE master_ranking_event
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	name VARCHAR(255) NOT NULL COMMENT "イベント名",
	reset_hour INT NOT NULL COMMENT "リセット時間",
	interval_hour INT NOT NULL COMMENT "間隔時間",
	repeat_setting TINYINT NOT NULL COMMENT "繰り返し設定",
	start_at TIMESTAMP NOT NULL COMMENT "開始日時",
	end_at TIMESTAMP DEFAULT NULL COMMENT "終了日時",
	PRIMARY KEY(id),
	UNIQUE KEY(id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
