
CREATE TABLE login_reward_model (

	id BIGINT NOT NULL AUTO_INCREMENT,

	name VARCHAR(255) NOT NULL,

	event_name VARCHAR(255) NOT NULL,

	created_at TIMESTAMP NOT NULL,

	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY(id),
	INDEX(name),INDEX(event_name)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
