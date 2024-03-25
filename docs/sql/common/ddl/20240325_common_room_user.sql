CREATE TABLE common_room_user
(
    room_id VARCHAR(255) NOT NULL COMMENT "ルームID",
	user_id VARCHAR(255) NOT NULL COMMENT "ユーザーID",
	room_user_position_type INT NOT NULL COMMENT "ルームユーザー立場タイプ",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(room_id,user_id),
	UNIQUE KEY(room_id,user_id),
	INDEX(room_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
