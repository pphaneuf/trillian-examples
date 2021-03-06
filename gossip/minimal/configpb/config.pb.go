// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package configpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	keyspb "github.com/google/trillian/crypto/keyspb"
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

// LogConfig describes the configuration options for a source Log,
// whose STH values are tracked/monitored.
type LogConfig struct {
	// Human-readable name for the log; must be unique
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Base URL for the log.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Log's public key. This is optional, but if omitted signatures from
	// the log will not be checked.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Minimum interval between requests to the log, for rate limiting.
	MinReqInterval       *duration.Duration `protobuf:"bytes,4,opt,name=min_req_interval,json=minReqInterval,proto3" json:"min_req_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LogConfig) Reset()         { *m = LogConfig{} }
func (m *LogConfig) String() string { return proto.CompactTextString(m) }
func (*LogConfig) ProtoMessage()    {}
func (*LogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *LogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogConfig.Unmarshal(m, b)
}
func (m *LogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogConfig.Marshal(b, m, deterministic)
}
func (m *LogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogConfig.Merge(m, src)
}
func (m *LogConfig) XXX_Size() int {
	return xxx_messageInfo_LogConfig.Size(m)
}
func (m *LogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LogConfig proto.InternalMessageInfo

func (m *LogConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogConfig) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *LogConfig) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *LogConfig) GetMinReqInterval() *duration.Duration {
	if m != nil {
		return m.MinReqInterval
	}
	return nil
}

// HubConfig describes the configuration options for a destination
// Gossip repository, which may be either a Gossip Hub or a regular
// CT log.
type HubConfig struct {
	// Human-readable name for the hub, must be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Base URL for the hub.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Hub's public key. This is optional, but if omitted signatures from
	// the log will not be checked.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Minimum interval between requests to the hub, for rate limiting.
	MinReqInterval *duration.Duration `protobuf:"bytes,4,opt,name=min_req_interval,json=minReqInterval,proto3" json:"min_req_interval,omitempty"`
	// Whether the hub is a true Gossip Hub.
	IsHub bool `protobuf:"varint,5,opt,name=is_hub,json=isHub,proto3" json:"is_hub,omitempty"`
	// Start point in a Hub beyond which entries of interest are found.
	// For example, for a Hub that is a regular CT Log this might be
	// a known point where the earliest encoded STH is to be found.
	// Specific to goshawk; not used by gosmin.
	StartIndex           int64    `protobuf:"varint,6,opt,name=start_index,json=startIndex,proto3" json:"start_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HubConfig) Reset()         { *m = HubConfig{} }
func (m *HubConfig) String() string { return proto.CompactTextString(m) }
func (*HubConfig) ProtoMessage()    {}
func (*HubConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

func (m *HubConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HubConfig.Unmarshal(m, b)
}
func (m *HubConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HubConfig.Marshal(b, m, deterministic)
}
func (m *HubConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HubConfig.Merge(m, src)
}
func (m *HubConfig) XXX_Size() int {
	return xxx_messageInfo_HubConfig.Size(m)
}
func (m *HubConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HubConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HubConfig proto.InternalMessageInfo

func (m *HubConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HubConfig) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HubConfig) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *HubConfig) GetMinReqInterval() *duration.Duration {
	if m != nil {
		return m.MinReqInterval
	}
	return nil
}

func (m *HubConfig) GetIsHub() bool {
	if m != nil {
		return m.IsHub
	}
	return false
}

func (m *HubConfig) GetStartIndex() int64 {
	if m != nil {
		return m.StartIndex
	}
	return 0
}

// GossipConfig describes the configuration of a gossiper.
type GossipConfig struct {
	// The source logs whose STHs will be logged.
	SourceLog []*LogConfig `protobuf:"bytes,1,rep,name=source_log,json=sourceLog,proto3" json:"source_log,omitempty"`
	// The destination hubs to which the minimal-gossip certificates will
	// be submitted.  These destination hubs need to be configured to accept
	// submissions.
	DestHub []*HubConfig `protobuf:"bytes,2,rep,name=dest_hub,json=destHub,proto3" json:"dest_hub,omitempty"`
	// The root certificate used for submissions to any CT logs, in PEM format;
	// this should include the public key corresponding to private_key below.
	// This field is unused if all destination hubs are true Gossip Hubs (rather
	// than CT logs that accept synthetic certificates).
	RootCert string `protobuf:"bytes,3,opt,name=root_cert,json=rootCert,proto3" json:"root_cert,omitempty"`
	// The private key that will be used to sign synthetic leaf certificates
	// that chain to the root_cert.
	PrivateKey *any.Any `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Number of buffered STHs allowed. Must not be negative.
	BufferSize           int32    `protobuf:"varint,5,opt,name=buffer_size,json=bufferSize,proto3" json:"buffer_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GossipConfig) Reset()         { *m = GossipConfig{} }
func (m *GossipConfig) String() string { return proto.CompactTextString(m) }
func (*GossipConfig) ProtoMessage()    {}
func (*GossipConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}

func (m *GossipConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GossipConfig.Unmarshal(m, b)
}
func (m *GossipConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GossipConfig.Marshal(b, m, deterministic)
}
func (m *GossipConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipConfig.Merge(m, src)
}
func (m *GossipConfig) XXX_Size() int {
	return xxx_messageInfo_GossipConfig.Size(m)
}
func (m *GossipConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GossipConfig proto.InternalMessageInfo

func (m *GossipConfig) GetSourceLog() []*LogConfig {
	if m != nil {
		return m.SourceLog
	}
	return nil
}

func (m *GossipConfig) GetDestHub() []*HubConfig {
	if m != nil {
		return m.DestHub
	}
	return nil
}

func (m *GossipConfig) GetRootCert() string {
	if m != nil {
		return m.RootCert
	}
	return ""
}

func (m *GossipConfig) GetPrivateKey() *any.Any {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *GossipConfig) GetBufferSize() int32 {
	if m != nil {
		return m.BufferSize
	}
	return 0
}

// GoshawkConfig describes the configuration of a gossiper.
type GoshawkConfig struct {
	// The source logs whose STHs will be checked.
	SourceLog []*LogConfig `protobuf:"bytes,1,rep,name=source_log,json=sourceLog,proto3" json:"source_log,omitempty"`
	// The destination hubs which will be scanned for minimal-gossip certificates.
	DestHub []*HubConfig `protobuf:"bytes,2,rep,name=dest_hub,json=destHub,proto3" json:"dest_hub,omitempty"`
	// Number of STHs pending verification that can be buffered up for each source log.
	BufferSize           int32    `protobuf:"varint,5,opt,name=buffer_size,json=bufferSize,proto3" json:"buffer_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoshawkConfig) Reset()         { *m = GoshawkConfig{} }
func (m *GoshawkConfig) String() string { return proto.CompactTextString(m) }
func (*GoshawkConfig) ProtoMessage()    {}
func (*GoshawkConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3}
}

func (m *GoshawkConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoshawkConfig.Unmarshal(m, b)
}
func (m *GoshawkConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoshawkConfig.Marshal(b, m, deterministic)
}
func (m *GoshawkConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoshawkConfig.Merge(m, src)
}
func (m *GoshawkConfig) XXX_Size() int {
	return xxx_messageInfo_GoshawkConfig.Size(m)
}
func (m *GoshawkConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GoshawkConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GoshawkConfig proto.InternalMessageInfo

func (m *GoshawkConfig) GetSourceLog() []*LogConfig {
	if m != nil {
		return m.SourceLog
	}
	return nil
}

func (m *GoshawkConfig) GetDestHub() []*HubConfig {
	if m != nil {
		return m.DestHub
	}
	return nil
}

func (m *GoshawkConfig) GetBufferSize() int32 {
	if m != nil {
		return m.BufferSize
	}
	return 0
}

func init() {
	proto.RegisterType((*LogConfig)(nil), "configpb.LogConfig")
	proto.RegisterType((*HubConfig)(nil), "configpb.HubConfig")
	proto.RegisterType((*GossipConfig)(nil), "configpb.GossipConfig")
	proto.RegisterType((*GoshawkConfig)(nil), "configpb.GoshawkConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0xcf, 0x6e, 0xd4, 0x3e,
	0x10, 0xc7, 0xe5, 0x6e, 0x77, 0x7f, 0x9b, 0x49, 0x7f, 0xa8, 0x18, 0x90, 0xd2, 0x45, 0xa2, 0xd1,
	0x9e, 0x72, 0x69, 0x82, 0x16, 0xf1, 0x00, 0xb0, 0xa0, 0x6d, 0x45, 0x0f, 0xc8, 0xdc, 0xb8, 0x44,
	0x76, 0xd6, 0xc9, 0x5a, 0x9b, 0xd8, 0xa9, 0xff, 0x14, 0xd2, 0xe7, 0xe0, 0x21, 0x78, 0x2c, 0x6e,
	0xbc, 0x06, 0x8a, 0x93, 0xed, 0xa1, 0x1c, 0xb8, 0x21, 0x71, 0xca, 0xf8, 0x3b, 0x33, 0x9a, 0xaf,
	0x3f, 0x13, 0xc3, 0x49, 0xa1, 0x64, 0x29, 0xaa, 0xb4, 0xd5, 0xca, 0x2a, 0x3c, 0x1f, 0x4e, 0x2d,
	0x5b, 0x2c, 0x0a, 0xdd, 0xb5, 0x56, 0x65, 0x7b, 0xde, 0x99, 0x96, 0x8d, 0x9f, 0xa1, 0x6a, 0x71,
	0x56, 0x29, 0x55, 0xd5, 0x3c, 0xf3, 0x27, 0xe6, 0xca, 0x8c, 0xca, 0x6e, 0x4c, 0xbd, 0x78, 0x98,
	0xda, 0x3a, 0x4d, 0xad, 0x50, 0x72, 0xc8, 0x2f, 0xbf, 0x23, 0x08, 0xae, 0x55, 0xb5, 0xf6, 0x63,
	0x30, 0x86, 0x63, 0x49, 0x1b, 0x1e, 0xa1, 0x18, 0x25, 0x01, 0xf1, 0x31, 0x3e, 0x85, 0x89, 0xd3,
	0x75, 0x74, 0xe4, 0xa5, 0x3e, 0xc4, 0x2f, 0x01, 0x5a, 0xc7, 0x6a, 0x51, 0xe4, 0x7b, 0xde, 0x45,
	0x93, 0x18, 0x25, 0xe1, 0xea, 0x71, 0x3a, 0x3a, 0xfa, 0xe8, 0x33, 0x1f, 0x78, 0x47, 0x82, 0xf6,
	0x10, 0xe2, 0x35, 0x9c, 0x36, 0x42, 0xe6, 0x9a, 0xdf, 0xe4, 0x42, 0x5a, 0xae, 0x6f, 0x69, 0x1d,
	0x1d, 0xfb, 0xbe, 0xb3, 0x74, 0x30, 0x98, 0x1e, 0x0c, 0xa6, 0xef, 0x46, 0x83, 0xe4, 0x51, 0x23,
	0x24, 0xe1, 0x37, 0x57, 0x63, 0xc3, 0xf2, 0x07, 0x82, 0xe0, 0xd2, 0xb1, 0x7f, 0xc0, 0x2a, 0x7e,
	0x06, 0x33, 0x61, 0xf2, 0x9d, 0x63, 0xd1, 0x34, 0x46, 0xc9, 0x9c, 0x4c, 0x85, 0xb9, 0x74, 0x0c,
	0x9f, 0x43, 0x68, 0x2c, 0xd5, 0x36, 0x17, 0x72, 0xcb, 0xbf, 0x46, 0xb3, 0x18, 0x25, 0x13, 0x02,
	0x5e, 0xba, 0xea, 0x95, 0xe5, 0x4f, 0x04, 0x27, 0x1b, 0x65, 0x8c, 0x68, 0xc7, 0x5b, 0xae, 0x00,
	0x8c, 0x72, 0xba, 0xe0, 0x79, 0xad, 0xaa, 0x08, 0xc5, 0x93, 0x24, 0x5c, 0x3d, 0x49, 0x0f, 0x3f,
	0x45, 0x7a, 0xbf, 0x39, 0x12, 0x0c, 0x65, 0xd7, 0xaa, 0xc2, 0x29, 0xcc, 0xb7, 0xdc, 0x58, 0x3f,
	0xfe, 0xe8, 0x61, 0xc7, 0x3d, 0x40, 0xf2, 0x5f, 0x5f, 0xd4, 0xbb, 0x7a, 0x0e, 0x81, 0x56, 0xca,
	0xe6, 0x05, 0xd7, 0xd6, 0x23, 0x0a, 0xc8, 0xbc, 0x17, 0xd6, 0x5c, 0x5b, 0xfc, 0x1a, 0xc2, 0x56,
	0x8b, 0x5b, 0x6a, 0xb9, 0x27, 0x38, 0x90, 0x78, 0xfa, 0x1b, 0x89, 0x37, 0xb2, 0x23, 0x30, 0x16,
	0xf6, 0x14, 0xcf, 0x21, 0x64, 0xae, 0x2c, 0xb9, 0xce, 0x8d, 0xb8, 0xe3, 0x9e, 0xc2, 0x94, 0xc0,
	0x20, 0x7d, 0x12, 0x77, 0x7c, 0xf9, 0x0d, 0xc1, 0xff, 0x1b, 0x65, 0x76, 0xf4, 0xcb, 0xfe, 0x2f,
	0x5e, 0xf5, 0x4f, 0xb6, 0xde, 0x6e, 0x3e, 0xbf, 0xaf, 0x84, 0xdd, 0x39, 0x96, 0x16, 0xaa, 0xc9,
	0xc6, 0xb7, 0xd3, 0x83, 0x11, 0xa5, 0x28, 0xa8, 0xe5, 0x17, 0x56, 0x53, 0x69, 0x5a, 0xaa, 0xb9,
	0x2c, 0xba, 0x8b, 0x4a, 0x65, 0x95, 0xdf, 0x56, 0xd6, 0x08, 0x29, 0x1a, 0x5a, 0x67, 0x87, 0xe1,
	0x6c, 0xe6, 0xd1, 0xbc, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x90, 0x18, 0xb7, 0x25, 0xcf, 0x03,
	0x00, 0x00,
}
