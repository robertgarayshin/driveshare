// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: cars.proto

package cars

import (
	users "github.com/robertgarayshin/driveshare/proto/users"
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

type Drive int32

const (
	Drive_DRIVE_INVALID Drive = 0
	Drive_DRIVE_FRONT   Drive = 1
	Drive_DRIVE_REAR    Drive = 2
	Drive_DRIVES_ALL    Drive = 3
)

// Enum value maps for Drive.
var (
	Drive_name = map[int32]string{
		0: "DRIVE_INVALID",
		1: "DRIVE_FRONT",
		2: "DRIVE_REAR",
		3: "DRIVES_ALL",
	}
	Drive_value = map[string]int32{
		"DRIVE_INVALID": 0,
		"DRIVE_FRONT":   1,
		"DRIVE_REAR":    2,
		"DRIVES_ALL":    3,
	}
)

func (x Drive) Enum() *Drive {
	p := new(Drive)
	*p = x
	return p
}

func (x Drive) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Drive) Descriptor() protoreflect.EnumDescriptor {
	return file_cars_proto_enumTypes[0].Descriptor()
}

func (Drive) Type() protoreflect.EnumType {
	return &file_cars_proto_enumTypes[0]
}

func (x Drive) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Drive.Descriptor instead.
func (Drive) EnumDescriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{0}
}

type Gear int32

const (
	Gear_GEAR_INVALID Gear = 0
	Gear_GEAR_AUTO    Gear = 1
	Gear_GEAR_MANUAL  Gear = 2
	Gear_GEAR_ROBOT   Gear = 3
)

// Enum value maps for Gear.
var (
	Gear_name = map[int32]string{
		0: "GEAR_INVALID",
		1: "GEAR_AUTO",
		2: "GEAR_MANUAL",
		3: "GEAR_ROBOT",
	}
	Gear_value = map[string]int32{
		"GEAR_INVALID": 0,
		"GEAR_AUTO":    1,
		"GEAR_MANUAL":  2,
		"GEAR_ROBOT":   3,
	}
)

func (x Gear) Enum() *Gear {
	p := new(Gear)
	*p = x
	return p
}

func (x Gear) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gear) Descriptor() protoreflect.EnumDescriptor {
	return file_cars_proto_enumTypes[1].Descriptor()
}

func (Gear) Type() protoreflect.EnumType {
	return &file_cars_proto_enumTypes[1]
}

func (x Gear) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gear.Descriptor instead.
func (Gear) EnumDescriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{1}
}

type Body int32

const (
	Body_BODY_INVALID   Body = 0
	Body_BODY_SALOON    Body = 1
	Body_BODY_HATCHBACK Body = 2
	Body_BODY_COUPE     Body = 3
	Body_BODY_ESTATE    Body = 4
	Body_BODY_CABRIOLET Body = 5
	Body_BODY_ROADSTER  Body = 6
	Body_BODY_SUV       Body = 7
	Body_BODY_PICKUP    Body = 8
	Body_BODY_CUV       Body = 9
	Body_BODY_LIMOUSINE Body = 10
	Body_BODY_MINIVAN   Body = 11
	Body_BODY_VAN       Body = 12
)

// Enum value maps for Body.
var (
	Body_name = map[int32]string{
		0:  "BODY_INVALID",
		1:  "BODY_SALOON",
		2:  "BODY_HATCHBACK",
		3:  "BODY_COUPE",
		4:  "BODY_ESTATE",
		5:  "BODY_CABRIOLET",
		6:  "BODY_ROADSTER",
		7:  "BODY_SUV",
		8:  "BODY_PICKUP",
		9:  "BODY_CUV",
		10: "BODY_LIMOUSINE",
		11: "BODY_MINIVAN",
		12: "BODY_VAN",
	}
	Body_value = map[string]int32{
		"BODY_INVALID":   0,
		"BODY_SALOON":    1,
		"BODY_HATCHBACK": 2,
		"BODY_COUPE":     3,
		"BODY_ESTATE":    4,
		"BODY_CABRIOLET": 5,
		"BODY_ROADSTER":  6,
		"BODY_SUV":       7,
		"BODY_PICKUP":    8,
		"BODY_CUV":       9,
		"BODY_LIMOUSINE": 10,
		"BODY_MINIVAN":   11,
		"BODY_VAN":       12,
	}
)

func (x Body) Enum() *Body {
	p := new(Body)
	*p = x
	return p
}

func (x Body) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Body) Descriptor() protoreflect.EnumDescriptor {
	return file_cars_proto_enumTypes[2].Descriptor()
}

func (Body) Type() protoreflect.EnumType {
	return &file_cars_proto_enumTypes[2]
}

