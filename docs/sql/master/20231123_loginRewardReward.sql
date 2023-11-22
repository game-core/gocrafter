
CREATE TABLE loginreward (

	id BIGINT NOT NULL AUTO_INCREMENT,

	login_reward_id BIGINT NOT NULL,

	item_id BIGINT NOT NULL,

	name VARCHAR(255) NOT NULL,

	created_at TIMESTAMP NOT NULL,

	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY(id),
	INDEX(name),INDEX(login_reward_id),INDEX(item_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
