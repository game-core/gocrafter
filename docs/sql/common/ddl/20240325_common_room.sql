CREATE TABLE common_room
(
    room_id VARCHAR(255) NOT NULL COMMENT "ルームID",
	host_user_id VARCHAR(255) NOT NULL COMMENT "作成したユーザーID",
	room_number INT NOT NULL COMMENT "ルーム番号",
	name VARCHAR(255) NOT NULL COMMENT "ルーム名",
	user_count INT NOT NULL COMMENT "ユーザー数",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(room_id),
	UNIQUE KEY(room_id),
	INDEX(host_user_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
