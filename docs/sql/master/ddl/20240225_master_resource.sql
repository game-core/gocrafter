CREATE TABLE master_resource
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	name VARCHAR(255) NOT NULL COMMENT "リソース名",
	resource_type INT NOT NULL COMMENT "リソースタイプ",
	PRIMARY KEY(id),
	UNIQUE KEY(id),
	INDEX(resource_type)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
