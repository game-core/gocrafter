CREATE TABLE user_action
(
    user_id VARCHAR(255) NOT NULL COMMENT "ユーザーID",
	name VARCHAR(255) NOT NULL COMMENT "ユーザー名",
	master_action_id BIGINT NOT NULL COMMENT "マスターアクションID",
	started_at TIMESTAMP NOT NULL COMMENT "開始日時",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(user_id,master_action_id),
	UNIQUE KEY(user_id,master_action_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
