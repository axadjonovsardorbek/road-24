// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: car.proto

package mobile

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CarRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StafId            *StafRes    `protobuf:"bytes,2,opt,name=staf_id,json=stafId,proto3" json:"staf_id,omitempty"`
	Type              *CarTypeRes `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Color             string      `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Year              int32       `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
	BodyNumber        string      `protobuf:"bytes,6,opt,name=body_number,json=bodyNumber,proto3" json:"body_number,omitempty"`
	HorsePower        int32       `protobuf:"varint,7,opt,name=horse_power,json=horsePower,proto3" json:"horse_power,omitempty"`
	Number            string      `protobuf:"bytes,8,opt,name=number,proto3" json:"number,omitempty"`
	TechnicalPassport string      `protobuf:"bytes,9,opt,name=technical_passport,json=technicalPassport,proto3" json:"technical_passport,omitempty"`
	Model             string      `protobuf:"bytes,10,opt,name=model,proto3" json:"model,omitempty"`
	ImageUrl          string      `protobuf:"bytes,11,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *CarRes) Reset() {
	*x = CarRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarRes) ProtoMessage() {}

func (x *CarRes) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarRes.ProtoReflect.Descriptor instead.
func (*CarRes) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{0}
}

func (x *CarRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CarRes) GetStafId() *StafRes {
	if x != nil {
		return x.StafId
	}
	return nil
}

func (x *CarRes) GetType() *CarTypeRes {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *CarRes) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *CarRes) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CarRes) GetBodyNumber() string {
	if x != nil {
		return x.BodyNumber
	}
	return ""
}

func (x *CarRes) GetHorsePower() int32 {
	if x != nil {
		return x.HorsePower
	}
	return 0
}

func (x *CarRes) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CarRes) GetTechnicalPassport() string {
	if x != nil {
		return x.TechnicalPassport
	}
	return ""
}

func (x *CarRes) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CarRes) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type CarCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StafId            string `protobuf:"bytes,1,opt,name=staf_id,json=stafId,proto3" json:"staf_id,omitempty"`
	Type              string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Color             string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Year              int32  `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
	BodyNumber        string `protobuf:"bytes,5,opt,name=body_number,json=bodyNumber,proto3" json:"body_number,omitempty"`
	HorsePower        int32  `protobuf:"varint,6,opt,name=horse_power,json=horsePower,proto3" json:"horse_power,omitempty"`
	Number            string `protobuf:"bytes,7,opt,name=number,proto3" json:"number,omitempty"`
	TechnicalPassport string `protobuf:"bytes,8,opt,name=technical_passport,json=technicalPassport,proto3" json:"technical_passport,omitempty"`
	Model             string `protobuf:"bytes,9,opt,name=model,proto3" json:"model,omitempty"`
	ImageUrl          string `protobuf:"bytes,10,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *CarCreateReq) Reset() {
	*x = CarCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarCreateReq) ProtoMessage() {}

func (x *CarCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarCreateReq.ProtoReflect.Descriptor instead.
func (*CarCreateReq) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{1}
}

func (x *CarCreateReq) GetStafId() string {
	if x != nil {
		return x.StafId
	}
	return ""
}

func (x *CarCreateReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CarCreateReq) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *CarCreateReq) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CarCreateReq) GetBodyNumber() string {
	if x != nil {
		return x.BodyNumber
	}
	return ""
}

func (x *CarCreateReq) GetHorsePower() int32 {
	if x != nil {
		return x.HorsePower
	}
	return 0
}

func (x *CarCreateReq) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CarCreateReq) GetTechnicalPassport() string {
	if x != nil {
		return x.TechnicalPassport
	}
	return ""
}

func (x *CarCreateReq) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CarCreateReq) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type CarUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarNumber         string `protobuf:"bytes,1,opt,name=car_number,json=carNumber,proto3" json:"car_number,omitempty"`
	TechnicalPassport string `protobuf:"bytes,2,opt,name=technical_passport,json=technicalPassport,proto3" json:"technical_passport,omitempty"`
	ImageUrl          string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *CarUpdateReq) Reset() {
	*x = CarUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarUpdateReq) ProtoMessage() {}

func (x *CarUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarUpdateReq.ProtoReflect.Descriptor instead.
func (*CarUpdateReq) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{2}
}

func (x *CarUpdateReq) GetCarNumber() string {
	if x != nil {
		return x.CarNumber
	}
	return ""
}

func (x *CarUpdateReq) GetTechnicalPassport() string {
	if x != nil {
		return x.TechnicalPassport
	}
	return ""
}

func (x *CarUpdateReq) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type CarGetByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *CarRes `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *CarGetByIdRes) Reset() {
	*x = CarGetByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarGetByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarGetByIdRes) ProtoMessage() {}

func (x *CarGetByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarGetByIdRes.ProtoReflect.Descriptor instead.
func (*CarGetByIdRes) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{3}
}

func (x *CarGetByIdRes) GetCar() *CarRes {
	if x != nil {
		return x.Car
	}
	return nil
}

type CarGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars []*CarRes `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
}

