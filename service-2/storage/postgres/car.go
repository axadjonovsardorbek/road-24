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

type CarRepo struct {
	db *sql.DB
}

func NewCarRepo(db *sql.DB) *CarRepo {
	return &CarRepo{db: db}
}

func (r *CarRepo) Create(req *mp.CarCreateReq) (*mp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO cars (
		id,
		staf_id,
		type,
		color,
		year,
		body_number,
		horse_power,
		number,
		technical_passport,
		model,
		image_url
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.Exec(query, id, req.StafId, req.TypeId, req.Color, req.Year, req.BodyNumber, req.HorsePower, req.Number, req.TechnicalPassport, req.Model, req.ImageUrl)
	if err != nil {
		log.Println("Error while creating car: ", err)
		return nil, err
	}

	log.Println("Successfully created car")
	return nil, nil
}

func (r *CarRepo) GetById(req *mp.ById) (*mp.CarGetByIdRes, error) {
	car := mp.CarGetByIdRes{
		Car: &mp.CarRes{},
	}

	query := `
	SELECT 
		id,
		staf_id,
		type,
		color,
		year,
		body_number,
		horse_power,
		number,
		technical_passport,
		model,
		image_url
	FROM 
		cars
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&car.Car.Id,
		&car.Car.StafId,
		&car.Car.TypeId,
		&car.Car.Color,
		&car.Car.Year,
		&car.Car.BodyNumber,
		&car.Car.HorsePower,
		&car.Car.Number,
		&car.Car.TechnicalPassport,
		&car.Car.Model,
		&car.Car.ImageUrl,
	)

	if err != nil {
		log.Println("Error while getting car by id: ", err)
		return nil, err
	}

	log.Println("Successfully got car by id")
	return &car, nil
}

func (r *CarRepo) GetAll(req *mp.CarGetAllReq) (*mp.CarGetAllRes, error) {
	cars := mp.CarGetAllRes{}

	query := `
	SELECT 
		id,
		staf_id,
		type,
		color,
		year,
		body_number,
		horse_power,
		number,
		technical_passport,
		model,
		image_url
	FROM 
		cars
	WHERE 
		deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.Type != "" && req.Type != "string" {
		conditions = append(conditions, " type = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Type)
	}
	if req.Model != "" && req.Model != "string" {
		conditions = append(conditions, " model = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Model)
	}
	if req.Year != "" && req.Year != "string" {
		conditions = append(conditions, " year = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Year)
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
		log.Println("Cars not found")
		return nil, errors.New("cars list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving cars: ", err)
		return nil, err
	}

	for rows.Next() {
		car := mp.CarRes{}

		err := rows.Scan(
			&car.Id,
			&car.StafId,
			&car.TypeId,
			&car.Color,
			&car.Year,
			&car.BodyNumber,
			&car.HorsePower,
			&car.Number,
			&car.TechnicalPassport,
			&car.Model,
			&car.ImageUrl,
		)

		if err != nil {
			log.Println("Error while scanning cars: ", err)
			return nil, err
		}

		cars.Cars = append(cars.Cars, &car)
	}

	log.Println("Successfully fetched all cars")
	return &cars, nil
}

func (r *CarRepo) Delete(req *mp.ById) (*mp.Void, error) {
	query := `
	UPDATE 
		cars
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting car: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("car with id %s not found", req.Id)
	}

	log.Println("Successfully deleted car")

	return nil, nil
}
