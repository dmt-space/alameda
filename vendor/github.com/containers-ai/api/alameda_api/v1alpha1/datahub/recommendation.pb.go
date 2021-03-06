// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/recommendation.proto

package containers_ai_alameda_v1alpha1_datahub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ControllerRecommendedType int32

const (
	ControllerRecommendedType_CRT_Undefined ControllerRecommendedType = 0
	ControllerRecommendedType_CRT_Primitive ControllerRecommendedType = 1
	ControllerRecommendedType_CRT_K8s       ControllerRecommendedType = 2
)

var ControllerRecommendedType_name = map[int32]string{
	0: "CRT_Undefined",
	1: "CRT_Primitive",
	2: "CRT_K8s",
}
var ControllerRecommendedType_value = map[string]int32{
	"CRT_Undefined": 0,
	"CRT_Primitive": 1,
	"CRT_K8s":       2,
}

func (x ControllerRecommendedType) String() string {
	return proto.EnumName(ControllerRecommendedType_name, int32(x))
}
func (ControllerRecommendedType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_33177d2ef962a49e, []int{0}
}

// *
// Represents a resource configuration recommendation
//
// It includes recommended limits and requests for the initial stage (a container which is just started) and after the initial stage
//
type ContainerRecommendation struct {
	Name                          string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitRecommendations          []*MetricData `protobuf:"bytes,2,rep,name=limit_recommendations,json=limitRecommendations,proto3" json:"limit_recommendations,omitempty"`
	RequestRecommendations        []*MetricData `protobuf:"bytes,3,rep,name=request_recommendations,json=requestRecommendations,proto3" json:"request_recommendations,omitempty"`
	InitialLimitRecommendations   []*MetricData `protobuf:"bytes,4,rep,name=initial_limit_recommendations,json=initialLimitRecommendations,proto3" json:"initial_limit_recommendations,omitempty"`
	InitialRequestRecommendations []*MetricData `protobuf:"bytes,5,rep,name=initial_request_recommendations,json=initialRequestRecommendations,proto3" json:"initial_request_recommendations,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}      `json:"-"`
	XXX_unrecognized              []byte        `json:"-"`
	XXX_sizecache                 int32         `json:"-"`
}

func (m *ContainerRecommendation) Reset()         { *m = ContainerRecommendation{} }
func (m *ContainerRecommendation) String() string { return proto.CompactTextString(m) }
func (*ContainerRecommendation) ProtoMessage()    {}
func (*ContainerRecommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_33177d2ef962a49e, []int{0}
}
func (m *ContainerRecommendation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerRecommendation.Unmarshal(m, b)
}
func (m *ContainerRecommendation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerRecommendation.Marshal(b, m, deterministic)
}
func (dst *ContainerRecommendation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerRecommendation.Merge(dst, src)
}
func (m *ContainerRecommendation) XXX_Size() int {
	return xxx_messageInfo_ContainerRecommendation.Size(m)
}
func (m *ContainerRecommendation) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerRecommendation.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerRecommendation proto.InternalMessageInfo

func (m *ContainerRecommendation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerRecommendation) GetLimitRecommendations() []*MetricData {
	if m != nil {
		return m.LimitRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetRequestRecommendations() []*MetricData {
	if m != nil {
		return m.RequestRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetInitialLimitRecommendations() []*MetricData {
	if m != nil {
		return m.InitialLimitRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetInitialRequestRecommendations() []*MetricData {
	if m != nil {
		return m.InitialRequestRecommendations
	}
	return nil
}

// *
// Represents a recommended pod-to-node assignment (i.e. pod placement)
//
type AssignPodPolicy struct {
	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Types that are valid to be assigned to Policy:
	//	*AssignPodPolicy_NodePriority
	//	*AssignPodPolicy_NodeSelector
	//	*AssignPodPolicy_NodeName
	Policy               isAssignPodPolicy_Policy `protobuf_oneof:"policy"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AssignPodPolicy) Reset()         { *m = AssignPodPolicy{} }
