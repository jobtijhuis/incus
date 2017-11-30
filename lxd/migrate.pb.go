// Code generated by protoc-gen-go.
// source: lxd/migrate.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	lxd/migrate.proto

It has these top-level messages:
	IDMapType
	Config
	Device
	Snapshot
	MigrationHeader
	MigrationControl
	MigrationSync
*/
package main

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

type MigrationFSType int32

const (
	MigrationFSType_RSYNC MigrationFSType = 0
	MigrationFSType_BTRFS MigrationFSType = 1
	MigrationFSType_ZFS   MigrationFSType = 2
	MigrationFSType_RBD   MigrationFSType = 3
)

var MigrationFSType_name = map[int32]string{
	0: "RSYNC",
	1: "BTRFS",
	2: "ZFS",
	3: "RBD",
}
var MigrationFSType_value = map[string]int32{
	"RSYNC": 0,
	"BTRFS": 1,
	"ZFS":   2,
	"RBD":   3,
}

func (x MigrationFSType) Enum() *MigrationFSType {
	p := new(MigrationFSType)
	*p = x
	return p
}
func (x MigrationFSType) String() string {
	return proto.EnumName(MigrationFSType_name, int32(x))
}
func (x *MigrationFSType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MigrationFSType_value, data, "MigrationFSType")
	if err != nil {
		return err
	}
	*x = MigrationFSType(value)
	return nil
}
func (MigrationFSType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CRIUType int32

const (
	CRIUType_CRIU_RSYNC CRIUType = 0
	CRIUType_PHAUL      CRIUType = 1
	CRIUType_NONE       CRIUType = 2
)

var CRIUType_name = map[int32]string{
	0: "CRIU_RSYNC",
	1: "PHAUL",
	2: "NONE",
}
var CRIUType_value = map[string]int32{
	"CRIU_RSYNC": 0,
	"PHAUL":      1,
	"NONE":       2,
}

func (x CRIUType) Enum() *CRIUType {
	p := new(CRIUType)
	*p = x
	return p
}
func (x CRIUType) String() string {
	return proto.EnumName(CRIUType_name, int32(x))
}
func (x *CRIUType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CRIUType_value, data, "CRIUType")
	if err != nil {
		return err
	}
	*x = CRIUType(value)
	return nil
}
func (CRIUType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type IDMapType struct {
	Isuid            *bool  `protobuf:"varint,1,req,name=isuid" json:"isuid,omitempty"`
	Isgid            *bool  `protobuf:"varint,2,req,name=isgid" json:"isgid,omitempty"`
	Hostid           *int32 `protobuf:"varint,3,req,name=hostid" json:"hostid,omitempty"`
	Nsid             *int32 `protobuf:"varint,4,req,name=nsid" json:"nsid,omitempty"`
	Maprange         *int32 `protobuf:"varint,5,req,name=maprange" json:"maprange,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IDMapType) Reset()                    { *m = IDMapType{} }
func (m *IDMapType) String() string            { return proto.CompactTextString(m) }
func (*IDMapType) ProtoMessage()               {}
func (*IDMapType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IDMapType) GetIsuid() bool {
	if m != nil && m.Isuid != nil {
		return *m.Isuid
	}
	return false
}

func (m *IDMapType) GetIsgid() bool {
	if m != nil && m.Isgid != nil {
		return *m.Isgid
	}
	return false
}

func (m *IDMapType) GetHostid() int32 {
	if m != nil && m.Hostid != nil {
		return *m.Hostid
	}
	return 0
}

func (m *IDMapType) GetNsid() int32 {
	if m != nil && m.Nsid != nil {
		return *m.Nsid
	}
	return 0
}

func (m *IDMapType) GetMaprange() int32 {
	if m != nil && m.Maprange != nil {
		return *m.Maprange
	}
	return 0
}

type Config struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Config) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Device struct {
	Name             *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Config           []*Config `protobuf:"bytes,2,rep,name=config" json:"config,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Device) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Device) GetConfig() []*Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type Snapshot struct {
	Name             *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	LocalConfig      []*Config `protobuf:"bytes,2,rep,name=localConfig" json:"localConfig,omitempty"`
	Profiles         []string  `protobuf:"bytes,3,rep,name=profiles" json:"profiles,omitempty"`
	Ephemeral        *bool     `protobuf:"varint,4,req,name=ephemeral" json:"ephemeral,omitempty"`
	LocalDevices     []*Device `protobuf:"bytes,5,rep,name=localDevices" json:"localDevices,omitempty"`
	Architecture     *int32    `protobuf:"varint,6,req,name=architecture" json:"architecture,omitempty"`
	Stateful         *bool     `protobuf:"varint,7,req,name=stateful" json:"stateful,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (m *Snapshot) String() string            { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Snapshot) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Snapshot) GetLocalConfig() []*Config {
	if m != nil {
		return m.LocalConfig
	}
	return nil
}

func (m *Snapshot) GetProfiles() []string {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *Snapshot) GetEphemeral() bool {
	if m != nil && m.Ephemeral != nil {
		return *m.Ephemeral
	}
	return false
}

func (m *Snapshot) GetLocalDevices() []*Device {
	if m != nil {
		return m.LocalDevices
	}
	return nil
}

func (m *Snapshot) GetArchitecture() int32 {
	if m != nil && m.Architecture != nil {
		return *m.Architecture
	}
	return 0
}

func (m *Snapshot) GetStateful() bool {
	if m != nil && m.Stateful != nil {
		return *m.Stateful
	}
	return false
}

type MigrationHeader struct {
	Fs               *MigrationFSType `protobuf:"varint,1,req,name=fs,enum=main.MigrationFSType" json:"fs,omitempty"`
	Criu             *CRIUType        `protobuf:"varint,2,opt,name=criu,enum=main.CRIUType" json:"criu,omitempty"`
	Idmap            []*IDMapType     `protobuf:"bytes,3,rep,name=idmap" json:"idmap,omitempty"`
	SnapshotNames    []string         `protobuf:"bytes,4,rep,name=snapshotNames" json:"snapshotNames,omitempty"`
	Snapshots        []*Snapshot      `protobuf:"bytes,5,rep,name=snapshots" json:"snapshots,omitempty"`
	Predump          *bool            `protobuf:"varint,7,opt,name=predump" json:"predump,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MigrationHeader) Reset()                    { *m = MigrationHeader{} }
func (m *MigrationHeader) String() string            { return proto.CompactTextString(m) }
func (*MigrationHeader) ProtoMessage()               {}
func (*MigrationHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MigrationHeader) GetFs() MigrationFSType {
	if m != nil && m.Fs != nil {
		return *m.Fs
	}
	return MigrationFSType_RSYNC
}

func (m *MigrationHeader) GetCriu() CRIUType {
	if m != nil && m.Criu != nil {
		return *m.Criu
	}
	return CRIUType_CRIU_RSYNC
}

func (m *MigrationHeader) GetIdmap() []*IDMapType {
	if m != nil {
		return m.Idmap
	}
	return nil
}

func (m *MigrationHeader) GetSnapshotNames() []string {
	if m != nil {
		return m.SnapshotNames
	}
	return nil
}

func (m *MigrationHeader) GetSnapshots() []*Snapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

func (m *MigrationHeader) GetPredump() bool {
	if m != nil && m.Predump != nil {
		return *m.Predump
	}
	return false
}

type MigrationControl struct {
	Success *bool `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	// optional failure message if sending a failure
	Message          *string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MigrationControl) Reset()                    { *m = MigrationControl{} }
func (m *MigrationControl) String() string            { return proto.CompactTextString(m) }
func (*MigrationControl) ProtoMessage()               {}
func (*MigrationControl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MigrationControl) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *MigrationControl) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type MigrationSync struct {
	FinalPreDump     *bool  `protobuf:"varint,1,req,name=finalPreDump" json:"finalPreDump,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MigrationSync) Reset()                    { *m = MigrationSync{} }
func (m *MigrationSync) String() string            { return proto.CompactTextString(m) }
func (*MigrationSync) ProtoMessage()               {}
func (*MigrationSync) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MigrationSync) GetFinalPreDump() bool {
	if m != nil && m.FinalPreDump != nil {
		return *m.FinalPreDump
	}
	return false
}

func init() {
	proto.RegisterType((*IDMapType)(nil), "main.IDMapType")
	proto.RegisterType((*Config)(nil), "main.Config")
	proto.RegisterType((*Device)(nil), "main.Device")
	proto.RegisterType((*Snapshot)(nil), "main.Snapshot")
	proto.RegisterType((*MigrationHeader)(nil), "main.MigrationHeader")
	proto.RegisterType((*MigrationControl)(nil), "main.MigrationControl")
	proto.RegisterType((*MigrationSync)(nil), "main.MigrationSync")
	proto.RegisterEnum("main.MigrationFSType", MigrationFSType_name, MigrationFSType_value)
	proto.RegisterEnum("main.CRIUType", CRIUType_name, CRIUType_value)
}

func init() { proto.RegisterFile("lxd/migrate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x6b, 0xdb, 0x4e,
	0x10, 0xfd, 0x49, 0xf2, 0x87, 0x34, 0x49, 0x1c, 0xfd, 0x96, 0xb6, 0x2c, 0xa5, 0x07, 0x21, 0x12,
	0x10, 0xa1, 0x38, 0xc1, 0x3d, 0xf5, 0xd8, 0xd8, 0x35, 0x09, 0x34, 0x6e, 0x58, 0x25, 0x87, 0xf6,
	0x52, 0x16, 0x69, 0x25, 0x2f, 0xd5, 0x17, 0xbb, 0x52, 0xa8, 0x4f, 0xfd, 0x9b, 0x7b, 0xed, 0xa9,
	0xec, 0xea, 0x23, 0x36, 0xb4, 0xb7, 0x79, 0x6f, 0x9e, 0xe6, 0x69, 0xde, 0x0e, 0xfc, 0x9f, 0xfd,
	0x88, 0x2f, 0x73, 0x9e, 0x0a, 0x5a, 0xb3, 0x79, 0x25, 0xca, 0xba, 0x44, 0xa3, 0x9c, 0xf2, 0xc2,
	0xff, 0x09, 0xce, 0xed, 0xea, 0x8e, 0x56, 0x0f, 0xbb, 0x8a, 0xa1, 0x17, 0x30, 0xe6, 0xb2, 0xe1,
	0x31, 0x36, 0x3c, 0x33, 0xb0, 0x49, 0x0b, 0x5a, 0x36, 0xe5, 0x31, 0x36, 0x7b, 0x36, 0xe5, 0x31,
	0x7a, 0x05, 0x93, 0x6d, 0x29, 0x6b, 0x1e, 0x63, 0xcb, 0x33, 0x83, 0x31, 0xe9, 0x10, 0x42, 0x30,
	0x2a, 0x24, 0x8f, 0xf1, 0x48, 0xb3, 0xba, 0x46, 0xaf, 0xc1, 0xce, 0x69, 0x25, 0x68, 0x91, 0x32,
	0x3c, 0xd6, 0xfc, 0x80, 0xfd, 0x2b, 0x98, 0x2c, 0xcb, 0x22, 0xe1, 0x29, 0x72, 0xc1, 0xfa, 0xce,
	0x76, 0xda, 0xdb, 0x21, 0xaa, 0x54, 0xce, 0x4f, 0x34, 0x6b, 0x98, 0x76, 0x76, 0x48, 0x0b, 0xfc,
	0x6b, 0x98, 0xac, 0xd8, 0x13, 0x8f, 0x98, 0xf6, 0xa2, 0x39, 0xeb, 0x3e, 0xd1, 0x35, 0x3a, 0x83,
	0x49, 0xa4, 0xe7, 0x61, 0xd3, 0xb3, 0x82, 0xa3, 0xc5, 0xf1, 0x5c, 0xed, 0x39, 0x6f, 0x3d, 0x48,
	0xd7, 0xf3, 0x7f, 0x1b, 0x60, 0x87, 0x05, 0xad, 0xe4, 0xb6, 0xac, 0xff, 0x3a, 0x66, 0x0e, 0x47,
	0x59, 0x19, 0xd1, 0x6c, 0xf9, 0xef, 0x59, 0xfb, 0x02, 0xb5, 0x62, 0x25, 0xca, 0x84, 0x67, 0x4c,
	0x62, 0xcb, 0xb3, 0x02, 0x87, 0x0c, 0x18, 0xbd, 0x01, 0x87, 0x55, 0x5b, 0x96, 0x33, 0x41, 0x33,
	0x9d, 0x8b, 0x4d, 0x9e, 0x09, 0x74, 0x05, 0xc7, 0x7a, 0x50, 0xbb, 0x93, 0xc4, 0xe3, 0x7d, 0xab,
	0x96, 0x24, 0x07, 0x0a, 0xe4, 0xc3, 0x31, 0x15, 0xd1, 0x96, 0xd7, 0x2c, 0xaa, 0x1b, 0xc1, 0xf0,
	0x44, 0x47, 0x7a, 0xc0, 0xa9, 0xff, 0x91, 0x35, 0xad, 0x59, 0xd2, 0x64, 0x78, 0xaa, 0x2d, 0x07,
	0xec, 0xff, 0x32, 0xe0, 0xf4, 0x4e, 0xdf, 0x02, 0x2f, 0x8b, 0x1b, 0x46, 0x63, 0x26, 0xd0, 0x39,
	0x98, 0x89, 0xd4, 0x09, 0xcc, 0x16, 0x2f, 0x5b, 0xef, 0x41, 0xb2, 0x0e, 0xd5, 0x75, 0x10, 0x33,
	0x51, 0xd6, 0xa3, 0x48, 0xf0, 0x06, 0x9b, 0x9e, 0x11, 0xcc, 0x16, 0xb3, 0x2e, 0x0f, 0x72, 0xfb,
	0xa8, 0x15, 0xba, 0x87, 0xce, 0x61, 0xcc, 0xe3, 0x9c, 0x56, 0x3a, 0x87, 0xa3, 0xc5, 0x69, 0x2b,
	0x1a, 0xae, 0x8c, 0xb4, 0x5d, 0x74, 0x06, 0x27, 0xb2, 0x7b, 0x81, 0x0d, 0xcd, 0x99, 0xc4, 0x23,
	0x1d, 0xdb, 0x21, 0x89, 0xde, 0x82, 0xd3, 0x13, 0x7d, 0x34, 0x9d, 0x6b, 0xff, 0x7c, 0xe4, 0x59,
	0x80, 0x30, 0x4c, 0x2b, 0xc1, 0xe2, 0x26, 0xaf, 0xf0, 0xd4, 0x33, 0x02, 0x9b, 0xf4, 0xd0, 0x5f,
	0x83, 0x3b, 0xec, 0xb3, 0x2c, 0x8b, 0x5a, 0x94, 0x99, 0x52, 0xcb, 0x26, 0x8a, 0x98, 0x94, 0xdd,
	0xc1, 0xf7, 0x50, 0x75, 0x72, 0x26, 0x25, 0x4d, 0x99, 0xde, 0xd4, 0x21, 0x3d, 0xf4, 0xdf, 0xc1,
	0xc9, 0x30, 0x27, 0xdc, 0x15, 0x91, 0x7a, 0x8c, 0x84, 0x17, 0x34, 0xbb, 0x17, 0x6c, 0xa5, 0x7c,
	0xdb, 0x49, 0x07, 0xdc, 0xc5, 0xfb, 0xbd, 0xbc, 0xdb, 0x30, 0x91, 0x03, 0x63, 0x12, 0x7e, 0xd9,
	0x2c, 0xdd, 0xff, 0x54, 0x79, 0xfd, 0x40, 0xd6, 0xa1, 0x6b, 0xa0, 0x29, 0x58, 0x5f, 0xd7, 0xa1,
	0x6b, 0xaa, 0x82, 0x5c, 0xaf, 0x5c, 0xeb, 0xe2, 0x12, 0xec, 0x3e, 0x5e, 0x34, 0x03, 0x50, 0xf5,
	0xb7, 0xbd, 0x0f, 0xef, 0x6f, 0x3e, 0x3c, 0x7e, 0x72, 0x0d, 0x64, 0xc3, 0x68, 0xf3, 0x79, 0xf3,
	0xd1, 0x35, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xe1, 0x62, 0xc7, 0xea, 0x03, 0x00, 0x00,
}
