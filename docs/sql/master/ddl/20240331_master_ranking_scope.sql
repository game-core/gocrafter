CREATE TABLE master_ranking_scope
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	name VARCHAR(255) NOT NULL COMMENT "ランキング範囲名",
	ranking_scope_type INT NOT NULL COMMENT "ランキング範囲タイプ",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(ranking_scope_type)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
