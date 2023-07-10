// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.15.6
// source: content_svr/protobuf/pbapi/api.proto

package pbapi

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

// *
// @Description :   http req and resp
type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbapi_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbapi_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbapi_api_proto_rawDescGZIP(), []int{0}
}

// debug api
type DebugUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"user_id" form:"user_id"
	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"user_id" form:"user_id"`
}

func (x *DebugUserInfoReq) Reset() {
	*x = DebugUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbapi_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugUserInfoReq) ProtoMessage() {}

func (x *DebugUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbapi_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugUserInfoReq.ProtoReflect.Descriptor instead.
func (*DebugUserInfoReq) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbapi_api_proto_rawDescGZIP(), []int{1}
}

func (x *DebugUserInfoReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type HttpHeaderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apptype     string `protobuf:"bytes,1,opt,name=apptype,proto3" json:"apptype,omitempty"`
	Appname     string `protobuf:"bytes,2,opt,name=appname,proto3" json:"appname,omitempty"`
	Token       string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Versioncode string `protobuf:"bytes,4,opt,name=versioncode,proto3" json:"versioncode,omitempty"`
	Debuguserid int64  `protobuf:"varint,5,opt,name=debuguserid,proto3" json:"debuguserid,omitempty"`
	Uk          string `protobuf:"bytes,6,opt,name=uk,proto3" json:"uk,omitempty"`
	Platform    string `protobuf:"bytes,7,opt,name=platform,proto3" json:"platform,omitempty"`
	Ip          string `protobuf:"bytes,8,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *HttpHeaderInfo) Reset() {
	*x = HttpHeaderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbapi_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpHeaderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpHeaderInfo) ProtoMessage() {}

func (x *HttpHeaderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbapi_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpHeaderInfo.ProtoReflect.Descriptor instead.
func (*HttpHeaderInfo) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbapi_api_proto_rawDescGZIP(), []int{2}
}

func (x *HttpHeaderInfo) GetApptype() string {
	if x != nil {
		return x.Apptype
	}
	return ""
}

func (x *HttpHeaderInfo) GetAppname() string {
	if x != nil {
		return x.Appname
	}
	return ""
}

func (x *HttpHeaderInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *HttpHeaderInfo) GetVersioncode() string {
	if x != nil {
		return x.Versioncode
	}
	return ""
}

func (x *HttpHeaderInfo) GetDebuguserid() int64 {
	if x != nil {
		return x.Debuguserid
	}
	return 0
}

func (x *HttpHeaderInfo) GetUk() string {
	if x != nil {
		return x.Uk
	}
	return ""
}

func (x *HttpHeaderInfo) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *HttpHeaderInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

