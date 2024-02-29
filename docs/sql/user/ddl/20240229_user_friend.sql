CREATE TABLE user_friend
(
    user_id VARCHAR(255) NOT NULL COMMENT "ユーザーID",
	friend_user_id VARCHAR(255) NOT NULL COMMENT "フレンドユーザーID",
	friend_type INT NOT NULL COMMENT "フレンドタイプ",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(user_id,friend_user_id),
	UNIQUE KEY(user_id,friend_user_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
