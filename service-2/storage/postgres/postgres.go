package postgres

import (
	"database/sql"
	"fmt"
	"mobile/config"
	"mobile/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db             *sql.DB
	CarTypeS       storage.CarTypeI
	CarS           storage.CarI
	DriverLicenceS storage.DriverLicenceI
	FineS          storage.FineI
	ServiceTypeS   storage.ServiceTypeI
	ServiceS       storage.ServiceI
	CarServiceS    storage.CarServiceI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	car_service := NewCarServiceRepo(db)
	car_type := NewCarTypeRepo(db)
	car := NewCarRepo(db)
	driver_licence := NewDriverLicenceRepo(db)
	fine := NewFineRepo(db)
	service_type := NewServiceTypeRepo(db)
	service := NewServiceRepo(db)

	return &Storage{
		Db:             db,
		CarServiceS:    car_service,
		CarTypeS:       car_type,
		CarS:           car,
		DriverLicenceS: driver_licence,
		ServiceS:       service,
		ServiceTypeS:   service_type,
		FineS:          fine,
	}, nil
}