func (x Body) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Body.Descriptor instead.
func (Body) EnumDescriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{2}
}

type Class int32

const (
	Class_CLASS_INVALID  Class = 0
	Class_CLASS_ECONOMY  Class = 1
	Class_CLASS_COMFORT  Class = 2
	Class_CLASS_BUSINESS Class = 3
	Class_CLASS_PREMIUM  Class = 4
	Class_CLASS_SUV      Class = 5
	Class_CLASS_VAN      Class = 6
	Class_CLASS_UNIQUE   Class = 7
)

// Enum value maps for Class.
var (
	Class_name = map[int32]string{
		0: "CLASS_INVALID",
		1: "CLASS_ECONOMY",
		2: "CLASS_COMFORT",
		3: "CLASS_BUSINESS",
		4: "CLASS_PREMIUM",
		5: "CLASS_SUV",
		6: "CLASS_VAN",
		7: "CLASS_UNIQUE",
	}
	Class_value = map[string]int32{
		"CLASS_INVALID":  0,
		"CLASS_ECONOMY":  1,
		"CLASS_COMFORT":  2,
		"CLASS_BUSINESS": 3,
		"CLASS_PREMIUM":  4,
		"CLASS_SUV":      5,
		"CLASS_VAN":      6,
		"CLASS_UNIQUE":   7,
	}
)

func (x Class) Enum() *Class {
	p := new(Class)
	*p = x
	return p
}

func (x Class) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Class) Descriptor() protoreflect.EnumDescriptor {
	return file_cars_proto_enumTypes[3].Descriptor()
}

func (Class) Type() protoreflect.EnumType {
	return &file_cars_proto_enumTypes[3]
}

func (x Class) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Class.Descriptor instead.
func (Class) EnumDescriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{3}
}

type Steering_Side int32

const (
	Steering_Side_STEERING_SIDE_INVALID Steering_Side = 0
	Steering_Side_STEERING_SIDE_LEFT    Steering_Side = 1
	Steering_Side_STEERING_SIDE_RIGHT   Steering_Side = 2
)

// Enum value maps for Steering_Side.
var (
	Steering_Side_name = map[int32]string{
		0: "STEERING_SIDE_INVALID",
		1: "STEERING_SIDE_LEFT",
		2: "STEERING_SIDE_RIGHT",
	}
	Steering_Side_value = map[string]int32{
		"STEERING_SIDE_INVALID": 0,
		"STEERING_SIDE_LEFT":    1,
		"STEERING_SIDE_RIGHT":   2,
	}
)

func (x Steering_Side) Enum() *Steering_Side {
	p := new(Steering_Side)
	*p = x
	return p
}

func (x Steering_Side) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Steering_Side) Descriptor() protoreflect.EnumDescriptor {
	return file_cars_proto_enumTypes[4].Descriptor()
}

func (Steering_Side) Type() protoreflect.EnumType {
	return &file_cars_proto_enumTypes[4]
}

func (x Steering_Side) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Steering_Side.Descriptor instead.
func (Steering_Side) EnumDescriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{4}
}

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand        string        `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Model        string        `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Engine       string        `protobuf:"bytes,4,opt,name=engine,proto3" json:"engine,omitempty"`
	Drive        Drive         `protobuf:"varint,5,opt,name=drive,proto3,enum=proto.cars.Drive" json:"drive,omitempty"`
	Year         string        `protobuf:"bytes,6,opt,name=year,proto3" json:"year,omitempty"`
	Seats        string        `protobuf:"bytes,7,opt,name=seats,proto3" json:"seats,omitempty"`
	Gear         Gear          `protobuf:"varint,8,opt,name=gear,proto3,enum=proto.cars.Gear" json:"gear,omitempty"`
	Mileage      int32         `protobuf:"varint,9,opt,name=mileage,proto3" json:"mileage,omitempty"`
	Body         Body          `protobuf:"varint,10,opt,name=body,proto3,enum=proto.cars.Body" json:"body,omitempty"`
	Class        Class         `protobuf:"varint,11,opt,name=class,proto3,enum=proto.cars.Class" json:"class,omitempty"`
	SteeringSide Steering_Side `protobuf:"varint,12,opt,name=steering_side,json=steeringSide,proto3,enum=proto.cars.Steering_Side" json:"steering_side,omitempty"`
	Owner        *users.User   `protobuf:"bytes,13,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_cars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Car) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Car) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Car) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *Car) GetDrive() Drive {
	if x != nil {
		return x.Drive
	}
	return Drive_DRIVE_INVALID
}

func (x *Car) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Car) GetSeats() string {
	if x != nil {
		return x.Seats
	}
	return ""
}

func (x *Car) GetGear() Gear {
	if x != nil {
		return x.Gear
	}
	return Gear_GEAR_INVALID
}

func (x *Car) GetMileage() int32 {
	if x != nil {
		return x.Mileage
	}
	return 0
}

func (x *Car) GetBody() Body {
	if x != nil {
		return x.Body
	}
	return Body_BODY_INVALID
}

func (x *Car) GetClass() Class {
	if x != nil {
		return x.Class
	}
	return Class_CLASS_INVALID
}

func (x *Car) GetSteeringSide() Steering_Side {
	if x != nil {
		return x.SteeringSide
	}
	return Steering_Side_STEERING_SIDE_INVALID
}

func (x *Car) GetOwner() *users.User {
	if x != nil {
		return x.Owner
	}
	return nil
}

type CarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car    *Car           `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
	Cars   []*Car         `protobuf:"bytes,2,rep,name=cars,proto3" json:"cars,omitempty"`
	Errors []*users.Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *CarResponse) Reset() {
	*x = CarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarResponse) ProtoMessage() {}

