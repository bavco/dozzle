// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.3
// source: types.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContainerAction int32

const (
	ContainerAction_Start   ContainerAction = 0
	ContainerAction_Stop    ContainerAction = 1
	ContainerAction_Restart ContainerAction = 2
)

// Enum value maps for ContainerAction.
var (
	ContainerAction_name = map[int32]string{
		0: "Start",
		1: "Stop",
		2: "Restart",
	}
	ContainerAction_value = map[string]int32{
		"Start":   0,
		"Stop":    1,
		"Restart": 2,
	}
)

func (x ContainerAction) Enum() *ContainerAction {
	p := new(ContainerAction)
	*p = x
	return p
}

func (x ContainerAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContainerAction) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[0].Descriptor()
}

func (ContainerAction) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[0]
}

func (x ContainerAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContainerAction.Descriptor instead.
func (ContainerAction) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image    string                 `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Status   string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"` // deprecated
	State    string                 `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	ImageId  string                 `protobuf:"bytes,6,opt,name=ImageId,proto3" json:"ImageId,omitempty"` // deprecated
	Created  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created,proto3" json:"created,omitempty"`
	Started  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=started,proto3" json:"started,omitempty"`
	Health   string                 `protobuf:"bytes,9,opt,name=health,proto3" json:"health,omitempty"`
	Host     string                 `protobuf:"bytes,10,opt,name=host,proto3" json:"host,omitempty"`
	Tty      bool                   `protobuf:"varint,11,opt,name=tty,proto3" json:"tty,omitempty"`
	Labels   map[string]string      `protobuf:"bytes,12,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Stats    []*ContainerStat       `protobuf:"bytes,13,rep,name=stats,proto3" json:"stats,omitempty"`
	Group    string                 `protobuf:"bytes,14,opt,name=group,proto3" json:"group,omitempty"`
	Command  string                 `protobuf:"bytes,15,opt,name=command,proto3" json:"command,omitempty"`
	Finished *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=finished,proto3" json:"finished,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	mi := &file_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *Container) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Container) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Container) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Container) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *Container) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Container) GetStarted() *timestamppb.Timestamp {
	if x != nil {
		return x.Started
	}
	return nil
}

func (x *Container) GetHealth() string {
	if x != nil {
		return x.Health
	}
	return ""
}

func (x *Container) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Container) GetTty() bool {
	if x != nil {
		return x.Tty
	}
	return false
}

func (x *Container) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Container) GetStats() []*ContainerStat {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Container) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Container) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Container) GetFinished() *timestamppb.Timestamp {
	if x != nil {
		return x.Finished
	}
	return nil
}

type ContainerStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CpuPercent    float64 `protobuf:"fixed64,2,opt,name=cpuPercent,proto3" json:"cpuPercent,omitempty"`
	MemoryUsage   float64 `protobuf:"fixed64,3,opt,name=memoryUsage,proto3" json:"memoryUsage,omitempty"`
	MemoryPercent float64 `protobuf:"fixed64,4,opt,name=memoryPercent,proto3" json:"memoryPercent,omitempty"`
}

func (x *ContainerStat) Reset() {
	*x = ContainerStat{}
	mi := &file_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerStat) ProtoMessage() {}

func (x *ContainerStat) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerStat.ProtoReflect.Descriptor instead.
func (*ContainerStat) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *ContainerStat) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContainerStat) GetCpuPercent() float64 {
	if x != nil {
		return x.CpuPercent
	}
	return 0
}

func (x *ContainerStat) GetMemoryUsage() float64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *ContainerStat) GetMemoryPercent() float64 {
	if x != nil {
		return x.MemoryPercent
	}
	return 0
}

type LogEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ContainerId string                 `protobuf:"bytes,2,opt,name=containerId,proto3" json:"containerId,omitempty"`
	Message     *anypb.Any             `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Level       string                 `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	Stream      string                 `protobuf:"bytes,6,opt,name=stream,proto3" json:"stream,omitempty"`
	Position    string                 `protobuf:"bytes,7,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *LogEvent) Reset() {
	*x = LogEvent{}
	mi := &file_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEvent) ProtoMessage() {}

func (x *LogEvent) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEvent.ProtoReflect.Descriptor instead.
func (*LogEvent) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *LogEvent) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LogEvent) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *LogEvent) GetMessage() *anypb.Any {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *LogEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogEvent) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *LogEvent) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *LogEvent) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

type SimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SimpleMessage) Reset() {
	*x = SimpleMessage{}
	mi := &file_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessage) ProtoMessage() {}

func (x *SimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessage.ProtoReflect.Descriptor instead.
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{3}
}

