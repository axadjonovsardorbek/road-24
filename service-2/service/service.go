package service

import (
	"context"
	mp "mobile/genproto/mobile"
	st "mobile/storage/postgres"
)

type ServiceService struct {
	storage st.Storage
	mp.UnimplementedServiceServiceServer
}

func NewServiceService(storage *st.Storage) *ServiceService {
	return &ServiceService{storage: *storage}
}

func (s *ServiceService) Create(ctx context.Context, req *mp.ServiceCreateReq) (*mp.Void, error) {
	return s.storage.ServiceS.Create(req)
}
func (s *ServiceService) GetById(ctx context.Context, req *mp.ById) (*mp.ServiceGetByIdRes, error) {
	return s.storage.ServiceS.GetById(req)
}
func (s *ServiceService) GetAll(ctx context.Context, req *mp.ServiceGetAllReq) (*mp.ServiceGetAllRes, error) {
	return s.storage.ServiceS.GetAll(req)
}
func (s *ServiceService) Delete(ctx context.Context, req *mp.ById) (*mp.Void, error) {
	return s.storage.ServiceS.Delete(req)
}
