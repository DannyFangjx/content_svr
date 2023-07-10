// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.15.6
// source: content_svr/protobuf/pbadmin/login.proto

package pbadmin

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

// req nil
type OpLoginInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"email" form:"email"
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email" form:"email"` //email
	// @inject_tag: json:"pwd" form:"pwd"
	Pwd string `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd" form:"pwd"` //pwd的md5
}

func (x *OpLoginInReq) Reset() {
	*x = OpLoginInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpLoginInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpLoginInReq) ProtoMessage() {}

func (x *OpLoginInReq) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpLoginInReq.ProtoReflect.Descriptor instead.
func (*OpLoginInReq) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbadmin_login_proto_rawDescGZIP(), []int{0}
}

func (x *OpLoginInReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OpLoginInReq) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

// 返回此对象数组
type OpLoginInResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"token" form:"token"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token" form:"token"`
}

func (x *OpLoginInResp) Reset() {
	*x = OpLoginInResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpLoginInResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpLoginInResp) ProtoMessage() {}

func (x *OpLoginInResp) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpLoginInResp.ProtoReflect.Descriptor instead.
func (*OpLoginInResp) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbadmin_login_proto_rawDescGZIP(), []int{1}
}

func (x *OpLoginInResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// req nil
type OpAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"email"
	Email *string `protobuf:"bytes,1,opt,name=email,proto3,oneof" json:"email"`
	// @inject_tag: json:"pwd"
	Pwd *string `protobuf:"bytes,2,opt,name=pwd,proto3,oneof" json:"pwd"`
	// @inject_tag: json:"op_type"
	OpType *int32 `protobuf:"varint,3,opt,name=op_type,json=opType,proto3,oneof" json:"op_type"` // 0-普通， 1-admin
	// @inject_tag: json:"name"
	Name *string `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name"`
	// @inject_tag: json:"phone"
	Phone *string `protobuf:"bytes,5,opt,name=phone,proto3,oneof" json:"phone"`
}

func (x *OpAddReq) Reset() {
	*x = OpAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpAddReq) ProtoMessage() {}

func (x *OpAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpAddReq.ProtoReflect.Descriptor instead.
func (*OpAddReq) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbadmin_login_proto_rawDescGZIP(), []int{2}
}

func (x *OpAddReq) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *OpAddReq) GetPwd() string {
	if x != nil && x.Pwd != nil {
		return *x.Pwd
	}
	return ""
}

func (x *OpAddReq) GetOpType() int32 {
	if x != nil && x.OpType != nil {
		return *x.OpType
	}
	return 0
}

func (x *OpAddReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *OpAddReq) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

// resp
type OpInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"operator_id"
	OperatorId int64 `protobuf:"varint,1,opt,name=operator_id,json=operatorId,proto3" json:"operator_id"`
	// @inject_tag: json:"email"
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email"`
	// @inject_tag: json:"op_type"
	OpType int32 `protobuf:"varint,3,opt,name=op_type,json=opType,proto3" json:"op_type"` // 0-普通， 1-admin
	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	// @inject_tag: json:"phone"
	Phone string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone"`
	// @inject_tag: json:"status"
	Status int32 `protobuf:"varint,6,opt,name=status,proto3" json:"status"` // 0:不可用 1:可用
}

func (x *OpInfo) Reset() {
	*x = OpInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpInfo) ProtoMessage() {}

func (x *OpInfo) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpInfo.ProtoReflect.Descriptor instead.
func (*OpInfo) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbadmin_login_proto_rawDescGZIP(), []int{3}
}

func (x *OpInfo) GetOperatorId() int64 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

func (x *OpInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OpInfo) GetOpType() int32 {
	if x != nil {
		return x.OpType
	}
	return 0
}

func (x *OpInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OpInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *OpInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// req
type OpInfoUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"original_pwd" form:"original_pwd"    //原密码
	OriginalPwd string `protobuf:"bytes,1,opt,name=original_pwd,json=originalPwd,proto3" json:"original_pwd" form:"original_pwd"`
	// @inject_tag: json:"pwd" form:"pwd"
	Pwd string `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd" form:"pwd"`
}

func (x *OpInfoUpdateReq) Reset() {
	*x = OpInfoUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpInfoUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpInfoUpdateReq) ProtoMessage() {}

func (x *OpInfoUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpInfoUpdateReq.ProtoReflect.Descriptor instead.
func (*OpInfoUpdateReq) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbadmin_login_proto_rawDescGZIP(), []int{4}
}

func (x *OpInfoUpdateReq) GetOriginalPwd() string {
	if x != nil {
		return x.OriginalPwd
	}
	return ""
}

func (x *OpInfoUpdateReq) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

// req
type OpInfoListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"page" form:"page"
	Page *int32 `protobuf:"varint,6,opt,name=page,proto3,oneof" json:"page" form:"page"` // 从1开始。
	// @inject_tag: json:"size" form:"size"
	Size *int32 `protobuf:"varint,7,opt,name=size,proto3,oneof" json:"size" form:"size"` // >0 ;建议default 10
}