func (x *SimpleMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ComplexMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ComplexMessage) Reset() {
	*x = ComplexMessage{}
	mi := &file_types_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplexMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexMessage) ProtoMessage() {}

func (x *ComplexMessage) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexMessage.ProtoReflect.Descriptor instead.
func (*ComplexMessage) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{4}
}

func (x *ComplexMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ContainerEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId   string                 `protobuf:"bytes,1,opt,name=actorId,proto3" json:"actorId,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Host      string                 `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ContainerEvent) Reset() {
	*x = ContainerEvent{}
	mi := &file_types_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerEvent) ProtoMessage() {}

func (x *ContainerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerEvent.ProtoReflect.Descriptor instead.
func (*ContainerEvent) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{5}
}

func (x *ContainerEvent) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

func (x *ContainerEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerEvent) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ContainerEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NodeAddress     string            `protobuf:"bytes,3,opt,name=nodeAddress,proto3" json:"nodeAddress,omitempty"`
	Swarm           bool              `protobuf:"varint,4,opt,name=swarm,proto3" json:"swarm,omitempty"`
	Labels          map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OperatingSystem string            `protobuf:"bytes,6,opt,name=operatingSystem,proto3" json:"operatingSystem,omitempty"`
	OsVersion       string            `protobuf:"bytes,7,opt,name=osVersion,proto3" json:"osVersion,omitempty"`
	OsType          string            `protobuf:"bytes,8,opt,name=osType,proto3" json:"osType,omitempty"`
	CpuCores        uint32            `protobuf:"varint,9,opt,name=cpuCores,proto3" json:"cpuCores,omitempty"`
	Memory          uint64            `protobuf:"varint,10,opt,name=memory,proto3" json:"memory,omitempty"`
	AgentVersion    string            `protobuf:"bytes,11,opt,name=agentVersion,proto3" json:"agentVersion,omitempty"`
	DockerVersion   string            `protobuf:"bytes,12,opt,name=dockerVersion,proto3" json:"dockerVersion,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	mi := &file_types_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{6}
}

func (x *Host) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Host) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Host) GetNodeAddress() string {
	if x != nil {
		return x.NodeAddress
	}
	return ""
}

func (x *Host) GetSwarm() bool {
	if x != nil {
		return x.Swarm
	}
	return false
}

func (x *Host) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Host) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *Host) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *Host) GetOsType() string {
	if x != nil {
		return x.OsType
	}
	return ""
}

func (x *Host) GetCpuCores() uint32 {
	if x != nil {
		return x.CpuCores
	}
	return 0
}

func (x *Host) GetMemory() uint64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Host) GetAgentVersion() string {
	if x != nil {
		return x.AgentVersion
	}
	return ""
}

func (x *Host) GetDockerVersion() string {
	if x != nil {
		return x.DockerVersion
	}
	return ""
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x04, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x74, 0x79, 0x12, 0x37, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x87, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x70,
	0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x22, 0xf0, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x24, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xaf, 0x03, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x12, 0x32, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x28,
	0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x73, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x73, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x33, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x10, 0x02, 0x42, 0x13, 0x5a, 0x11,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_types_proto_goTypes = []any{
	(ContainerAction)(0),          // 0: protobuf.ContainerAction
	(*Container)(nil),             // 1: protobuf.Container
	(*ContainerStat)(nil),         // 2: protobuf.ContainerStat
	(*LogEvent)(nil),              // 3: protobuf.LogEvent
	(*SimpleMessage)(nil),         // 4: protobuf.SimpleMessage
	(*ComplexMessage)(nil),        // 5: protobuf.ComplexMessage
	(*ContainerEvent)(nil),        // 6: protobuf.ContainerEvent
	(*Host)(nil),                  // 7: protobuf.Host
	nil,                           // 8: protobuf.Container.LabelsEntry
	nil,                           // 9: protobuf.Host.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 11: google.protobuf.Any
}
var file_types_proto_depIdxs = []int32{
	10, // 0: protobuf.Container.created:type_name -> google.protobuf.Timestamp
	10, // 1: protobuf.Container.started:type_name -> google.protobuf.Timestamp
	8,  // 2: protobuf.Container.labels:type_name -> protobuf.Container.LabelsEntry
	2,  // 3: protobuf.Container.stats:type_name -> protobuf.ContainerStat
	10, // 4: protobuf.Container.finished:type_name -> google.protobuf.Timestamp
	11, // 5: protobuf.LogEvent.message:type_name -> google.protobuf.Any
	10, // 6: protobuf.LogEvent.timestamp:type_name -> google.protobuf.Timestamp
	10, // 7: protobuf.ContainerEvent.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 8: protobuf.Host.labels:type_name -> protobuf.Host.LabelsEntry
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		EnumInfos:         file_types_proto_enumTypes,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
