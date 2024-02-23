CREATE TABLE common_shard
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
	shard_key VARCHAR(255) NOT NULL COMMENT "シャードキー",
	name VARCHAR(255) NOT NULL COMMENT "シャード名",
	count INT NOT NULL COMMENT "シャードされたユーザー数",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(id)
	UNIQUE KEY(id)
	INDEX(shard_key)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
