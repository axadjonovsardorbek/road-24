package postgres

import (
	"booking/config"
	"booking/storage"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db               *sql.DB
	ProviderS        storage.ProviderI
	BookingS         storage.BookingI
	PaymentS         storage.PaymentI
	ReviewS          storage.ReviewI
	ServiceS         storage.ServiceI
	ProviderServiceS storage.ProviderServiceI
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

	rAddr := fmt.Sprintf("%s%s", config.REDIS_HOST, config.REDIS_PORT)

	rdb := redis.NewClient(&redis.Options{
		Addr: rAddr,
	})

	provider := NewProviderRepo(db)
	booking := NewBookingRepo(db, rdb)
	payment := NewPaymentRepo(db)
	review := NewReviewRepo(db)
	service := NewServiceRepo(db)
	provider_service := NewProviderServiceRepo(db)

	return &Storage{
		Db:               db,
		ProviderS:        provider,
		BookingS:         booking,
		PaymentS:         payment,
		ReviewS:          review,
		ServiceS:         service,
		ProviderServiceS: provider_service,
	}, nil
}
