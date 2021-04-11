package servies

import (
	"cab-management-portal/app/models"
	"cab-management-portal/app/utilEntities"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Service struct {
	Sequel  *sqlx.DB
	logger *log.Logger
}

func GetServiceObject(dependencies *utilEntities.Dependencies) (Services, error) {
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
	return &Service{
		Sequel: db,
		logger: dependencies.Logger,
	}, nil

}

type Services interface {
	CreateCab(ctx context.Context, cab *models.Cab) error
	GetCab(ctx context.Context, id int) (*models.Cab, error)
	GetCabActivities(ctx context.Context, id int) ([]models.CabAudit, error)
	GetCabForUpdate(ctx context.Context, id int, tx *sqlx.Tx) (*models.Cab, error)
	GetMostIdleCabOfCity(ctx context.Context, cityId int, tx *sqlx.Tx) ([]models.Cab, error)
	UpdateCabCity(ctx context.Context, cab *models.Cab, tx *sqlx.Tx) error
	UpdateCabCityTxn(ctx context.Context, CabId int, CurrentCityId int) error
	UpdateCabState(ctx context.Context, cab *models.Cab, tx *sqlx.Tx) error
	UpdateCab(ctx context.Context, cab *models.Cab, tx *sqlx.Tx) error
	GetAllCabs(ctx context.Context) ([]models.Cab, error)

	CreateCity(ctx context.Context, city *models.City) error
	GetCity(ctx context.Context, id int) (*models.City, error)
	GetAllCities(ctx context.Context) ([]models.City, error)

	UpdateRide(ctx context.Context, ride *models.Ride, tx *sqlx.Tx) error
	CreateRideRequest(ctx context.Context, rideRequest *models.RideRequest) error
	GetCityWiseRideInsight(ctx context.Context) ([]RideInsight, error)
	BookCabTxn(ctx context.Context, cityId int) (*models.Cab, *models.Ride, error)
	FinishRideTxn(ctx context.Context, rideId int) (*models.Cab, *models.Ride, error)
	GetAllRides(ctx context.Context) ([]models.Ride, error)
}