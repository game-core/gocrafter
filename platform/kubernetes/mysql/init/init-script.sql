CREATE DATABASE IF NOT EXISTS `gocrafter_common` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
GRANT ALL ON gocrafter_common.* TO 'mysql_user'@'%';

CREATE DATABASE IF NOT EXISTS `gocrafter_master` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
GRANT ALL ON gocrafter_master.* TO 'mysql_user'@'%';

CREATE DATABASE IF NOT EXISTS `gocrafter_user_0` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
GRANT ALL ON gocrafter_user_0.* TO 'mysql_user'@'%';

CREATE DATABASE IF NOT EXISTS `gocrafter_user_1` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
GRANT ALL ON gocrafter_user_1.* TO 'mysql_user'@'%';
