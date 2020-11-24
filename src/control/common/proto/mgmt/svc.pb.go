// Code generated by protoc-gen-go. DO NOT EDIT.
// source: svc.proto

package mgmt

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

type JoinResp_State int32

const (
	JoinResp_IN  JoinResp_State = 0
	JoinResp_OUT JoinResp_State = 1
)

var JoinResp_State_name = map[int32]string{
	0: "IN",
	1: "OUT",
}

var JoinResp_State_value = map[string]int32{
	"IN":  0,
	"OUT": 1,
}

func (x JoinResp_State) String() string {
	return proto.EnumName(JoinResp_State_name, int32(x))
}

func (JoinResp_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{4, 0}
}

// Generic response just containing DER from IO server.
type DaosResp struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DaosResp) Reset()         { *m = DaosResp{} }
func (m *DaosResp) String() string { return proto.CompactTextString(m) }
func (*DaosResp) ProtoMessage()    {}
func (*DaosResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{0}
}

func (m *DaosResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DaosResp.Unmarshal(m, b)
}
func (m *DaosResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DaosResp.Marshal(b, m, deterministic)
}
func (m *DaosResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DaosResp.Merge(m, src)
}
func (m *DaosResp) XXX_Size() int {
	return xxx_messageInfo_DaosResp.Size(m)
}
func (m *DaosResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DaosResp.DiscardUnknown(m)
}

var xxx_messageInfo_DaosResp proto.InternalMessageInfo

func (m *DaosResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type GroupUpdateReq struct {
	MapVersion           uint32                   `protobuf:"varint,1,opt,name=map_version,json=mapVersion,proto3" json:"map_version,omitempty"`
	Servers              []*GroupUpdateReq_Server `protobuf:"bytes,2,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GroupUpdateReq) Reset()         { *m = GroupUpdateReq{} }
func (m *GroupUpdateReq) String() string { return proto.CompactTextString(m) }
func (*GroupUpdateReq) ProtoMessage()    {}
func (*GroupUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{1}
}

func (m *GroupUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupUpdateReq.Unmarshal(m, b)
}
func (m *GroupUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupUpdateReq.Marshal(b, m, deterministic)
}
func (m *GroupUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupUpdateReq.Merge(m, src)
}
func (m *GroupUpdateReq) XXX_Size() int {
	return xxx_messageInfo_GroupUpdateReq.Size(m)
}
func (m *GroupUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_GroupUpdateReq proto.InternalMessageInfo

func (m *GroupUpdateReq) GetMapVersion() uint32 {
	if m != nil {
		return m.MapVersion
	}
	return 0
}

func (m *GroupUpdateReq) GetServers() []*GroupUpdateReq_Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

type GroupUpdateReq_Server struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Uri                  string   `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupUpdateReq_Server) Reset()         { *m = GroupUpdateReq_Server{} }
func (m *GroupUpdateReq_Server) String() string { return proto.CompactTextString(m) }
func (*GroupUpdateReq_Server) ProtoMessage()    {}
func (*GroupUpdateReq_Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{1, 0}
}

func (m *GroupUpdateReq_Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupUpdateReq_Server.Unmarshal(m, b)
}
func (m *GroupUpdateReq_Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupUpdateReq_Server.Marshal(b, m, deterministic)
}
func (m *GroupUpdateReq_Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupUpdateReq_Server.Merge(m, src)
}
func (m *GroupUpdateReq_Server) XXX_Size() int {
	return xxx_messageInfo_GroupUpdateReq_Server.Size(m)
}
func (m *GroupUpdateReq_Server) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupUpdateReq_Server.DiscardUnknown(m)
}

var xxx_messageInfo_GroupUpdateReq_Server proto.InternalMessageInfo

func (m *GroupUpdateReq_Server) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *GroupUpdateReq_Server) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type GroupUpdateResp struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupUpdateResp) Reset()         { *m = GroupUpdateResp{} }
func (m *GroupUpdateResp) String() string { return proto.CompactTextString(m) }
func (*GroupUpdateResp) ProtoMessage()    {}
func (*GroupUpdateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{2}
}

