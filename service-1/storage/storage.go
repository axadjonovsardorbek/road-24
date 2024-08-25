package storage

import (
	ap "auth/genproto/auth"
)

type UserI interface {
	Register(*ap.UserCreateReq) (*ap.Void, error)
	Login(*ap.UserLoginReq) (*ap.TokenRes, error)
	Profile(*ap.ById) (*ap.UserRes, error)
	DeleteProfile(*ap.ById) (*ap.Void, error)
	GetAllUsers(*ap.GetAllUsersReq) (*ap.GetAllUsersRes, error)
	RefreshToken(*ap.ById) (*ap.TokenRes, error)
}
