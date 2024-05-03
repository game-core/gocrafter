CREATE TABLE user_account_token
(
    user_id VARCHAR(255) NOT NULL COMMENT "ユーザーID",
	token VARCHAR(255) NOT NULL COMMENT "トークン",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(user_id),
	UNIQUE KEY(user_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
