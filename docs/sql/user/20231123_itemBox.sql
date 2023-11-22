
CREATE TABLE item (

	id BIGINT NOT NULL AUTO_INCREMENT,

	shard_key INT NOT NULL,

	item_id BIGINT NOT NULL,

	item_id BIGINT NOT NULL,

	count INT NOT NULL,

	created_at TIMESTAMP NOT NULL,

	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY(id),
	INDEX(item_id),INDEX(item_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
