// Code generated by protoc-gen-go.
// source: api/login.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LoginInitRequest struct {
	Msg     string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *LoginInitRequest) Reset()                    { *m = LoginInitRequest{} }
func (m *LoginInitRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginInitRequest) ProtoMessage()               {}
func (*LoginInitRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *LoginInitRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LoginInitRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type LoginInitResponse struct {
	Salt      []byte `protobuf:"bytes,1,opt,name=salt,proto3" json:"salt,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	IsCaptcha bool   `protobuf:"varint,3,opt,name=is_captcha,json=isCaptcha" json:"is_captcha,omitempty"`
	Error     *Error `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *LoginInitResponse) Reset()                    { *m = LoginInitResponse{} }
func (m *LoginInitResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginInitResponse) ProtoMessage()               {}
func (*LoginInitResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *LoginInitResponse) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *LoginInitResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LoginInitResponse) GetIsCaptcha() bool {
	if m != nil {
		return m.IsCaptcha
	}
	return false
}

func (m *LoginInitResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type LoginAuthRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Captcha  string `protobuf:"bytes,3,opt,name=captcha" json:"captcha,omitempty"`
}

func (m *LoginAuthRequest) Reset()                    { *m = LoginAuthRequest{} }
func (m *LoginAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginAuthRequest) ProtoMessage()               {}
func (*LoginAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *LoginAuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginAuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginAuthRequest) GetCaptcha() string {
	if m != nil {
		return m.Captcha
	}
	return ""
}

type LoginAuthResponse struct {
	Msg       string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Privilege string `protobuf:"bytes,3,opt,name=privilege" json:"privilege,omitempty"`
	Error     *Error `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *LoginAuthResponse) Reset()                    { *m = LoginAuthResponse{} }
func (m *LoginAuthResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginAuthResponse) ProtoMessage()               {}
func (*LoginAuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *LoginAuthResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LoginAuthResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginAuthResponse) GetPrivilege() string {
	if m != nil {
		return m.Privilege
	}
	return ""
}

func (m *LoginAuthResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginInitRequest)(nil), "api.LoginInitRequest")
	proto.RegisterType((*LoginInitResponse)(nil), "api.LoginInitResponse")
	proto.RegisterType((*LoginAuthRequest)(nil), "api.LoginAuthRequest")
	proto.RegisterType((*LoginAuthResponse)(nil), "api.LoginAuthResponse")
}

func init() { proto.RegisterFile("api/login.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xe9, 0xee, 0xaa, 0xdb, 0x51, 0xb0, 0xe6, 0x14, 0x16, 0x85, 0xd2, 0x53, 0x4f, 0x15,
	0xf4, 0x2e, 0x88, 0x78, 0x10, 0x3c, 0xe5, 0x0f, 0x48, 0xec, 0x86, 0xee, 0x40, 0x9b, 0xc4, 0x4c,
	0xba, 0x9e, 0xf5, 0x97, 0x4b, 0xda, 0x74, 0x2d, 0x8a, 0x78, 0x9b, 0xf7, 0xa6, 0xbc, 0x79, 0x5f,
	0x03, 0xe7, 0xd2, 0xe2, 0x75, 0x6b, 0x1a, 0xd4, 0x95, 0x75, 0xc6, 0x1b, 0xb6, 0x94, 0x16, 0x37,
	0x59, 0x70, 0x6b, 0xd3, 0x75, 0x26, 0xda, 0xc5, 0x1d, 0x64, 0xcf, 0xe1, 0xab, 0x27, 0x8d, 0x5e,
	0xa8, 0xb7, 0x5e, 0x91, 0x67, 0x19, 0x2c, 0x3b, 0x6a, 0x78, 0x92, 0x27, 0x65, 0x2a, 0xc2, 0xc8,
	0x38, 0x9c, 0xec, 0x95, 0x23, 0x34, 0x9a, 0x2f, 0x06, 0x77, 0x92, 0xc5, 0x67, 0x02, 0x17, 0xb3,
	0x00, 0xb2, 0x46, 0x93, 0x62, 0x0c, 0x56, 0x24, 0x5b, 0x3f, 0x44, 0x9c, 0x89, 0x61, 0xfe, 0x3b,
	0x83, 0x5d, 0x01, 0x20, 0xbd, 0xd4, 0xd2, 0xfa, 0x7a, 0x27, 0xf9, 0x32, 0x4f, 0xca, 0xb5, 0x48,
	0x91, 0x1e, 0x46, 0x83, 0xe5, 0x70, 0xa4, 0x9c, 0x33, 0x8e, 0xaf, 0xf2, 0xa4, 0x3c, 0xbd, 0x81,
	0x4a, 0x5a, 0xac, 0x1e, 0x83, 0x23, 0xc6, 0x45, 0xb1, 0x8d, 0x10, 0xf7, 0xbd, 0xdf, 0x4d, 0x10,
	0x1b, 0x58, 0xf7, 0xa4, 0x9c, 0x96, 0x9d, 0x8a, 0x24, 0x07, 0x1d, 0x76, 0x56, 0x12, 0xbd, 0x1b,
	0xb7, 0x8d, 0x5d, 0x0e, 0x3a, 0xd4, 0x9c, 0x37, 0x49, 0xc5, 0x24, 0x8b, 0x8f, 0x09, 0x75, 0x3c,
	0x13, 0x51, 0x7f, 0xff, 0xac, 0xf9, 0xe5, 0xc5, 0x8f, 0xcb, 0x97, 0x90, 0x5a, 0x87, 0x7b, 0x6c,
	0x55, 0xa3, 0x62, 0xfe, 0xb7, 0xf1, 0x3f, 0xe9, 0xeb, 0xf1, 0xf0, 0x6a, 0xb7, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xf9, 0xf7, 0x84, 0x9f, 0xdf, 0x01, 0x00, 0x00,
}
