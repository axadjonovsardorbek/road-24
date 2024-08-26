package service

import (
	"context"
	mp "mobile/genproto/mobile"
	st "mobile/storage/postgres"
)

type CarServiceService struct {
	storage st.Storage
	mp.UnimplementedCarServiceServiceServer
}

func NewCarServiceService(storage *st.Storage) *CarServiceService {
	return &CarServiceService{storage: *storage}
}

func (s *CarServiceService) Create(ctx context.Context, req *mp.CarServiceCreateReq) (*mp.Void, error) {
	return s.storage.CarServiceS.Create(req)
}
func (s *CarServiceService) GetById(ctx context.Context, req *mp.ById) (*mp.CarServiceGetByIdRes, error) {
	return s.storage.CarServiceS.GetById(req)
}
func (s *CarServiceService) GetAll(ctx context.Context, req *mp.CarServiceGetAllReq) (*mp.CarServiceGetAllRes, error) {
	return s.storage.CarServiceS.GetAll(req)
}
func (s *CarServiceService) Delete(ctx context.Context, req *mp.ById) (*mp.Void, error) {
	return s.storage.CarServiceS.Delete(req)
}
