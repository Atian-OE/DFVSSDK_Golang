// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfvssdk/model/message.proto

package model

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

type MsgID int32

const (
	MsgID_ConnectID          MsgID = 0
	MsgID_DisconnectID       MsgID = 1
	MsgID_AlarmNotifyID      MsgID = 2
	MsgID_FiberStateNotifyID MsgID = 3
	MsgID_HeartBeatID        MsgID = 250
)

var MsgID_name = map[int32]string{
	0:   "ConnectID",
	1:   "DisconnectID",
	2:   "AlarmNotifyID",
	3:   "FiberStateNotifyID",
	250: "HeartBeatID",
}

var MsgID_value = map[string]int32{
	"ConnectID":          0,
	"DisconnectID":       1,
	"AlarmNotifyID":      2,
	"FiberStateNotifyID": 3,
	"HeartBeatID":        250,
}

func (x MsgID) String() string {
	return proto.EnumName(MsgID_name, int32(x))
}

func (MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_03a67b221dbaa966, []int{0}
}

//防区状态
type FiberState int32

const (
	FiberState_None    FiberState = 0
	FiberState_NoFiber FiberState = 1
	FiberState_Break   FiberState = 2
	FiberState_TooLong FiberState = 3
)

var FiberState_name = map[int32]string{
	0: "None",
	1: "NoFiber",
	2: "Break",
	3: "TooLong",
}

var FiberState_value = map[string]int32{
	"None":    0,
	"NoFiber": 1,
	"Break":   2,
	"TooLong": 3,
}

func (x FiberState) String() string {
	return proto.EnumName(FiberState_name, int32(x))
}

func (FiberState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_03a67b221dbaa966, []int{1}
}

