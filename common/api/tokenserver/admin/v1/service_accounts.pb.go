// Code generated by protoc-gen-go.
// source: service_accounts.proto
// DO NOT EDIT!

package admin

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tokenserver "github.com/luci/luci-go/common/api/tokenserver"
import tokenserver1 "github.com/luci/luci-go/common/api/tokenserver"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// CreateServiceAccountRequest is parameters for CreateServiceAccount call.
type CreateServiceAccountRequest struct {
	Ca    string `protobuf:"bytes,1,opt,name=ca" json:"ca,omitempty"`
	Fqdn  string `protobuf:"bytes,2,opt,name=fqdn" json:"fqdn,omitempty"`
	Force bool   `protobuf:"varint,3,opt,name=force" json:"force,omitempty"`
}

func (m *CreateServiceAccountRequest) Reset()                    { *m = CreateServiceAccountRequest{} }
func (m *CreateServiceAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateServiceAccountRequest) ProtoMessage()               {}
func (*CreateServiceAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// CreateServiceAccountResponse is returned by CreateServiceAccount call.
type CreateServiceAccountResponse struct {
	ServiceAccount *tokenserver.ServiceAccount `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount" json:"service_account,omitempty"`
}

func (m *CreateServiceAccountResponse) Reset()                    { *m = CreateServiceAccountResponse{} }
func (m *CreateServiceAccountResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateServiceAccountResponse) ProtoMessage()               {}
func (*CreateServiceAccountResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *CreateServiceAccountResponse) GetServiceAccount() *tokenserver.ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

// MintAccessTokenRequest is parameters for MintAccessToken call.
type MintAccessTokenRequest struct {
	Ca     string   `protobuf:"bytes,1,opt,name=ca" json:"ca,omitempty"`
	Fqdn   string   `protobuf:"bytes,2,opt,name=fqdn" json:"fqdn,omitempty"`
	Scopes []string `protobuf:"bytes,3,rep,name=scopes" json:"scopes,omitempty"`
}

func (m *MintAccessTokenRequest) Reset()                    { *m = MintAccessTokenRequest{} }
func (m *MintAccessTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*MintAccessTokenRequest) ProtoMessage()               {}
func (*MintAccessTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

// MintAccessTokenResponse is returned by MintAccessToken call.
type MintAccessTokenResponse struct {
	ServiceAccount    *tokenserver.ServiceAccount     `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount" json:"service_account,omitempty"`
	Oauth2AccessToken *tokenserver1.OAuth2AccessToken `protobuf:"bytes,2,opt,name=oauth2_access_token,json=oauth2AccessToken" json:"oauth2_access_token,omitempty"`
}

func (m *MintAccessTokenResponse) Reset()                    { *m = MintAccessTokenResponse{} }
func (m *MintAccessTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*MintAccessTokenResponse) ProtoMessage()               {}
func (*MintAccessTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *MintAccessTokenResponse) GetServiceAccount() *tokenserver.ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

func (m *MintAccessTokenResponse) GetOauth2AccessToken() *tokenserver1.OAuth2AccessToken {
	if m != nil {
		return m.Oauth2AccessToken
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateServiceAccountRequest)(nil), "tokenserver.admin.CreateServiceAccountRequest")
	proto.RegisterType((*CreateServiceAccountResponse)(nil), "tokenserver.admin.CreateServiceAccountResponse")
	proto.RegisterType((*MintAccessTokenRequest)(nil), "tokenserver.admin.MintAccessTokenRequest")
	proto.RegisterType((*MintAccessTokenResponse)(nil), "tokenserver.admin.MintAccessTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for ServiceAccounts service

type ServiceAccountsClient interface {
	// CreateServiceAccount creates Google Cloud IAM service account associated
	// with given host.
	//
	// It uses token server configuration to pick a cloud project and to derive
	// service account ID. See documentation for CertificateAuthorityConfig proto
	// message for more info.
	//
	// This operation is idempotent.
	CreateServiceAccount(ctx context.Context, in *CreateServiceAccountRequest, opts ...grpc.CallOption) (*CreateServiceAccountResponse, error)
	// MintAccessToken generates a new access token for a service account
	// associated with the given host.
	//
	// It will register the service account first, if necessary.
	MintAccessToken(ctx context.Context, in *MintAccessTokenRequest, opts ...grpc.CallOption) (*MintAccessTokenResponse, error)
}
type serviceAccountsPRPCClient struct {
	client *prpccommon.Client
}

func NewServiceAccountsPRPCClient(client *prpccommon.Client) ServiceAccountsClient {
	return &serviceAccountsPRPCClient{client}
}

func (c *serviceAccountsPRPCClient) CreateServiceAccount(ctx context.Context, in *CreateServiceAccountRequest, opts ...grpc.CallOption) (*CreateServiceAccountResponse, error) {
	out := new(CreateServiceAccountResponse)
	err := c.client.Call(ctx, "tokenserver.admin.ServiceAccounts", "CreateServiceAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAccountsPRPCClient) MintAccessToken(ctx context.Context, in *MintAccessTokenRequest, opts ...grpc.CallOption) (*MintAccessTokenResponse, error) {
	out := new(MintAccessTokenResponse)
	err := c.client.Call(ctx, "tokenserver.admin.ServiceAccounts", "MintAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type serviceAccountsClient struct {
	cc *grpc.ClientConn
}

func NewServiceAccountsClient(cc *grpc.ClientConn) ServiceAccountsClient {
	return &serviceAccountsClient{cc}
}

func (c *serviceAccountsClient) CreateServiceAccount(ctx context.Context, in *CreateServiceAccountRequest, opts ...grpc.CallOption) (*CreateServiceAccountResponse, error) {
	out := new(CreateServiceAccountResponse)
	err := grpc.Invoke(ctx, "/tokenserver.admin.ServiceAccounts/CreateServiceAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAccountsClient) MintAccessToken(ctx context.Context, in *MintAccessTokenRequest, opts ...grpc.CallOption) (*MintAccessTokenResponse, error) {
	out := new(MintAccessTokenResponse)
	err := grpc.Invoke(ctx, "/tokenserver.admin.ServiceAccounts/MintAccessToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServiceAccounts service

type ServiceAccountsServer interface {
	// CreateServiceAccount creates Google Cloud IAM service account associated
	// with given host.
	//
	// It uses token server configuration to pick a cloud project and to derive
	// service account ID. See documentation for CertificateAuthorityConfig proto
	// message for more info.
	//
	// This operation is idempotent.
	CreateServiceAccount(context.Context, *CreateServiceAccountRequest) (*CreateServiceAccountResponse, error)
	// MintAccessToken generates a new access token for a service account
	// associated with the given host.
	//
	// It will register the service account first, if necessary.
	MintAccessToken(context.Context, *MintAccessTokenRequest) (*MintAccessTokenResponse, error)
}

func RegisterServiceAccountsServer(s prpc.Registrar, srv ServiceAccountsServer) {
	s.RegisterService(&_ServiceAccounts_serviceDesc, srv)
}

func _ServiceAccounts_CreateServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateServiceAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ServiceAccountsServer).CreateServiceAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ServiceAccounts_MintAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(MintAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ServiceAccountsServer).MintAccessToken(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ServiceAccounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenserver.admin.ServiceAccounts",
	HandlerType: (*ServiceAccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServiceAccount",
			Handler:    _ServiceAccounts_CreateServiceAccount_Handler,
		},
		{
			MethodName: "MintAccessToken",
			Handler:    _ServiceAccounts_MintAccessToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor2 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0x51, 0x4f, 0xc2, 0x30,
	0x10, 0xc7, 0x33, 0x50, 0xa2, 0x47, 0x02, 0xa1, 0x12, 0x24, 0xc3, 0x18, 0xb3, 0x27, 0x35, 0xb1,
	0x4b, 0xe6, 0xa3, 0x4f, 0x44, 0x5e, 0xd5, 0x04, 0x49, 0x7c, 0x24, 0xa5, 0x14, 0x68, 0x94, 0x76,
	0xac, 0x9d, 0x7e, 0x2c, 0xbf, 0x9c, 0x1f, 0xc0, 0xae, 0xe5, 0x61, 0x1b, 0x8b, 0x81, 0xc4, 0x17,
	0x72, 0x77, 0xdc, 0xff, 0x77, 0xff, 0x5d, 0x0f, 0x7a, 0x8a, 0x25, 0x9f, 0x9c, 0xb2, 0x29, 0xa1,
	0x54, 0xa6, 0x42, 0x2b, 0x1c, 0x27, 0x52, 0x4b, 0xd4, 0xd1, 0xf2, 0x9d, 0x89, 0xec, 0x4f, 0x96,
	0x60, 0x32, 0x5f, 0x73, 0xe1, 0x8f, 0x96, 0x5c, 0xaf, 0xd2, 0x19, 0xa6, 0x72, 0x1d, 0x7e, 0xa4,
	0x94, 0xdb, 0x9f, 0xbb, 0xa5, 0x0c, 0x4d, 0x61, 0x2d, 0x45, 0x48, 0x62, 0x1e, 0xe6, 0x54, 0x61,
	0x89, 0xec, 0xc0, 0xfe, 0xc3, 0x81, 0x14, 0x17, 0x3b, 0x71, 0xf0, 0x06, 0x83, 0xc7, 0x84, 0x11,
	0xcd, 0x5e, 0x1d, 0x7b, 0xe8, 0xd0, 0x63, 0xb6, 0x49, 0x99, 0xd2, 0xa8, 0x05, 0x35, 0x4a, 0xfa,
	0xde, 0x95, 0x77, 0x7d, 0x3a, 0x36, 0x11, 0x42, 0x70, 0xb4, 0xd8, 0xcc, 0x45, 0xbf, 0x66, 0x2b,
	0x36, 0x46, 0x5d, 0x38, 0x5e, 0xc8, 0x84, 0xb2, 0x7e, 0xdd, 0x14, 0x4f, 0xc6, 0x2e, 0x09, 0xe6,
	0x70, 0x51, 0x0d, 0x56, 0xb1, 0x34, 0x4e, 0xd0, 0x08, 0xda, 0xa5, 0xcf, 0xb1, 0x63, 0x9a, 0xd1,
	0x00, 0xe7, 0x17, 0x55, 0x52, 0xb7, 0x54, 0x21, 0x0f, 0x26, 0xd0, 0x7b, 0xe2, 0x42, 0x9b, 0x94,
	0x29, 0x35, 0xc9, 0x74, 0x87, 0x38, 0xef, 0x41, 0x43, 0x51, 0x19, 0x33, 0x65, 0xac, 0xd7, 0x4d,
	0x75, 0x9b, 0x05, 0xdf, 0x1e, 0x9c, 0xef, 0x60, 0xff, 0xd3, 0x37, 0x7a, 0x86, 0x33, 0x49, 0x52,
	0xbd, 0x8a, 0x32, 0x88, 0x99, 0x31, 0xb5, 0x5a, 0x6b, 0xae, 0x19, 0x5d, 0x16, 0x48, 0x2f, 0xc3,
	0xac, 0x2f, 0x6f, 0xa5, 0xe3, 0xa4, 0xb9, 0x52, 0xf4, 0xe3, 0x41, 0xbb, 0x38, 0x52, 0xa1, 0x2f,
	0xe8, 0x56, 0xbd, 0x00, 0xc2, 0x78, 0xe7, 0x12, 0xf1, 0x1f, 0x37, 0xe0, 0x87, 0x7b, 0xf7, 0x6f,
	0x57, 0xb4, 0x82, 0x76, 0x69, 0x7b, 0xe8, 0xa6, 0x82, 0x51, 0xfd, 0x70, 0xfe, 0xed, 0x3e, 0xad,
	0x6e, 0xd2, 0xac, 0x61, 0x8f, 0xf8, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x35, 0xb4, 0x86,
	0x74, 0x03, 0x00, 0x00,
}