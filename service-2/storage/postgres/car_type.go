package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	mp "mobile/genproto/mobile"

	"github.com/google/uuid"
)

type CarTypeRepo struct {
	db *sql.DB
}

func NewCarTypeRepo(db *sql.DB) *CarTypeRepo {
	return &CarTypeRepo{db: db}
}

func (r *CarTypeRepo) Create(req *mp.CarTypeCreateReq) (*mp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO car_types (
		id,
		name
	) VALUES ($1, $2)
	`

	_, err := r.db.Exec(query, id, req.Name)
	if err != nil {
		log.Println("Error while creating car type: ", err)
		return nil, err
	}

	log.Println("Successfully created car type")
	return nil, nil
}

func (r *CarTypeRepo) GetById(req *mp.ById) (*mp.CarTypeGetByIdRes, error) {
	car := mp.CarTypeGetByIdRes{
		Car: &mp.CarTypeRes{},
	}

	query := `
	SELECT 
		id,
		name
	FROM 
		car_types
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&car.Car.Id,
		&car.Car.Name,
	)

	if err != nil {
		log.Println("Error while getting car type by id: ", err)
		return nil, err
	}

	log.Println("Successfully got car type by id")
	return &car, nil
}

func (r *CarTypeRepo) GetAll(req *mp.CarTypeGetAllReq) (*mp.CarTypeGetAllRes, error) {
	cars := mp.CarTypeGetAllRes{}

	query := `
	SELECT 
		id,
		name
	FROM 
		car_types
	WHERE 
		deleted_at = 0
	`

	var args []interface{}

	var limit int32 = 10
	var offset int32 = (req.Filter.Page - 1) * limit

	args = append(args, limit, offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)
	if err == sql.ErrNoRows {
		log.Println("Car types not found")
		return nil, errors.New("car types list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving car types: ", err)
		return nil, err
	}

	for rows.Next() {
		car := mp.CarTypeRes{}

		err := rows.Scan(
			&car.Id,
			&car.Name,
		)

		if err != nil {
			log.Println("Error while scanning car types: ", err)
			return nil, err
		}

		cars.Cars = append(cars.Cars, &car)
	}

	log.Println("Successfully fetched all car types")
	return &cars, nil
}

func (r *CarTypeRepo) Delete(req *mp.ById) (*mp.Void, error) {
	query := `
	UPDATE 
		car_types
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting car type: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("car type with id %s not found", req.Id)
	}

	log.Println("Successfully deleted car type")

	return nil, nil
}
