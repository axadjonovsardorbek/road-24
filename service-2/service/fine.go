package service

import (
	"context"
	mp "mobile/genproto/mobile"
	st "mobile/storage/postgres"
)

type FineService struct {
	storage st.Storage
	mp.UnimplementedFineServiceServer
}

func NewFineService(storage *st.Storage) *FineService {
	return &FineService{storage: *storage}
}

func (s *FineService) Create(ctx context.Context, req *mp.FineCreateReq) (*mp.Void, error) {
	return s.storage.FineS.Create(req)
}
func (s *FineService) GetById(ctx context.Context, req *mp.ById) (*mp.FineGetByIdRes, error) {
	return s.storage.FineS.GetById(req)
}
func (s *FineService) GetAll(ctx context.Context, req *mp.FineGetAllReq) (*mp.FineGetAllRes, error) {
	return s.storage.FineS.GetAll(req)
}
func (s *FineService) Update(ctx context.Context, req *mp.FineUpdateReq) (*mp.Void, error){
	return s.storage.FineS.Update(req)
}
