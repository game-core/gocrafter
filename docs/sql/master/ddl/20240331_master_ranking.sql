CREATE TABLE master_ranking
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	master_ranking_event_id BIGINT NOT NULL COMMENT "ランキングID",
	name VARCHAR(255) NOT NULL COMMENT "ランキング名",
	ranking_scope_type INT NOT NULL COMMENT "ランキング範囲タイプ",
	limit INT NOT NULL COMMENT "ランキング上限",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(master_ranking_event_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
