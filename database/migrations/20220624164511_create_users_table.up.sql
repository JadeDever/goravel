CREATE TABLE users (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  username varchar(20) NOT NULL,
  name varchar(20) NOT NULL,
  password varchar(20) NOT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  created_at datetime(3) DEFAULT NULL,
  updated_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_users_created_at (created_at),
  KEY idx_users_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
