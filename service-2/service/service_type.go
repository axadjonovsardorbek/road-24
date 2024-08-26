package service

import (
	"context"
	mp "mobile/genproto/mobile"
	st "mobile/storage/postgres"
)

type ServiceTypeService struct {
	storage st.Storage
	mp.UnimplementedServiceTypeServiceServer
}

func NewServiceTypeService(storage *st.Storage) *ServiceTypeService {
	return &ServiceTypeService{storage: *storage}
}

func (s *ServiceTypeService) Create(ctx context.Context, req *mp.ServiceTypeCreateReq) (*mp.Void, error) {
	return s.storage.ServiceTypeS.Create(req)
}
func (s *ServiceTypeService) GetById(ctx context.Context, req *mp.ById) (*mp.ServiceTypeGetByIdRes, error) {
	return s.storage.ServiceTypeS.GetById(req)
}
func (s *ServiceTypeService) GetAll(ctx context.Context, req *mp.ServiceTypeGetAllReq) (*mp.ServiceTypeGetAllRes, error) {
	return s.storage.ServiceTypeS.GetAll(req)
}
func (s *ServiceTypeService) Delete(ctx context.Context, req *mp.ById) (*mp.Void, error) {
	return s.storage.ServiceTypeS.Delete(req)
}
