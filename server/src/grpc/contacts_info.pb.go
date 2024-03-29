// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contacts_info.proto

package com_bc_contacts_protobuffer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//群
type GroupInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon                 string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupInfo) Reset()         { *m = GroupInfo{} }
func (m *GroupInfo) String() string { return proto.CompactTextString(m) }
func (*GroupInfo) ProtoMessage()    {}
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{0}
}

func (m *GroupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupInfo.Unmarshal(m, b)
}
func (m *GroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupInfo.Marshal(b, m, deterministic)
}
func (m *GroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupInfo.Merge(m, src)
}
func (m *GroupInfo) XXX_Size() int {
	return xxx_messageInfo_GroupInfo.Size(m)
}
func (m *GroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GroupInfo proto.InternalMessageInfo

func (m *GroupInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GroupInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GroupInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

//联系人
type Contacts struct {
	NickName             string   `protobuf:"bytes,1,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Remarks              string   `protobuf:"bytes,2,opt,name=remarks,proto3" json:"remarks,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contacts) Reset()         { *m = Contacts{} }
func (m *Contacts) String() string { return proto.CompactTextString(m) }
func (*Contacts) ProtoMessage()    {}
func (*Contacts) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{1}
}

func (m *Contacts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contacts.Unmarshal(m, b)
}
func (m *Contacts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contacts.Marshal(b, m, deterministic)
}
func (m *Contacts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contacts.Merge(m, src)
}
func (m *Contacts) XXX_Size() int {
	return xxx_messageInfo_Contacts.Size(m)
}
func (m *Contacts) XXX_DiscardUnknown() {
	xxx_messageInfo_Contacts.DiscardUnknown(m)
}

var xxx_messageInfo_Contacts proto.InternalMessageInfo

func (m *Contacts) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *Contacts) GetRemarks() string {
	if m != nil {
		return m.Remarks
	}
	return ""
}

func (m *Contacts) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Contacts) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//新朋友
type NewFriend struct {
	NickName             string   `protobuf:"bytes,1,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Avatar               string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewFriend) Reset()         { *m = NewFriend{} }
func (m *NewFriend) String() string { return proto.CompactTextString(m) }
func (*NewFriend) ProtoMessage()    {}
func (*NewFriend) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{2}
}

func (m *NewFriend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewFriend.Unmarshal(m, b)
}
func (m *NewFriend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewFriend.Marshal(b, m, deterministic)
}
func (m *NewFriend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewFriend.Merge(m, src)
}
func (m *NewFriend) XXX_Size() int {
	return xxx_messageInfo_NewFriend.Size(m)
}
func (m *NewFriend) XXX_DiscardUnknown() {
	xxx_messageInfo_NewFriend.DiscardUnknown(m)
}

var xxx_messageInfo_NewFriend proto.InternalMessageInfo

func (m *NewFriend) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *NewFriend) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *NewFriend) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//关注人
type FocusPeople struct {
	NickName             string   `protobuf:"bytes,1,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Avatar               string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FocusPeople) Reset()         { *m = FocusPeople{} }
func (m *FocusPeople) String() string { return proto.CompactTextString(m) }
func (*FocusPeople) ProtoMessage()    {}
func (*FocusPeople) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{3}
}

func (m *FocusPeople) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FocusPeople.Unmarshal(m, b)
}
func (m *FocusPeople) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FocusPeople.Marshal(b, m, deterministic)
}
func (m *FocusPeople) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FocusPeople.Merge(m, src)
}
func (m *FocusPeople) XXX_Size() int {
	return xxx_messageInfo_FocusPeople.Size(m)
}
func (m *FocusPeople) XXX_DiscardUnknown() {
	xxx_messageInfo_FocusPeople.DiscardUnknown(m)
}

var xxx_messageInfo_FocusPeople proto.InternalMessageInfo

