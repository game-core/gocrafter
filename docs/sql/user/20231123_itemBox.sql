
CREATE TABLE item_box (

	id BIGINT NOT NULL AUTO_INCREMENT,

	shard_key INT NOT NULL,

	account_id BIGINT NOT NULL,

	item_name VARCHAR(255) NOT NULL,

	count INT NOT NULL,

	created_at TIMESTAMP NOT NULL,

	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY(id),
	INDEX(account_id),INDEX(item_name)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
