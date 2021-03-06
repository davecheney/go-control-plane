// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/grpc_service.proto

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// gRPC service configuration. This is used by :ref:`ApiConfigSource
// <envoy_api_msg_ApiConfigSource>` and filter configurations.
type GrpcService struct {
	// Types that are valid to be assigned to TargetSpecifier:
	//	*GrpcService_EnvoyGrpc_
	//	*GrpcService_GoogleGrpc_
	TargetSpecifier isGrpcService_TargetSpecifier `protobuf_oneof:"target_specifier"`
	// The timeout for the gRPC request. This is the timeout for a specific
	// request.
	Timeout *google_protobuf2.Duration `protobuf:"bytes,3,opt,name=timeout" json:"timeout,omitempty"`
	// A set of credentials that will be composed to form the `channel credentials
	// <https://grpc.io/docs/guides/auth.html#credential-types>`_.
	Credentials []*GrpcService_Credentials `protobuf:"bytes,4,rep,name=credentials" json:"credentials,omitempty"`
}

func (m *GrpcService) Reset()                    { *m = GrpcService{} }
func (m *GrpcService) String() string            { return proto.CompactTextString(m) }
func (*GrpcService) ProtoMessage()               {}
func (*GrpcService) Descriptor() ([]byte, []int) { return fileDescriptorGrpcService, []int{0} }

type isGrpcService_TargetSpecifier interface {
	isGrpcService_TargetSpecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type GrpcService_EnvoyGrpc_ struct {
	EnvoyGrpc *GrpcService_EnvoyGrpc `protobuf:"bytes,1,opt,name=envoy_grpc,json=envoyGrpc,oneof"`
}
type GrpcService_GoogleGrpc_ struct {
	GoogleGrpc *GrpcService_GoogleGrpc `protobuf:"bytes,2,opt,name=google_grpc,json=googleGrpc,oneof"`
}

func (*GrpcService_EnvoyGrpc_) isGrpcService_TargetSpecifier()  {}
func (*GrpcService_GoogleGrpc_) isGrpcService_TargetSpecifier() {}

func (m *GrpcService) GetTargetSpecifier() isGrpcService_TargetSpecifier {
	if m != nil {
		return m.TargetSpecifier
	}
	return nil
}

func (m *GrpcService) GetEnvoyGrpc() *GrpcService_EnvoyGrpc {
	if x, ok := m.GetTargetSpecifier().(*GrpcService_EnvoyGrpc_); ok {
		return x.EnvoyGrpc
	}
	return nil
}

func (m *GrpcService) GetGoogleGrpc() *GrpcService_GoogleGrpc {
	if x, ok := m.GetTargetSpecifier().(*GrpcService_GoogleGrpc_); ok {
		return x.GoogleGrpc
	}
	return nil
}

func (m *GrpcService) GetTimeout() *google_protobuf2.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *GrpcService) GetCredentials() []*GrpcService_Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GrpcService) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GrpcService_OneofMarshaler, _GrpcService_OneofUnmarshaler, _GrpcService_OneofSizer, []interface{}{
		(*GrpcService_EnvoyGrpc_)(nil),
		(*GrpcService_GoogleGrpc_)(nil),
	}
}

func _GrpcService_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GrpcService)
	// target_specifier
	switch x := m.TargetSpecifier.(type) {
	case *GrpcService_EnvoyGrpc_:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EnvoyGrpc); err != nil {
			return err
		}
	case *GrpcService_GoogleGrpc_:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GoogleGrpc); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GrpcService.TargetSpecifier has unexpected type %T", x)
	}
	return nil
}

