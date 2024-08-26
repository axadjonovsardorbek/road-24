package storage

import (
	mp "mobile/genproto/mobile"
)

type CarTypeI interface {
	Create(*mp.CarTypeCreateReq) (*mp.Void, error)
	GetById(*mp.ById) (*mp.CarTypeGetByIdRes, error)
	GetAll(*mp.CarTypeGetAllReq) (*mp.CarTypeGetAllRes, error)
	Delete(*mp.ById) (*mp.Void, error)
}

type CarServiceI interface {
	Create(*mp.CarServiceCreateReq) (*mp.Void, error)
	GetById(*mp.ById) (*mp.CarServiceGetByIdRes, error)
	GetAll(*mp.CarServiceGetAllReq) (*mp.CarServiceGetAllRes, error)
	Delete(*mp.ById) (*mp.Void, error)
}

type CarI interface {
	Create(*mp.CarCreateReq) (*mp.Void, error)
	GetById(*mp.ById) (*mp.CarGetByIdRes, error)
	GetAll(*mp.CarGetAllReq) (*mp.CarGetAllRes, error)
	Update(*mp.CarUpdateReq) (*mp.Void, error)
	Delete(*mp.ById) (*mp.Void, error)
}

type DriverLicenceI interface {
	Create(*mp.DriverLicenceCreateReq) (*mp.Void, error)
	GetById(*mp.ById) (*mp.DriverLicenceGetByIdRes, error)
	GetAll(*mp.DriverLicenceGetAllReq) (*mp.DriverLicenceGetAllRes, error)
	Delete(*mp.ById) (*mp.Void, error)
}

type FineI interface {
	Create(*mp.FineCreateReq) (*mp.Void, error)
	GetById(*mp.ById) (*mp.FineGetByIdRes, error)
	GetAll(*mp.FineGetAllReq) (*mp.FineGetAllRes, error)
	Update(*mp.FineUpdateReq) (*mp.Void, error)
}

type ServiceTypeI interface {
	Create(*mp.ServiceTypeCreateReq) (*mp.Void, error)
	GetById(*mp.ById) (*mp.ServiceTypeGetByIdRes, error)
	GetAll(*mp.ServiceTypeGetAllReq) (*mp.ServiceTypeGetAllRes, error)
	Delete(*mp.ById) (*mp.Void, error)
}

type ServiceI interface {
	Create(*mp.ServiceCreateReq) (*mp.Void, error)
	GetById(*mp.ById) (*mp.ServiceGetByIdRes, error)
	GetAll(*mp.ServiceGetAllReq) (*mp.ServiceGetAllRes, error)
	Delete(*mp.ById) (*mp.Void, error)
}
