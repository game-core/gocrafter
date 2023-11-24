
CREATE TABLE login_reward_status (

	id BIGINT NOT NULL AUTO_INCREMENT,

	shard_key VARCHAR(255) NOT NULL,

	account_id BIGINT NOT NULL,

	login_reward_model_Name VARCHAR(255) NOT NULL,

	last_received_at TIMESTAMP NOT NULL,

	created_at TIMESTAMP NOT NULL,

	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY(id),
	INDEX(account_id),INDEX(login_reward_model_Name)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