func _GrpcService_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GrpcService)
	switch tag {
	case 1: // target_specifier.envoy_grpc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GrpcService_EnvoyGrpc)
		err := b.DecodeMessage(msg)
		m.TargetSpecifier = &GrpcService_EnvoyGrpc_{msg}
		return true, err
	case 2: // target_specifier.google_grpc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GrpcService_GoogleGrpc)
		err := b.DecodeMessage(msg)
		m.TargetSpecifier = &GrpcService_GoogleGrpc_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GrpcService_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GrpcService)
	// target_specifier
	switch x := m.TargetSpecifier.(type) {
	case *GrpcService_EnvoyGrpc_:
		s := proto.Size(x.EnvoyGrpc)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GrpcService_GoogleGrpc_:
		s := proto.Size(x.GoogleGrpc)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GrpcService_EnvoyGrpc struct {
	// The name of the upstream gRPC cluster. SSL credentials will be supplied
	// in the :ref:`Cluster <envoy_api_msg_Cluster>` :ref:`tls_context
	// <envoy_api_field_Cluster.tls_context>`.
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (m *GrpcService_EnvoyGrpc) Reset()         { *m = GrpcService_EnvoyGrpc{} }
func (m *GrpcService_EnvoyGrpc) String() string { return proto.CompactTextString(m) }
func (*GrpcService_EnvoyGrpc) ProtoMessage()    {}
func (*GrpcService_EnvoyGrpc) Descriptor() ([]byte, []int) {
	return fileDescriptorGrpcService, []int{0, 0}
}

func (m *GrpcService_EnvoyGrpc) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type GrpcService_GoogleGrpc struct {
	// The target URI when using the `Google C++ gRPC client
	// <https://github.com/grpc/grpc>`_. SSL credentials will be supplied in
	// :ref:`credentials <envoy_api_field_GrpcService.credentials>`.
	TargetUri      string                                 `protobuf:"bytes,1,opt,name=target_uri,json=targetUri,proto3" json:"target_uri,omitempty"`
	SslCredentials *GrpcService_GoogleGrpc_SslCredentials `protobuf:"bytes,2,opt,name=ssl_credentials,json=sslCredentials" json:"ssl_credentials,omitempty"`
	// The human readable prefix to use when emitting statistics for the gRPC
	// service.
	//
	// .. csv-table::
	//    :header: Name, Type, Description
	//    :widths: 1, 1, 2
	//
	//    streams_total, Counter, Total number of streams opened
	//    streams_closed_<gRPC status code>, Counter, Total streams closed with <gRPC status code>
	StatPrefix string `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
}

func (m *GrpcService_GoogleGrpc) Reset()         { *m = GrpcService_GoogleGrpc{} }
func (m *GrpcService_GoogleGrpc) String() string { return proto.CompactTextString(m) }
func (*GrpcService_GoogleGrpc) ProtoMessage()    {}
func (*GrpcService_GoogleGrpc) Descriptor() ([]byte, []int) {
	return fileDescriptorGrpcService, []int{0, 1}
}

func (m *GrpcService_GoogleGrpc) GetTargetUri() string {
	if m != nil {
		return m.TargetUri
	}
	return ""
}

func (m *GrpcService_GoogleGrpc) GetSslCredentials() *GrpcService_GoogleGrpc_SslCredentials {
	if m != nil {
		return m.SslCredentials
	}
	return nil
}

func (m *GrpcService_GoogleGrpc) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

// See https://grpc.io/grpc/cpp/structgrpc_1_1_ssl_credentials_options.html.
type GrpcService_GoogleGrpc_SslCredentials struct {
	// PEM encoded server root certificates.
	RootCerts *DataSource `protobuf:"bytes,1,opt,name=root_certs,json=rootCerts" json:"root_certs,omitempty"`
	// PEM encoded client private key.
	PrivateKey *DataSource `protobuf:"bytes,2,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	// PEM encoded client certificate chain.
	CertChain *DataSource `protobuf:"bytes,3,opt,name=cert_chain,json=certChain" json:"cert_chain,omitempty"`
}

func (m *GrpcService_GoogleGrpc_SslCredentials) Reset()         { *m = GrpcService_GoogleGrpc_SslCredentials{} }
func (m *GrpcService_GoogleGrpc_SslCredentials) String() string { return proto.CompactTextString(m) }
func (*GrpcService_GoogleGrpc_SslCredentials) ProtoMessage()    {}
func (*GrpcService_GoogleGrpc_SslCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptorGrpcService, []int{0, 1, 0}
}

func (m *GrpcService_GoogleGrpc_SslCredentials) GetRootCerts() *DataSource {
	if m != nil {
		return m.RootCerts
	}
	return nil
}

func (m *GrpcService_GoogleGrpc_SslCredentials) GetPrivateKey() *DataSource {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *GrpcService_GoogleGrpc_SslCredentials) GetCertChain() *DataSource {
	if m != nil {
		return m.CertChain
	}
	return nil
}

// gRPC credentials as described at
// https://grpc.io/docs/guides/auth.html#credential-types.
//
// .. note::
//
//   Credentials are only currently implemented for the Google gRPC client.
type GrpcService_Credentials struct {
	// Types that are valid to be assigned to CredentialSpecifier:
	//	*GrpcService_Credentials_AccessToken
	CredentialSpecifier isGrpcService_Credentials_CredentialSpecifier `protobuf_oneof:"credential_specifier"`
}

func (m *GrpcService_Credentials) Reset()         { *m = GrpcService_Credentials{} }
func (m *GrpcService_Credentials) String() string { return proto.CompactTextString(m) }
func (*GrpcService_Credentials) ProtoMessage()    {}
func (*GrpcService_Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptorGrpcService, []int{0, 2}
}

type isGrpcService_Credentials_CredentialSpecifier interface {
	isGrpcService_Credentials_CredentialSpecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type GrpcService_Credentials_AccessToken struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3,oneof"`
}

func (*GrpcService_Credentials_AccessToken) isGrpcService_Credentials_CredentialSpecifier() {}

func (m *GrpcService_Credentials) GetCredentialSpecifier() isGrpcService_Credentials_CredentialSpecifier {
	if m != nil {
		return m.CredentialSpecifier
	}
	return nil
}

func (m *GrpcService_Credentials) GetAccessToken() string {
	if x, ok := m.GetCredentialSpecifier().(*GrpcService_Credentials_AccessToken); ok {
		return x.AccessToken
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GrpcService_Credentials) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GrpcService_Credentials_OneofMarshaler, _GrpcService_Credentials_OneofUnmarshaler, _GrpcService_Credentials_OneofSizer, []interface{}{
		(*GrpcService_Credentials_AccessToken)(nil),
	}
}

func _GrpcService_Credentials_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GrpcService_Credentials)
	// credential_specifier
	switch x := m.CredentialSpecifier.(type) {
	case *GrpcService_Credentials_AccessToken:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.AccessToken)
	case nil:
	default:
		return fmt.Errorf("GrpcService_Credentials.CredentialSpecifier has unexpected type %T", x)
	}
	return nil
}

