package postgres

import (
	ap "auth/genproto/auth"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	t "auth/api/token"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewUsersRepo(db *sql.DB, rdb *redis.Client) *UsersRepo {
	return &UsersRepo{db: db, rdb: rdb}
}

func (u *UsersRepo) Register(req *ap.UserCreateReq) (*ap.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO users(
		id, 
		first_name,
		last_name,
		username,
		password,
		role,
		is_admin
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := u.db.Exec(query, id, req.FirstName, req.LastName, req.Username, req.Password, req.Role, req.IsAdmin)

	if err != nil {
		log.Println("Error while registering user: ", err)
		return nil, err
	}

	log.Println("Successfully registered user")

	return nil, nil
}

func (u *UsersRepo) Login(req *ap.UserLoginReq) (*ap.TokenRes, error) {
	var id string
	var username string
	var password string
	var role string
	var is_admin string

	query := `
	SELECT 
		id,
		username,
		password,
		role,
		is_admin
	FROM 
		users
	WHERE
		username = $1
	AND 
		deleted_at = 0
	`

	row := u.db.QueryRow(query, req.Username)

	err := row.Scan(
		&id,
		&username,
		&password,
		&role,
		&is_admin,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		log.Println("Error while login user: ", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token := t.GenerateJWTToken(id, username, role, is_admin)
	tokens := ap.TokenRes{
		Token: token.AccessToken,
		ExpAt: "3 hours",
	}

	return &tokens, nil
}

func (u *UsersRepo) Profile(req *ap.ById) (*ap.UserRes, error) {
	userData, err := u.rdb.Get(context.Background(), req.Id).Result()
	if err == redis.Nil {
		user := ap.UserRes{}

	query := `
	SELECT 
		id, 
		first_name,
		last_name,
		username,
		role,
		is_admin
	FROM 	
		users
	WHERE
		id = $1
	AND 
		deleted_at = 0
	`
		row := u.db.QueryRow(query, req.Id)

		err := row.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Role,
			&user.IsAdmin,
		)

		if err != nil {
			log.Println("Error while getting user profile: ", err)
			return nil, err
		}

		fmt.Println("Successfully got profile")

		jsonData, err := json.Marshal(user)
		if err != nil {
			log.Printf("JSON marshalling error: %v", err)
			return nil, err
		}
		err = u.rdb.Set(context.Background(), req.Id, jsonData, 5*time.Minute).Err()
		if err != nil {
			log.Printf("Redis set error: %v", err)
			return nil, err
		}

		return &user, nil

	} else if err != nil {
		log.Printf("Redis get error: %v", err)
		return nil, err
	}

	user := &ap.UserRes{}

	err = json.Unmarshal([]byte(userData), user)
	if err != nil {
		log.Printf("JSON unmarshalling error: %v", err)
		return nil, err
	}

	return user, nil
}

func (u *UsersRepo) DeleteProfile(req *ap.ById) (*ap.Void, error) {
	query := `
	UPDATE 
		users
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := u.db.Exec(query, req.Id)
	if err != nil {
		log.Println("Error while deleting user: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("user with id %s not found", req.Id)
	}

	log.Println("Successfully deleted user")

	return nil, nil
}


func (u *UsersRepo) GetAllUsers(req *ap.GetAllUsersReq) (*ap.GetAllUsersRes, error) {
	users := ap.GetAllUsersRes{}

	query := `
	SELECT 
		id, 
		first_name,
		last_name,
		username,
		role,
		is_admin
	FROM 	
		users
	WHERE
		deleted_at = 0
	`
	var args []interface{}
	var conditions []string

	if req.Role != "" && req.Role != "string" {
		conditions = append(conditions, " role = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Role)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var limit int32
	var offset int32

	limit = 10
	offset = (req.Page - 1) * limit

	args = append(args, limit, offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := u.db.Query(query, args...)

	if err == sql.ErrNoRows {
		log.Println("Users not found")
		return nil, errors.New("users not found")
	}

	if err != nil {
		log.Println("Error while retriving users: ", err)
		return nil, err
	}

	for rows.Next() {
		user := ap.UserRes{}

		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Role,
			&user.IsAdmin,
		)

		if err != nil {
			log.Println("Error while scanning all users: ", err)
			return nil, err
		}

		users.Users = append(users.Users, &user)
	}

	log.Println("Successfully fetched all users")

	return &users, nil
}


func (u *UsersRepo) RefreshToken(req *ap.ById) (*ap.TokenRes, error) {
	var id string
	var username string
	var password string
	var role string
	var is_admin string

	query := `
	SELECT
		id,
		username,
		password,
		role,
		is_admin
	FROM
		users
	WHERE
		id = $1
	AND
		deleted_at = 0
	`

	row := u.db.QueryRow(query, req.Id)

	err := row.Scan(
		&id,
		&username,
		&password,
		&role,
		&is_admin,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		log.Println("Error while getting user: ", err)
		return nil, err
	}

	token := t.GenerateJWTToken(id, username, role, is_admin)
	tokens := ap.TokenRes{
		Token: token.RefreshToken,
		ExpAt: "24 hours",
	}

	return &tokens, nil
}
