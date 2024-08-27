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

type FineRepo struct {
	db *sql.DB
}

func NewFineRepo(db *sql.DB) *FineRepo {
	return &FineRepo{db: db}
}

func (r *FineRepo) Create(req *mp.FineCreateReq) (*mp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO fines (
		id,
		staf_id,
		technical_number,
		car_number,
		licence_number,
		fine_date,
		fine_amount
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(query, id, req.StafId, req.TechnicalNumber, req.CarNumber, req.LicenceNumber, req.FineDate, req.FineAmount)
	if err != nil {
		log.Println("Error while creating fine: ", err)
		return nil, err
	}

	log.Println("Successfully created fine")
	return nil, nil
}

func (r *FineRepo) GetById(req *mp.ById) (*mp.FineGetByIdRes, error) {
	fine := mp.FineGetByIdRes{
		Fine: &mp.FineRes{
			StafId: &mp.StafRes{},
		},
	}

	query := `
	SELECT 
		f.id,
		u.id,
		u.first_name,
		u.last_name,
		f.technical_number,
		f.car_number,
		f.licence_number,
		f.fine_date,
		f.fine_amount,
		f.payment_date,
		CASE 
        	WHEN f.payment_date IS NULL THEN 'not yet paid' 
        	ELSE f.payment_date::text 
    	END as payment_status,
		f.status
	FROM 
		fines f
	JOIN 
		users u 
	ON 
		u.id = f.staf_id AND u.deleted_at = 0
	WHERE 
		f.deleted_at = 0 AND f.id = $1
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&fine.Fine.Id,
		&fine.Fine.StafId.Id,
		&fine.Fine.StafId.FirstName,
		&fine.Fine.StafId.LastName,
		&fine.Fine.TechnicalNumber,
		&fine.Fine.CarNumber,
		&fine.Fine.LicenceNumber,
		&fine.Fine.FineDate,
		&fine.Fine.FineAmount,
		&fine.Fine.PaymentDate,
		&fine.Fine.Status,
	)

	if err != nil {
		log.Println("Error while getting fine by id: ", err)
		return nil, err
	}

	log.Println("Successfully got fine by id")
	return &fine, nil
}

func (r *FineRepo) GetAll(req *mp.FineGetAllReq) (*mp.FineGetAllRes, error) {
	fines := mp.FineGetAllRes{}

	query := `
	SELECT 
		f.id,
		u.id,
		u.first_name,
		u.last_name,
		f.technical_number,
		f.car_number,
		f.licence_number,
		f.fine_date,
		f.fine_amount,
		CASE 
        	WHEN f.payment_date IS NULL THEN 'not yet paid' 
        	ELSE f.payment_date::text 
    	END as payment_status,
		f.status
	FROM 
		fines f
	JOIN 
		users u 
	ON 
		u.id = f.staf_id AND u.deleted_at = 0
	WHERE 
		f.deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.CarNumber != "" && req.CarNumber != "string" {
		conditions = append(conditions, " LOWER(f.car_number) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%"+req.CarNumber+"%")
	}

	if req.Status != "" && req.Status != "string" {
		conditions = append(conditions, " LOWER(f.status) = LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, req.Status)
	}

	if req.StafId != "" && req.StafId != "string" {
		conditions = append(conditions, " f.staf_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.StafId)
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
		log.Println("Fines not found")
		return nil, errors.New("fines list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving fines: ", err)
		return nil, err
	}

	for rows.Next() {
		fine := mp.FineRes{
			StafId: &mp.StafRes{},
		}

		err := rows.Scan(
			&fine.Id,
			&fine.StafId.Id,
			&fine.StafId.FirstName,
			&fine.StafId.LastName,
			&fine.TechnicalNumber,
			&fine.CarNumber,
			&fine.LicenceNumber,
			&fine.FineDate,
			&fine.FineAmount,
			&fine.PaymentDate,
			&fine.Status,
		)

		if err != nil {
			log.Println("Error while scanning fines: ", err)
			return nil, err
		}

		fines.Fines = append(fines.Fines, &fine)
	}

	log.Println("Successfully fetched all fines")
	return &fines, nil
}

func (r *FineRepo) Update(req *mp.FineUpdateReq) (*mp.Void, error) {
	query := `
	UPDATE
		fines
	SET 
		payment_date = now()
	`

	var status_query string

	var conditions []string
	var args []interface{}


	if req.Status != "" && req.Status != "string" {
		conditions = append(conditions, " status = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Status)
		status_query = " AND status <> 'paid'"
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"
	query += status_query

	args = append(args, req.Id)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating fine: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("fine with id %s couldn't update", req.Id)
	}

	log.Println("Successfully updated fine")

	return nil, nil
}

