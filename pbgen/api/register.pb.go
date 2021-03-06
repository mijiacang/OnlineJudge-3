// Code generated by protoc-gen-go.
// source: api/register.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// username, golang `^[a-zA-Z][0-9a-zA-Z_]{3,19}$`
// email, golang `^[a-zA-Z0-9._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$`
// phone,
// school,
// motto, less than 60 words
type RegisterRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	School   string `protobuf:"bytes,5,opt,name=school" json:"school,omitempty"`
	Motto    string `protobuf:"bytes,6,opt,name=motto" json:"motto,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegisterRequest) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *RegisterRequest) GetMotto() string {
	if m != nil {
		return m.Motto
	}
	return ""
}

// if not empty, something is wrong.
type RegisterResponse struct {
	UserId        int64  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Username      string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	CheckUsername string `protobuf:"bytes,3,opt,name=check_username,json=checkUsername" json:"check_username,omitempty"`
	CheckPassword string `protobuf:"bytes,4,opt,name=check_password,json=checkPassword" json:"check_password,omitempty"`
	CheckEmail    string `protobuf:"bytes,5,opt,name=check_email,json=checkEmail" json:"check_email,omitempty"`
	CheckPhone    string `protobuf:"bytes,6,opt,name=check_phone,json=checkPhone" json:"check_phone,omitempty"`
	CheckSchool   string `protobuf:"bytes,7,opt,name=check_school,json=checkSchool" json:"check_school,omitempty"`
	CheckMotto    string `protobuf:"bytes,8,opt,name=check_motto,json=checkMotto" json:"check_motto,omitempty"`
	Error         *Error `protobuf:"bytes,9,opt,name=error" json:"error,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *RegisterResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RegisterResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterResponse) GetCheckUsername() string {
	if m != nil {
		return m.CheckUsername
	}
	return ""
}

func (m *RegisterResponse) GetCheckPassword() string {
	if m != nil {
		return m.CheckPassword
	}
	return ""
}

func (m *RegisterResponse) GetCheckEmail() string {
	if m != nil {
		return m.CheckEmail
	}
	return ""
}

func (m *RegisterResponse) GetCheckPhone() string {
	if m != nil {
		return m.CheckPhone
	}
	return ""
}

func (m *RegisterResponse) GetCheckSchool() string {
	if m != nil {
		return m.CheckSchool
	}
	return ""
}

func (m *RegisterResponse) GetCheckMotto() string {
	if m != nil {
		return m.CheckMotto
	}
	return ""
}

func (m *RegisterResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "api.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "api.RegisterResponse")
}

func init() { proto.RegisterFile("api/register.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x49, 0xd2, 0xa4, 0xed, 0xad, 0x3f, 0x65, 0x10, 0x1d, 0xba, 0x31, 0x16, 0x84, 0xae,
	0x22, 0xe8, 0x33, 0x74, 0xe1, 0x42, 0x28, 0x11, 0xd7, 0x65, 0x4c, 0x06, 0x33, 0xd8, 0xe4, 0x8e,
	0x33, 0x29, 0x3e, 0x90, 0xcf, 0xe1, 0xbb, 0xc9, 0xcc, 0x4d, 0x63, 0xb4, 0xcb, 0x73, 0xee, 0x97,
	0xe1, 0x9c, 0x13, 0x60, 0x42, 0xab, 0x3b, 0x23, 0xdf, 0x94, 0x6d, 0xa5, 0xc9, 0xb4, 0xc1, 0x16,
	0x59, 0x24, 0xb4, 0x5a, 0xcc, 0xdd, 0xa1, 0xc0, 0xba, 0xc6, 0x86, 0xec, 0xe5, 0x57, 0x00, 0xe7,
	0x79, 0x47, 0xe6, 0xf2, 0x63, 0x2f, 0x6d, 0xcb, 0x16, 0x30, 0xd9, 0x5b, 0x69, 0x1a, 0x51, 0x4b,
	0x1e, 0xa4, 0xc1, 0x6a, 0x9a, 0xf7, 0xda, 0xdd, 0xb4, 0xb0, 0xf6, 0x13, 0x4d, 0xc9, 0x43, 0xba,
	0x1d, 0x34, 0xbb, 0x80, 0x58, 0xd6, 0x42, 0xed, 0x78, 0xe4, 0x0f, 0x24, 0x9c, 0xab, 0x2b, 0x6c,
	0x24, 0x1f, 0x91, 0xeb, 0x05, 0xbb, 0x84, 0xc4, 0x16, 0x15, 0xe2, 0x8e, 0xc7, 0xde, 0xee, 0x94,
	0xa3, 0x6b, 0x6c, 0x5b, 0xe4, 0x09, 0xd1, 0x5e, 0x2c, 0xbf, 0x43, 0x98, 0xff, 0xa6, 0xb4, 0x1a,
	0x1b, 0x2b, 0xd9, 0x15, 0x8c, 0x5d, 0xac, 0xad, 0x2a, 0x7d, 0xca, 0x28, 0x4f, 0x9c, 0x7c, 0x2c,
	0xff, 0xe4, 0x0f, 0xff, 0xe5, 0xbf, 0x85, 0xb3, 0xa2, 0x92, 0xc5, 0xfb, 0xb6, 0x27, 0x28, 0xec,
	0xa9, 0x77, 0x5f, 0x8e, 0xb0, 0xbe, 0xec, 0x68, 0x80, 0x6d, 0x0e, 0x8d, 0xaf, 0x61, 0x46, 0x18,
	0xf5, 0xa6, 0x2a, 0xe0, 0xad, 0xb5, 0x2f, 0xdf, 0x03, 0x34, 0x41, 0x32, 0x00, 0x36, 0x7e, 0x87,
	0x1b, 0x38, 0x21, 0xa0, 0x5b, 0x63, 0xec, 0x09, 0xfa, 0xe8, 0x99, 0x26, 0xe9, 0xdf, 0xa0, 0x61,
	0x26, 0x83, 0x37, 0x9e, 0x9c, 0xc3, 0x52, 0x88, 0xa5, 0x31, 0x68, 0xf8, 0x34, 0x0d, 0x56, 0xb3,
	0x7b, 0xc8, 0x84, 0x56, 0xd9, 0xda, 0x39, 0x39, 0x1d, 0x5e, 0x13, 0xff, 0xb3, 0x1f, 0x7e, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa1, 0xd3, 0xd5, 0x26, 0x19, 0x02, 0x00, 0x00,
}