func (x *CarGetAllRes) Reset() {
	*x = CarGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarGetAllRes) ProtoMessage() {}

func (x *CarGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarGetAllRes.ProtoReflect.Descriptor instead.
func (*CarGetAllRes) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{4}
}

func (x *CarGetAllRes) GetCars() []*CarRes {
	if x != nil {
		return x.Cars
	}
	return nil
}

type CarGetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Model  string  `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Year   int32   `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	Filter *Filter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *CarGetAllReq) Reset() {
	*x = CarGetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarGetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarGetAllReq) ProtoMessage() {}

func (x *CarGetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarGetAllReq.ProtoReflect.Descriptor instead.
func (*CarGetAllReq) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{5}
}

func (x *CarGetAllReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CarGetAllReq) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CarGetAllReq) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CarGetAllReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

var File_car_proto protoreflect.FileDescriptor

var file_car_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x32, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd0, 0x02, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a,
	0x07, 0x73, 0x74, 0x61, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x52, 0x65, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x66, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x43,
	0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x64,
	0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x6f, 0x64, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x6f,
	0x72, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xa1, 0x02, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x66, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x69, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x50, 0x61,
	0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x79, 0x0a, 0x0c, 0x43, 0x61, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x50,
	0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x22, 0x31, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22, 0x32, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x43,
	0x61, 0x72, 0x52, 0x65, 0x73, 0x52, 0x04, 0x63, 0x61, 0x72, 0x73, 0x22, 0x74, 0x0a, 0x0c, 0x43,
	0x61, 0x72, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x32, 0xf4, 0x01, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x0c, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2e,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0c, 0x2e, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x2e, 0x43, 0x61, 0x72, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x34,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x43, 0x61, 0x72, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_car_proto_rawDescOnce sync.Once
	file_car_proto_rawDescData = file_car_proto_rawDesc
)

func file_car_proto_rawDescGZIP() []byte {
	file_car_proto_rawDescOnce.Do(func() {
		file_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_proto_rawDescData)
	})
	return file_car_proto_rawDescData
}

var file_car_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_car_proto_goTypes = []interface{}{
	(*CarRes)(nil),        // 0: mobile.CarRes
	(*CarCreateReq)(nil),  // 1: mobile.CarCreateReq
	(*CarUpdateReq)(nil),  // 2: mobile.CarUpdateReq
	(*CarGetByIdRes)(nil), // 3: mobile.CarGetByIdRes
	(*CarGetAllRes)(nil),  // 4: mobile.CarGetAllRes
	(*CarGetAllReq)(nil),  // 5: mobile.CarGetAllReq
	(*StafRes)(nil),       // 6: mobile.StafRes
	(*CarTypeRes)(nil),    // 7: mobile.CarTypeRes
	(*Filter)(nil),        // 8: mobile.Filter
	(*ById)(nil),          // 9: mobile.ById
	(*Void)(nil),          // 10: mobile.Void
}
var file_car_proto_depIdxs = []int32{
	6,  // 0: mobile.CarRes.staf_id:type_name -> mobile.StafRes
	7,  // 1: mobile.CarRes.type:type_name -> mobile.CarTypeRes
	0,  // 2: mobile.CarGetByIdRes.car:type_name -> mobile.CarRes
	0,  // 3: mobile.CarGetAllRes.cars:type_name -> mobile.CarRes
	8,  // 4: mobile.CarGetAllReq.filter:type_name -> mobile.Filter
	1,  // 5: mobile.CarService.Create:input_type -> mobile.CarCreateReq
	9,  // 6: mobile.CarService.GetById:input_type -> mobile.ById
	5,  // 7: mobile.CarService.GetAll:input_type -> mobile.CarGetAllReq
	2,  // 8: mobile.CarService.Update:input_type -> mobile.CarUpdateReq
	9,  // 9: mobile.CarService.Delete:input_type -> mobile.ById
	10, // 10: mobile.CarService.Create:output_type -> mobile.Void
	3,  // 11: mobile.CarService.GetById:output_type -> mobile.CarGetByIdRes
	4,  // 12: mobile.CarService.GetAll:output_type -> mobile.CarGetAllRes
	10, // 13: mobile.CarService.Update:output_type -> mobile.Void
	10, // 14: mobile.CarService.Delete:output_type -> mobile.Void
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_car_proto_init() }
func file_car_proto_init() {
	if File_car_proto != nil {
		return
	}
	file_common2_proto_init()
	file_car_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarCreateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarUpdateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_car_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarGetByIdRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_car_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarGetAllRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_car_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarGetAllReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_car_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_proto_goTypes,
		DependencyIndexes: file_car_proto_depIdxs,
		MessageInfos:      file_car_proto_msgTypes,
	}.Build()
	File_car_proto = out.File
	file_car_proto_rawDesc = nil
	file_car_proto_goTypes = nil
	file_car_proto_depIdxs = nil
}