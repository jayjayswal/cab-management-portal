package servies

import (
	"cab-management-portal/app/utilEntities"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Services struct {
	Sequel  *sqlx.DB
	//utils   common.UtilHelpers
	logger *log.Logger
}

func GetServiceObject(dependencies *utilEntities.Dependencies) (*Services, error) {
	sequelConfig := dependencies.ConfigHelper.SequelConfig
	dependencies.Logger.Print("Trying to connect to mysql.....")
	db, err := sqlx.Connect("mysql", fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		sequelConfig.Username,
		sequelConfig.Password,
		sequelConfig.Host,
		sequelConfig.Database,
	))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	dependencies.Logger.Print("Mysql connection successful.")
	return &Services{
		Sequel: db,
		logger: dependencies.Logger,
	}, nil

}