func (m *AssignPodPolicy) String() string { return proto.CompactTextString(m) }
func (*AssignPodPolicy) ProtoMessage()    {}
func (*AssignPodPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_33177d2ef962a49e, []int{1}
}
func (m *AssignPodPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignPodPolicy.Unmarshal(m, b)
}
func (m *AssignPodPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignPodPolicy.Marshal(b, m, deterministic)
}
func (dst *AssignPodPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignPodPolicy.Merge(dst, src)
}
func (m *AssignPodPolicy) XXX_Size() int {
	return xxx_messageInfo_AssignPodPolicy.Size(m)
}
func (m *AssignPodPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignPodPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_AssignPodPolicy proto.InternalMessageInfo

func (m *AssignPodPolicy) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type isAssignPodPolicy_Policy interface {
	isAssignPodPolicy_Policy()
}

type AssignPodPolicy_NodePriority struct {
	NodePriority *NodePriority `protobuf:"bytes,2,opt,name=node_priority,json=nodePriority,proto3,oneof"`
}

type AssignPodPolicy_NodeSelector struct {
	NodeSelector *Selector `protobuf:"bytes,3,opt,name=node_selector,json=nodeSelector,proto3,oneof"`
}

type AssignPodPolicy_NodeName struct {
	NodeName string `protobuf:"bytes,4,opt,name=node_name,json=nodeName,proto3,oneof"`
}

func (*AssignPodPolicy_NodePriority) isAssignPodPolicy_Policy() {}

func (*AssignPodPolicy_NodeSelector) isAssignPodPolicy_Policy() {}

func (*AssignPodPolicy_NodeName) isAssignPodPolicy_Policy() {}

func (m *AssignPodPolicy) GetPolicy() isAssignPodPolicy_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *AssignPodPolicy) GetNodePriority() *NodePriority {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodePriority); ok {
		return x.NodePriority
	}
	return nil
}

func (m *AssignPodPolicy) GetNodeSelector() *Selector {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodeSelector); ok {
		return x.NodeSelector
	}
	return nil
}

func (m *AssignPodPolicy) GetNodeName() string {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodeName); ok {
		return x.NodeName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AssignPodPolicy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AssignPodPolicy_OneofMarshaler, _AssignPodPolicy_OneofUnmarshaler, _AssignPodPolicy_OneofSizer, []interface{}{
		(*AssignPodPolicy_NodePriority)(nil),
		(*AssignPodPolicy_NodeSelector)(nil),
		(*AssignPodPolicy_NodeName)(nil),
	}
}

func _AssignPodPolicy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AssignPodPolicy)
	// policy
	switch x := m.Policy.(type) {
	case *AssignPodPolicy_NodePriority:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodePriority); err != nil {
			return err
		}
	case *AssignPodPolicy_NodeSelector:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodeSelector); err != nil {
			return err
		}
	case *AssignPodPolicy_NodeName:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NodeName)
	case nil:
	default:
		return fmt.Errorf("AssignPodPolicy.Policy has unexpected type %T", x)
	}
	return nil
}

func _AssignPodPolicy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AssignPodPolicy)
	switch tag {
	case 2: // policy.node_priority
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NodePriority)
		err := b.DecodeMessage(msg)
		m.Policy = &AssignPodPolicy_NodePriority{msg}
		return true, err
	case 3: // policy.node_selector
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Selector)
		err := b.DecodeMessage(msg)
		m.Policy = &AssignPodPolicy_NodeSelector{msg}
		return true, err
	case 4: // policy.node_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Policy = &AssignPodPolicy_NodeName{x}
		return true, err
	default:
		return false, nil
	}
}

