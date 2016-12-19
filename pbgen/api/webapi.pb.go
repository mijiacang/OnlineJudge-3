// Code generated by protoc-gen-go.
// source: api/webapi.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WebRequest struct {
	CsrfToken              string                  `protobuf:"bytes,1,opt,name=csrf_token,json=csrfToken" json:"csrf_token,omitempty"`
	Captcha                string                  `protobuf:"bytes,2,opt,name=captcha" json:"captcha,omitempty"`
	ListProblemsRequest    *ListProblemsRequest    `protobuf:"bytes,3,opt,name=list_problems_request,json=listProblemsRequest" json:"list_problems_request,omitempty"`
	ListSubmissionsRequest *ListSubmissionsRequest `protobuf:"bytes,4,opt,name=list_submissions_request,json=listSubmissionsRequest" json:"list_submissions_request,omitempty"`
	LoginInitRequest       *LoginInitRequest       `protobuf:"bytes,5,opt,name=login_init_request,json=loginInitRequest" json:"login_init_request,omitempty"`
	LoginAuthRequest       *LoginAuthRequest       `protobuf:"bytes,6,opt,name=login_auth_request,json=loginAuthRequest" json:"login_auth_request,omitempty"`
	RegisterRequest        *RegisterRequest        `protobuf:"bytes,7,opt,name=register_request,json=registerRequest" json:"register_request,omitempty"`
	ShowProblemRequest     *ShowProblemRequest     `protobuf:"bytes,8,opt,name=show_problem_request,json=showProblemRequest" json:"show_problem_request,omitempty"`
	SubmitRequest          *SubmitRequest          `protobuf:"bytes,9,opt,name=submit_request,json=submitRequest" json:"submit_request,omitempty"`
}

func (m *WebRequest) Reset()                    { *m = WebRequest{} }
func (m *WebRequest) String() string            { return proto.CompactTextString(m) }
func (*WebRequest) ProtoMessage()               {}
func (*WebRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *WebRequest) GetCsrfToken() string {
	if m != nil {
		return m.CsrfToken
	}
	return ""
}

func (m *WebRequest) GetCaptcha() string {
	if m != nil {
		return m.Captcha
	}
	return ""
}

func (m *WebRequest) GetListProblemsRequest() *ListProblemsRequest {
	if m != nil {
		return m.ListProblemsRequest
	}
	return nil
}

func (m *WebRequest) GetListSubmissionsRequest() *ListSubmissionsRequest {
	if m != nil {
		return m.ListSubmissionsRequest
	}
	return nil
}

func (m *WebRequest) GetLoginInitRequest() *LoginInitRequest {
	if m != nil {
		return m.LoginInitRequest
	}
	return nil
}

func (m *WebRequest) GetLoginAuthRequest() *LoginAuthRequest {
	if m != nil {
		return m.LoginAuthRequest
	}
	return nil
}

func (m *WebRequest) GetRegisterRequest() *RegisterRequest {
	if m != nil {
		return m.RegisterRequest
	}
	return nil
}

func (m *WebRequest) GetShowProblemRequest() *ShowProblemRequest {
	if m != nil {
		return m.ShowProblemRequest
	}
	return nil
}

func (m *WebRequest) GetSubmitRequest() *SubmitRequest {
	if m != nil {
		return m.SubmitRequest
	}
	return nil
}

type WebResponse struct {
	SetCsrfToken            string                   `protobuf:"bytes,1,opt,name=set_csrf_token,json=setCsrfToken" json:"set_csrf_token,omitempty"`
	CaptchaUrl              string                   `protobuf:"bytes,2,opt,name=captcha_url,json=captchaUrl" json:"captcha_url,omitempty"`
	ListProblemsResponse    *ListProblemsResponse    `protobuf:"bytes,3,opt,name=list_problems_response,json=listProblemsResponse" json:"list_problems_response,omitempty"`
	ListSubmissionsResponse *ListSubmissionsResponse `protobuf:"bytes,4,opt,name=list_submissions_response,json=listSubmissionsResponse" json:"list_submissions_response,omitempty"`
	LoginInitResponse       *LoginInitResponse       `protobuf:"bytes,5,opt,name=login_init_response,json=loginInitResponse" json:"login_init_response,omitempty"`
	LoginAuthResponse       *LoginAuthResponse       `protobuf:"bytes,6,opt,name=login_auth_response,json=loginAuthResponse" json:"login_auth_response,omitempty"`
	RegisterResponse        *RegisterResponse        `protobuf:"bytes,7,opt,name=register_response,json=registerResponse" json:"register_response,omitempty"`
	ShowProblemResponse     *ShowProblemResponse     `protobuf:"bytes,8,opt,name=show_problem_response,json=showProblemResponse" json:"show_problem_response,omitempty"`
	SubmitResponse          *SubmitResponse          `protobuf:"bytes,9,opt,name=submit_response,json=submitResponse" json:"submit_response,omitempty"`
	Error                   *Error                   `protobuf:"bytes,10,opt,name=error" json:"error,omitempty"`
}

