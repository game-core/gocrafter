
CREATE TABLE login_reward_reward (

	id BIGINT NOT NULL AUTO_INCREMENT,

	login_reward_model_name VARCHAR(255) NOT NULL,

	item_name VARCHAR(255) NOT NULL,

	name VARCHAR(255) NOT NULL,

	step_number INT NOT NULL,

	created_at TIMESTAMP NOT NULL,

	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY(id),
	INDEX(name),INDEX(login_reward_model_name),INDEX(item_name)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