func (x *CarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarResponse.ProtoReflect.Descriptor instead.
func (*CarResponse) Descriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{1}
}

func (x *CarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

func (x *CarResponse) GetCars() []*Car {
	if x != nil {
		return x.Cars
	}
	return nil
}

func (x *CarResponse) GetErrors() []*users.Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

type CarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CarRequest) Reset() {
	*x = CarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarRequest) ProtoMessage() {}

func (x *CarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarRequest.ProtoReflect.Descriptor instead.
func (*CarRequest) Descriptor() ([]byte, []int) {
	return file_cars_proto_rawDescGZIP(), []int{2}
}

var File_cars_proto protoreflect.FileDescriptor

var file_cars_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x1a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x03, 0x0a, 0x03,
	0x43, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x61, 0x72, 0x73, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x52, 0x05, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x67, 0x65,
	0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x61, 0x72, 0x52, 0x04, 0x67, 0x65, 0x61, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x27, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x73, 0x74, 0x65,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x53, 0x74,
	0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x53, 0x69, 0x64, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x65,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72,
	0x52, 0x03, 0x63, 0x61, 0x72, 0x12, 0x23, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73,
	0x2e, 0x43, 0x61, 0x72, 0x52, 0x04, 0x63, 0x61, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2a, 0x4b, 0x0a, 0x05, 0x44, 0x72, 0x69, 0x76, 0x65, 0x12, 0x11, 0x0a,
	0x0d, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x46, 0x52, 0x4f, 0x4e, 0x54, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x52, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x52, 0x49, 0x56, 0x45, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x10,
	0x03, 0x2a, 0x48, 0x0a, 0x04, 0x47, 0x65, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x45, 0x41,
	0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x47,
	0x45, 0x41, 0x52, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45,
	0x41, 0x52, 0x5f, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x47,
	0x45, 0x41, 0x52, 0x5f, 0x52, 0x4f, 0x42, 0x4f, 0x54, 0x10, 0x03, 0x2a, 0xe6, 0x01, 0x0a, 0x04,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x53,
	0x41, 0x4c, 0x4f, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x4f, 0x44, 0x59, 0x5f,
	0x48, 0x41, 0x54, 0x43, 0x48, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x42,
	0x4f, 0x44, 0x59, 0x5f, 0x43, 0x4f, 0x55, 0x50, 0x45, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x42,
	0x4f, 0x44, 0x59, 0x5f, 0x45, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e,
	0x42, 0x4f, 0x44, 0x59, 0x5f, 0x43, 0x41, 0x42, 0x52, 0x49, 0x4f, 0x4c, 0x45, 0x54, 0x10, 0x05,
	0x12, 0x11, 0x0a, 0x0d, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x52, 0x4f, 0x41, 0x44, 0x53, 0x54, 0x45,
	0x52, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x53, 0x55, 0x56, 0x10,
	0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x50, 0x49, 0x43, 0x4b, 0x55, 0x50,
	0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x43, 0x55, 0x56, 0x10, 0x09,
	0x12, 0x12, 0x0a, 0x0e, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x4c, 0x49, 0x4d, 0x4f, 0x55, 0x53, 0x49,
	0x4e, 0x45, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x4d, 0x49, 0x4e,
	0x49, 0x56, 0x41, 0x4e, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x56,
	0x41, 0x4e, 0x10, 0x0c, 0x2a, 0x97, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x45, 0x43, 0x4f, 0x4e, 0x4f,
	0x4d, 0x59, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x43, 0x4f,
	0x4d, 0x46, 0x4f, 0x52, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4c, 0x41, 0x53, 0x53,
	0x5f, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x4c, 0x41, 0x53, 0x53, 0x5f, 0x50, 0x52, 0x45, 0x4d, 0x49, 0x55, 0x4d, 0x10, 0x04, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x53, 0x55, 0x56, 0x10, 0x05, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x56, 0x41, 0x4e, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c,
	0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x10, 0x07, 0x2a, 0x5b,
	0x0a, 0x0d, 0x53, 0x74, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x53, 0x69, 0x64, 0x65, 0x12,
	0x19, 0x0a, 0x15, 0x53, 0x54, 0x45, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x49, 0x44, 0x45,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54,
	0x45, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x49, 0x44, 0x45, 0x5f, 0x4c, 0x45, 0x46, 0x54,
	0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x45, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x49, 0x44, 0x45, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x02, 0x32, 0x97, 0x02, 0x0a, 0x04,
	0x43, 0x61, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61,
	0x72, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43,
	0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61,
	0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x74, 0x67, 0x61, 0x72, 0x61, 0x79, 0x73,
	0x68, 0x69, 0x6e, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cars_proto_rawDescOnce sync.Once
	file_cars_proto_rawDescData = file_cars_proto_rawDesc
)

