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

type ServiceI interface {
	Create(*bp.ServiceRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.ServiceGetByIdRes, error)
	GetAll(*bp.ServiceGetAllReq) (*bp.ServiceGetAllRes, error)
	Update(*bp.ServiceUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type BookingI interface {
	Create(*bp.BookingRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.BookingGetByIdRes, error)
	GetAll(*bp.BookingGetAllReq) (*bp.BookingGetAllRes, error)
	Update(*bp.BookingUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
	GetPopularService(*bp.ById) (*bp.GetPopularServiceRes, error)
}

type PaymentI interface {
	Create(*bp.PaymentRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.PaymentGetByIdRes, error)
	GetAll(*bp.PaymentGetAllReq) (*bp.PaymentGetAllRes, error)
	Update(*bp.PaymentUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
	GetBookingId(*bp.ById) (*bp.ById, error)
	GetBookingAmount(*bp.ById) (*bp.GetAmountRes, error)
}

type ReviewI interface {
	Create(*bp.ReviewRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.ReviewGetByIdRes, error)
	GetAll(*bp.ReviewGetAllReq) (*bp.ReviewGetAllRes, error)
	Update(*bp.ReviewUpdateReq) (*bp.Void, error)
	Delete(*bp.ById) (*bp.Void, error)
	GetProviderRating(*bp.ById) (*bp.GetProviderRatingRes, error)
}

type ProviderServiceI interface {
	Create(*bp.ProviderServiceRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.ProviderServiceGetByIdRes, error)
	GetAll(*bp.ProviderServiceGetAllReq) (*bp.ProviderServiceGetAllRes, error)
	Delete(*bp.ById) (*bp.Void, error)
}

type NotificationI interface {
	Create(*bp.NotificationRes) (*bp.Void, error)
	GetById(*bp.ById) (*bp.NotificationGetByIdRes, error)
	GetAll(*bp.NotificationGetAllReq) (*bp.NotificationGetAllRes, error)
	Update(*bp.NotificationUpdateReq) (*bp.Void, error)
}
