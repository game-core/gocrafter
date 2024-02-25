CREATE TABLE master_rarity
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	name VARCHAR(255) NOT NULL COMMENT "レアリティ名",
	rarity_type INT NOT NULL COMMENT "レアリティタイプ",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(rarity_type)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
