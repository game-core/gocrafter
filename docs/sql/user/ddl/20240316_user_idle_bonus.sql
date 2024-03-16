CREATE TABLE user_idle_bonus
(
    user_id VARCHAR(255) NOT NULL COMMENT "ユーザーID",
	master_idle_bonus_id BIGINT NOT NULL COMMENT "放置ボーナスID",
	received_at TIMESTAMP NOT NULL COMMENT "受け取り日時",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(user_id,master_idle_bonus_id),
	UNIQUE KEY(user_id,master_idle_bonus_id),
	INDEX(user_id),
	INDEX(user_id,master_idle_bonus_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