func (m *FocusPeople) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *FocusPeople) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *FocusPeople) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//群成员信息
type GroupMember struct {
	NickName             string     `protobuf:"bytes,1,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Avatar               string     `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Id                   string     `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Status               UserStatus `protobuf:"varint,4,opt,name=status,proto3,enum=com.bc.enum.protobuffer.UserStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GroupMember) Reset()         { *m = GroupMember{} }
func (m *GroupMember) String() string { return proto.CompactTextString(m) }
func (*GroupMember) ProtoMessage()    {}
func (*GroupMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{4}
}

func (m *GroupMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupMember.Unmarshal(m, b)
}
func (m *GroupMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupMember.Marshal(b, m, deterministic)
}
func (m *GroupMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupMember.Merge(m, src)
}
func (m *GroupMember) XXX_Size() int {
	return xxx_messageInfo_GroupMember.Size(m)
}
func (m *GroupMember) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupMember.DiscardUnknown(m)
}

var xxx_messageInfo_GroupMember proto.InternalMessageInfo

func (m *GroupMember) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *GroupMember) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *GroupMember) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GroupMember) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_ONLINE
}

//请求通讯录列表页
type RequestContacts struct {
	Token                *UserToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RequestContacts) Reset()         { *m = RequestContacts{} }
func (m *RequestContacts) String() string { return proto.CompactTextString(m) }
func (*RequestContacts) ProtoMessage()    {}
func (*RequestContacts) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{5}
}

func (m *RequestContacts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestContacts.Unmarshal(m, b)
}
func (m *RequestContacts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestContacts.Marshal(b, m, deterministic)
}
func (m *RequestContacts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestContacts.Merge(m, src)
}
func (m *RequestContacts) XXX_Size() int {
	return xxx_messageInfo_RequestContacts.Size(m)
}
func (m *RequestContacts) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestContacts.DiscardUnknown(m)
}

var xxx_messageInfo_RequestContacts proto.InternalMessageInfo

func (m *RequestContacts) GetToken() *UserToken {
	if m != nil {
		return m.Token
	}
	return nil
}

//通讯录列表页返回结果
type ResponseContacts struct {
	Code                 *ResponseCode  `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	NewFriendList        []*NewFriend   `protobuf:"bytes,2,rep,name=new_friend_list,json=newFriendList,proto3" json:"new_friend_list,omitempty"`
	GroupList            []*GroupInfo   `protobuf:"bytes,3,rep,name=group_list,json=groupList,proto3" json:"group_list,omitempty"`
	FocusPeopleList      []*FocusPeople `protobuf:"bytes,4,rep,name=focus_people_list,json=focusPeopleList,proto3" json:"focus_people_list,omitempty"`
	ContactsList         []*Contacts    `protobuf:"bytes,5,rep,name=contacts_list,json=contactsList,proto3" json:"contacts_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResponseContacts) Reset()         { *m = ResponseContacts{} }
func (m *ResponseContacts) String() string { return proto.CompactTextString(m) }
func (*ResponseContacts) ProtoMessage()    {}
func (*ResponseContacts) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{6}
}

func (m *ResponseContacts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseContacts.Unmarshal(m, b)
}
func (m *ResponseContacts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseContacts.Marshal(b, m, deterministic)
}
func (m *ResponseContacts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseContacts.Merge(m, src)
}
func (m *ResponseContacts) XXX_Size() int {
	return xxx_messageInfo_ResponseContacts.Size(m)
}
func (m *ResponseContacts) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseContacts.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseContacts proto.InternalMessageInfo

func (m *ResponseContacts) GetCode() *ResponseCode {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ResponseContacts) GetNewFriendList() []*NewFriend {
	if m != nil {
		return m.NewFriendList
	}
	return nil
}

func (m *ResponseContacts) GetGroupList() []*GroupInfo {
	if m != nil {
		return m.GroupList
	}
	return nil
}

func (m *ResponseContacts) GetFocusPeopleList() []*FocusPeople {
	if m != nil {
		return m.FocusPeopleList
	}
	return nil
}

func (m *ResponseContacts) GetContactsList() []*Contacts {
	if m != nil {
		return m.ContactsList
	}
	return nil
}