func file_cars_proto_rawDescGZIP() []byte {
	file_cars_proto_rawDescOnce.Do(func() {
		file_cars_proto_rawDescData = protoimpl.X.CompressGZIP(file_cars_proto_rawDescData)
	})
	return file_cars_proto_rawDescData
}

var file_cars_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_cars_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cars_proto_goTypes = []interface{}{
	(Drive)(0),          // 0: proto.cars.Drive
	(Gear)(0),           // 1: proto.cars.Gear
	(Body)(0),           // 2: proto.cars.Body
	(Class)(0),          // 3: proto.cars.Class
	(Steering_Side)(0),  // 4: proto.cars.Steering_Side
	(*Car)(nil),         // 5: proto.cars.Car
	(*CarResponse)(nil), // 6: proto.cars.CarResponse
	(*CarRequest)(nil),  // 7: proto.cars.CarRequest
	(*users.User)(nil),  // 8: proto.users.User
	(*users.Error)(nil), // 9: proto.users.Error
}
var file_cars_proto_depIdxs = []int32{
	0,  // 0: proto.cars.Car.drive:type_name -> proto.cars.Drive
	1,  // 1: proto.cars.Car.gear:type_name -> proto.cars.Gear
	2,  // 2: proto.cars.Car.body:type_name -> proto.cars.Body
	3,  // 3: proto.cars.Car.class:type_name -> proto.cars.Class
	4,  // 4: proto.cars.Car.steering_side:type_name -> proto.cars.Steering_Side
	8,  // 5: proto.cars.Car.owner:type_name -> proto.users.User
	5,  // 6: proto.cars.CarResponse.car:type_name -> proto.cars.Car
	5,  // 7: proto.cars.CarResponse.cars:type_name -> proto.cars.Car
	9,  // 8: proto.cars.CarResponse.errors:type_name -> proto.users.Error
	5,  // 9: proto.cars.Cars.Create:input_type -> proto.cars.Car
	5,  // 10: proto.cars.Cars.Edit:input_type -> proto.cars.Car
	5,  // 11: proto.cars.Cars.Delete:input_type -> proto.cars.Car
	7,  // 12: proto.cars.Cars.GetAllCars:input_type -> proto.cars.CarRequest
	5,  // 13: proto.cars.Cars.GetCarById:input_type -> proto.cars.Car
	6,  // 14: proto.cars.Cars.Create:output_type -> proto.cars.CarResponse
	6,  // 15: proto.cars.Cars.Edit:output_type -> proto.cars.CarResponse
	6,  // 16: proto.cars.Cars.Delete:output_type -> proto.cars.CarResponse
	6,  // 17: proto.cars.Cars.GetAllCars:output_type -> proto.cars.CarResponse
	6,  // 18: proto.cars.Cars.GetCarById:output_type -> proto.cars.CarResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_cars_proto_init() }
func file_cars_proto_init() {
	if File_cars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_cars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarResponse); i {
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
		file_cars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarRequest); i {
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
			RawDescriptor: file_cars_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cars_proto_goTypes,
		DependencyIndexes: file_cars_proto_depIdxs,
		EnumInfos:         file_cars_proto_enumTypes,
		MessageInfos:      file_cars_proto_msgTypes,
	}.Build()
	File_cars_proto = out.File
	file_cars_proto_rawDesc = nil
	file_cars_proto_goTypes = nil
	file_cars_proto_depIdxs = nil
}