func _AssignPodPolicy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AssignPodPolicy)
	// policy
	switch x := m.Policy.(type) {
	case *AssignPodPolicy_NodePriority:
		s := proto.Size(x.NodePriority)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AssignPodPolicy_NodeSelector:
		s := proto.Size(x.NodeSelector)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AssignPodPolicy_NodeName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.NodeName)))
		n += len(x.NodeName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// *
// Represents a set of container resource configuration recommendations of a pod
//
type PodRecommendation struct {
	NamespacedName           *NamespacedName            `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ApplyRecommendationNow   bool                       `protobuf:"varint,2,opt,name=apply_recommendation_now,json=applyRecommendationNow,proto3" json:"apply_recommendation_now,omitempty"`
	AssignPodPolicy          *AssignPodPolicy           `protobuf:"bytes,3,opt,name=assign_pod_policy,json=assignPodPolicy,proto3" json:"assign_pod_policy,omitempty"`
	ContainerRecommendations []*ContainerRecommendation `protobuf:"bytes,4,rep,name=container_recommendations,json=containerRecommendations,proto3" json:"container_recommendations,omitempty"`
	StartTime                *timestamp.Timestamp       `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime                  *timestamp.Timestamp       `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	TopController            *TopController             `protobuf:"bytes,7,opt,name=top_controller,json=topController,proto3" json:"top_controller,omitempty"`
	RecommendationId         string                     `protobuf:"bytes,8,opt,name=recommendation_id,json=recommendationId,proto3" json:"recommendation_id,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *PodRecommendation) Reset()         { *m = PodRecommendation{} }
func (m *PodRecommendation) String() string { return proto.CompactTextString(m) }
func (*PodRecommendation) ProtoMessage()    {}
func (*PodRecommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_33177d2ef962a49e, []int{2}
}
func (m *PodRecommendation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodRecommendation.Unmarshal(m, b)
}
func (m *PodRecommendation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodRecommendation.Marshal(b, m, deterministic)
}
func (dst *PodRecommendation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodRecommendation.Merge(dst, src)
}
func (m *PodRecommendation) XXX_Size() int {
	return xxx_messageInfo_PodRecommendation.Size(m)
}
func (m *PodRecommendation) XXX_DiscardUnknown() {
	xxx_messageInfo_PodRecommendation.DiscardUnknown(m)
}

var xxx_messageInfo_PodRecommendation proto.InternalMessageInfo

func (m *PodRecommendation) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *PodRecommendation) GetApplyRecommendationNow() bool {
	if m != nil {
		return m.ApplyRecommendationNow
	}
	return false
}

func (m *PodRecommendation) GetAssignPodPolicy() *AssignPodPolicy {
	if m != nil {
		return m.AssignPodPolicy
	}
	return nil
}

func (m *PodRecommendation) GetContainerRecommendations() []*ContainerRecommendation {
	if m != nil {
		return m.ContainerRecommendations
	}
	return nil
}

func (m *PodRecommendation) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *PodRecommendation) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *PodRecommendation) GetTopController() *TopController {
	if m != nil {
		return m.TopController
	}
	return nil
}

func (m *PodRecommendation) GetRecommendationId() string {
	if m != nil {
		return m.RecommendationId
	}
	return ""
}

