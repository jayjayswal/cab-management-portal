package servies

import (
	"context"
)

func (s *Service) CreateSchema(ctx context.Context) {
	var schema = `
CREATE TABLE cabs (
  id int unsigned NOT NULL AUTO_INCREMENT,
  cab_number varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  current_state varchar(10) NOT NULL DEFAULT 'IDLE',
  current_city_id int unsigned DEFAULT NULL,
  is_active tinyint(1) NOT NULL DEFAULT '1',
  last_ride_end_time timestamp NULL DEFAULT NULL,
  created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY cab_number (cab_number),
  KEY current_state (current_state),
  KEY is_active (is_active),
  KEY current_city_id (current_city_id),
  CONSTRAINT cabs_ibfk_1 FOREIGN KEY (current_city_id) REFERENCES cities (id) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE cabs_audit (
  id int unsigned NOT NULL,
  cab_number varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  current_state varchar(10) NOT NULL DEFAULT 'IDLE',
  current_city_id int unsigned DEFAULT NULL,
  is_active tinyint(1) NOT NULL DEFAULT '1',
  last_ride_end_time timestamp NULL DEFAULT NULL,
  created timestamp NOT NULL,
  updated timestamp NOT NULL,
  audit_added_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY id (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE cabs_idle_duration (
  id int unsigned NOT NULL AUTO_INCREMENT,
  cab_id int NOT NULL,
  idle_start_time timestamp NOT NULL,
  idle_end_time timestamp NOT NULL,
  total_duration float NOT NULL DEFAULT '0',
  created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY cab_id (cab_id),
  KEY idle_start_time (idle_start_time),
  KEY idle_end_time (idle_end_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE cities (
  id int unsigned NOT NULL AUTO_INCREMENT,
  name varchar(100) NOT NULL DEFAULT '',
  is_active tinyint(1) NOT NULL DEFAULT '1',
  created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY name (name),
  KEY is_active (is_active)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE ride_requests (
  id int unsigned NOT NULL AUTO_INCREMENT,
  start_city_id int NOT NULL,
  current_state varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

	`
	s.Sequel.MustExec(schema)

}
