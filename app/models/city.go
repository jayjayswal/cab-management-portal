package models

/**
CREATE TABLE `cities` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `last_updated_by` int DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `is_active` (`is_active`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/

type City struct {
	Id            int    `db:"id"`
	Name          string `db:"name"`
	IsActive      int    `db:"is_active"`
	LastUpdatedBy int    `db:"last_updated_by"`
	Created       string `db:"created"`
	Updated       string `db:"updated"`
}

const (
	CitiesTableName = "cities"
)
