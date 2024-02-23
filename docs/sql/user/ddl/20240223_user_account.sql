CREATE TABLE user_account
(
    user_id VARCHAR(255) NOT NULL COMMENT "ユーザーID",
	name VARCHAR(255) NOT NULL COMMENT "ユーザー名",
	password VARCHAR(255) NOT NULL COMMENT "パスワード",
	login_at TIMESTAMP NOT NULL COMMENT "ログイン日時",
	logout_at TIMESTAMP NOT NULL COMMENT "ログアウト日時",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(user_id)
	UNIQUE KEY(user_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
