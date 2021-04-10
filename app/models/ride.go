package models

/**
CREATE TABLE `rides` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `cab_id` int unsigned NOT NULL,
  `start_city_id` int unsigned NOT NULL,
  `start_time` timestamp NOT NULL,
  `end_time` timestamp NULL DEFAULT NULL,
  `current_state` varchar(20) NOT NULL DEFAULT '',
  `last_updated_by` int DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `cab_id` (`cab_id`),
  KEY `start_city_id` (`start_city_id`),
  KEY `start_time` (`start_time`),
  KEY `end_time` (`end_time`),
  KEY `current_state` (`current_state`),
  CONSTRAINT `rides_ibfk_1` FOREIGN KEY (`cab_id`) REFERENCES `cabs` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `rides_ibfk_2` FOREIGN KEY (`start_city_id`) REFERENCES `cities` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/

type Ride struct {
	Id            int     `db:"id"`
	CabId         int     `db:"cab_id"`
	StartCityId   int     `db:"start_city_id"`
	StartTime     string  `db:"start_time"`
	EndTime       *string `db:"end_time"`
	CurrentState  string  `db:"current_state"`
	LastUpdatedBy int     `db:"last_updated_by"`
	Created       string  `db:"created"`
	Updated       string  `db:"updated"`
}

const (
	RidesTableName = "rides"
)