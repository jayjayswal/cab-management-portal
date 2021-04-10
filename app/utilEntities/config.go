package utilEntities

import (
	"cab-management-portal/app/configs"
	"cab-management-portal/app/constants"
)

type ConfigHelper struct {
	SequelConfig configs.SequelConfig
}

func GetConfigHelper(environment *Environment) *ConfigHelper {
	var sequelConfig configs.SequelConfig
	if environment.Tier == constants.DevTier {
		sequelConfig = configs.SequelConfig{
			Host:     "127.0.0.1:3306",
			Username: "root",
			Password: "mysqlroot",
			Database: "test",
		}
	} else if environment.Tier == constants.ProdTier {
		sequelConfig = configs.SequelConfig{
			Host:     "127.0.0.1:3306",
			Username: "root",
			Password: "mysqlroot",
			Database: "test",
		}
	}
	return &ConfigHelper{
		SequelConfig: sequelConfig,
	}
}