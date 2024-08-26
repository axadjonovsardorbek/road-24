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

type ServiceRepo struct {
	db *sql.DB
}

func NewServiceRepo(db *sql.DB) *ServiceRepo {
	return &ServiceRepo{db: db}
}

func (r *ServiceRepo) Create(req *mp.ServiceCreateReq) (*mp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO services (
		id,
		type,
		sertificat_number,
		owner_name,
		address,
		phone_number,
		name
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(query, id, req.Type, req.SertificatNumber, req.OwnerName, req.Address, req.PhoneNumber, req.Name)
	if err != nil {
		log.Println("Error while creating service: ", err)
		return nil, err
	}

	log.Println("Successfully created service")
	return nil, nil
}

func (r *ServiceRepo) GetById(req *mp.ById) (*mp.ServiceGetByIdRes, error) {
	service := mp.ServiceGetByIdRes{
		Service: &mp.ServiceRes{
			Type: &mp.ServiceTypeRes{},
		},
	}

	query := `
	SELECT 
		s.id,
		t.id,
		t.name,
		s.name,
		s.sertificat_number,
		s.owner_name,
		s.address,
		s.phone_number
	FROM 
		services s
	JOIN 
		service_types t 
	ON 
		t.name = s.type AND t.deleted_at = 0
	WHERE 
		s.deleted_at = 0 AND s.id = $1
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&service.Service.Id,
		&service.Service.Type.Id,
		&service.Service.Type.Name,
		&service.Service.Name,
		&service.Service.SertificatNumber,
		&service.Service.OwnerName,
		&service.Service.Address,
		&service.Service.PhoneNumber,
	)

	if err != nil {
		log.Println("Error while getting service by id: ", err)
		return nil, err
	}

	log.Println("Successfully got service by id")
	return &service, nil
}

func (r *ServiceRepo) GetAll(req *mp.ServiceGetAllReq) (*mp.ServiceGetAllRes, error) {
	services := mp.ServiceGetAllRes{}

	query := `
	SELECT 
		s.id,
		t.id,
		t.name,
		s.name,
		s.sertificat_number,
		s.owner_name,
		s.address,
		s.phone_number
	FROM 
		services s
	JOIN 
		service_types t 
	ON 
		t.name = s.type AND t.deleted_at = 0
	WHERE 
		s.deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.Address != "" && req.Address != "string" {
		conditions = append(conditions, " LOWER(s.address) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%"+req.Address+"%")
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
		log.Println("Services not found")
		return nil, errors.New("services list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving services: ", err)
		return nil, err
	}

	for rows.Next() {
		service := mp.ServiceRes{
			Type: &mp.ServiceTypeRes{},
		}

		err := rows.Scan(
			&service.Id,
			&service.Type.Id,
			&service.Type.Name,
			&service.Name,
			&service.SertificatNumber,
			&service.OwnerName,
			&service.Address,
			&service.PhoneNumber,
		)

		if err != nil {
			log.Println("Error while scanning services: ", err)
			return nil, err
		}

		services.Services = append(services.Services, &service)
	}

	log.Println("Successfully fetched all services")
	return &services, nil
}

func (r *ServiceRepo) Delete(req *mp.ById) (*mp.Void, error) {
	query := `
	UPDATE 
		services
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting service: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service with id %s not found", req.Id)
	}

	log.Println("Successfully deleted service")

	return nil, nil
}