func (m *GroupUpdateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupUpdateResp.Unmarshal(m, b)
}
func (m *GroupUpdateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupUpdateResp.Marshal(b, m, deterministic)
}
func (m *GroupUpdateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupUpdateResp.Merge(m, src)
}
func (m *GroupUpdateResp) XXX_Size() int {
	return xxx_messageInfo_GroupUpdateResp.Size(m)
}
func (m *GroupUpdateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupUpdateResp.DiscardUnknown(m)
}

var xxx_messageInfo_GroupUpdateResp proto.InternalMessageInfo

func (m *GroupUpdateResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type JoinReq struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Rank                 uint32   `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Uri                  string   `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Nctxs                uint32   `protobuf:"varint,4,opt,name=nctxs,proto3" json:"nctxs,omitempty"`
	Addr                 string   `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	SrvFaultDomain       string   `protobuf:"bytes,6,opt,name=srvFaultDomain,proto3" json:"srvFaultDomain,omitempty"`
	Idx                  uint32   `protobuf:"varint,7,opt,name=idx,proto3" json:"idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinReq) Reset()         { *m = JoinReq{} }
func (m *JoinReq) String() string { return proto.CompactTextString(m) }
func (*JoinReq) ProtoMessage()    {}
func (*JoinReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{3}
}

func (m *JoinReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinReq.Unmarshal(m, b)
}
func (m *JoinReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinReq.Marshal(b, m, deterministic)
}
func (m *JoinReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinReq.Merge(m, src)
}
func (m *JoinReq) XXX_Size() int {
	return xxx_messageInfo_JoinReq.Size(m)
}
func (m *JoinReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinReq.DiscardUnknown(m)
}

var xxx_messageInfo_JoinReq proto.InternalMessageInfo

func (m *JoinReq) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *JoinReq) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *JoinReq) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *JoinReq) GetNctxs() uint32 {
	if m != nil {
		return m.Nctxs
	}
	return 0
}

func (m *JoinReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *JoinReq) GetSrvFaultDomain() string {
	if m != nil {
		return m.SrvFaultDomain
	}
	return ""
}

func (m *JoinReq) GetIdx() uint32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

