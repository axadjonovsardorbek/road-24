package service

import (
	ap "auth/genproto/auth"
	st "auth/storage/postgres"
	"context"
)

type UsersService struct {
	storage st.Storage
	ap.UnimplementedUserServiceServer
}

func NewUsersService(storage *st.Storage) *UsersService {
	return &UsersService{storage: *storage}
}

func (u *UsersService) Register(ctx context.Context, req *ap.UserCreateReq) (*ap.Void, error) {
	return u.storage.UserS.Register(req)
}

func (u *UsersService) Login(ctx context.Context, req *ap.UserLoginReq) (*ap.TokenRes, error) {
	return u.storage.UserS.Login(req)
}

func (u *UsersService) LoginDriver(ctx context.Context, req *ap.DriverLoginReq) (*ap.TokenRes, error) {
	return u.storage.UserS.LoginDriver(req)
}

func (u *UsersService) Profile(ctx context.Context, req *ap.ById) (*ap.UserRes, error) {
	return u.storage.UserS.Profile(req)
}

func (u *UsersService) DeleteProfile(ctx context.Context, req *ap.ById) (*ap.Void, error) {
	return u.storage.UserS.DeleteProfile(req)
}

func (u *UsersService) GetAllUsers(ctx context.Context, req *ap.GetAllUsersReq) (*ap.GetAllUsersRes, error){
	return u.storage.UserS.GetAllUsers(req)
}

func (u *UsersService) RefreshToken(ctx context.Context, req *ap.ById) (*ap.TokenRes, error) {
	return u.storage.UserS.RefreshToken(req)
}

func (u *UsersService) RefreshTokenForDriver(ctx context.Context, req *ap.ById) (*ap.TokenRes, error) {
	return u.storage.UserS.RefreshTokenForDriver(req)
}