//设备
type FiberStateNotify struct {
	ChannelID            int32      `protobuf:"varint,1,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	DeviceID             string     `protobuf:"bytes,2,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Timestamp            int64      `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	State                FiberState `protobuf:"varint,4,opt,name=State,proto3,enum=model.FiberState" json:"State,omitempty"`
	ChannelLength        float32    `protobuf:"fixed32,5,opt,name=ChannelLength,proto3" json:"ChannelLength,omitempty"`
	ShowLength           float32    `protobuf:"fixed32,6,opt,name=ShowLength,proto3" json:"ShowLength,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FiberStateNotify) Reset()         { *m = FiberStateNotify{} }
func (m *FiberStateNotify) String() string { return proto.CompactTextString(m) }
func (*FiberStateNotify) ProtoMessage()    {}
func (*FiberStateNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a67b221dbaa966, []int{0}
}

func (m *FiberStateNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiberStateNotify.Unmarshal(m, b)
}
func (m *FiberStateNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiberStateNotify.Marshal(b, m, deterministic)
}
func (m *FiberStateNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiberStateNotify.Merge(m, src)
}
func (m *FiberStateNotify) XXX_Size() int {
	return xxx_messageInfo_FiberStateNotify.Size(m)
}
func (m *FiberStateNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_FiberStateNotify.DiscardUnknown(m)
}

var xxx_messageInfo_FiberStateNotify proto.InternalMessageInfo

func (m *FiberStateNotify) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *FiberStateNotify) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *FiberStateNotify) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *FiberStateNotify) GetState() FiberState {
	if m != nil {
		return m.State
	}
	return FiberState_None
}

func (m *FiberStateNotify) GetChannelLength() float32 {
	if m != nil {
		return m.ChannelLength
	}
	return 0
}

func (m *FiberStateNotify) GetShowLength() float32 {
	if m != nil {
		return m.ShowLength
	}
	return 0
}

//警报事件
type AlarmEvent struct {
	Id                          int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MaxIntensity                int32    `protobuf:"varint,3,opt,name=max_intensity,json=maxIntensity,proto3" json:"max_intensity,omitempty"`
	MaxIntensityMatch           float32  `protobuf:"fixed32,4,opt,name=max_intensity_match,json=maxIntensityMatch,proto3" json:"max_intensity_match,omitempty"`
	Area                        int32    `protobuf:"varint,5,opt,name=area,proto3" json:"area,omitempty"`
	AreaMatch                   float32  `protobuf:"fixed32,6,opt,name=area_match,json=areaMatch,proto3" json:"area_match,omitempty"`
	EventWidth                  float32  `protobuf:"fixed32,7,opt,name=event_width,json=eventWidth,proto3" json:"event_width,omitempty"`
	EventWidthMatch             float32  `protobuf:"fixed32,8,opt,name=event_width_match,json=eventWidthMatch,proto3" json:"event_width_match,omitempty"`
	KeepTimeDutyRatio           float32  `protobuf:"fixed32,9,opt,name=keep_time_duty_ratio,json=keepTimeDutyRatio,proto3" json:"keep_time_duty_ratio,omitempty"`
	KeepTimeDutyRatioMatch      float32  `protobuf:"fixed32,10,opt,name=keep_time_duty_ratio_match,json=keepTimeDutyRatioMatch,proto3" json:"keep_time_duty_ratio_match,omitempty"`
	StdDeviationX               float32  `protobuf:"fixed32,11,opt,name=std_deviation_x,json=stdDeviationX,proto3" json:"std_deviation_x,omitempty"`
	StdDeviationXMatch          float32  `protobuf:"fixed32,12,opt,name=std_deviation_x_match,json=stdDeviationXMatch,proto3" json:"std_deviation_x_match,omitempty"`
	IntensityGenerateRatio      float32  `protobuf:"fixed32,13,opt,name=intensity_generate_ratio,json=intensityGenerateRatio,proto3" json:"intensity_generate_ratio,omitempty"`
	IntensityGenerateRatioMatch float32  `protobuf:"fixed32,14,opt,name=intensity_generate_ratio_match,json=intensityGenerateRatioMatch,proto3" json:"intensity_generate_ratio_match,omitempty"`
	EventStartTime              int64    `protobuf:"varint,15,opt,name=event_start_time,json=eventStartTime,proto3" json:"event_start_time,omitempty"`
	EventKeepTime               int64    `protobuf:"varint,16,opt,name=event_keep_time,json=eventKeepTime,proto3" json:"event_keep_time,omitempty"`
	EventCentre                 float32  `protobuf:"fixed32,17,opt,name=event_centre,json=eventCentre,proto3" json:"event_centre,omitempty"`
	MatchLevel                  int32    `protobuf:"varint,18,opt,name=match_level,json=matchLevel,proto3" json:"match_level,omitempty"`
	MatchRatio                  float32  `protobuf:"fixed32,19,opt,name=match_ratio,json=matchRatio,proto3" json:"match_ratio,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *AlarmEvent) Reset()         { *m = AlarmEvent{} }
func (m *AlarmEvent) String() string { return proto.CompactTextString(m) }
func (*AlarmEvent) ProtoMessage()    {}
func (*AlarmEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a67b221dbaa966, []int{1}
}

func (m *AlarmEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmEvent.Unmarshal(m, b)
}
func (m *AlarmEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmEvent.Marshal(b, m, deterministic)
}
func (m *AlarmEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmEvent.Merge(m, src)
}
func (m *AlarmEvent) XXX_Size() int {
	return xxx_messageInfo_AlarmEvent.Size(m)
}
func (m *AlarmEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmEvent proto.InternalMessageInfo

func (m *AlarmEvent) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AlarmEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlarmEvent) GetMaxIntensity() int32 {
	if m != nil {
		return m.MaxIntensity
	}
	return 0
}

func (m *AlarmEvent) GetMaxIntensityMatch() float32 {
	if m != nil {
		return m.MaxIntensityMatch
	}
	return 0
}

func (m *AlarmEvent) GetArea() int32 {
	if m != nil {
		return m.Area
	}
	return 0
}

func (m *AlarmEvent) GetAreaMatch() float32 {
	if m != nil {
		return m.AreaMatch
	}
	return 0
}

func (m *AlarmEvent) GetEventWidth() float32 {
	if m != nil {
		return m.EventWidth
	}
	return 0
}

func (m *AlarmEvent) GetEventWidthMatch() float32 {
	if m != nil {
		return m.EventWidthMatch
	}
	return 0
}

func (m *AlarmEvent) GetKeepTimeDutyRatio() float32 {
	if m != nil {
		return m.KeepTimeDutyRatio
	}
	return 0
}

func (m *AlarmEvent) GetKeepTimeDutyRatioMatch() float32 {
	if m != nil {
		return m.KeepTimeDutyRatioMatch
	}
	return 0
}

func (m *AlarmEvent) GetStdDeviationX() float32 {
	if m != nil {
		return m.StdDeviationX
	}
	return 0
}

func (m *AlarmEvent) GetStdDeviationXMatch() float32 {
	if m != nil {
		return m.StdDeviationXMatch
	}
	return 0
}

func (m *AlarmEvent) GetIntensityGenerateRatio() float32 {
	if m != nil {
		return m.IntensityGenerateRatio
	}
	return 0
}

func (m *AlarmEvent) GetIntensityGenerateRatioMatch() float32 {
	if m != nil {
		return m.IntensityGenerateRatioMatch
	}
	return 0
}

func (m *AlarmEvent) GetEventStartTime() int64 {
	if m != nil {
		return m.EventStartTime
	}
	return 0
}

func (m *AlarmEvent) GetEventKeepTime() int64 {
	if m != nil {
		return m.EventKeepTime
	}
	return 0
}

func (m *AlarmEvent) GetEventCentre() float32 {
	if m != nil {
		return m.EventCentre
	}
	return 0
}

func (m *AlarmEvent) GetMatchLevel() int32 {
	if m != nil {
		return m.MatchLevel
	}
	return 0
}

func (m *AlarmEvent) GetMatchRatio() float32 {
	if m != nil {
		return m.MatchRatio
	}
	return 0
}

//警报事件通知
type AlarmEventNotify struct {
	ChannelID            int32         `protobuf:"varint,1,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	DeviceID             string        `protobuf:"bytes,2,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Timestamp            int64         `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	ShowLength           float32       `protobuf:"fixed32,4,opt,name=ShowLength,proto3" json:"ShowLength,omitempty"`
	Alarms               []*AlarmEvent `protobuf:"bytes,5,rep,name=Alarms,proto3" json:"Alarms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AlarmEventNotify) Reset()         { *m = AlarmEventNotify{} }
func (m *AlarmEventNotify) String() string { return proto.CompactTextString(m) }
func (*AlarmEventNotify) ProtoMessage()    {}
func (*AlarmEventNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a67b221dbaa966, []int{2}
}

func (m *AlarmEventNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmEventNotify.Unmarshal(m, b)
}
func (m *AlarmEventNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmEventNotify.Marshal(b, m, deterministic)
}
func (m *AlarmEventNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmEventNotify.Merge(m, src)
}
func (m *AlarmEventNotify) XXX_Size() int {
	return xxx_messageInfo_AlarmEventNotify.Size(m)
}
func (m *AlarmEventNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmEventNotify.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmEventNotify proto.InternalMessageInfo

func (m *AlarmEventNotify) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *AlarmEventNotify) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *AlarmEventNotify) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AlarmEventNotify) GetShowLength() float32 {
	if m != nil {
		return m.ShowLength
	}
	return 0
}

func (m *AlarmEventNotify) GetAlarms() []*AlarmEvent {
	if m != nil {
		return m.Alarms
	}
	return nil
}

//心跳
type HeartBeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeat) Reset()         { *m = HeartBeat{} }
func (m *HeartBeat) String() string { return proto.CompactTextString(m) }
func (*HeartBeat) ProtoMessage()    {}
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a67b221dbaa966, []int{3}
}

func (m *HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeat.Unmarshal(m, b)
}
func (m *HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeat.Marshal(b, m, deterministic)
}
func (m *HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeat.Merge(m, src)
}
func (m *HeartBeat) XXX_Size() int {
	return xxx_messageInfo_HeartBeat.Size(m)
}
func (m *HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeat proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("model.MsgID", MsgID_name, MsgID_value)
	proto.RegisterEnum("model.FiberState", FiberState_name, FiberState_value)
	proto.RegisterType((*FiberStateNotify)(nil), "model.FiberStateNotify")
	proto.RegisterType((*AlarmEvent)(nil), "model.AlarmEvent")
	proto.RegisterType((*AlarmEventNotify)(nil), "model.AlarmEventNotify")
	proto.RegisterType((*HeartBeat)(nil), "model.HeartBeat")
}

func init() { proto.RegisterFile("dfvssdk/model/message.proto", fileDescriptor_03a67b221dbaa966) }

var fileDescriptor_03a67b221dbaa966 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4d, 0x53, 0xd3, 0x5c,
	0x14, 0xc7, 0x49, 0xdb, 0x14, 0x7a, 0xfa, 0x42, 0x7a, 0x78, 0x1e, 0x26, 0x03, 0x8a, 0xb5, 0x3a,
	0x1a, 0x59, 0x94, 0x51, 0x37, 0x8e, 0xae, 0xa4, 0xf1, 0x25, 0x23, 0xb0, 0x08, 0xcc, 0xe8, 0x2e,
	0x73, 0x69, 0x0e, 0x6d, 0x86, 0x26, 0x61, 0x92, 0x4b, 0xa1, 0x6b, 0x3f, 0x93, 0x1f, 0xc6, 0xaf,
	0xe2, 0xca, 0xb9, 0xe7, 0x86, 0x86, 0x22, 0x2e, 0x5d, 0x35, 0xfd, 0x9d, 0xff, 0xb9, 0xe7, 0x7f,
	0xce, 0x7d, 0x81, 0xed, 0xf0, 0x6c, 0x96, 0xe7, 0xe1, 0xf9, 0x5e, 0x9c, 0x86, 0x34, 0xdd, 0x8b,
	0x29, 0xcf, 0xc5, 0x98, 0x06, 0x17, 0x59, 0x2a, 0x53, 0x34, 0x19, 0xf6, 0x7f, 0x1a, 0x60, 0x7d,
	0x8c, 0x4e, 0x29, 0x3b, 0x96, 0x42, 0xd2, 0x51, 0x2a, 0xa3, 0xb3, 0x39, 0x3e, 0x80, 0xc6, 0x70,
	0x22, 0x92, 0x84, 0xa6, 0x9e, 0x6b, 0x1b, 0x3d, 0xc3, 0x31, 0xfd, 0x12, 0xe0, 0x16, 0xac, 0xb9,
	0x34, 0x8b, 0x46, 0xe4, 0xb9, 0x76, 0xa5, 0x67, 0x38, 0x0d, 0x7f, 0xf1, 0x5f, 0x65, 0x9e, 0x44,
	0x31, 0xe5, 0x52, 0xc4, 0x17, 0x76, 0xb5, 0x67, 0x38, 0x55, 0xbf, 0x04, 0xf8, 0x1c, 0x4c, 0x2e,
	0x63, 0xd7, 0x7a, 0x86, 0xd3, 0x79, 0xd5, 0x1d, 0xb0, 0x87, 0x41, 0x59, 0xdf, 0xd7, 0x71, 0x7c,
	0x0a, 0xed, 0xa2, 0xde, 0x01, 0x25, 0x63, 0x39, 0xb1, 0xcd, 0x9e, 0xe1, 0x54, 0xfc, 0x65, 0x88,
	0x3b, 0x00, 0xc7, 0x93, 0xf4, 0xaa, 0x90, 0xd4, 0x59, 0x72, 0x8b, 0xf4, 0xbf, 0xd7, 0x01, 0xde,
	0x4f, 0x45, 0x16, 0x7f, 0x98, 0x51, 0x22, 0xb1, 0x03, 0x95, 0x28, 0x2c, 0xda, 0xa9, 0x44, 0x21,
	0x22, 0xd4, 0x12, 0x11, 0x53, 0xd1, 0x03, 0x7f, 0xe3, 0x13, 0x68, 0xc7, 0xe2, 0x3a, 0x88, 0x12,
	0x49, 0x49, 0x1e, 0xc9, 0x39, 0xf7, 0x60, 0xfa, 0xad, 0x58, 0x5c, 0x7b, 0x37, 0x0c, 0x07, 0xb0,
	0xb1, 0x24, 0x0a, 0x62, 0x21, 0x47, 0x13, 0x6e, 0xaa, 0xe2, 0x77, 0x6f, 0x4b, 0x0f, 0x55, 0x40,
	0x15, 0x12, 0x19, 0x09, 0x6e, 0xc2, 0xf4, 0xf9, 0x1b, 0x1f, 0x02, 0xa8, 0xdf, 0x22, 0x55, 0x7b,
	0x6f, 0x28, 0xa2, 0x53, 0x1e, 0x41, 0x93, 0x94, 0xe9, 0xe0, 0x2a, 0x0a, 0xe5, 0xc4, 0x5e, 0xd5,
	0xbd, 0x31, 0xfa, 0xaa, 0x08, 0xee, 0x42, 0xf7, 0x96, 0xa0, 0x58, 0x66, 0x8d, 0x65, 0xeb, 0xa5,
	0x4c, 0x2f, 0xb6, 0x07, 0xff, 0x9d, 0x13, 0x5d, 0x04, 0x32, 0x8a, 0x29, 0x08, 0x2f, 0xe5, 0x3c,
	0xc8, 0x84, 0x8c, 0x52, 0xbb, 0xa1, 0x0d, 0xab, 0x98, 0xda, 0x23, 0xf7, 0x52, 0xce, 0x7d, 0x15,
	0xc0, 0xb7, 0xb0, 0x75, 0x5f, 0x42, 0x51, 0x05, 0x38, 0x6d, 0xf3, 0x8f, 0x34, 0x5d, 0xec, 0x19,
	0xac, 0xe7, 0x32, 0x0c, 0x42, 0x9a, 0x45, 0x0a, 0x26, 0xc1, 0xb5, 0xdd, 0xd4, 0x9b, 0x97, 0xcb,
	0xd0, 0xbd, 0xa1, 0xdf, 0xf0, 0x25, 0xfc, 0x7f, 0x47, 0x57, 0x2c, 0xdf, 0x62, 0x35, 0x2e, 0xa9,
	0xf5, 0xd2, 0x6f, 0xc0, 0x2e, 0x67, 0x3e, 0xa6, 0x84, 0x32, 0x21, 0xa9, 0xe8, 0xa5, 0xad, 0x4d,
	0x2d, 0xe2, 0x9f, 0x8a, 0xb0, 0x6e, 0x68, 0x08, 0x3b, 0x7f, 0xcb, 0x2c, 0xaa, 0x76, 0x38, 0x7f,
	0xfb, 0xfe, 0x7c, 0x5d, 0xde, 0x01, 0x4b, 0x8f, 0x3c, 0x97, 0x22, 0x93, 0x3c, 0x1c, 0x7b, 0x9d,
	0x8f, 0x78, 0x87, 0xf9, 0xb1, 0xc2, 0x6a, 0x22, 0x6a, 0x06, 0x5a, 0xb9, 0x98, 0xa2, 0x6d, 0xb1,
	0xb0, 0xcd, 0xf8, 0x4b, 0x31, 0x39, 0x7c, 0x0c, 0x2d, 0xad, 0x1b, 0x51, 0x22, 0x33, 0xb2, 0xbb,
	0x6c, 0x42, 0xef, 0xfc, 0x90, 0x91, 0x3a, 0x08, 0x6c, 0x30, 0x98, 0xd2, 0x8c, 0xa6, 0x36, 0xf2,
	0x11, 0x02, 0x46, 0x07, 0x8a, 0x94, 0x02, 0x3d, 0x87, 0x0d, 0x7d, 0x52, 0x18, 0xb1, 0xf7, 0xfe,
	0x0f, 0x03, 0xac, 0xf2, 0x16, 0xfc, 0xe3, 0x1b, 0xbe, 0x7c, 0x25, 0x6b, 0x77, 0xaf, 0x24, 0xbe,
	0x80, 0x3a, 0x7b, 0xc9, 0x6d, 0xb3, 0x57, 0x75, 0x9a, 0x8b, 0x27, 0xa0, 0x34, 0xe8, 0x17, 0x82,
	0x7e, 0x13, 0x1a, 0x9f, 0x49, 0x64, 0x72, 0x9f, 0x84, 0xdc, 0x25, 0x30, 0x0f, 0xf3, 0xb1, 0xe7,
	0x62, 0x1b, 0x1a, 0xc3, 0x34, 0x49, 0x68, 0x24, 0x3d, 0xd7, 0x5a, 0x41, 0x0b, 0x5a, 0x6e, 0x94,
	0x8f, 0x16, 0xc4, 0xc0, 0x2e, 0xb4, 0x79, 0x01, 0xdd, 0xa8, 0xe7, 0x5a, 0x15, 0xdc, 0x04, 0xbc,
	0xfb, 0xc4, 0x79, 0xae, 0x55, 0x45, 0x0b, 0x9a, 0x8b, 0x0a, 0x9e, 0x6b, 0xfd, 0x32, 0x76, 0xdf,
	0x01, 0x94, 0x4a, 0x5c, 0x83, 0xda, 0x51, 0x9a, 0x90, 0xb5, 0x82, 0x4d, 0x58, 0x3d, 0x4a, 0x39,
	0x62, 0x19, 0xd8, 0x00, 0x73, 0x3f, 0x23, 0x71, 0x6e, 0x55, 0x14, 0x3f, 0x49, 0xd3, 0x83, 0x34,
	0x19, 0x5b, 0xd5, 0xd3, 0x3a, 0x3f, 0xac, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x13, 0xa3,
	0x56, 0xb7, 0x77, 0x05, 0x00, 0x00,
}
