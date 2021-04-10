package models

import "time"

/**
CREATE TABLE `admin` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(25) NOT NULL DEFAULT '',
  `email` varchar(300) NOT NULL DEFAULT '',
  `is_active` tinyint NOT NULL DEFAULT '0',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `is_active` (`is_active`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `auth_token` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` int unsigned NOT NULL,
  `auth_token` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `is_active` tinyint(1) NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_token` (`auth_token`),
  KEY `admin_id_fk` (`admin_id`),
  KEY `is_active` (`is_active`),
  CONSTRAINT `admin_id_fk` FOREIGN KEY (`admin_id`) REFERENCES `admin` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/

type Admin struct {
	Id       int       `db:"id"`
	Name     string    `db:"name"`
	Email    string    `db:"email"`
	IsActive int       `db:"is_active"`
	Created  time.Time `db:"created"`
	Updated  time.Time `db:"updated"`
}

type AuthToken struct {
	Id        int       `db:"id"`
	AdminId   int       `db:"admin_id"`
	AuthToken string    `db:"auth_token"`
	IsActive  int       `db:"is_active"`
	Created   time.Time `db:"created"`
	Updated   time.Time `db:"updated"`
}

const (
	AdminTableName     = "admin"
	AuthTokenTableName = "auth_token"
)
