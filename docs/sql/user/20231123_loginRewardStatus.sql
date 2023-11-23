
CREATE TABLE loginreward (

	id BIGINT NOT NULL AUTO_INCREMENT,

	shard_key INT NOT NULL,

	user_id BIGINT NOT NULL,

	login_reward_model_Name VARCHAR(255) NOT NULL,

	created_at TIMESTAMP NOT NULL,

	created_at TIMESTAMP NOT NULL,

	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY(id),
	INDEX(user_id),INDEX(login_reward_model_Name)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
