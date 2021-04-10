package models

/**
CREATE TABLE `cabs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `current_state` varchar(10) NOT NULL DEFAULT 'IDLE',
  `current_city_id` int unsigned NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `last_updated_by` int DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `current_state` (`current_state`),
  KEY `is_active` (`is_active`),
  KEY `current_city_id` (`current_city_id`),
  CONSTRAINT `cabs_ibfk_1` FOREIGN KEY (`current_city_id`) REFERENCES `cities` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `cabs_idle_duration` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `cab_id` int NOT NULL,
  `idle_start_time` timestamp NOT NULL,
  `idle_end_time` timestamp NOT NULL,
  `total_duration` float NOT NULL DEFAULT '0',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `cab_id` (`cab_id`),
  KEY `idle_start_time` (`idle_start_time`),
  KEY `idle_end_time` (`idle_end_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/

type Cab struct {
	Id            int    `db:"id"`
	CurrentState  string `db:"current_state"`
	CurrentCityId int    `db:"current_city_id"`
	IsActive      int    `db:"is_active"`
	LastUpdatedBy int    `db:"last_updated_by"`
	Created       string `db:"created"`
	Updated       string `db:"updated"`
}

type CabAudit struct {
	Id            int    `db:"id"`
	CurrentState  string `db:"current_state"`
	CurrentCityId int    `db:"current_city_id"`
	IsActive      int    `db:"is_active"`
	LastUpdatedBy int    `db:"last_updated_by"`
	Created       string `db:"created"`
	Updated       string `db:"updated"`
	AuditAddedAt  string `db:"audit_added_at"`
}

type CabIdleDuration struct {
	Id            int     `db:"id"`
	CabId         int     `db:"cab_id"`
	IdleStartTime string  `db:"idle_start_time"`
	IdleEndTime   string  `db:"idle_end_time"`
	TotalDuration float32 `db:"total_duration"`
	Created       string  `db:"created"`
	Updated       string  `db:"updated"`
}

const (
	CabsTableName             = "cabs"
	CabsAuditTableName        = "cabs_audit"
	CabsIdleDurationTableName = "cabs_idle_duration"
	CabIdleState              = "IDLE"
	CabOnTripState            = "ON_TRIP"
)
