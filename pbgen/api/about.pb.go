// Code generated by protoc-gen-go.
// source: api/about.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/about.proto
	api/common.proto
	api/edit_problem.proto
	api/language.proto
	api/list_contests.proto
	api/list_problems.proto
	api/list_submissions.proto
	api/login.proto
	api/logout.proto
	api/ojinfo.proto
	api/register.proto
	api/show_problem.proto
	api/submit.proto
	api/upload_problem.proto
	api/webapi.proto

It has these top-level messages:
	AboutRequest
	AboutResponse
	Error
	EditProblemRequest
	EditProblemResponse
	Language
	ListContestsRequest
	ListContestsResponse
	ListProblemsRequest
	ListProblemsResponse
	ListSubmissionsRequest
	ListSubmissionsResponse
	LoginInitRequest
	LoginInitResponse
	LoginAuthRequest
	LoginAuthResponse
	LogoutRequest
	LogoutResponse
	OJInfo
	RegisterRequest
	RegisterResponse
	ShowProblemRequest
	Problem
	ShowProblemResponse
	SubmitRequest
	SubmitResponse
	UploadProblemRequest
	UploadProblemResponse
	WebRequest
	WebResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AboutRequest struct {
	NeedOjsList       bool `protobuf:"varint,1,opt,name=need_ojs_list,json=needOjsList" json:"need_ojs_list,omitempty"`
	NeedLanguagesList bool `protobuf:"varint,2,opt,name=need_languages_list,json=needLanguagesList" json:"need_languages_list,omitempty"`
}

func (m *AboutRequest) Reset()                    { *m = AboutRequest{} }
func (m *AboutRequest) String() string            { return proto.CompactTextString(m) }
func (*AboutRequest) ProtoMessage()               {}
func (*AboutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AboutRequest) GetNeedOjsList() bool {
	if m != nil {
		return m.NeedOjsList
	}
	return false
}

func (m *AboutRequest) GetNeedLanguagesList() bool {
	if m != nil {
		return m.NeedLanguagesList
	}
	return false
}

type AboutResponse struct {
	OjsList       []*OJInfo   `protobuf:"bytes,1,rep,name=ojs_list,json=ojsList" json:"ojs_list,omitempty"`
	LanguagesList []*Language `protobuf:"bytes,2,rep,name=languages_list,json=languagesList" json:"languages_list,omitempty"`
	Error         *Error      `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *AboutResponse) Reset()                    { *m = AboutResponse{} }
func (m *AboutResponse) String() string            { return proto.CompactTextString(m) }
func (*AboutResponse) ProtoMessage()               {}
func (*AboutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AboutResponse) GetOjsList() []*OJInfo {
	if m != nil {
		return m.OjsList
	}
	return nil
}

func (m *AboutResponse) GetLanguagesList() []*Language {
	if m != nil {
		return m.LanguagesList
	}
	return nil
}

func (m *AboutResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*AboutRequest)(nil), "api.AboutRequest")
	proto.RegisterType((*AboutResponse)(nil), "api.AboutResponse")
}

func init() { proto.RegisterFile("api/about.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8f, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0x87, 0xa9, 0x45, 0x5d, 0xa6, 0xc6, 0x3f, 0xf1, 0xb2, 0xec, 0xa9, 0xf4, 0x20, 0x3d, 0x45,
	0x58, 0x7d, 0x01, 0x0f, 0x1e, 0x94, 0x85, 0x85, 0xbc, 0xc0, 0x92, 0x6a, 0x76, 0x49, 0xe9, 0x66,
	0x62, 0x92, 0x3e, 0x87, 0xaf, 0x2c, 0x93, 0xa6, 0x50, 0xf0, 0xfa, 0xfd, 0x3e, 0xe6, 0x63, 0xe0,
	0x4e, 0x39, 0xf3, 0xac, 0x3a, 0x1c, 0xa3, 0x70, 0x1e, 0x23, 0xf2, 0x52, 0x39, 0xb3, 0xb9, 0x27,
	0xfa, 0x85, 0xe7, 0x33, 0xda, 0x09, 0x6f, 0x38, 0x91, 0x41, 0xd9, 0xd3, 0xa8, 0x4e, 0x3a, 0xb3,
	0x64, 0x61, 0x6f, 0xec, 0x11, 0x27, 0xd2, 0x74, 0x70, 0xf3, 0x46, 0xb7, 0xa4, 0xfe, 0x19, 0x75,
	0x88, 0xbc, 0x01, 0x66, 0xb5, 0xfe, 0x3e, 0x60, 0x1f, 0x0e, 0x83, 0x09, 0x71, 0x5d, 0xd4, 0x45,
	0xbb, 0x92, 0x15, 0xc1, 0x7d, 0x1f, 0x76, 0x26, 0x44, 0x2e, 0xe0, 0x31, 0x39, 0xf3, 0xf1, 0x6c,
	0x5e, 0x24, 0xf3, 0x81, 0xa6, 0xdd, 0xbc, 0x90, 0xdf, 0xfc, 0x16, 0xc0, 0x72, 0x24, 0x38, 0xb4,
	0x41, 0xf3, 0x27, 0x58, 0x2d, 0x02, 0x65, 0x5b, 0x6d, 0x2b, 0xa1, 0x9c, 0x11, 0xfb, 0xcf, 0x0f,
	0x7b, 0x44, 0x79, 0x8d, 0xb9, 0xf4, 0x0a, 0xb7, 0xff, 0x22, 0x64, 0xb3, 0x64, 0xcf, 0x15, 0xc9,
	0x86, 0x65, 0x8f, 0xd7, 0x70, 0xa9, 0xbd, 0x47, 0xbf, 0x2e, 0xeb, 0xa2, 0xad, 0xb6, 0x90, 0xe4,
	0x77, 0x22, 0x72, 0x1a, 0xba, 0xab, 0xf4, 0xfc, 0xcb, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2,
	0x79, 0x5b, 0x8b, 0x4c, 0x01, 0x00, 0x00,
}