// DB   public class UserInfo {
type UserinfoDbModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"id"
	Id *int64 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty" gorm:"id"` // 主键
	// @inject_tag: gorm:"user_id"
	UserId *int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty" gorm:"user_id"` // 用户id，全局唯一标识
	// @inject_tag: gorm:"app_flag"
	AppFlag *int32 `protobuf:"varint,3,opt,name=app_flag,json=appFlag,proto3,oneof" json:"app_flag,omitempty" gorm:"app_flag"` // 应用标识，1:笑话；2:趣图；3:壁纸；4:星座，5:正能量，6:求赞，7；速配，14:猫爪运动
	// @inject_tag: gorm:"phone"
	Phone *string `protobuf:"bytes,4,opt,name=phone,proto3,oneof" json:"phone,omitempty" gorm:"phone"` // 手机号码
	// @inject_tag: gorm:"password"
	Password *string `protobuf:"bytes,5,opt,name=password,proto3,oneof" json:"password,omitempty" gorm:"password"` // 登陆密码，加盐后使用md5存储
	// @inject_tag: gorm:"nick_name"
	NickName *string `protobuf:"bytes,6,opt,name=nick_name,json=nickName,proto3,oneof" json:"nick_name,omitempty" gorm:"nick_name"` // 昵称
	// @inject_tag: gorm:"open_nick_name"
	OpenNickName *string `protobuf:"bytes,7,opt,name=open_nick_name,json=openNickName,proto3,oneof" json:"open_nick_name,omitempty" gorm:"open_nick_name"` // 第三方昵称
	// @inject_tag: gorm:"contact"
	Contact *string `protobuf:"bytes,8,opt,name=contact,proto3,oneof" json:"contact,omitempty" gorm:"contact"` // 联系方式
	// @inject_tag: gorm:"photo"
	Photo *string `protobuf:"bytes,9,opt,name=photo,proto3,oneof" json:"photo,omitempty" gorm:"photo"` // 头像
	// @inject_tag: gorm:"open_photo"
	OpenPhoto *string `protobuf:"bytes,10,opt,name=open_photo,json=openPhoto,proto3,oneof" json:"open_photo,omitempty" gorm:"open_photo"` // 第三方头像
	// @inject_tag: gorm:"source"
	Source *int32 `protobuf:"varint,11,opt,name=source,proto3,oneof" json:"source,omitempty" gorm:"source"` // 用户来源 1:微信；2:qq；3:新浪微博；0:系统注册
	// @inject_tag: gorm:"birth"
	Birth *string `protobuf:"bytes,12,opt,name=birth,proto3,oneof" json:"birth,omitempty" gorm:"birth"` // 生日
	// @inject_tag: gorm:"gender"
	Gender *int32 `protobuf:"varint,13,opt,name=gender,proto3,oneof" json:"gender,omitempty" gorm:"gender"` // 性别 2女 1男
	// @inject_tag: gorm:"type
	Type *int32 `protobuf:"varint,14,opt,name=type,proto3,oneof" json:"type,omitempty"` // 用户类型 1:普通用户；2:大神陪玩
	// @inject_tag: gorm:"level"
	Level *string `protobuf:"bytes,15,opt,name=level,proto3,oneof" json:"level,omitempty" gorm:"level"` // 普通用户等级
	// @inject_tag: gorm:"focus"
	Focus *int32 `protobuf:"varint,16,opt,name=focus,proto3,oneof" json:"focus,omitempty" gorm:"focus"` // 关注用户数量
	// @inject_tag: gorm:"fans"
	Fans *int32 `protobuf:"varint,17,opt,name=fans,proto3,oneof" json:"fans,omitempty" gorm:"fans"` // 被关注次数
	// @inject_tag: gorm:"province"
	Province *string `protobuf:"bytes,18,opt,name=province,proto3,oneof" json:"province,omitempty" gorm:"province"` // 所在省份
	// @inject_tag: gorm:"city"
	City *string `protobuf:"bytes,19,opt,name=city,proto3,oneof" json:"city,omitempty" gorm:"city"` // 所在城市
	// @inject_tag: gorm:"constellation"
	Constellation *string `protobuf:"bytes,20,opt,name=constellation,proto3,oneof" json:"constellation,omitempty" gorm:"constellation"` // 星座
	// @inject_tag: gorm:"description"
	Description *string `protobuf:"bytes,21,opt,name=description,proto3,oneof" json:"description,omitempty" gorm:"description"` // 签名
	// @inject_tag: gorm:"score"
	Score *int32 `protobuf:"varint,22,opt,name=score,proto3,oneof" json:"score,omitempty" gorm:"score"` // 评价总分
	// @inject_tag: gorm:"order_quantity"
	OrderQuantity *int32 `protobuf:"varint,23,opt,name=order_quantity,json=orderQuantity,proto3,oneof" json:"order_quantity,omitempty" gorm:"order_quantity"` // 完成总单数
	// @inject_tag: gorm:"enabled"
	Enabled *int32 `protobuf:"varint,24,opt,name=enabled,proto3,oneof" json:"enabled,omitempty" gorm:"enabled"` // 启用状态,1:启用，2:限时封禁，3:永久封禁
	// @inject_tag: gorm:"works_lock"
	WorksLock *int32 `protobuf:"varint,25,opt,name=works_lock,json=worksLock,proto3,oneof" json:"works_lock,omitempty" gorm:"works_lock"` // 作品发布锁定，1:启用，2:限时封禁，3:永久封禁
	// @inject_tag: gorm:"comment_lock"
	CommentLock *int32 `protobuf:"varint,26,opt,name=comment_lock,json=commentLock,proto3,oneof" json:"comment_lock,omitempty" gorm:"comment_lock"` // 作品发布锁定，1:启用，2:限时封禁，3:永久封禁
	// @inject_tag: gorm:"status"
	Status *int32 `protobuf:"varint,27,opt,name=status,proto3,oneof" json:"status,omitempty" gorm:"status"` // 用户状态,0:无效；1:有效
	// @inject_tag: gorm:"create_time"
	CreateTime *string `protobuf:"bytes,28,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty" gorm:"create_time"` // 创建时间
	// @inject_tag: gorm:"update_time"
	UpdateTime *string `protobuf:"bytes,29,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty" gorm:"update_time"` // 最后修改时间
	// @inject_tag: gorm:"update_gender_count"
	UpdateGenderCount *int32 `protobuf:"varint,30,opt,name=update_gender_count,json=updateGenderCount,proto3,oneof" json:"update_gender_count,omitempty" gorm:"update_gender_count"` // 性别修改次数
	// @inject_tag: gorm:"friend_lock"
	FriendLock *int32 `protobuf:"varint,31,opt,name=friend_lock,json=friendLock,proto3,oneof" json:"friend_lock,omitempty" gorm:"friend_lock"` // 好友开关 1：开，0：关
	// @inject_tag: gorm:"-" json:"-"
	EmbeddedNamer *int32 `protobuf:"varint,1000,opt,name=embeddedNamer,proto3,oneof" json:"-" gorm:"-"` // 系统字段，配置tablename
}

