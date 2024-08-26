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

type DriverLicenceRepo struct {
	db *sql.DB
}

func NewDriverLicenceRepo(db *sql.DB) *DriverLicenceRepo {
	return &DriverLicenceRepo{db: db}
}

func (r *DriverLicenceRepo) Create(req *mp.DriverLicenceCreateReq) (*mp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO driver_licences (
		id,
		first_name,
		last_name,
		third_name,
		address,
		got_date,
		exp_date,
		category,
		got_address,
		licence_number
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.Exec(query, id, req.FirstName, req.LastName, req.ThirdName, req.Address, req.GotDate, req.ExpDate, req.Category, req.GotAddress, req.LicenceNumber)
	if err != nil {
		log.Println("Error while creating driver licence: ", err)
		return nil, err
	}

	log.Println("Successfully created driver licence")
	return nil, nil
}

func (r *DriverLicenceRepo) GetById(req *mp.ById) (*mp.DriverLicenceGetByIdRes, error) {
	licence := mp.DriverLicenceGetByIdRes{
		Licence: &mp.DriverLicenceRes{},
	}

	query := `
	SELECT 
		id,
		first_name,
		last_name,
		third_name,
		address,
		to_char(got_date, 'YYYY-MM-DD'),
		to_char(exp_date, 'YYYY-MM-DD'),
		category,
		got_address,
		licence_number
	FROM 
		driver_licences
	WHERE 
		id = $1 AND deleted_at = 0
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&licence.Licence.Id,
		&licence.Licence.FirstName,
		&licence.Licence.LastName,
		&licence.Licence.ThirdName,
		&licence.Licence.Address,
		&licence.Licence.GotDate,
		&licence.Licence.ExpDate,
		&licence.Licence.Category,
		&licence.Licence.GotAddress,
		&licence.Licence.LicenceNumber,
	)

	if err != nil {
		log.Println("Error while getting driver licence by id: ", err)
		return nil, err
	}

	log.Println("Successfully got driver licence by id")
	return &licence, nil
}

func (r *DriverLicenceRepo) GetAll(req *mp.DriverLicenceGetAllReq) (*mp.DriverLicenceGetAllRes, error) {
	licences := mp.DriverLicenceGetAllRes{}

	query := `
	SELECT 
		id,
		first_name,
		last_name,
		third_name,
		address,
		to_char(got_date, 'YYYY-MM-DD'),
		to_char(exp_date, 'YYYY-MM-DD'),
		category,
		got_address,
		licence_number
	FROM 
		driver_licences
	WHERE 
		deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.Category != "" {
		conditions = append(conditions, " LOWER(category) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%"+req.Category+"%")
	}
	if req.ExpAt != "" {
		conditions = append(conditions, " exp_date <= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.ExpAt)
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
		log.Println("Driver licences not found")
		return nil, errors.New("driver licences list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving driver licences: ", err)
		return nil, err
	}

	for rows.Next() {
		licence := mp.DriverLicenceRes{}

		err := rows.Scan(
			&licence.Id,
			&licence.FirstName,
			&licence.LastName,
			&licence.ThirdName,
			&licence.Address,
			&licence.GotDate,
			&licence.ExpDate,
			&licence.Category,
			&licence.GotAddress,
			&licence.LicenceNumber,
		)

		if err != nil {
			log.Println("Error while scanning driver licences: ", err)
			return nil, err
		}

		licences.Licences = append(licences.Licences, &licence)
	}

	log.Println("Successfully fetched all driver licences")
	return &licences, nil
}

func (r *DriverLicenceRepo) Delete(req *mp.ById) (*mp.Void, error) {
	query := `
	UPDATE 
		driver_licences
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting driver licence: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("driver licence with id %s not found", req.Id)
	}

	log.Println("Successfully deleted driver licence")

	return nil, nil
}
