
CREATE TABLE account (

	id BIGINT NOT NULL AUTO_INCREMENT,

	uuid VARCHAR(255) NOT NULL,

	name VARCHAR(255) NOT NULL,

	password VARCHAR(255) NOT NULL,

	created_at TIMESTAMP NOT NULL,

	updated_at TIMESTAMP NOT NULL,

	PRIMARY KEY(id),
	INDEX(uuid)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;