package app

import (
	"log"
	"mobile/config"
	"mobile/service"
	mp "mobile/genproto/mobile"
	"mobile/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

func Run(config config.Config) {
	db, err := postgres.NewPostgresStorage(config)

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", config.MOBILE_PORT)

	if err != nil {
		log.Fatalf("Failed to listen tcp: %v", err)
	}

	s := grpc.NewServer()

	mp.RegisterCarServiceServer(s, service.NewCarService(db))
	mp.RegisterCarServiceServiceServer(s, service.NewCarServiceService(db))
	mp.RegisterCarTypeServiceServer(s, service.NewCarTypeService(db))
	mp.RegisterDriverLicenceServiceServer(s, service.NewDriverLicenceService(db))
	mp.RegisterFineServiceServer(s, service.NewFineService(db))
	mp.RegisterServiceServiceServer(s, service.NewServiceService(db))
	mp.RegisterServiceTypeServiceServer(s, service.NewServiceTypeService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