func _GrpcService_Credentials_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GrpcService_Credentials)
	switch tag {
	case 1: // credential_specifier.access_token
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.CredentialSpecifier = &GrpcService_Credentials_AccessToken{x}
		return true, err
	default:
		return false, nil
	}
}

func _GrpcService_Credentials_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GrpcService_Credentials)
	// credential_specifier
	switch x := m.CredentialSpecifier.(type) {
	case *GrpcService_Credentials_AccessToken:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AccessToken)))
		n += len(x.AccessToken)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*GrpcService)(nil), "envoy.api.v2.GrpcService")
	proto.RegisterType((*GrpcService_EnvoyGrpc)(nil), "envoy.api.v2.GrpcService.EnvoyGrpc")
	proto.RegisterType((*GrpcService_GoogleGrpc)(nil), "envoy.api.v2.GrpcService.GoogleGrpc")
	proto.RegisterType((*GrpcService_GoogleGrpc_SslCredentials)(nil), "envoy.api.v2.GrpcService.GoogleGrpc.SslCredentials")
	proto.RegisterType((*GrpcService_Credentials)(nil), "envoy.api.v2.GrpcService.Credentials")
}
func (m *GrpcService) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcService) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TargetSpecifier != nil {
		nn1, err := m.TargetSpecifier.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.Timeout != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGrpcService(dAtA, i, uint64(m.Timeout.Size()))
		n2, err := m.Timeout.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Credentials) > 0 {
		for _, msg := range m.Credentials {
			dAtA[i] = 0x22
			i++
			i = encodeVarintGrpcService(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *GrpcService_EnvoyGrpc_) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.EnvoyGrpc != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGrpcService(dAtA, i, uint64(m.EnvoyGrpc.Size()))
		n3, err := m.EnvoyGrpc.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *GrpcService_GoogleGrpc_) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.GoogleGrpc != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGrpcService(dAtA, i, uint64(m.GoogleGrpc.Size()))
		n4, err := m.GoogleGrpc.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *GrpcService_EnvoyGrpc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcService_EnvoyGrpc) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ClusterName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGrpcService(dAtA, i, uint64(len(m.ClusterName)))
		i += copy(dAtA[i:], m.ClusterName)
	}
	return i, nil
}