func (m *WebResponse) Reset()                    { *m = WebResponse{} }
func (m *WebResponse) String() string            { return proto.CompactTextString(m) }
func (*WebResponse) ProtoMessage()               {}
func (*WebResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *WebResponse) GetSetCsrfToken() string {
	if m != nil {
		return m.SetCsrfToken
	}
	return ""
}

func (m *WebResponse) GetCaptchaUrl() string {
	if m != nil {
		return m.CaptchaUrl
	}
	return ""
}

func (m *WebResponse) GetListProblemsResponse() *ListProblemsResponse {
	if m != nil {
		return m.ListProblemsResponse
	}
	return nil
}

func (m *WebResponse) GetListSubmissionsResponse() *ListSubmissionsResponse {
	if m != nil {
		return m.ListSubmissionsResponse
	}
	return nil
}

func (m *WebResponse) GetLoginInitResponse() *LoginInitResponse {
	if m != nil {
		return m.LoginInitResponse
	}
	return nil
}

func (m *WebResponse) GetLoginAuthResponse() *LoginAuthResponse {
	if m != nil {
		return m.LoginAuthResponse
	}
	return nil
}

func (m *WebResponse) GetRegisterResponse() *RegisterResponse {
	if m != nil {
		return m.RegisterResponse
	}
	return nil
}

func (m *WebResponse) GetShowProblemResponse() *ShowProblemResponse {
	if m != nil {
		return m.ShowProblemResponse
	}
	return nil
}

func (m *WebResponse) GetSubmitResponse() *SubmitResponse {
	if m != nil {
		return m.SubmitResponse
	}
	return nil
}

func (m *WebResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*WebRequest)(nil), "api.WebRequest")
	proto.RegisterType((*WebResponse)(nil), "api.WebResponse")
}

