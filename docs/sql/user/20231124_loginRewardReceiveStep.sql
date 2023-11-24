
CREATE TABLE login_reward_receive_step (

	id BIGINT NOT NULL AUTO_INCREMENT,

	shard_key VARCHAR(255) NOT NULL,

	user_id BIGINT NOT NULL,

	login_reward_model_id BIGINT NOT NULL,

	step_number INT NOT NULL,

	created_at TIMESTAMP NOT NULL,

	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY(id),
	INDEX(user_id),INDEX(login_reward_model_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