type ControllerRecommendation struct {
	RecommendedType      ControllerRecommendedType  `protobuf:"varint,1,opt,name=recommended_type,json=recommendedType,proto3,enum=containers_ai.alameda.v1alpha1.datahub.ControllerRecommendedType" json:"recommended_type,omitempty"`
	RecommendedSpec      *ControllerRecommendedSpec `protobuf:"bytes,2,opt,name=recommended_spec,json=recommendedSpec,proto3" json:"recommended_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ControllerRecommendation) Reset()         { *m = ControllerRecommendation{} }
func (m *ControllerRecommendation) String() string { return proto.CompactTextString(m) }
func (*ControllerRecommendation) ProtoMessage()    {}
func (*ControllerRecommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_33177d2ef962a49e, []int{3}
}
func (m *ControllerRecommendation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerRecommendation.Unmarshal(m, b)
}
func (m *ControllerRecommendation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerRecommendation.Marshal(b, m, deterministic)
}
func (dst *ControllerRecommendation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerRecommendation.Merge(dst, src)
}
func (m *ControllerRecommendation) XXX_Size() int {
	return xxx_messageInfo_ControllerRecommendation.Size(m)
}
func (m *ControllerRecommendation) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerRecommendation.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerRecommendation proto.InternalMessageInfo

func (m *ControllerRecommendation) GetRecommendedType() ControllerRecommendedType {
	if m != nil {
		return m.RecommendedType
	}
	return ControllerRecommendedType_CRT_Undefined
}

func (m *ControllerRecommendation) GetRecommendedSpec() *ControllerRecommendedSpec {
	if m != nil {
		return m.RecommendedSpec
	}
	return nil
}

type ControllerRecommendedSpec struct {
	NamespacedName       *NamespacedName      `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	CurrentReplicas      int32                `protobuf:"varint,2,opt,name=current_replicas,json=currentReplicas,proto3" json:"current_replicas,omitempty"`
	DesiredReplicas      int32                `protobuf:"varint,3,opt,name=desired_replicas,json=desiredReplicas,proto3" json:"desired_replicas,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ControllerRecommendedSpec) Reset()         { *m = ControllerRecommendedSpec{} }
func (m *ControllerRecommendedSpec) String() string { return proto.CompactTextString(m) }
func (*ControllerRecommendedSpec) ProtoMessage()    {}
func (*ControllerRecommendedSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_33177d2ef962a49e, []int{4}
}
func (m *ControllerRecommendedSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerRecommendedSpec.Unmarshal(m, b)
}
func (m *ControllerRecommendedSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerRecommendedSpec.Marshal(b, m, deterministic)
}
func (dst *ControllerRecommendedSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerRecommendedSpec.Merge(dst, src)
}
func (m *ControllerRecommendedSpec) XXX_Size() int {
	return xxx_messageInfo_ControllerRecommendedSpec.Size(m)
}
func (m *ControllerRecommendedSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerRecommendedSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerRecommendedSpec proto.InternalMessageInfo

func (m *ControllerRecommendedSpec) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *ControllerRecommendedSpec) GetCurrentReplicas() int32 {
	if m != nil {
		return m.CurrentReplicas
	}
	return 0
}

func (m *ControllerRecommendedSpec) GetDesiredReplicas() int32 {
	if m != nil {
		return m.DesiredReplicas
	}
	return 0
}

func (m *ControllerRecommendedSpec) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerRecommendation)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation")
	proto.RegisterType((*AssignPodPolicy)(nil), "containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy")
	proto.RegisterType((*PodRecommendation)(nil), "containers_ai.alameda.v1alpha1.datahub.PodRecommendation")
	proto.RegisterType((*ControllerRecommendation)(nil), "containers_ai.alameda.v1alpha1.datahub.ControllerRecommendation")
	proto.RegisterType((*ControllerRecommendedSpec)(nil), "containers_ai.alameda.v1alpha1.datahub.ControllerRecommendedSpec")
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.ControllerRecommendedType", ControllerRecommendedType_name, ControllerRecommendedType_value)
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/recommendation.proto", fileDescriptor_recommendation_33177d2ef962a49e)
}

var fileDescriptor_recommendation_33177d2ef962a49e = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5f, 0x6f, 0xfb, 0x34,
	0x14, 0x5d, 0xff, 0x6d, 0x9d, 0xcb, 0xd6, 0x36, 0x82, 0x2d, 0x1b, 0x9a, 0x36, 0xf5, 0x01, 0x6d,
	0x4c, 0x4a, 0x59, 0x61, 0x30, 0xc4, 0x03, 0x1a, 0xe3, 0x61, 0x08, 0xa8, 0x2a, 0xaf, 0x88, 0x07,
	0x90, 0x22, 0x2f, 0xbe, 0xeb, 0x2c, 0x12, 0xdb, 0xd8, 0xee, 0xa6, 0x22, 0x3e, 0x03, 0x7c, 0x05,
	0x3e, 0x25, 0xbc, 0xa2, 0x38, 0x49, 0xb7, 0x84, 0x76, 0x0d, 0xfb, 0xe9, 0xf7, 0x16, 0x5f, 0xdf,
	0x73, 0xce, 0xbd, 0xd7, 0xc7, 0x0e, 0x3a, 0x23, 0x21, 0x89, 0x80, 0x12, 0x9f, 0x48, 0xd6, 0x7f,
	0x38, 0x23, 0xa1, 0xbc, 0x27, 0x67, 0x7d, 0x4a, 0x0c, 0xb9, 0x9f, 0xde, 0xf6, 0x15, 0x04, 0x22,
	0x8a, 0x80, 0x53, 0x62, 0x98, 0xe0, 0x9e, 0x54, 0xc2, 0x08, 0xe7, 0x83, 0x40, 0x70, 0x43, 0x18,
	0x07, 0xa5, 0x7d, 0xc2, 0xbc, 0x94, 0xc0, 0xcb, 0xc0, 0x5e, 0x0a, 0xde, 0x3f, 0x9c, 0x08, 0x31,
	0x09, 0xa1, 0x6f, 0x51, 0xb7, 0xd3, 0xbb, 0xbe, 0x61, 0x11, 0x68, 0x43, 0x22, 0x99, 0x10, 0xed,
	0x9f, 0xbe, 0xa8, 0x1d, 0x81, 0x21, 0xf1, 0x77, 0x9a, 0xfc, 0x72, 0xa1, 0x52, 0x50, 0x9f, 0x68,
	0xcd, 0x26, 0x3c, 0x02, 0x6e, 0x52, 0xc8, 0xc9, 0x2a, 0x7e, 0xc5, 0x82, 0x52, 0xa5, 0x28, 0xd0,
	0x62, 0xaa, 0x02, 0x48, 0x92, 0x7b, 0xff, 0xd4, 0xd0, 0xee, 0x55, 0x36, 0x03, 0x9c, 0x1b, 0x91,
	0xe3, 0xa0, 0x3a, 0x27, 0x11, 0xb8, 0x95, 0xa3, 0xca, 0xf1, 0x26, 0xb6, 0xdf, 0xce, 0x04, 0xbd,
	0x17, 0xb2, 0x88, 0x19, 0x3f, 0x3f, 0x4e, 0xed, 0x56, 0x8f, 0x6a, 0xc7, 0xad, 0xc1, 0xc0, 0x2b,
	0x37, 0x50, 0xef, 0x7b, 0x5b, 0xf1, 0xd7, 0xc4, 0x10, 0xfc, 0xae, 0x25, 0xcc, 0x6b, 0x6b, 0xe7,
	0x17, 0xb4, 0xab, 0xe0, 0xd7, 0x29, 0xe8, 0xff, 0x4a, 0xd5, 0x5e, 0x2d, 0xb5, 0x93, 0x52, 0x16,
	0xc5, 0x1e, 0xd0, 0x01, 0xe3, 0xcc, 0x30, 0x12, 0xfa, 0x8b, 0xbb, 0xab, 0xbf, 0x5a, 0xf2, 0xfd,
	0x94, 0xf8, 0xbb, 0x45, 0x4d, 0xfe, 0x86, 0x0e, 0x33, 0xdd, 0x65, 0xcd, 0x36, 0x5e, 0xad, 0x9c,
	0xb5, 0x84, 0x17, 0xf6, 0xdc, 0xfb, 0xab, 0x8a, 0xda, 0x97, 0xd6, 0x66, 0x23, 0x41, 0x47, 0x22,
	0x64, 0xc1, 0xcc, 0xf1, 0x50, 0x3d, 0x36, 0xb6, 0x3d, 0xf1, 0xd6, 0x60, 0xdf, 0x4b, 0x5c, 0xef,
	0x65, 0xae, 0xf7, 0xc6, 0x99, 0xeb, 0xb1, 0xcd, 0x73, 0x7e, 0x42, 0x5b, 0x5c, 0x50, 0xf0, 0xa5,
	0x62, 0x42, 0x31, 0x33, 0x73, 0xab, 0x16, 0xf8, 0x49, 0xd9, 0x6a, 0x87, 0x82, 0xc2, 0x28, 0xc5,
	0x5e, 0xaf, 0xe1, 0x77, 0xf8, 0xb3, 0xb5, 0xf3, 0x63, 0x4a, 0xae, 0x21, 0x84, 0xc0, 0x08, 0xe5,
	0xd6, 0x2c, 0xf9, 0x47, 0x65, 0xc9, 0x6f, 0x52, 0x5c, 0x46, 0x9c, 0xad, 0x9d, 0x03, 0xb4, 0x69,
	0x89, 0xad, 0xb9, 0xeb, 0xb1, 0xb9, 0xaf, 0xd7, 0x70, 0x33, 0x0e, 0x0d, 0x49, 0x04, 0x5f, 0x35,
	0xd1, 0xba, 0xb4, 0xe3, 0xe8, 0xfd, 0xd1, 0x40, 0xdd, 0x91, 0xa0, 0x85, 0x6b, 0xe1, 0xa3, 0x76,
	0x8c, 0xd4, 0x92, 0x04, 0x40, 0xfd, 0xf9, 0x0d, 0x69, 0x0d, 0x3e, 0x2d, 0xdd, 0xf6, 0x1c, 0x1e,
	0x7f, 0xe1, 0x6d, 0x9e, 0x5b, 0x3b, 0x17, 0xc8, 0x25, 0x52, 0x86, 0xb3, 0x82, 0x17, 0x7c, 0x2e,
	0x1e, 0xed, 0x80, 0x9b, 0x78, 0xc7, 0xee, 0xe7, 0xeb, 0x1a, 0x8a, 0x47, 0x27, 0x40, 0xdd, 0xe4,
	0xe5, 0xf0, 0xe3, 0x47, 0x24, 0xe9, 0x22, 0x1d, 0xdb, 0x67, 0x65, 0x8b, 0x2b, 0x78, 0x02, 0xb7,
	0x49, 0xc1, 0x24, 0xbf, 0xa3, 0xbd, 0x39, 0xd5, 0x92, 0x8b, 0xf2, 0x65, 0x59, 0xb1, 0x25, 0x4f,
	0x0f, 0x76, 0x83, 0xc5, 0x1b, 0xda, 0xf9, 0x1c, 0x21, 0x6d, 0x88, 0x32, 0xbe, 0x35, 0x6a, 0x63,
	0xa5, 0x51, 0x37, 0x6d, 0x76, 0xbc, 0x76, 0xce, 0x51, 0x13, 0x38, 0x4d, 0x80, 0xeb, 0x2b, 0x81,
	0x1b, 0xc0, 0xa9, 0x85, 0xfd, 0x8c, 0xb6, 0x8d, 0x90, 0x7e, 0x5c, 0x91, 0x12, 0x61, 0x08, 0xca,
	0xdd, 0xb0, 0xe0, 0xf3, 0xb2, 0x4d, 0x8e, 0x85, 0xbc, 0x9a, 0x83, 0xf1, 0x96, 0x79, 0xbe, 0x74,
	0x4e, 0x51, 0xb7, 0x70, 0xcc, 0x8c, 0xba, 0x4d, 0xfb, 0xe2, 0x76, 0xf2, 0x1b, 0xdf, 0xd0, 0xde,
	0xdf, 0x15, 0xe4, 0x3e, 0xa3, 0xca, 0xfb, 0x32, 0x44, 0x4f, 0x00, 0xa0, 0xbe, 0x99, 0xc9, 0xc4,
	0x98, 0xdb, 0x83, 0xcb, 0xff, 0x73, 0x1c, 0x05, 0x6e, 0xa0, 0xe3, 0x99, 0x04, 0xdc, 0x56, 0xf9,
	0x40, 0x51, 0x4d, 0x4b, 0x08, 0xd2, 0xdb, 0xff, 0x66, 0x6a, 0x37, 0x12, 0x82, 0x9c, 0x5a, 0x1c,
	0xe8, 0xfd, 0x59, 0x45, 0x7b, 0x4b, 0xd3, 0xdf, 0xfe, 0x8d, 0x3c, 0x41, 0x9d, 0x60, 0xaa, 0x14,
	0xf0, 0xf8, 0x7d, 0x96, 0x21, 0x0b, 0x88, 0xb6, 0xcd, 0x36, 0x70, 0x3b, 0x8d, 0xe3, 0x34, 0x1c,
	0xa7, 0x52, 0xd0, 0x4c, 0x01, 0x7d, 0x4a, 0xad, 0x25, 0xa9, 0x69, 0x7c, 0x9e, 0xfa, 0x05, 0x6a,
	0x05, 0x0a, 0x88, 0x81, 0xc4, 0x92, 0xf5, 0x95, 0x96, 0x44, 0x49, 0x7a, 0x1c, 0xf8, 0x70, 0xb8,
	0x64, 0x20, 0xf6, 0x70, 0xba, 0x68, 0xeb, 0x0a, 0x8f, 0xfd, 0x1f, 0x38, 0x85, 0x3b, 0xc6, 0x81,
	0x76, 0xd6, 0xb2, 0xd0, 0x48, 0xc5, 0xff, 0x21, 0xf6, 0x00, 0x9d, 0x8a, 0xd3, 0x42, 0x1b, 0x71,
	0xe8, 0xdb, 0x0b, 0xdd, 0xa9, 0xde, 0xae, 0x5b, 0xbd, 0x8f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xa8, 0x6c, 0xea, 0x29, 0x45, 0x09, 0x00, 0x00,
}