func (m *GrpcService_GoogleGrpc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcService_GoogleGrpc) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TargetUri) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGrpcService(dAtA, i, uint64(len(m.TargetUri)))
		i += copy(dAtA[i:], m.TargetUri)
	}
	if m.SslCredentials != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGrpcService(dAtA, i, uint64(m.SslCredentials.Size()))
		n5, err := m.SslCredentials.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.StatPrefix) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGrpcService(dAtA, i, uint64(len(m.StatPrefix)))
		i += copy(dAtA[i:], m.StatPrefix)
	}
	return i, nil
}

func (m *GrpcService_GoogleGrpc_SslCredentials) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcService_GoogleGrpc_SslCredentials) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RootCerts != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGrpcService(dAtA, i, uint64(m.RootCerts.Size()))
		n6, err := m.RootCerts.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.PrivateKey != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGrpcService(dAtA, i, uint64(m.PrivateKey.Size()))
		n7, err := m.PrivateKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.CertChain != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGrpcService(dAtA, i, uint64(m.CertChain.Size()))
		n8, err := m.CertChain.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}

func (m *GrpcService_Credentials) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcService_Credentials) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CredentialSpecifier != nil {
		nn9, err := m.CredentialSpecifier.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn9
	}
	return i, nil
}

func (m *GrpcService_Credentials_AccessToken) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintGrpcService(dAtA, i, uint64(len(m.AccessToken)))
	i += copy(dAtA[i:], m.AccessToken)
	return i, nil
}
func encodeVarintGrpcService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GrpcService) Size() (n int) {
	var l int
	_ = l
	if m.TargetSpecifier != nil {
		n += m.TargetSpecifier.Size()
	}
	if m.Timeout != nil {
		l = m.Timeout.Size()
		n += 1 + l + sovGrpcService(uint64(l))
	}
	if len(m.Credentials) > 0 {
		for _, e := range m.Credentials {
			l = e.Size()
			n += 1 + l + sovGrpcService(uint64(l))
		}
	}
	return n
}

func (m *GrpcService_EnvoyGrpc_) Size() (n int) {
	var l int
	_ = l
	if m.EnvoyGrpc != nil {
		l = m.EnvoyGrpc.Size()
		n += 1 + l + sovGrpcService(uint64(l))
	}
	return n
}
func (m *GrpcService_GoogleGrpc_) Size() (n int) {
	var l int
	_ = l
	if m.GoogleGrpc != nil {
		l = m.GoogleGrpc.Size()
		n += 1 + l + sovGrpcService(uint64(l))
	}
	return n
}
func (m *GrpcService_EnvoyGrpc) Size() (n int) {
	var l int
	_ = l
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sovGrpcService(uint64(l))
	}
	return n
}

func (m *GrpcService_GoogleGrpc) Size() (n int) {
	var l int
	_ = l
	l = len(m.TargetUri)
	if l > 0 {
		n += 1 + l + sovGrpcService(uint64(l))
	}
	if m.SslCredentials != nil {
		l = m.SslCredentials.Size()
		n += 1 + l + sovGrpcService(uint64(l))
	}
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovGrpcService(uint64(l))
	}
	return n
}

func (m *GrpcService_GoogleGrpc_SslCredentials) Size() (n int) {
	var l int
	_ = l
	if m.RootCerts != nil {
		l = m.RootCerts.Size()
		n += 1 + l + sovGrpcService(uint64(l))
	}
	if m.PrivateKey != nil {
		l = m.PrivateKey.Size()
		n += 1 + l + sovGrpcService(uint64(l))
	}
	if m.CertChain != nil {
		l = m.CertChain.Size()
		n += 1 + l + sovGrpcService(uint64(l))
	}
	return n
}

func (m *GrpcService_Credentials) Size() (n int) {
	var l int
	_ = l
	if m.CredentialSpecifier != nil {
		n += m.CredentialSpecifier.Size()
	}
	return n
}

func (m *GrpcService_Credentials_AccessToken) Size() (n int) {
	var l int
	_ = l
	l = len(m.AccessToken)
	n += 1 + l + sovGrpcService(uint64(l))
	return n
}