//请求群成员列表
type RequestGroupMember struct {
	Token                *UserToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	GroupId              string     `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RequestGroupMember) Reset()         { *m = RequestGroupMember{} }
func (m *RequestGroupMember) String() string { return proto.CompactTextString(m) }
func (*RequestGroupMember) ProtoMessage()    {}
func (*RequestGroupMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{7}
}

func (m *RequestGroupMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestGroupMember.Unmarshal(m, b)
}
func (m *RequestGroupMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestGroupMember.Marshal(b, m, deterministic)
}
func (m *RequestGroupMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestGroupMember.Merge(m, src)
}
func (m *RequestGroupMember) XXX_Size() int {
	return xxx_messageInfo_RequestGroupMember.Size(m)
}
func (m *RequestGroupMember) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestGroupMember.DiscardUnknown(m)
}

var xxx_messageInfo_RequestGroupMember proto.InternalMessageInfo

func (m *RequestGroupMember) GetToken() *UserToken {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *RequestGroupMember) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

//返回群成员列表
type ResponseGroupMember struct {
	Code                 *ResponseCode  `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Members              []*GroupMember `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResponseGroupMember) Reset()         { *m = ResponseGroupMember{} }
func (m *ResponseGroupMember) String() string { return proto.CompactTextString(m) }
func (*ResponseGroupMember) ProtoMessage()    {}
func (*ResponseGroupMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{8}
}

func (m *ResponseGroupMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseGroupMember.Unmarshal(m, b)
}
func (m *ResponseGroupMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseGroupMember.Marshal(b, m, deterministic)
}
func (m *ResponseGroupMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseGroupMember.Merge(m, src)
}
func (m *ResponseGroupMember) XXX_Size() int {
	return xxx_messageInfo_ResponseGroupMember.Size(m)
}
func (m *ResponseGroupMember) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseGroupMember.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseGroupMember proto.InternalMessageInfo

func (m *ResponseGroupMember) GetCode() *ResponseCode {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ResponseGroupMember) GetMembers() []*GroupMember {
	if m != nil {
		return m.Members
	}
	return nil
}

//请求添加好友
type RequestAddFriend struct {
	Token                *UserToken   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Type                 AccountsType `protobuf:"varint,2,opt,name=type,proto3,enum=com.bc.enum.protobuffer.AccountsType" json:"type,omitempty"`
	Accounts             string       `protobuf:"bytes,3,opt,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RequestAddFriend) Reset()         { *m = RequestAddFriend{} }
func (m *RequestAddFriend) String() string { return proto.CompactTextString(m) }
func (*RequestAddFriend) ProtoMessage()    {}
func (*RequestAddFriend) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{9}
}

func (m *RequestAddFriend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAddFriend.Unmarshal(m, b)
}
func (m *RequestAddFriend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAddFriend.Marshal(b, m, deterministic)
}
func (m *RequestAddFriend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAddFriend.Merge(m, src)
}
func (m *RequestAddFriend) XXX_Size() int {
	return xxx_messageInfo_RequestAddFriend.Size(m)
}
func (m *RequestAddFriend) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAddFriend.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAddFriend proto.InternalMessageInfo

func (m *RequestAddFriend) GetToken() *UserToken {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *RequestAddFriend) GetType() AccountsType {
	if m != nil {
		return m.Type
	}
	return AccountsType_PHONE_NUMBER
}

func (m *RequestAddFriend) GetAccounts() string {
	if m != nil {
		return m.Accounts
	}
	return ""
}