func (x *OpInfoListReq) Reset() {
	*x = OpInfoListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpInfoListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpInfoListReq) ProtoMessage() {}

func (x *OpInfoListReq) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpInfoListReq.ProtoReflect.Descriptor instead.
func (*OpInfoListReq) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbadmin_login_proto_rawDescGZIP(), []int{5}
}

func (x *OpInfoListReq) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *OpInfoListReq) GetSize() int32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

// resp
type OpInfoListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"op_info_list"
	OpInfoList []*OpInfo `protobuf:"bytes,1,rep,name=op_info_list,json=opInfoList,proto3" json:"op_info_list"`
	// @inject_tag: json:"total"
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *OpInfoListResp) Reset() {
	*x = OpInfoListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpInfoListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpInfoListResp) ProtoMessage() {}

func (x *OpInfoListResp) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbadmin_login_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpInfoListResp.ProtoReflect.Descriptor instead.
func (*OpInfoListResp) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbadmin_login_proto_rawDescGZIP(), []int{6}
}

func (x *OpInfoListResp) GetOpInfoList() []*OpInfo {
	if x != nil {
		return x.OpInfoList
	}
	return nil
}

func (x *OpInfoListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_content_svr_protobuf_pbadmin_login_proto protoreflect.FileDescriptor

var file_content_svr_protobuf_pbadmin_login_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x76, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0c, 0x4f, 0x70,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x77, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x4f, 0x70, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xbf, 0x01, 0x0a, 0x08, 0x4f, 0x70,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x15, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x03, 0x70, 0x77, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x06, 0x6f, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x70, 0x77, 0x64, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x06,
	0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a,
	0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x46, 0x0a, 0x0f, 0x6f, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x77, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x77, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x77, 0x64,
	0x22, 0x53, 0x0a, 0x0d, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x51, 0x0a, 0x0e, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x0c, 0x6f, 0x70, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x26, 0x5a, 0x24, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x76, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x62, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x3b, 0x70, 0x62, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_content_svr_protobuf_pbadmin_login_proto_rawDescOnce sync.Once
	file_content_svr_protobuf_pbadmin_login_proto_rawDescData = file_content_svr_protobuf_pbadmin_login_proto_rawDesc
)

func file_content_svr_protobuf_pbadmin_login_proto_rawDescGZIP() []byte {
	file_content_svr_protobuf_pbadmin_login_proto_rawDescOnce.Do(func() {
		file_content_svr_protobuf_pbadmin_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_svr_protobuf_pbadmin_login_proto_rawDescData)
	})
	return file_content_svr_protobuf_pbadmin_login_proto_rawDescData
}

var file_content_svr_protobuf_pbadmin_login_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_content_svr_protobuf_pbadmin_login_proto_goTypes = []interface{}{
	(*OpLoginInReq)(nil),    // 0: OpLoginInReq
	(*OpLoginInResp)(nil),   // 1: OpLoginInResp
	(*OpAddReq)(nil),        // 2: OpAddReq
	(*OpInfo)(nil),          // 3: opInfo
	(*OpInfoUpdateReq)(nil), // 4: opInfoUpdateReq
	(*OpInfoListReq)(nil),   // 5: opInfoListReq
	(*OpInfoListResp)(nil),  // 6: opInfoListResp
}
var file_content_svr_protobuf_pbadmin_login_proto_depIdxs = []int32{
	3, // 0: opInfoListResp.op_info_list:type_name -> opInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_content_svr_protobuf_pbadmin_login_proto_init() }
func file_content_svr_protobuf_pbadmin_login_proto_init() {
	if File_content_svr_protobuf_pbadmin_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_content_svr_protobuf_pbadmin_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpLoginInReq); i {
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
		file_content_svr_protobuf_pbadmin_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpLoginInResp); i {
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
		file_content_svr_protobuf_pbadmin_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpAddReq); i {
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
		file_content_svr_protobuf_pbadmin_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpInfo); i {
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
		file_content_svr_protobuf_pbadmin_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpInfoUpdateReq); i {
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
		file_content_svr_protobuf_pbadmin_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpInfoListReq); i {
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
		file_content_svr_protobuf_pbadmin_login_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpInfoListResp); i {
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
	file_content_svr_protobuf_pbadmin_login_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_content_svr_protobuf_pbadmin_login_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_content_svr_protobuf_pbadmin_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_content_svr_protobuf_pbadmin_login_proto_goTypes,
		DependencyIndexes: file_content_svr_protobuf_pbadmin_login_proto_depIdxs,
		MessageInfos:      file_content_svr_protobuf_pbadmin_login_proto_msgTypes,
	}.Build()
	File_content_svr_protobuf_pbadmin_login_proto = out.File
	file_content_svr_protobuf_pbadmin_login_proto_rawDesc = nil
	file_content_svr_protobuf_pbadmin_login_proto_goTypes = nil
	file_content_svr_protobuf_pbadmin_login_proto_depIdxs = nil
}