func (x *UserinfoDbModel) Reset() {
	*x = UserinfoDbModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_svr_protobuf_pbapi_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserinfoDbModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserinfoDbModel) ProtoMessage() {}

func (x *UserinfoDbModel) ProtoReflect() protoreflect.Message {
	mi := &file_content_svr_protobuf_pbapi_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserinfoDbModel.ProtoReflect.Descriptor instead.
func (*UserinfoDbModel) Descriptor() ([]byte, []int) {
	return file_content_svr_protobuf_pbapi_api_proto_rawDescGZIP(), []int{3}
}

func (x *UserinfoDbModel) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *UserinfoDbModel) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *UserinfoDbModel) GetAppFlag() int32 {
	if x != nil && x.AppFlag != nil {
		return *x.AppFlag
	}
	return 0
}

func (x *UserinfoDbModel) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UserinfoDbModel) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *UserinfoDbModel) GetNickName() string {
	if x != nil && x.NickName != nil {
		return *x.NickName
	}
	return ""
}

func (x *UserinfoDbModel) GetOpenNickName() string {
	if x != nil && x.OpenNickName != nil {
		return *x.OpenNickName
	}
	return ""
}

func (x *UserinfoDbModel) GetContact() string {
	if x != nil && x.Contact != nil {
		return *x.Contact
	}
	return ""
}

func (x *UserinfoDbModel) GetPhoto() string {
	if x != nil && x.Photo != nil {
		return *x.Photo
	}
	return ""
}

func (x *UserinfoDbModel) GetOpenPhoto() string {
	if x != nil && x.OpenPhoto != nil {
		return *x.OpenPhoto
	}
	return ""
}

func (x *UserinfoDbModel) GetSource() int32 {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return 0
}

func (x *UserinfoDbModel) GetBirth() string {
	if x != nil && x.Birth != nil {
		return *x.Birth
	}
	return ""
}

func (x *UserinfoDbModel) GetGender() int32 {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return 0
}

func (x *UserinfoDbModel) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *UserinfoDbModel) GetLevel() string {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return ""
}

func (x *UserinfoDbModel) GetFocus() int32 {
	if x != nil && x.Focus != nil {
		return *x.Focus
	}
	return 0
}

func (x *UserinfoDbModel) GetFans() int32 {
	if x != nil && x.Fans != nil {
		return *x.Fans
	}
	return 0
}

func (x *UserinfoDbModel) GetProvince() string {
	if x != nil && x.Province != nil {
		return *x.Province
	}
	return ""
}

func (x *UserinfoDbModel) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *UserinfoDbModel) GetConstellation() string {
	if x != nil && x.Constellation != nil {
		return *x.Constellation
	}
	return ""
}

func (x *UserinfoDbModel) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UserinfoDbModel) GetScore() int32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *UserinfoDbModel) GetOrderQuantity() int32 {
	if x != nil && x.OrderQuantity != nil {
		return *x.OrderQuantity
	}
	return 0
}

func (x *UserinfoDbModel) GetEnabled() int32 {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return 0
}

func (x *UserinfoDbModel) GetWorksLock() int32 {
	if x != nil && x.WorksLock != nil {
		return *x.WorksLock
	}
	return 0
}

func (x *UserinfoDbModel) GetCommentLock() int32 {
	if x != nil && x.CommentLock != nil {
		return *x.CommentLock
	}
	return 0
}

func (x *UserinfoDbModel) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *UserinfoDbModel) GetCreateTime() string {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return ""
}

func (x *UserinfoDbModel) GetUpdateTime() string {
	if x != nil && x.UpdateTime != nil {
		return *x.UpdateTime
	}
	return ""
}

