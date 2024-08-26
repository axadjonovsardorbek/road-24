package service

import (
	"context"
	mp "mobile/genproto/mobile"
	st "mobile/storage/postgres"
)

type DriverLicenceService struct {
	storage st.Storage
	mp.UnimplementedDriverLicenceServiceServer
}

func NewDriverLicenceService(storage *st.Storage) *DriverLicenceService {
	return &DriverLicenceService{storage: *storage}
}

func (s *DriverLicenceService) Create(ctx context.Context, req *mp.DriverLicenceCreateReq) (*mp.Void, error) {
	return s.storage.DriverLicenceS.Create(req)
}
func (s *DriverLicenceService) GetById(ctx context.Context, req *mp.ById) (*mp.DriverLicenceGetByIdRes, error) {
	return s.storage.DriverLicenceS.GetById(req)
}
func (s *DriverLicenceService) GetAll(ctx context.Context, req *mp.DriverLicenceGetAllReq) (*mp.DriverLicenceGetAllRes, error) {
	return s.storage.DriverLicenceS.GetAll(req)
}
func (s *DriverLicenceService) Delete(ctx context.Context, req *mp.ById) (*mp.Void, error) {
	return s.storage.DriverLicenceS.Delete(req)
}