func init() { proto.RegisterFile("api/webapi.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x80, 0x55, 0xb6, 0xb6, 0xeb, 0x2b, 0xb4, 0xdd, 0x6b, 0xd7, 0x66, 0x05, 0x44, 0x35, 0x71,
	0xd8, 0x69, 0x48, 0x70, 0x42, 0x42, 0x42, 0x30, 0x81, 0x34, 0x69, 0x12, 0x28, 0x63, 0x82, 0x5b,
	0x94, 0x56, 0x66, 0xb5, 0x70, 0xe3, 0x60, 0xbb, 0xea, 0x99, 0xdf, 0xc5, 0x9f, 0x9b, 0x62, 0x3f,
	0xbb, 0x49, 0x9a, 0x63, 0xbe, 0x67, 0x7f, 0x69, 0x9c, 0x2f, 0x85, 0x51, 0x9a, 0xf3, 0x37, 0x3b,
	0xb6, 0x4c, 0x73, 0x7e, 0x95, 0x2b, 0x69, 0x24, 0x1e, 0xa5, 0x39, 0x9f, 0x5b, 0xbc, 0x92, 0x9b,
	0x8d, 0xcc, 0x1c, 0x9e, 0xcf, 0x0a, 0x22, 0xb8, 0x36, 0x49, 0xae, 0xe4, 0x52, 0xb0, 0x8d, 0xa6,
	0xc1, 0x3c, 0x0c, 0xf4, 0x76, 0xb9, 0xe1, 0x5a, 0x73, 0x99, 0xf9, 0xd9, 0xd0, 0xce, 0xe4, 0x03,
	0xf7, 0x16, 0x2c, 0x80, 0x62, 0x0f, 0x5c, 0x1b, 0xa6, 0x88, 0x4d, 0x0b, 0xa6, 0xd7, 0x72, 0xe7,
	0xcd, 0xc4, 0xed, 0x6f, 0xb0, 0x4e, 0xe3, 0xc8, 0xc5, 0xff, 0x63, 0x80, 0x9f, 0x6c, 0x19, 0xb3,
	0xbf, 0x5b, 0xa6, 0x0d, 0xbe, 0x04, 0x58, 0x69, 0xf5, 0x3b, 0x31, 0xf2, 0x0f, 0xcb, 0xa2, 0xd6,
	0xa2, 0x75, 0xd9, 0x8b, 0x7b, 0x05, 0xf9, 0x51, 0x00, 0x8c, 0xa0, 0xbb, 0x4a, 0x73, 0xb3, 0x5a,
	0xa7, 0xd1, 0x13, 0x3b, 0xf3, 0x97, 0x78, 0x0b, 0x67, 0x95, 0x27, 0x49, 0x94, 0x33, 0x46, 0x47,
	0x8b, 0xd6, 0x65, 0xff, 0x6d, 0x74, 0x55, 0x9c, 0xc6, 0x2d, 0xd7, 0xe6, 0x3b, 0x2d, 0xa0, 0x3b,
	0xc6, 0x63, 0x71, 0x08, 0xf1, 0x1e, 0xa2, 0xfa, 0xe3, 0x07, 0xe1, 0xb1, 0x15, 0x3e, 0x0f, 0xc2,
	0xbb, 0xfd, 0x1a, 0xef, 0x9c, 0x8a, 0x46, 0x8e, 0xd7, 0x80, 0xf6, 0xe4, 0x12, 0x9e, 0x71, 0x13,
	0x84, 0x6d, 0x2b, 0x3c, 0x73, 0xc2, 0x62, 0x7c, 0x93, 0x71, 0xe3, 0x55, 0x23, 0x51, 0x23, 0x7b,
	0x49, 0xba, 0x35, 0xeb, 0x20, 0xe9, 0xd4, 0x25, 0x9f, 0xb6, 0x66, 0x5d, 0x95, 0x94, 0x08, 0x7e,
	0x84, 0x91, 0x7f, 0x65, 0x41, 0xd1, 0xb5, 0x8a, 0x89, 0x55, 0xc4, 0x34, 0xf4, 0x86, 0xa1, 0xaa,
	0x02, 0xbc, 0x81, 0x49, 0xf9, 0xfd, 0x06, 0xc9, 0x89, 0x95, 0xcc, 0xac, 0xe4, 0x6e, 0x2d, 0x77,
	0x74, 0xb2, 0xde, 0x83, 0xfa, 0x80, 0xe1, 0x7b, 0x18, 0xb8, 0x24, 0x82, 0xa4, 0x67, 0x25, 0xe8,
	0x24, 0x76, 0xe4, 0xf7, 0x3f, 0xd3, 0xe5, 0xcb, 0x8b, 0x7f, 0x6d, 0xe8, 0xdb, 0x7a, 0x74, 0x2e,
	0x33, 0xcd, 0xf0, 0x35, 0x0c, 0x34, 0x33, 0xc9, 0x41, 0x42, 0x4f, 0x35, 0x33, 0xd7, 0xa1, 0xa2,
	0x57, 0xd0, 0xa7, 0x6c, 0x92, 0xad, 0x12, 0x54, 0x12, 0x10, 0xba, 0x57, 0x02, 0xbf, 0xc1, 0xb4,
	0x1e, 0x93, 0xbb, 0x01, 0xd5, 0x74, 0xde, 0x50, 0x93, 0x5b, 0x10, 0x4f, 0x44, 0x03, 0xc5, 0x5f,
	0x70, 0xde, 0xd0, 0x13, 0x39, 0x5d, 0x50, 0x2f, 0x9a, 0x83, 0x22, 0xed, 0x4c, 0x34, 0x0f, 0xf0,
	0x2b, 0x8c, 0x2b, 0x49, 0x91, 0xd3, 0x35, 0x35, 0xad, 0x37, 0x45, 0xb6, 0x53, 0x51, 0x47, 0x7b,
	0x0f, 0x55, 0x45, 0x9e, 0x4e, 0xdd, 0xe3, 0x22, 0xaa, 0x78, 0xca, 0x08, 0x3f, 0xc3, 0x69, 0x29,
	0x2c, 0xb2, 0x74, 0x4b, 0x71, 0xee, 0xcb, 0x22, 0xc9, 0x48, 0xd5, 0x48, 0xf1, 0x2d, 0xd7, 0xda,
	0x22, 0xcf, 0x49, 0xe9, 0x5b, 0xae, 0xc4, 0x45, 0xaa, 0xb1, 0x3e, 0x84, 0xf8, 0x01, 0x86, 0x21,
	0x2f, 0xf2, 0xb8, 0xbe, 0xc6, 0x95, 0xbe, 0x48, 0x31, 0xd0, 0x95, 0x6b, 0x5c, 0x40, 0x9b, 0x29,
	0x25, 0x55, 0x04, 0x76, 0x0f, 0xd8, 0x3d, 0x5f, 0x0a, 0x12, 0xbb, 0xc1, 0xb2, 0x63, 0xff, 0xc8,
	0xde, 0x3d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x63, 0xd5, 0x0e, 0x93, 0x77, 0x05, 0x00, 0x00,
}