type JoinResp struct {
	Status               int32          `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Rank                 uint32         `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	State                JoinResp_State `protobuf:"varint,3,opt,name=state,proto3,enum=mgmt.JoinResp_State" json:"state,omitempty"`
	FaultDomain          string         `protobuf:"bytes,4,opt,name=faultDomain,proto3" json:"faultDomain,omitempty"`
	LocalJoin            bool           `protobuf:"varint,5,opt,name=localJoin,proto3" json:"localJoin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *JoinResp) Reset()         { *m = JoinResp{} }
func (m *JoinResp) String() string { return proto.CompactTextString(m) }
func (*JoinResp) ProtoMessage()    {}
func (*JoinResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{4}
}

func (m *JoinResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResp.Unmarshal(m, b)
}
func (m *JoinResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResp.Marshal(b, m, deterministic)
}
func (m *JoinResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResp.Merge(m, src)
}
func (m *JoinResp) XXX_Size() int {
	return xxx_messageInfo_JoinResp.Size(m)
}
func (m *JoinResp) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResp.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResp proto.InternalMessageInfo

func (m *JoinResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *JoinResp) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *JoinResp) GetState() JoinResp_State {
	if m != nil {
		return m.State
	}
	return JoinResp_IN
}

func (m *JoinResp) GetFaultDomain() string {
	if m != nil {
		return m.FaultDomain
	}
	return ""
}

func (m *JoinResp) GetLocalJoin() bool {
	if m != nil {
		return m.LocalJoin
	}
	return false
}

type LeaderQueryReq struct {
	System               string   `protobuf:"bytes,1,opt,name=system,proto3" json:"system,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaderQueryReq) Reset()         { *m = LeaderQueryReq{} }
func (m *LeaderQueryReq) String() string { return proto.CompactTextString(m) }
func (*LeaderQueryReq) ProtoMessage()    {}
func (*LeaderQueryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{5}
}

func (m *LeaderQueryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaderQueryReq.Unmarshal(m, b)
}
func (m *LeaderQueryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaderQueryReq.Marshal(b, m, deterministic)
}
func (m *LeaderQueryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaderQueryReq.Merge(m, src)
}
func (m *LeaderQueryReq) XXX_Size() int {
	return xxx_messageInfo_LeaderQueryReq.Size(m)
}
func (m *LeaderQueryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaderQueryReq.DiscardUnknown(m)
}

var xxx_messageInfo_LeaderQueryReq proto.InternalMessageInfo

func (m *LeaderQueryReq) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

type LeaderQueryResp struct {
	CurrentLeader        string   `protobuf:"bytes,1,opt,name=currentLeader,proto3" json:"currentLeader,omitempty"`
	Replicas             []string `protobuf:"bytes,2,rep,name=replicas,proto3" json:"replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaderQueryResp) Reset()         { *m = LeaderQueryResp{} }
func (m *LeaderQueryResp) String() string { return proto.CompactTextString(m) }
func (*LeaderQueryResp) ProtoMessage()    {}
func (*LeaderQueryResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{6}
}

func (m *LeaderQueryResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaderQueryResp.Unmarshal(m, b)
}
func (m *LeaderQueryResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaderQueryResp.Marshal(b, m, deterministic)
}
func (m *LeaderQueryResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaderQueryResp.Merge(m, src)
}
func (m *LeaderQueryResp) XXX_Size() int {
	return xxx_messageInfo_LeaderQueryResp.Size(m)
}
func (m *LeaderQueryResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaderQueryResp.DiscardUnknown(m)
}

var xxx_messageInfo_LeaderQueryResp proto.InternalMessageInfo

func (m *LeaderQueryResp) GetCurrentLeader() string {
	if m != nil {
		return m.CurrentLeader
	}
	return ""
}

func (m *LeaderQueryResp) GetReplicas() []string {
	if m != nil {
		return m.Replicas
	}
	return nil
}

// RASEvent describes a RAS event in the DAOS system.
type RASEvent struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Severity             uint32   `protobuf:"varint,3,opt,name=severity,proto3" json:"severity,omitempty"`
	Msg                  string   `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	Rank                 uint32   `protobuf:"varint,5,opt,name=rank,proto3" json:"rank,omitempty"`
	Hostname             string   `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Data                 []byte   `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	Eid                  uint32   `protobuf:"varint,8,opt,name=eid,proto3" json:"eid,omitempty"`
	Type                 uint32   `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	InstanceIdx          uint32   `protobuf:"varint,10,opt,name=instance_idx,json=instanceIdx,proto3" json:"instance_idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RASEvent) Reset()         { *m = RASEvent{} }
func (m *RASEvent) String() string { return proto.CompactTextString(m) }
func (*RASEvent) ProtoMessage()    {}
func (*RASEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{7}
}

func (m *RASEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RASEvent.Unmarshal(m, b)
}
func (m *RASEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RASEvent.Marshal(b, m, deterministic)
}
func (m *RASEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RASEvent.Merge(m, src)
}
func (m *RASEvent) XXX_Size() int {
	return xxx_messageInfo_RASEvent.Size(m)
}
func (m *RASEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RASEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RASEvent proto.InternalMessageInfo

func (m *RASEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RASEvent) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *RASEvent) GetSeverity() uint32 {
	if m != nil {
		return m.Severity
	}
	return 0
}

func (m *RASEvent) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RASEvent) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *RASEvent) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *RASEvent) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RASEvent) GetEid() uint32 {
	if m != nil {
		return m.Eid
	}
	return 0
}

func (m *RASEvent) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RASEvent) GetInstanceIdx() uint32 {
	if m != nil {
		return m.InstanceIdx
	}
	return 0
}

// ClusterEventReq wraps an event subtype e.g. RASEvent.
type ClusterEventReq struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*ClusterEventReq_Ras
	Event                isClusterEventReq_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClusterEventReq) Reset()         { *m = ClusterEventReq{} }
func (m *ClusterEventReq) String() string { return proto.CompactTextString(m) }
func (*ClusterEventReq) ProtoMessage()    {}
func (*ClusterEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{8}
}

func (m *ClusterEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterEventReq.Unmarshal(m, b)
}
func (m *ClusterEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterEventReq.Marshal(b, m, deterministic)
}
func (m *ClusterEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterEventReq.Merge(m, src)
}
func (m *ClusterEventReq) XXX_Size() int {
	return xxx_messageInfo_ClusterEventReq.Size(m)
}
func (m *ClusterEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterEventReq proto.InternalMessageInfo

func (m *ClusterEventReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isClusterEventReq_Event interface {
	isClusterEventReq_Event()
}

type ClusterEventReq_Ras struct {
	Ras *RASEvent `protobuf:"bytes,2,opt,name=ras,proto3,oneof"`
}

func (*ClusterEventReq_Ras) isClusterEventReq_Event() {}

func (m *ClusterEventReq) GetEvent() isClusterEventReq_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ClusterEventReq) GetRas() *RASEvent {
	if x, ok := m.GetEvent().(*ClusterEventReq_Ras); ok {
		return x.Ras
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClusterEventReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClusterEventReq_Ras)(nil),
	}
}

