package service

import (
	"context"
	mp "mobile/genproto/mobile"
	st "mobile/storage/postgres"
)

type CarTypeService struct {
	storage st.Storage
	mp.UnimplementedCarTypeServiceServer
}

func NewCarTypeService(storage *st.Storage) *CarTypeService {
	return &CarTypeService{storage: *storage}
}

func (s *CarTypeService) Create(ctx context.Context, req *mp.CarTypeCreateReq) (*mp.Void, error) {
	return s.storage.CarTypeS.Create(req)
}
func (s *CarTypeService) GetById(ctx context.Context, req *mp.ById) (*mp.CarTypeGetByIdRes, error) {
	return s.storage.CarTypeS.GetById(req)
}
func (s *CarTypeService) GetAll(ctx context.Context, req *mp.CarTypeGetAllReq) (*mp.CarTypeGetAllRes, error) {
	return s.storage.CarTypeS.GetAll(req)
}
func (s *CarTypeService) Delete(ctx context.Context, req *mp.ById) (*mp.Void, error) {
	return s.storage.CarTypeS.Delete(req)
}