func sovGrpcService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGrpcService(x uint64) (n int) {
	return sovGrpcService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GrpcService) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GrpcService: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrpcService: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnvoyGrpc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &GrpcService_EnvoyGrpc{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.TargetSpecifier = &GrpcService_EnvoyGrpc_{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleGrpc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &GrpcService_GoogleGrpc{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.TargetSpecifier = &GrpcService_GoogleGrpc_{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = &google_protobuf2.Duration{}
			}
			if err := m.Timeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credentials = append(m.Credentials, &GrpcService_Credentials{})
			if err := m.Credentials[len(m.Credentials)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GrpcService_EnvoyGrpc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EnvoyGrpc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnvoyGrpc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GrpcService_GoogleGrpc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GoogleGrpc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoogleGrpc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SslCredentials", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SslCredentials == nil {
				m.SslCredentials = &GrpcService_GoogleGrpc_SslCredentials{}
			}
			if err := m.SslCredentials.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GrpcService_GoogleGrpc_SslCredentials) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SslCredentials: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SslCredentials: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootCerts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RootCerts == nil {
				m.RootCerts = &DataSource{}
			}
			if err := m.RootCerts.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivateKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrivateKey == nil {
				m.PrivateKey = &DataSource{}
			}
			if err := m.PrivateKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CertChain == nil {
				m.CertChain = &DataSource{}
			}
			if err := m.CertChain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GrpcService_Credentials) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Credentials: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credentials: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpcService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CredentialSpecifier = &GrpcService_Credentials_AccessToken{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGrpcService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrpcService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGrpcService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGrpcService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGrpcService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGrpcService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGrpcService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrpcService   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/grpc_service.proto", fileDescriptorGrpcService) }

var fileDescriptorGrpcService = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x51, 0x8b, 0xd3, 0x4c,
	0x14, 0xdd, 0x6c, 0xf7, 0xdb, 0x8f, 0xdc, 0x94, 0x2a, 0x83, 0x68, 0x37, 0x60, 0x29, 0xae, 0x42,
	0x11, 0x49, 0xa1, 0x7d, 0x90, 0x7d, 0x6d, 0x2b, 0x5d, 0x10, 0x44, 0x53, 0x7d, 0x13, 0x86, 0xe9,
	0xf4, 0xb6, 0x0e, 0x9b, 0x66, 0xc2, 0xcc, 0xa4, 0xd8, 0x37, 0x7f, 0x96, 0x08, 0xc2, 0x3e, 0x89,
	0x8f, 0xfe, 0x04, 0xe9, 0xdb, 0xfe, 0x0b, 0x99, 0x64, 0xba, 0x4d, 0x59, 0x2a, 0xbe, 0xe5, 0xce,
	0x39, 0xf7, 0x9c, 0x7b, 0x6e, 0x2e, 0x3c, 0x64, 0x99, 0xe8, 0x2e, 0x54, 0xc6, 0xa9, 0x46, 0xb5,
	0x12, 0x1c, 0xa3, 0x4c, 0x49, 0x23, 0x49, 0x1d, 0xd3, 0x95, 0x5c, 0x47, 0x2c, 0x13, 0xd1, 0xaa,
	0x17, 0x36, 0x2c, 0x6b, 0xca, 0xb4, 0x43, 0xc3, 0xd6, 0x42, 0xca, 0x45, 0x82, 0xdd, 0xa2, 0x9a,
	0xe6, 0xf3, 0xee, 0x2c, 0x57, 0xcc, 0x08, 0x99, 0x3a, 0xfc, 0xd1, 0x8a, 0x25, 0x62, 0xc6, 0x0c,
	0x76, 0xb7, 0x1f, 0x25, 0xf0, 0xe4, 0xc7, 0x29, 0x04, 0x63, 0x95, 0xf1, 0x49, 0x69, 0x46, 0x46,
	0x00, 0x85, 0x11, 0xb5, 0x23, 0x34, 0xbd, 0xb6, 0xd7, 0x09, 0x7a, 0xe7, 0x51, 0xd5, 0x3b, 0xaa,
	0xd0, 0xa3, 0x57, 0x16, 0xb0, 0x0f, 0x97, 0x47, 0xb1, 0x8f, 0xdb, 0x82, 0x8c, 0x21, 0x28, 0x07,
	0x2a, 0x65, 0x8e, 0x0b, 0x99, 0xa7, 0x87, 0x65, 0xc6, 0x05, 0xd9, 0xe9, 0xc0, 0xe2, 0xb6, 0x22,
	0x7d, 0xf8, 0xdf, 0x88, 0x25, 0xca, 0xdc, 0x34, 0x6b, 0x85, 0xc8, 0x59, 0x54, 0xa2, 0xd1, 0x36,
	0x69, 0x34, 0x72, 0x49, 0xe3, 0x2d, 0xd3, 0xba, 0x73, 0x85, 0x33, 0x4c, 0x8d, 0x60, 0x89, 0x6e,
	0x9e, 0xb4, 0x6b, 0x9d, 0xa0, 0xf7, 0xec, 0xb0, 0xfb, 0x70, 0x47, 0x8e, 0xab, 0x9d, 0xe1, 0x05,
	0xf8, 0xb7, 0x01, 0xc9, 0x0b, 0xa8, 0xf3, 0x24, 0xd7, 0x06, 0x15, 0x4d, 0xd9, 0x12, 0x8b, 0xdd,
	0xf8, 0x03, 0xff, 0xdb, 0xcd, 0x75, 0xed, 0x44, 0x1d, 0xb7, 0xbd, 0x38, 0x70, 0xf0, 0x1b, 0xb6,
	0xc4, 0xf0, 0x4b, 0x0d, 0x60, 0x97, 0x8a, 0x74, 0x00, 0x0c, 0x53, 0x0b, 0x34, 0x34, 0x57, 0xe2,
	0x6e, 0xab, 0x5f, 0x82, 0x1f, 0x94, 0x20, 0x1f, 0xe1, 0x9e, 0xd6, 0x09, 0xad, 0x06, 0x28, 0xd7,
	0xd7, 0xff, 0x97, 0xf5, 0x45, 0x13, 0x9d, 0x54, 0xe3, 0x34, 0xf4, 0x5e, 0x4d, 0x9e, 0x43, 0xa0,
	0x0d, 0x33, 0x34, 0x53, 0x38, 0x17, 0x9f, 0x8b, 0x9d, 0xee, 0x0d, 0x02, 0x16, 0x7d, 0x5b, 0x80,
	0xe1, 0x77, 0x0f, 0x1a, 0xfb, 0x72, 0xe4, 0x25, 0x80, 0x92, 0xd2, 0x50, 0x8e, 0xca, 0x68, 0x77,
	0x1d, 0xcd, 0xfd, 0xb9, 0x46, 0xcc, 0xb0, 0x89, 0xcc, 0x15, 0xc7, 0xd8, 0xb7, 0xdc, 0xa1, 0xa5,
	0x92, 0x0b, 0x08, 0x32, 0x25, 0x56, 0xcc, 0x20, 0xbd, 0xc2, 0xb5, 0x4b, 0x74, 0xb8, 0x13, 0x1c,
	0xf9, 0x35, 0xae, 0xad, 0xa7, 0xb5, 0xa3, 0xfc, 0x13, 0x13, 0xa9, 0xbb, 0x82, 0xbf, 0x78, 0x5a,
	0xee, 0xd0, 0x52, 0xc3, 0x77, 0x10, 0x54, 0x67, 0x3f, 0x87, 0x3a, 0xe3, 0x1c, 0xb5, 0xa6, 0x46,
	0x5e, 0x61, 0x5a, 0xfe, 0x84, 0xcb, 0xa3, 0x38, 0x28, 0x5f, 0xdf, 0xdb, 0xc7, 0xc1, 0x63, 0x78,
	0xb0, 0xdb, 0x3c, 0xd5, 0x19, 0x72, 0x31, 0x17, 0xa8, 0xc8, 0x7f, 0x5f, 0x6f, 0xae, 0x6b, 0xde,
	0xe0, 0x0c, 0xee, 0xbb, 0xdf, 0x78, 0x07, 0xaa, 0xff, 0xdc, 0xb4, 0xbc, 0x5f, 0x9b, 0x96, 0xf7,
	0x7b, 0xd3, 0xf2, 0xa6, 0xa7, 0xc5, 0x79, 0xf6, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xab, 0x54,
	0xa1, 0x41, 0xce, 0x03, 0x00, 0x00,
}