// ClusterEventResp contains response to an event notification.
type ClusterEventResp struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterEventResp) Reset()         { *m = ClusterEventResp{} }
func (m *ClusterEventResp) String() string { return proto.CompactTextString(m) }
func (*ClusterEventResp) ProtoMessage()    {}
func (*ClusterEventResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{9}
}

func (m *ClusterEventResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterEventResp.Unmarshal(m, b)
}
func (m *ClusterEventResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterEventResp.Marshal(b, m, deterministic)
}
func (m *ClusterEventResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterEventResp.Merge(m, src)
}
func (m *ClusterEventResp) XXX_Size() int {
	return xxx_messageInfo_ClusterEventResp.Size(m)
}
func (m *ClusterEventResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterEventResp.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterEventResp proto.InternalMessageInfo

func (m *ClusterEventResp) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAttachInfoReq struct {
	Sys                  string   `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty"`
	AllRanks             bool     `protobuf:"varint,2,opt,name=allRanks,proto3" json:"allRanks,omitempty"`
	Jobid                string   `protobuf:"bytes,3,opt,name=jobid,proto3" json:"jobid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAttachInfoReq) Reset()         { *m = GetAttachInfoReq{} }
func (m *GetAttachInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetAttachInfoReq) ProtoMessage()    {}
func (*GetAttachInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{10}
}

func (m *GetAttachInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttachInfoReq.Unmarshal(m, b)
}
func (m *GetAttachInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttachInfoReq.Marshal(b, m, deterministic)
}
func (m *GetAttachInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttachInfoReq.Merge(m, src)
}
func (m *GetAttachInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetAttachInfoReq.Size(m)
}
func (m *GetAttachInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttachInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttachInfoReq proto.InternalMessageInfo

func (m *GetAttachInfoReq) GetSys() string {
	if m != nil {
		return m.Sys
	}
	return ""
}

func (m *GetAttachInfoReq) GetAllRanks() bool {
	if m != nil {
		return m.AllRanks
	}
	return false
}

func (m *GetAttachInfoReq) GetJobid() string {
	if m != nil {
		return m.Jobid
	}
	return ""
}

type GetAttachInfoResp struct {
	Status int32                    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Psrs   []*GetAttachInfoResp_Psr `protobuf:"bytes,2,rep,name=psrs,proto3" json:"psrs,omitempty"`
	// These CaRT settings are shared with the
	// libdaos client to aid in CaRT initialization.
	Provider             string   `protobuf:"bytes,3,opt,name=Provider,proto3" json:"Provider,omitempty"`
	Interface            string   `protobuf:"bytes,4,opt,name=Interface,proto3" json:"Interface,omitempty"`
	Domain               string   `protobuf:"bytes,5,opt,name=Domain,proto3" json:"Domain,omitempty"`
	CrtCtxShareAddr      uint32   `protobuf:"varint,6,opt,name=CrtCtxShareAddr,proto3" json:"CrtCtxShareAddr,omitempty"`
	CrtTimeout           uint32   `protobuf:"varint,7,opt,name=CrtTimeout,proto3" json:"CrtTimeout,omitempty"`
	NetDevClass          uint32   `protobuf:"varint,8,opt,name=NetDevClass,proto3" json:"NetDevClass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAttachInfoResp) Reset()         { *m = GetAttachInfoResp{} }
func (m *GetAttachInfoResp) String() string { return proto.CompactTextString(m) }
func (*GetAttachInfoResp) ProtoMessage()    {}
func (*GetAttachInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{11}
}

func (m *GetAttachInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttachInfoResp.Unmarshal(m, b)
}
func (m *GetAttachInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttachInfoResp.Marshal(b, m, deterministic)
}
func (m *GetAttachInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttachInfoResp.Merge(m, src)
}
func (m *GetAttachInfoResp) XXX_Size() int {
	return xxx_messageInfo_GetAttachInfoResp.Size(m)
}
func (m *GetAttachInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttachInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttachInfoResp proto.InternalMessageInfo

func (m *GetAttachInfoResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetAttachInfoResp) GetPsrs() []*GetAttachInfoResp_Psr {
	if m != nil {
		return m.Psrs
	}
	return nil
}

func (m *GetAttachInfoResp) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *GetAttachInfoResp) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *GetAttachInfoResp) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *GetAttachInfoResp) GetCrtCtxShareAddr() uint32 {
	if m != nil {
		return m.CrtCtxShareAddr
	}
	return 0
}

