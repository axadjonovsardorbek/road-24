package clients

import (
	"auth/config"
	bp "auth/genproto/mobile"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Car           bp.CarServiceClient
	CarService    bp.CarServiceServiceClient
	CarType       bp.CarTypeServiceClient
	DriverLicence bp.DriverLicenceServiceClient
	Fine          bp.FineServiceClient
	ServiceType   bp.ServiceTypeServiceClient
	Service       bp.ServiceServiceClient
}

func NewGrpcClients(cfg *config.Config) (*GrpcClients, error) {
	connB, err := grpc.NewClient(cfg.MOBILE_HOST+cfg.MOBILE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		Car:           bp.NewCarServiceClient(connB),
		CarService:    bp.NewCarServiceServiceClient(connB),
		CarType:       bp.NewCarTypeServiceClient(connB),
		DriverLicence: bp.NewDriverLicenceServiceClient(connB),
		Fine:          bp.NewFineServiceClient(connB),
		ServiceType:   bp.NewServiceTypeServiceClient(connB),
		Service:       bp.NewServiceServiceClient(connB),
	}, nil
}
