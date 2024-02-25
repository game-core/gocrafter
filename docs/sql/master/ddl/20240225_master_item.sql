CREATE TABLE master_item
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	name VARCHAR(255) NOT NULL COMMENT "アイテム名",
	resource_type INT NOT NULL COMMENT "リソースタイプ",
	rarity_type INT NOT NULL COMMENT "レアリティタイプ",
	content VARCHAR(255) NOT NULL COMMENT "コンテンツ",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(name)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
