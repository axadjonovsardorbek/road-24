package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	mp "mobile/genproto/mobile"

	"github.com/google/uuid"
)

type ServiceTypeRepo struct {
	db *sql.DB
}

func NewServiceTypeRepo(db *sql.DB) *ServiceTypeRepo {
	return &ServiceTypeRepo{db: db}
}

func (r *ServiceTypeRepo) Create(req *mp.ServiceTypeCreateReq) (*mp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO service_types (
		id,
		name
	) VALUES ($1, $2)
	`

	_, err := r.db.Exec(query, id, req.Name)
	if err != nil {
		log.Println("Error while creating service type: ", err)
		return nil, err
	}

	log.Println("Successfully created service type")
	return nil, nil
}

func (r *ServiceTypeRepo) GetById(req *mp.ById) (*mp.ServiceTypeGetByIdRes, error) {
	serviceType := mp.ServiceTypeGetByIdRes{
		Service: &mp.ServiceTypeRes{},
	}

	query := `
	SELECT 
		id,
		name
	FROM 
		service_types
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&serviceType.Service.Id,
		&serviceType.Service.Name,
	)

	if err != nil {
		log.Println("Error while getting service type by id: ", err)
		return nil, err
	}

	log.Println("Successfully got service type by id")
	return &serviceType, nil
}

func (r *ServiceTypeRepo) GetAll(req *mp.ServiceTypeGetAllReq) (*mp.ServiceTypeGetAllRes, error) {
	serviceTypes := mp.ServiceTypeGetAllRes{}

	query := `
	SELECT 
		id,
		name
	FROM 
		service_types
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
		log.Println("Service types not found")
		return nil, errors.New("service types list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving service types: ", err)
		return nil, err
	}

	for rows.Next() {
		serviceType := mp.ServiceTypeRes{}

		err := rows.Scan(
			&serviceType.Id,
			&serviceType.Name,
		)

		if err != nil {
			log.Println("Error while scanning service types: ", err)
			return nil, err
		}

		serviceTypes.Services = append(serviceTypes.Services, &serviceType)
	}

	log.Println("Successfully fetched all service types")
	return &serviceTypes, nil
}

func (r *ServiceTypeRepo) Delete(req *mp.ById) (*mp.Void, error) {
	query := `
	UPDATE 
		service_types
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting service type: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service type with id %s not found", req.Id)
	}

	log.Println("Successfully deleted service type")

	return nil, nil
}