//返回添加结果
type ResponseAddFriend struct {
	Code                 *ResponseCode `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ResponseAddFriend) Reset()         { *m = ResponseAddFriend{} }
func (m *ResponseAddFriend) String() string { return proto.CompactTextString(m) }
func (*ResponseAddFriend) ProtoMessage()    {}
func (*ResponseAddFriend) Descriptor() ([]byte, []int) {
	return fileDescriptor_f820a11efb13ada5, []int{10}
}

func (m *ResponseAddFriend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseAddFriend.Unmarshal(m, b)
}
func (m *ResponseAddFriend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseAddFriend.Marshal(b, m, deterministic)
}
func (m *ResponseAddFriend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseAddFriend.Merge(m, src)
}
func (m *ResponseAddFriend) XXX_Size() int {
	return xxx_messageInfo_ResponseAddFriend.Size(m)
}
func (m *ResponseAddFriend) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseAddFriend.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseAddFriend proto.InternalMessageInfo

func (m *ResponseAddFriend) GetCode() *ResponseCode {
	if m != nil {
		return m.Code
	}
	return nil
}

func init() {
	proto.RegisterType((*GroupInfo)(nil), "com.bc.contacts.protobuffer.group_info")
	proto.RegisterType((*Contacts)(nil), "com.bc.contacts.protobuffer.contacts")
	proto.RegisterType((*NewFriend)(nil), "com.bc.contacts.protobuffer.new_friend")
	proto.RegisterType((*FocusPeople)(nil), "com.bc.contacts.protobuffer.focus_people")
	proto.RegisterType((*GroupMember)(nil), "com.bc.contacts.protobuffer.group_member")
	proto.RegisterType((*RequestContacts)(nil), "com.bc.contacts.protobuffer.request_contacts")
	proto.RegisterType((*ResponseContacts)(nil), "com.bc.contacts.protobuffer.response_contacts")
	proto.RegisterType((*RequestGroupMember)(nil), "com.bc.contacts.protobuffer.request_group_member")
	proto.RegisterType((*ResponseGroupMember)(nil), "com.bc.contacts.protobuffer.response_group_member")
	proto.RegisterType((*RequestAddFriend)(nil), "com.bc.contacts.protobuffer.request_add_friend")
	proto.RegisterType((*ResponseAddFriend)(nil), "com.bc.contacts.protobuffer.response_add_friend")
}

func init() { proto.RegisterFile("contacts_info.proto", fileDescriptor_f820a11efb13ada5) }

var fileDescriptor_f820a11efb13ada5 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x95, 0x8f, 0xa6, 0xc9, 0x34, 0x4d, 0xc9, 0x06, 0x50, 0x48, 0x2f, 0x95, 0xc5, 0x47,
	0xcb, 0xc1, 0x87, 0x70, 0xab, 0x10, 0x17, 0x50, 0x25, 0x10, 0x2a, 0x60, 0xd4, 0xb3, 0xe5, 0xac,
	0xd7, 0xd5, 0x2a, 0xf1, 0xae, 0xd9, 0x5d, 0x53, 0xf1, 0x14, 0x88, 0x3b, 0x0f, 0xc0, 0x63, 0x22,
	0xcf, 0xee, 0xda, 0x89, 0xa0, 0x51, 0xa5, 0xf4, 0xb6, 0x33, 0xde, 0xff, 0x6f, 0x47, 0xf3, 0x9f,
	0x31, 0x4c, 0xa8, 0x14, 0x26, 0xa1, 0x46, 0xc7, 0x5c, 0x64, 0x32, 0x2c, 0x94, 0x34, 0x92, 0x1c,
	0x53, 0x99, 0x87, 0x0b, 0x1a, 0xfa, 0x6f, 0x36, 0xbd, 0x28, 0xb3, 0x8c, 0xa9, 0xd9, 0x04, 0x03,
	0x2a, 0x57, 0x31, 0x95, 0x29, 0xb3, 0x9f, 0x66, 0x64, 0x25, 0xaf, 0xb9, 0x88, 0xb5, 0x49, 0x4c,
	0xe9, 0xae, 0xcf, 0xc6, 0x4c, 0x94, 0x79, 0x9c, 0xb2, 0x8c, 0x0b, 0x77, 0x2d, 0x78, 0x07, 0x70,
	0xad, 0x64, 0x59, 0xe0, 0x63, 0x64, 0x04, 0x6d, 0x9e, 0x4e, 0x5b, 0x27, 0xad, 0xd3, 0x41, 0xd4,
	0xe6, 0x29, 0x21, 0xd0, 0x15, 0x49, 0xce, 0xa6, 0x6d, 0xcc, 0xe0, 0xb9, 0xca, 0x71, 0x2a, 0xc5,
	0xb4, 0x63, 0x73, 0xd5, 0x39, 0xe0, 0xd0, 0xf7, 0x95, 0x91, 0x63, 0x18, 0x08, 0x4e, 0x97, 0x31,
	0x0a, 0x2d, 0xaa, 0x5f, 0x25, 0x2e, 0x2b, 0xf1, 0x14, 0xf6, 0x15, 0xcb, 0x13, 0xb5, 0xd4, 0x8e,
	0xe9, 0x43, 0xf2, 0x18, 0x7a, 0xc9, 0xf7, 0xc4, 0x24, 0xca, 0x81, 0x5d, 0xe4, 0x4a, 0xea, 0xfa,
	0x92, 0x82, 0x2f, 0x00, 0x82, 0xdd, 0xc4, 0x99, 0xe2, 0x4c, 0xa4, 0xdb, 0x1f, 0x6b, 0x90, 0xed,
	0xff, 0x20, 0x3b, 0x35, 0xf2, 0x2b, 0x0c, 0x33, 0x49, 0x4b, 0x1d, 0x17, 0x4c, 0x16, 0x2b, 0x76,
	0x3f, 0xd0, 0x5f, 0x2d, 0x18, 0xda, 0xce, 0xe6, 0x2c, 0x5f, 0x30, 0x75, 0x2f, 0x54, 0xf2, 0x1a,
	0x7a, 0xd6, 0x51, 0xec, 0xc8, 0x68, 0xfe, 0x34, 0x74, 0x83, 0x51, 0x39, 0xbb, 0x3e, 0x14, 0x61,
	0xa9, 0x99, 0x72, 0xee, 0x47, 0x4e, 0x13, 0x5c, 0xc2, 0x03, 0xc5, 0xbe, 0x95, 0x4c, 0x9b, 0xb8,
	0xb6, 0xeb, 0x1c, 0xf6, 0x8c, 0x5c, 0x32, 0x81, 0x25, 0x1d, 0x34, 0x40, 0x1c, 0x9f, 0x7f, 0x89,
	0x78, 0x37, 0xb2, 0x92, 0xe0, 0x67, 0x07, 0xc6, 0x8a, 0xe9, 0x42, 0x0a, 0xcd, 0x1a, 0xe2, 0x1b,
	0xe8, 0x56, 0x73, 0xe8, 0x80, 0x2f, 0x3d, 0xd0, 0x0f, 0xe9, 0x06, 0x73, 0x4d, 0x9d, 0xb2, 0x08,
	0x75, 0xe4, 0x13, 0x1c, 0x35, 0x0e, 0xc7, 0x2b, 0xae, 0xcd, 0xb4, 0x7d, 0xd2, 0x39, 0x3d, 0x98,
	0xbf, 0x08, 0xb7, 0x6c, 0x41, 0xd8, 0x68, 0xa2, 0x43, 0xc1, 0x6e, 0x2e, 0xf0, 0xf8, 0x91, 0x6b,
	0x43, 0x2e, 0xfc, 0x8c, 0x23, 0xab, 0x73, 0x07, 0x56, 0xb3, 0x12, 0xd1, 0x00, 0xcf, 0xc8, 0xb9,
	0x82, 0xf1, 0xfa, 0x9c, 0x58, 0x5c, 0x17, 0x71, 0x67, 0x5b, 0x71, 0xeb, 0xaa, 0xe8, 0x08, 0xa3,
	0xcf, 0x18, 0x20, 0xf6, 0x03, 0x1c, 0xd6, 0x2b, 0x8f, 0xc8, 0x3d, 0x44, 0x3e, 0xdb, 0x8a, 0xf4,
	0xc9, 0x68, 0xe8, 0x4f, 0x15, 0x2b, 0xc8, 0xe1, 0xa1, 0x77, 0x78, 0x63, 0xf8, 0x76, 0x70, 0x99,
	0x3c, 0x81, 0xbe, 0xeb, 0x47, 0xea, 0x97, 0x16, 0xe3, 0xf7, 0x69, 0xf0, 0xbb, 0x05, 0x8f, 0x6a,
	0x0b, 0x37, 0x1e, 0xdc, 0x75, 0x08, 0xde, 0xc2, 0xbe, 0x25, 0x69, 0x67, 0xfe, 0xd9, 0x1d, 0x0c,
	0xb3, 0x8a, 0xc8, 0x2b, 0x83, 0x3f, 0x2d, 0x20, 0xbe, 0x1d, 0x49, 0x9a, 0xfa, 0x9f, 0xc6, 0x2e,
	0xcd, 0x38, 0x87, 0xae, 0xf9, 0x51, 0xd8, 0x3f, 0xe2, 0x68, 0xfe, 0xfc, 0xd6, 0xf5, 0x4b, 0x28,
	0x95, 0xa5, 0x30, 0x3a, 0xae, 0x6e, 0x47, 0xa8, 0x21, 0x33, 0xe8, 0xfb, 0xb4, 0x5b, 0xe9, 0x3a,
	0x0e, 0xae, 0x60, 0x52, 0xb7, 0x61, 0xad, 0xd4, 0x1d, 0xdb, 0xb8, 0xe8, 0xe1, 0x85, 0x57, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x2b, 0xc8, 0x17, 0x55, 0x06, 0x00, 0x00,
}
