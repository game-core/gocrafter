CREATE TABLE common_ranking_room
(
    master_ranking_id BIGINT NOT NULL COMMENT "マスターランキングID",
	room_id VARCHAR(255) NOT NULL COMMENT "ルームID",
	user_id VARCHAR(255) NOT NULL COMMENT "ユーザーID",
	score INT NOT NULL COMMENT "スコア",
	ranked_at TIMESTAMP NOT NULL COMMENT "ランク付け日時",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(master_ranking_id,room_id,user_id),
	UNIQUE KEY(master_ranking_id,room_id,user_id),
	INDEX(master_ranking_id,room_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
