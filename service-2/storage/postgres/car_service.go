package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	mp "mobile/genproto/mobile"

	"github.com/google/uuid"
)

type CarServiceRepo struct {
	db *sql.DB
}

func NewCarServiceRepo(db *sql.DB) *CarServiceRepo {
	return &CarServiceRepo{db: db}
}

func (r *CarServiceRepo) Create(req *mp.CarServiceCreateReq) (*mp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO car_services (
		id,
		car_number,
		service_type,
		service_date,
		exp_date
	) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(query, id, req.CarNumber, req.ServiceType, req.ServiceDate, req.ExpDate)
	if err != nil {
		log.Println("Error while creating car service: ", err)
		return nil, err
	}

	log.Println("Successfully created car service")
	return nil, nil
}

func (r *CarServiceRepo) GetById(req *mp.ById) (*mp.CarServiceGetByIdRes, error) {
	carService := mp.CarServiceGetByIdRes{
		Car: &mp.CarServiceRes{
			ServiceType: &mp.ServiceTypeRes{},
		},
	}

	query := `
	SELECT 
		cs.id,
		cs.car_number,
		t.id,
		t.name,
		to_char(cs.service_date, 'YYYY-MM-DD'),
		to_char(cs.exp_date, 'YYYY-MM-DD')
	FROM 
		car_services cs
	JOIN 
		service_types t
	ON
		cs.service_type = t.name AND t.deleted_at = 0
	WHERE 
		cs.deleted_at = 0 AND cs.id = $1
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&carService.Car.Id,
		&carService.Car.CarNumber,
		&carService.Car.ServiceType.Id,
		&carService.Car.ServiceType.Name,
		&carService.Car.ServiceDate,
		&carService.Car.ExpDate,
	)

	if err != nil {
		log.Println("Error while getting car service by id: ", err)
		return nil, err
	}

	log.Println("Successfully got car service by id")
	return &carService, nil
}

func (r *CarServiceRepo) GetAll(req *mp.CarServiceGetAllReq) (*mp.CarServiceGetAllRes, error) {
	carServices := mp.CarServiceGetAllRes{}

	query := `
	SELECT 
		cs.id,
		cs.car_number,
		t.id,
		t.name,
		to_char(cs.service_date, 'YYYY-MM-DD'),
		to_char(cs.exp_date, 'YYYY-MM-DD')
	FROM 
		car_services cs
	JOIN 
		service_types t
	ON
		cs.service_type = t.name AND t.deleted_at = 0
	WHERE 
		cs.deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.CarNumber != "" {
		conditions = append(conditions, " LOWER(cs.car_number) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%"+req.CarNumber+"%")
	}
	if req.ServiceType != "" {
		conditions = append(conditions, " LOWER(cs.service_type) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%"+req.ServiceType+"%")
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var limit int32 = 10
	var offset int32 = (req.Filter.Page - 1) * limit

	args = append(args, limit, offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)
	if err == sql.ErrNoRows {
		log.Println("Car services not found")
		return nil, errors.New("car services list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving car services: ", err)
		return nil, err
	}

	for rows.Next() {
		carService := mp.CarServiceRes{
			ServiceType: &mp.ServiceTypeRes{},
		}

		err := rows.Scan(
			&carService.Id,
			&carService.CarNumber,
			&carService.ServiceType.Id,
			&carService.ServiceType.Name,
			&carService.ServiceDate,
			&carService.ExpDate,
		)

		if err != nil {
			log.Println("Error while scanning car services: ", err)
			return nil, err
		}

		carServices.Cars = append(carServices.Cars, &carService)
	}

	log.Println("Successfully fetched all car services")
	return &carServices, nil
}

func (r *CarServiceRepo) Delete(req *mp.ById) (*mp.Void, error) {
	query := `
	UPDATE 
		car_services
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting car service: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("car service with id %s not found", req.Id)
	}

	log.Println("Successfully deleted car service")

	return nil, nil
}
