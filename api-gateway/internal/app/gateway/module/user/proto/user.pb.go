// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: app/module/user/proto/user.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_module_user_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_app_module_user_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_app_module_user_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_module_user_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_app_module_user_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_app_module_user_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *Users) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type VoidParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VoidParam) Reset() {
	*x = VoidParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_module_user_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoidParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoidParam) ProtoMessage() {}

func (x *VoidParam) ProtoReflect() protoreflect.Message {
	mi := &file_app_module_user_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoidParam.ProtoReflect.Descriptor instead.
func (*VoidParam) Descriptor() ([]byte, []int) {
	return file_app_module_user_proto_user_proto_rawDescGZIP(), []int{2}
}

type UpdateUserParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserParam) Reset() {
	*x = UpdateUserParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_module_user_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserParam) ProtoMessage() {}

func (x *UpdateUserParam) ProtoReflect() protoreflect.Message {
	mi := &file_app_module_user_proto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserParam.ProtoReflect.Descriptor instead.
func (*UpdateUserParam) Descriptor() ([]byte, []int) {
	return file_app_module_user_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserParam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUserParam) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_app_module_user_proto_user_proto protoreflect.FileDescriptor

var file_app_module_user_proto_user_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa0, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x30, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x0b, 0x0a, 0x09, 0x76, 0x6f, 0x69, 0x64, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x22, 0x48, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xfd, 0x02, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x6f, 0x69, 0x64,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x3d, 0x0a, 0x0e, 0x72, 0x65, 0x61,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x6f, 0x69, 0x64, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x19, 0x5a, 0x17,
	0x2e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_module_user_proto_user_proto_rawDescOnce sync.Once
	file_app_module_user_proto_user_proto_rawDescData = file_app_module_user_proto_user_proto_rawDesc
)

func file_app_module_user_proto_user_proto_rawDescGZIP() []byte {
	file_app_module_user_proto_user_proto_rawDescOnce.Do(func() {
		file_app_module_user_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_module_user_proto_user_proto_rawDescData)
	})
	return file_app_module_user_proto_user_proto_rawDescData
}

var file_app_module_user_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_module_user_proto_user_proto_goTypes = []interface{}{
	(*User)(nil),                   // 0: userPackage.User
	(*Users)(nil),                  // 1: userPackage.Users
	(*VoidParam)(nil),              // 2: userPackage.voidParam
	(*UpdateUserParam)(nil),        // 3: userPackage.updateUserParam
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),   // 6: google.protobuf.BoolValue
}
var file_app_module_user_proto_user_proto_depIdxs = []int32{
	4,  // 0: userPackage.User.created_at:type_name -> google.protobuf.Timestamp
	4,  // 1: userPackage.User.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: userPackage.Users.users:type_name -> userPackage.User
	0,  // 3: userPackage.updateUserParam.user:type_name -> userPackage.User
	0,  // 4: userPackage.UserService.createUser:input_type -> userPackage.User
	2,  // 5: userPackage.UserService.readUsers:input_type -> userPackage.voidParam
	2,  // 6: userPackage.UserService.readUserStream:input_type -> userPackage.voidParam
	5,  // 7: userPackage.UserService.readUser:input_type -> google.protobuf.StringValue
	3,  // 8: userPackage.UserService.updateUser:input_type -> userPackage.updateUserParam
	5,  // 9: userPackage.UserService.deleteUser:input_type -> google.protobuf.StringValue
	0,  // 10: userPackage.UserService.createUser:output_type -> userPackage.User
	1,  // 11: userPackage.UserService.readUsers:output_type -> userPackage.Users
	0,  // 12: userPackage.UserService.readUserStream:output_type -> userPackage.User
	0,  // 13: userPackage.UserService.readUser:output_type -> userPackage.User
	0,  // 14: userPackage.UserService.updateUser:output_type -> userPackage.User
	6,  // 15: userPackage.UserService.deleteUser:output_type -> google.protobuf.BoolValue
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_app_module_user_proto_user_proto_init() }
func file_app_module_user_proto_user_proto_init() {
	if File_app_module_user_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_module_user_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_app_module_user_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_app_module_user_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoidParam); i {
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
		file_app_module_user_proto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserParam); i {
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
			RawDescriptor: file_app_module_user_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_module_user_proto_user_proto_goTypes,
		DependencyIndexes: file_app_module_user_proto_user_proto_depIdxs,
		MessageInfos:      file_app_module_user_proto_user_proto_msgTypes,
	}.Build()
	File_app_module_user_proto_user_proto = out.File
	file_app_module_user_proto_user_proto_rawDesc = nil
	file_app_module_user_proto_user_proto_goTypes = nil
	file_app_module_user_proto_user_proto_depIdxs = nil
}