func (m *GetAttachInfoResp) GetCrtTimeout() uint32 {
	if m != nil {
		return m.CrtTimeout
	}
	return 0
}

func (m *GetAttachInfoResp) GetNetDevClass() uint32 {
	if m != nil {
		return m.NetDevClass
	}
	return 0
}

type GetAttachInfoResp_Psr struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Uri                  string   `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAttachInfoResp_Psr) Reset()         { *m = GetAttachInfoResp_Psr{} }
func (m *GetAttachInfoResp_Psr) String() string { return proto.CompactTextString(m) }
func (*GetAttachInfoResp_Psr) ProtoMessage()    {}
func (*GetAttachInfoResp_Psr) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{11, 0}
}

func (m *GetAttachInfoResp_Psr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttachInfoResp_Psr.Unmarshal(m, b)
}
func (m *GetAttachInfoResp_Psr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttachInfoResp_Psr.Marshal(b, m, deterministic)
}
func (m *GetAttachInfoResp_Psr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttachInfoResp_Psr.Merge(m, src)
}
func (m *GetAttachInfoResp_Psr) XXX_Size() int {
	return xxx_messageInfo_GetAttachInfoResp_Psr.Size(m)
}
func (m *GetAttachInfoResp_Psr) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttachInfoResp_Psr.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttachInfoResp_Psr proto.InternalMessageInfo

func (m *GetAttachInfoResp_Psr) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *GetAttachInfoResp_Psr) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type PrepShutdownReq struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepShutdownReq) Reset()         { *m = PrepShutdownReq{} }
func (m *PrepShutdownReq) String() string { return proto.CompactTextString(m) }
func (*PrepShutdownReq) ProtoMessage()    {}
func (*PrepShutdownReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{12}
}

func (m *PrepShutdownReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepShutdownReq.Unmarshal(m, b)
}
func (m *PrepShutdownReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepShutdownReq.Marshal(b, m, deterministic)
}
func (m *PrepShutdownReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepShutdownReq.Merge(m, src)
}
func (m *PrepShutdownReq) XXX_Size() int {
	return xxx_messageInfo_PrepShutdownReq.Size(m)
}
func (m *PrepShutdownReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepShutdownReq.DiscardUnknown(m)
}

var xxx_messageInfo_PrepShutdownReq proto.InternalMessageInfo

func (m *PrepShutdownReq) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

type PingRankReq struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRankReq) Reset()         { *m = PingRankReq{} }
func (m *PingRankReq) String() string { return proto.CompactTextString(m) }
func (*PingRankReq) ProtoMessage()    {}
func (*PingRankReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{13}
}

func (m *PingRankReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRankReq.Unmarshal(m, b)
}
func (m *PingRankReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRankReq.Marshal(b, m, deterministic)
}
func (m *PingRankReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRankReq.Merge(m, src)
}
func (m *PingRankReq) XXX_Size() int {
	return xxx_messageInfo_PingRankReq.Size(m)
}
func (m *PingRankReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRankReq.DiscardUnknown(m)
}

var xxx_messageInfo_PingRankReq proto.InternalMessageInfo

func (m *PingRankReq) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

type SetRankReq struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRankReq) Reset()         { *m = SetRankReq{} }
func (m *SetRankReq) String() string { return proto.CompactTextString(m) }
func (*SetRankReq) ProtoMessage()    {}
func (*SetRankReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5747b2e02f0c537, []int{14}
}

func (m *SetRankReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRankReq.Unmarshal(m, b)
}
func (m *SetRankReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRankReq.Marshal(b, m, deterministic)
}
func (m *SetRankReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRankReq.Merge(m, src)
}
func (m *SetRankReq) XXX_Size() int {
	return xxx_messageInfo_SetRankReq.Size(m)
}
func (m *SetRankReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRankReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetRankReq proto.InternalMessageInfo

func (m *SetRankReq) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func init() {
	proto.RegisterEnum("mgmt.JoinResp_State", JoinResp_State_name, JoinResp_State_value)
	proto.RegisterType((*DaosResp)(nil), "mgmt.DaosResp")
	proto.RegisterType((*GroupUpdateReq)(nil), "mgmt.GroupUpdateReq")
	proto.RegisterType((*GroupUpdateReq_Server)(nil), "mgmt.GroupUpdateReq.Server")
	proto.RegisterType((*GroupUpdateResp)(nil), "mgmt.GroupUpdateResp")
	proto.RegisterType((*JoinReq)(nil), "mgmt.JoinReq")
	proto.RegisterType((*JoinResp)(nil), "mgmt.JoinResp")
	proto.RegisterType((*LeaderQueryReq)(nil), "mgmt.LeaderQueryReq")
	proto.RegisterType((*LeaderQueryResp)(nil), "mgmt.LeaderQueryResp")
	proto.RegisterType((*RASEvent)(nil), "mgmt.RASEvent")
	proto.RegisterType((*ClusterEventReq)(nil), "mgmt.ClusterEventReq")
	proto.RegisterType((*ClusterEventResp)(nil), "mgmt.ClusterEventResp")
	proto.RegisterType((*GetAttachInfoReq)(nil), "mgmt.GetAttachInfoReq")
	proto.RegisterType((*GetAttachInfoResp)(nil), "mgmt.GetAttachInfoResp")
	proto.RegisterType((*GetAttachInfoResp_Psr)(nil), "mgmt.GetAttachInfoResp.Psr")
	proto.RegisterType((*PrepShutdownReq)(nil), "mgmt.PrepShutdownReq")
	proto.RegisterType((*PingRankReq)(nil), "mgmt.PingRankReq")
	proto.RegisterType((*SetRankReq)(nil), "mgmt.SetRankReq")
}

func init() {
	proto.RegisterFile("svc.proto", fileDescriptor_e5747b2e02f0c537)
}

var fileDescriptor_e5747b2e02f0c537 = []byte{
	// 778 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0x8e, 0x1b, 0x45,
	0x10, 0xce, 0xf8, 0x77, 0x5c, 0xce, 0xda, 0xa6, 0x15, 0xa1, 0xd1, 0x82, 0xc0, 0x69, 0x01, 0x32,
	0x20, 0x19, 0x69, 0x11, 0x07, 0x58, 0xbc, 0x10, 0x16, 0xa1, 0xc5, 0xb4, 0x93, 0xbc, 0x46, 0x9d,
	0x99, 0xda, 0x75, 0x93, 0xf9, 0xa3, 0xbb, 0xc7, 0xd8, 0x37, 0xe1, 0x06, 0x5c, 0x82, 0x77, 0x8e,
	0xc4, 0x2b, 0xaa, 0xee, 0xf1, 0xac, 0xed, 0x90, 0x15, 0x6f, 0xf5, 0xd7, 0x5f, 0x55, 0x7d, 0x55,
	0x35, 0x03, 0x03, 0xb3, 0x89, 0xe7, 0xa5, 0x2e, 0x6c, 0xc1, 0x3a, 0xd9, 0x5d, 0x66, 0x39, 0x87,
	0xf0, 0x4a, 0x16, 0x46, 0xa0, 0x29, 0xd9, 0xfb, 0xd0, 0x33, 0x56, 0xda, 0xca, 0x44, 0xc1, 0x34,
	0x98, 0x75, 0x45, 0xad, 0xf1, 0x3f, 0x02, 0x18, 0x3d, 0xd3, 0x45, 0x55, 0xbe, 0x28, 0x13, 0x69,
	0x51, 0xe0, 0x6f, 0xec, 0x63, 0x18, 0x66, 0xb2, 0x7c, 0xb5, 0x41, 0x6d, 0x54, 0x91, 0xbb, 0xf8,
	0x33, 0x01, 0x99, 0x2c, 0x5f, 0x7a, 0x0b, 0xfb, 0x06, 0xfa, 0x06, 0x35, 0xf9, 0xa3, 0xd6, 0xb4,
	0x3d, 0x1b, 0x5e, 0x7c, 0x30, 0xa7, 0x7c, 0xf3, 0x63, 0x9c, 0xf9, 0xca, 0xc5, 0x88, 0x7d, 0xec,
	0xf9, 0x1c, 0x7a, 0xde, 0xc4, 0x18, 0x74, 0xb4, 0xcc, 0xdf, 0xd4, 0xd0, 0x4e, 0x66, 0x13, 0x68,
	0x57, 0x5a, 0x45, 0xad, 0x69, 0x30, 0x1b, 0x08, 0x12, 0xf9, 0xe7, 0x30, 0x3e, 0x42, 0x7c, 0xa0,
	0x8b, 0x3f, 0x03, 0xe8, 0xff, 0x58, 0xa8, 0x9c, 0xca, 0x67, 0xd0, 0xa9, 0x2a, 0x95, 0xb8, 0x88,
	0x81, 0x70, 0x72, 0x93, 0xb0, 0xf5, 0x76, 0xc2, 0x76, 0x93, 0x90, 0x3d, 0x81, 0x6e, 0x1e, 0xdb,
	0xad, 0x89, 0x3a, 0x2e, 0xcc, 0x2b, 0xf4, 0x56, 0x26, 0x89, 0x8e, 0xba, 0x1e, 0x8f, 0x64, 0xf6,
	0x19, 0x8c, 0x8c, 0xde, 0x7c, 0x2f, 0xab, 0xd4, 0x5e, 0x15, 0x99, 0x54, 0x79, 0xd4, 0x73, 0xde,
	0x13, 0x2b, 0xe5, 0x50, 0xc9, 0x36, 0xea, 0x3b, 0x3c, 0x12, 0xf9, 0x5f, 0x01, 0x84, 0xbe, 0xd2,
	0x77, 0xb7, 0xf3, 0x9f, 0xe5, 0x7e, 0x01, 0x5d, 0xf2, 0xa2, 0x2b, 0x78, 0x74, 0xf1, 0xc4, 0x53,
	0xbe, 0x87, 0x9a, 0xaf, 0xc8, 0x27, 0x7c, 0x08, 0x9b, 0xc2, 0xf0, 0xf6, 0xa0, 0xb6, 0x8e, 0xab,
	0xed, 0xd0, 0xc4, 0x3e, 0x84, 0x41, 0x5a, 0xc4, 0x32, 0xa5, 0xf7, 0xae, 0xb3, 0x50, 0xdc, 0x1b,
	0x78, 0x04, 0x5d, 0x87, 0xc7, 0x7a, 0xd0, 0xba, 0xbe, 0x99, 0x3c, 0x62, 0x7d, 0x68, 0xff, 0xfc,
	0xe2, 0xf9, 0x24, 0xe0, 0x33, 0x18, 0xfd, 0x84, 0x32, 0x41, 0xfd, 0x4b, 0x85, 0x7a, 0x47, 0x74,
	0x53, 0x0f, 0x3b, 0x63, 0x31, 0xab, 0x09, 0xaf, 0x35, 0xbe, 0x82, 0xf1, 0x51, 0xa4, 0x29, 0xd9,
	0x27, 0x70, 0x16, 0x57, 0x5a, 0x63, 0x6e, 0xbd, 0xa7, 0x7e, 0x71, 0x6c, 0x64, 0xe7, 0x10, 0x6a,
	0x2c, 0x53, 0x15, 0x4b, 0xbf, 0x5e, 0x03, 0xd1, 0xe8, 0xfc, 0x9f, 0x00, 0x42, 0x71, 0xb9, 0xfa,
	0x6e, 0x83, 0xb9, 0x25, 0x96, 0x72, 0x99, 0xe1, 0x7e, 0xd0, 0x24, 0x53, 0x5f, 0x56, 0x65, 0x68,
	0xac, 0xcc, 0xca, 0x7a, 0x97, 0xee, 0x0d, 0x04, 0x6d, 0x70, 0x83, 0x5a, 0xd9, 0x9d, 0xa3, 0xf1,
	0x4c, 0x34, 0x3a, 0x8d, 0x2a, 0x33, 0x77, 0x35, 0x57, 0x24, 0x36, 0x53, 0xe8, 0x1e, 0x4c, 0xe1,
	0x1c, 0xc2, 0x75, 0x61, 0xac, 0xcb, 0xeb, 0x47, 0xde, 0xe8, 0x14, 0x9f, 0x48, 0x2b, 0xdd, 0xb4,
	0x1f, 0x0b, 0x27, 0x13, 0x2a, 0xaa, 0x24, 0x0a, 0xfd, 0x02, 0xa0, 0x5f, 0x45, 0xbb, 0x2b, 0x31,
	0x1a, 0x78, 0x54, 0x92, 0xd9, 0x53, 0x78, 0xac, 0x72, 0x63, 0x65, 0x1e, 0xe3, 0x2b, 0xda, 0x17,
	0x70, 0xbe, 0xe1, 0xde, 0x76, 0x9d, 0x6c, 0xf9, 0x0d, 0x8c, 0x17, 0x69, 0x65, 0x2c, 0x6a, 0xd7,
	0x3c, 0x31, 0x3f, 0x82, 0x56, 0xbd, 0xe6, 0x67, 0xa2, 0xa5, 0x12, 0xc6, 0xa1, 0xad, 0x1d, 0x67,
	0xc1, 0x6c, 0x78, 0x31, 0xf2, 0xfb, 0xb1, 0x27, 0xeb, 0x87, 0x47, 0x82, 0x9c, 0xdf, 0xf6, 0xa1,
	0x8b, 0xa4, 0x73, 0x0e, 0x93, 0x63, 0x3c, 0x53, 0x9e, 0x02, 0xf2, 0x97, 0x30, 0x79, 0x86, 0xf6,
	0xd2, 0x5a, 0x19, 0xaf, 0xaf, 0xf3, 0xdb, 0x82, 0x92, 0x4e, 0xa0, 0x6d, 0x76, 0xa6, 0xe6, 0x9c,
	0x44, 0xa2, 0x44, 0xa6, 0xa9, 0x90, 0xf9, 0x1b, 0x9f, 0x3b, 0x14, 0x8d, 0x4e, 0x17, 0xf5, 0x6b,
	0xf1, 0x5a, 0x25, 0xf5, 0x95, 0x79, 0x85, 0xff, 0xdd, 0x82, 0xf7, 0x4e, 0x80, 0x1f, 0x38, 0x86,
	0xaf, 0xa0, 0x53, 0x9a, 0xb7, 0x3e, 0x35, 0xa7, 0xcf, 0xe7, 0x4b, 0xa3, 0x85, 0x0b, 0xa4, 0x82,
	0x96, 0xba, 0xd8, 0x28, 0xda, 0x30, 0x9f, 0xb7, 0xd1, 0x69, 0x3f, 0xae, 0x73, 0x8b, 0xfa, 0x56,
	0xc6, 0x58, 0xcf, 0xfa, 0xde, 0x40, 0x25, 0xd4, 0x27, 0xe3, 0x8f, 0xbd, 0xd6, 0xd8, 0x0c, 0xc6,
	0x0b, 0x6d, 0x17, 0x76, 0xbb, 0x5a, 0x4b, 0x8d, 0x97, 0xf4, 0x35, 0xe8, 0x39, 0x96, 0x4e, 0xcd,
	0xec, 0x23, 0x80, 0x85, 0xb6, 0xcf, 0x55, 0x86, 0x45, 0x65, 0xeb, 0xbb, 0x3f, 0xb0, 0xd0, 0x65,
	0xde, 0xa0, 0xbd, 0xc2, 0xcd, 0x22, 0x95, 0xc6, 0xd4, 0x7b, 0x71, 0x68, 0x3a, 0xff, 0x12, 0xda,
	0x4b, 0xf3, 0x7f, 0x3f, 0x91, 0x9f, 0xc2, 0x78, 0xa9, 0xb1, 0x5c, 0xad, 0x2b, 0x9b, 0x14, 0xbf,
	0xef, 0x3f, 0x7f, 0xa7, 0x0f, 0xf9, 0x53, 0x18, 0x2e, 0x55, 0x7e, 0x47, 0x33, 0x79, 0x57, 0xc8,
	0x14, 0x60, 0x85, 0xf6, 0x81, 0x88, 0xd7, 0x3d, 0xf7, 0x6b, 0xf9, 0xfa, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xbc, 0x42, 0x76, 0xab, 0x67, 0x06, 0x00, 0x00,
}