func (x *UserinfoDbModel) GetUpdateGenderCount() int32 {
	if x != nil && x.UpdateGenderCount != nil {
		return *x.UpdateGenderCount
	}
	return 0
}

func (x *UserinfoDbModel) GetFriendLock() int32 {
	if x != nil && x.FriendLock != nil {
		return *x.FriendLock
	}
	return 0
}

func (x *UserinfoDbModel) GetEmbeddedNamer() int32 {
	if x != nil && x.EmbeddedNamer != nil {
		return *x.EmbeddedNamer
	}
	return 0
}

var File_content_svr_protobuf_pbapi_api_proto protoreflect.FileDescriptor

var file_content_svr_protobuf_pbapi_api_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x76, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0a, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x2a, 0x0a, 0x10, 0x44, 0x65, 0x62, 0x75, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xda,
	0x01, 0x0a, 0x0e, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x70, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70,
	0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x75, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x75, 0x6b, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0xdf, 0x0b, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x44, 0x62, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x07, 0x61, 0x70, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x03, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x04, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20,
	0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x05, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e,
	0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x6e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0a, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x0c, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0d, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0e, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x0f, 0x52, 0x05, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a,
	0x04, 0x66, 0x61, 0x6e, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x48, 0x10, 0x52, 0x04, 0x66,
	0x61, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x48, 0x11, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x48, 0x12, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x13, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x14, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x15, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a,
	0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x05, 0x48, 0x16, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x48, 0x17, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x19, 0x20, 0x01, 0x28, 0x05, 0x48, 0x18, 0x52, 0x09,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x19, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63,
	0x6b, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x1a, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x1b, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x1c, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a,
	0x13, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x48, 0x1d, 0x52, 0x11, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x48, 0x1e, 0x52, 0x0a, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0d, 0x65, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x72, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x1f, 0x52, 0x0d, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x70, 0x70, 0x5f,
	0x66, 0x6c, 0x61, 0x67, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x66, 0x61, 0x6e, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x72, 0x42, 0x22, 0x5a,
	0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x76, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x61, 0x70, 0x69, 0x3b, 0x70, 0x62, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_content_svr_protobuf_pbapi_api_proto_rawDescOnce sync.Once
	file_content_svr_protobuf_pbapi_api_proto_rawDescData = file_content_svr_protobuf_pbapi_api_proto_rawDesc
)

func file_content_svr_protobuf_pbapi_api_proto_rawDescGZIP() []byte {
	file_content_svr_protobuf_pbapi_api_proto_rawDescOnce.Do(func() {
		file_content_svr_protobuf_pbapi_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_svr_protobuf_pbapi_api_proto_rawDescData)
	})
	return file_content_svr_protobuf_pbapi_api_proto_rawDescData
}

var file_content_svr_protobuf_pbapi_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_content_svr_protobuf_pbapi_api_proto_goTypes = []interface{}{
	(*BaseResp)(nil),         // 0: BaseResp
	(*DebugUserInfoReq)(nil), // 1: DebugUserInfoReq
	(*HttpHeaderInfo)(nil),   // 2: HttpHeaderInfo
	(*UserinfoDbModel)(nil),  // 3: UserinfoDbModel
}
var file_content_svr_protobuf_pbapi_api_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_content_svr_protobuf_pbapi_api_proto_init() }
func file_content_svr_protobuf_pbapi_api_proto_init() {
	if File_content_svr_protobuf_pbapi_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_content_svr_protobuf_pbapi_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_content_svr_protobuf_pbapi_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugUserInfoReq); i {
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
		file_content_svr_protobuf_pbapi_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpHeaderInfo); i {
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
		file_content_svr_protobuf_pbapi_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserinfoDbModel); i {
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
	file_content_svr_protobuf_pbapi_api_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_content_svr_protobuf_pbapi_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_content_svr_protobuf_pbapi_api_proto_goTypes,
		DependencyIndexes: file_content_svr_protobuf_pbapi_api_proto_depIdxs,
		MessageInfos:      file_content_svr_protobuf_pbapi_api_proto_msgTypes,
	}.Build()
	File_content_svr_protobuf_pbapi_api_proto = out.File
	file_content_svr_protobuf_pbapi_api_proto_rawDesc = nil
	file_content_svr_protobuf_pbapi_api_proto_goTypes = nil
	file_content_svr_protobuf_pbapi_api_proto_depIdxs = nil
}
