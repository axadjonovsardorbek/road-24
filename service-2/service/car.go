package service

import (
	"context"
	mp "mobile/genproto/mobile"
	st "mobile/storage/postgres"
)

type CarService struct {
	storage st.Storage
	mp.UnimplementedCarServiceServer
}

func NewCarService(storage *st.Storage) *CarService {
	return &CarService{storage: *storage}
}

func (s *CarService) Create(ctx context.Context, req *mp.CarCreateReq) (*mp.Void, error) {
	return s.storage.CarS.Create(req)
}
func (s *CarService) GetById(ctx context.Context, req *mp.ById) (*mp.CarGetByIdRes, error) {
	return s.storage.CarS.GetById(req)
}
func (s *CarService) GetAll(ctx context.Context, req *mp.CarGetAllReq) (*mp.CarGetAllRes, error) {
	return s.storage.CarS.GetAll(req)
}
func (s *CarService) Update(ctx context.Context, req *mp.CarUpdateReq) (*mp.Void, error){
	return s.storage.CarS.Update(req)
}
func (s *CarService) Delete(ctx context.Context, req *mp.ById) (*mp.Void, error) {
	return s.storage.CarS.Delete(req)
}
