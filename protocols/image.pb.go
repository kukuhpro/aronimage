// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image.proto

package protocols

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProcessingImageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Base64String         string   `protobuf:"bytes,2,opt,name=base64string,proto3" json:"base64string,omitempty"`
	Imagename            string   `protobuf:"bytes,3,opt,name=imagename,proto3" json:"imagename,omitempty"`
	Prefixpath           string   `protobuf:"bytes,4,opt,name=prefixpath,proto3" json:"prefixpath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessingImageRequest) Reset()         { *m = ProcessingImageRequest{} }
func (m *ProcessingImageRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessingImageRequest) ProtoMessage()    {}
func (*ProcessingImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cb8c039d77e86688, []int{0}
}
func (m *ProcessingImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessingImageRequest.Unmarshal(m, b)
}
func (m *ProcessingImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessingImageRequest.Marshal(b, m, deterministic)
}
func (dst *ProcessingImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessingImageRequest.Merge(dst, src)
}
func (m *ProcessingImageRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessingImageRequest.Size(m)
}
func (m *ProcessingImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessingImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessingImageRequest proto.InternalMessageInfo

func (m *ProcessingImageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProcessingImageRequest) GetBase64String() string {
	if m != nil {
		return m.Base64String
	}
	return ""
}

func (m *ProcessingImageRequest) GetImagename() string {
	if m != nil {
		return m.Imagename
	}
	return ""
}

func (m *ProcessingImageRequest) GetPrefixpath() string {
	if m != nil {
		return m.Prefixpath
	}
	return ""
}

type ImageListPathRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Prefixpath           string   `protobuf:"bytes,2,opt,name=prefixpath,proto3" json:"prefixpath,omitempty"`
	Filename             string   `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageListPathRequest) Reset()         { *m = ImageListPathRequest{} }
func (m *ImageListPathRequest) String() string { return proto.CompactTextString(m) }
func (*ImageListPathRequest) ProtoMessage()    {}
func (*ImageListPathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cb8c039d77e86688, []int{1}
}
func (m *ImageListPathRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageListPathRequest.Unmarshal(m, b)
}
func (m *ImageListPathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageListPathRequest.Marshal(b, m, deterministic)
}
func (dst *ImageListPathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageListPathRequest.Merge(dst, src)
}
func (m *ImageListPathRequest) XXX_Size() int {
	return xxx_messageInfo_ImageListPathRequest.Size(m)
}
func (m *ImageListPathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageListPathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImageListPathRequest proto.InternalMessageInfo

func (m *ImageListPathRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageListPathRequest) GetPrefixpath() string {
	if m != nil {
		return m.Prefixpath
	}
	return ""
}

func (m *ImageListPathRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type ImageListPathResponse struct {
	Bucketname           string      `protobuf:"bytes,1,opt,name=bucketname,proto3" json:"bucketname,omitempty"`
	Images               []*ListPath `protobuf:"bytes,2,rep,name=images,proto3" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ImageListPathResponse) Reset()         { *m = ImageListPathResponse{} }
func (m *ImageListPathResponse) String() string { return proto.CompactTextString(m) }
func (*ImageListPathResponse) ProtoMessage()    {}
func (*ImageListPathResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cb8c039d77e86688, []int{2}
}
func (m *ImageListPathResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageListPathResponse.Unmarshal(m, b)
}
func (m *ImageListPathResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageListPathResponse.Marshal(b, m, deterministic)
}
func (dst *ImageListPathResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageListPathResponse.Merge(dst, src)
}
func (m *ImageListPathResponse) XXX_Size() int {
	return xxx_messageInfo_ImageListPathResponse.Size(m)
}
func (m *ImageListPathResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageListPathResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImageListPathResponse proto.InternalMessageInfo

func (m *ImageListPathResponse) GetBucketname() string {
	if m != nil {
		return m.Bucketname
	}
	return ""
}

func (m *ImageListPathResponse) GetImages() []*ListPath {
	if m != nil {
		return m.Images
	}
	return nil
}

type ListPath struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPath) Reset()         { *m = ListPath{} }
func (m *ListPath) String() string { return proto.CompactTextString(m) }
func (*ListPath) ProtoMessage()    {}
func (*ListPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cb8c039d77e86688, []int{3}
}
func (m *ListPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPath.Unmarshal(m, b)
}
func (m *ListPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPath.Marshal(b, m, deterministic)
}
func (dst *ListPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPath.Merge(dst, src)
}
func (m *ListPath) XXX_Size() int {
	return xxx_messageInfo_ListPath.Size(m)
}
func (m *ListPath) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPath.DiscardUnknown(m)
}

var xxx_messageInfo_ListPath proto.InternalMessageInfo

func (m *ListPath) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListPath) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ProcessingImageResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessingImageResponse) Reset()         { *m = ProcessingImageResponse{} }
func (m *ProcessingImageResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessingImageResponse) ProtoMessage()    {}
func (*ProcessingImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cb8c039d77e86688, []int{4}
}
func (m *ProcessingImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessingImageResponse.Unmarshal(m, b)
}
func (m *ProcessingImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessingImageResponse.Marshal(b, m, deterministic)
}
func (dst *ProcessingImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessingImageResponse.Merge(dst, src)
}
func (m *ProcessingImageResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessingImageResponse.Size(m)
}
func (m *ProcessingImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessingImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessingImageResponse proto.InternalMessageInfo

func (m *ProcessingImageResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type SetupConfigurationRequest struct {
	Config                       *GlobalConfig                   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Imageprocessingconfiguration []*ImageProcessingConfiguration `protobuf:"bytes,2,rep,name=imageprocessingconfiguration,proto3" json:"imageprocessingconfiguration,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                        `json:"-"`
	XXX_unrecognized             []byte                          `json:"-"`
	XXX_sizecache                int32                           `json:"-"`
}

func (m *SetupConfigurationRequest) Reset()         { *m = SetupConfigurationRequest{} }
func (m *SetupConfigurationRequest) String() string { return proto.CompactTextString(m) }
func (*SetupConfigurationRequest) ProtoMessage()    {}
func (*SetupConfigurationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cb8c039d77e86688, []int{5}
}
func (m *SetupConfigurationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetupConfigurationRequest.Unmarshal(m, b)
}
func (m *SetupConfigurationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetupConfigurationRequest.Marshal(b, m, deterministic)
}
func (dst *SetupConfigurationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetupConfigurationRequest.Merge(dst, src)
}
func (m *SetupConfigurationRequest) XXX_Size() int {
	return xxx_messageInfo_SetupConfigurationRequest.Size(m)
}
func (m *SetupConfigurationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetupConfigurationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetupConfigurationRequest proto.InternalMessageInfo

func (m *SetupConfigurationRequest) GetConfig() *GlobalConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *SetupConfigurationRequest) GetImageprocessingconfiguration() []*ImageProcessingConfiguration {
	if m != nil {
		return m.Imageprocessingconfiguration
	}
	return nil
}

type GlobalConfig struct {
	Bucketname           string   `protobuf:"bytes,1,opt,name=bucketname,proto3" json:"bucketname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlobalConfig) Reset()         { *m = GlobalConfig{} }
func (m *GlobalConfig) String() string { return proto.CompactTextString(m) }
func (*GlobalConfig) ProtoMessage()    {}
func (*GlobalConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cb8c039d77e86688, []int{6}
}
func (m *GlobalConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalConfig.Unmarshal(m, b)
}
func (m *GlobalConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalConfig.Marshal(b, m, deterministic)
}
func (dst *GlobalConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalConfig.Merge(dst, src)
}
func (m *GlobalConfig) XXX_Size() int {
	return xxx_messageInfo_GlobalConfig.Size(m)
}
func (m *GlobalConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalConfig proto.InternalMessageInfo

func (m *GlobalConfig) GetBucketname() string {
	if m != nil {
		return m.Bucketname
	}
	return ""
}

type ImageProcessingConfiguration struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Imageconfigs         []*ImageConfig `protobuf:"bytes,2,rep,name=imageconfigs,proto3" json:"imageconfigs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ImageProcessingConfiguration) Reset()         { *m = ImageProcessingConfiguration{} }
func (m *ImageProcessingConfiguration) String() string { return proto.CompactTextString(m) }
func (*ImageProcessingConfiguration) ProtoMessage()    {}
func (*ImageProcessingConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cb8c039d77e86688, []int{7}
}
func (m *ImageProcessingConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageProcessingConfiguration.Unmarshal(m, b)
}
func (m *ImageProcessingConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageProcessingConfiguration.Marshal(b, m, deterministic)
}
func (dst *ImageProcessingConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageProcessingConfiguration.Merge(dst, src)
}
func (m *ImageProcessingConfiguration) XXX_Size() int {
	return xxx_messageInfo_ImageProcessingConfiguration.Size(m)
}
func (m *ImageProcessingConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageProcessingConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ImageProcessingConfiguration proto.InternalMessageInfo

func (m *ImageProcessingConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageProcessingConfiguration) GetImageconfigs() []*ImageConfig {
	if m != nil {
		return m.Imageconfigs
	}
	return nil
}

type ImageConfig struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Width                int32    `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageConfig) Reset()         { *m = ImageConfig{} }
func (m *ImageConfig) String() string { return proto.CompactTextString(m) }
func (*ImageConfig) ProtoMessage()    {}
func (*ImageConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cb8c039d77e86688, []int{8}
}
func (m *ImageConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageConfig.Unmarshal(m, b)
}
func (m *ImageConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageConfig.Marshal(b, m, deterministic)
}
func (dst *ImageConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageConfig.Merge(dst, src)
}
func (m *ImageConfig) XXX_Size() int {
	return xxx_messageInfo_ImageConfig.Size(m)
}
func (m *ImageConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ImageConfig proto.InternalMessageInfo

func (m *ImageConfig) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ImageConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageConfig) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ImageConfig) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

type SetupConfigurationResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetupConfigurationResponse) Reset()         { *m = SetupConfigurationResponse{} }
func (m *SetupConfigurationResponse) String() string { return proto.CompactTextString(m) }
func (*SetupConfigurationResponse) ProtoMessage()    {}
func (*SetupConfigurationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cb8c039d77e86688, []int{9}
}
func (m *SetupConfigurationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetupConfigurationResponse.Unmarshal(m, b)
}
func (m *SetupConfigurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetupConfigurationResponse.Marshal(b, m, deterministic)
}
func (dst *SetupConfigurationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetupConfigurationResponse.Merge(dst, src)
}
func (m *SetupConfigurationResponse) XXX_Size() int {
	return xxx_messageInfo_SetupConfigurationResponse.Size(m)
}
func (m *SetupConfigurationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetupConfigurationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetupConfigurationResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProcessingImageRequest)(nil), "protocols.ProcessingImageRequest")
	proto.RegisterType((*ImageListPathRequest)(nil), "protocols.ImageListPathRequest")
	proto.RegisterType((*ImageListPathResponse)(nil), "protocols.ImageListPathResponse")
	proto.RegisterType((*ListPath)(nil), "protocols.ListPath")
	proto.RegisterType((*ProcessingImageResponse)(nil), "protocols.ProcessingImageResponse")
	proto.RegisterType((*SetupConfigurationRequest)(nil), "protocols.SetupConfigurationRequest")
	proto.RegisterType((*GlobalConfig)(nil), "protocols.GlobalConfig")
	proto.RegisterType((*ImageProcessingConfiguration)(nil), "protocols.ImageProcessingConfiguration")
	proto.RegisterType((*ImageConfig)(nil), "protocols.ImageConfig")
	proto.RegisterType((*SetupConfigurationResponse)(nil), "protocols.SetupConfigurationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImageClient is the client API for Image service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageClient interface {
	ProcessingImage(ctx context.Context, in *ProcessingImageRequest, opts ...grpc.CallOption) (*ProcessingImageResponse, error)
	GetImageListPath(ctx context.Context, in *ImageListPathRequest, opts ...grpc.CallOption) (*ImageListPathResponse, error)
}

type imageClient struct {
	cc *grpc.ClientConn
}

func NewImageClient(cc *grpc.ClientConn) ImageClient {
	return &imageClient{cc}
}

func (c *imageClient) ProcessingImage(ctx context.Context, in *ProcessingImageRequest, opts ...grpc.CallOption) (*ProcessingImageResponse, error) {
	out := new(ProcessingImageResponse)
	err := c.cc.Invoke(ctx, "/protocols.Image/ProcessingImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) GetImageListPath(ctx context.Context, in *ImageListPathRequest, opts ...grpc.CallOption) (*ImageListPathResponse, error) {
	out := new(ImageListPathResponse)
	err := c.cc.Invoke(ctx, "/protocols.Image/GetImageListPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServer is the server API for Image service.
type ImageServer interface {
	ProcessingImage(context.Context, *ProcessingImageRequest) (*ProcessingImageResponse, error)
	GetImageListPath(context.Context, *ImageListPathRequest) (*ImageListPathResponse, error)
}

func RegisterImageServer(s *grpc.Server, srv ImageServer) {
	s.RegisterService(&_Image_serviceDesc, srv)
}

func _Image_ProcessingImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessingImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).ProcessingImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.Image/ProcessingImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).ProcessingImage(ctx, req.(*ProcessingImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_GetImageListPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageListPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).GetImageListPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.Image/GetImageListPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).GetImageListPath(ctx, req.(*ImageListPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Image_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocols.Image",
	HandlerType: (*ImageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessingImage",
			Handler:    _Image_ProcessingImage_Handler,
		},
		{
			MethodName: "GetImageListPath",
			Handler:    _Image_GetImageListPath_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image.proto",
}

func init() { proto.RegisterFile("image.proto", fileDescriptor_image_cb8c039d77e86688) }

var fileDescriptor_image_cb8c039d77e86688 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x25, 0xed, 0x26, 0xda, 0x4e, 0x2b, 0xb1, 0x32, 0x4b, 0xb7, 0x54, 0x15, 0x14, 0x5f, 0x58,
	0x09, 0x29, 0x40, 0x41, 0x1c, 0xb8, 0x72, 0x58, 0x21, 0x71, 0x58, 0x85, 0x03, 0x12, 0x37, 0x27,
	0xeb, 0x24, 0xd6, 0x66, 0xe3, 0x10, 0x3b, 0x02, 0x7e, 0x82, 0xff, 0x41, 0xe2, 0xe3, 0x50, 0x26,
	0x4e, 0xea, 0x84, 0x10, 0x4e, 0xf6, 0x8c, 0x9f, 0xdf, 0x7b, 0x33, 0x63, 0xc3, 0x52, 0xdc, 0xb1,
	0x84, 0xfb, 0x45, 0x29, 0xb5, 0x24, 0x0b, 0x5c, 0x22, 0x99, 0x29, 0xfa, 0xd3, 0x81, 0xf5, 0x75,
	0x29, 0x23, 0xae, 0x94, 0xc8, 0x93, 0x0f, 0x35, 0x28, 0xe0, 0x5f, 0x2b, 0xae, 0x34, 0x21, 0x70,
	0x92, 0xb3, 0x3b, 0xbe, 0x71, 0xf6, 0xce, 0xe5, 0x22, 0xc0, 0x3d, 0xa1, 0xb0, 0x0a, 0x99, 0xe2,
	0x6f, 0xdf, 0x28, 0x5d, 0x8a, 0x3c, 0xd9, 0xcc, 0xf0, 0xac, 0x97, 0x23, 0x3b, 0x58, 0xa0, 0x18,
	0x5e, 0x9e, 0x23, 0xe0, 0x98, 0x20, 0x8f, 0x01, 0x8a, 0x92, 0xc7, 0xe2, 0x7b, 0xc1, 0x74, 0xba,
	0x39, 0xc1, 0x63, 0x2b, 0x43, 0x63, 0x38, 0x47, 0x17, 0x1f, 0x85, 0xd2, 0xd7, 0x4c, 0xa7, 0x53,
	0x6e, 0xfa, 0x5c, 0xb3, 0x21, 0x17, 0xd9, 0xc2, 0x69, 0x2c, 0x32, 0xdb, 0x48, 0x17, 0xd3, 0x1b,
	0x78, 0x38, 0xd0, 0x51, 0x85, 0xcc, 0x15, 0x92, 0x86, 0x55, 0x74, 0xcb, 0xb5, 0x25, 0x67, 0x65,
	0xc8, 0x73, 0xf0, 0xb0, 0x1a, 0xb5, 0x99, 0xed, 0xe7, 0x97, 0xcb, 0xc3, 0x03, 0xbf, 0xeb, 0xa6,
	0xdf, 0x91, 0x19, 0x08, 0x7d, 0x09, 0xa7, 0x6d, 0x6e, 0xb4, 0x82, 0x33, 0x98, 0x57, 0x65, 0x66,
	0xac, 0xd7, 0x5b, 0xfa, 0x0a, 0x2e, 0xfe, 0x9a, 0x87, 0x71, 0xb6, 0x06, 0x4f, 0x69, 0xa6, 0x2b,
	0x85, 0x14, 0x6e, 0x60, 0x22, 0xfa, 0xcb, 0x81, 0x47, 0x9f, 0xb8, 0xae, 0x8a, 0xf7, 0x32, 0x8f,
	0x45, 0x52, 0x95, 0x4c, 0x0b, 0x99, 0xb7, 0x8d, 0x7b, 0x01, 0x5e, 0x84, 0x79, 0xbc, 0xb5, 0x3c,
	0x5c, 0x58, 0x7e, 0xaf, 0x32, 0x19, 0xb2, 0xac, 0xb9, 0x16, 0x18, 0x18, 0xb9, 0x85, 0x1d, 0xba,
	0x2f, 0x3a, 0x1b, 0x91, 0xcd, 0x6b, 0xca, 0x7e, 0x66, 0xd1, 0xa0, 0xcd, 0xa3, 0xeb, 0xbe, 0x8d,
	0x49, 0x32, 0xea, 0xc3, 0xca, 0x36, 0xf1, 0xbf, 0xee, 0xd3, 0x1c, 0x76, 0x53, 0x6a, 0xa3, 0x4d,
	0x7e, 0x07, 0x2b, 0xf4, 0xd0, 0x28, 0xb7, 0x73, 0x5b, 0x0f, 0x0b, 0x30, 0x6d, 0xe8, 0x61, 0x69,
	0x04, 0x4b, 0xeb, 0xb0, 0xa6, 0xd7, 0x3f, 0x8a, 0x8e, 0xbe, 0xde, 0x77, 0x92, 0x33, 0x4b, 0x72,
	0x0d, 0x5e, 0xca, 0x45, 0x92, 0x6a, 0x7c, 0x77, 0x6e, 0x60, 0x22, 0x72, 0x0e, 0xee, 0x37, 0x71,
	0x63, 0x1e, 0xbe, 0x1b, 0x34, 0x01, 0xdd, 0xc1, 0x76, 0x6c, 0x7e, 0xcd, 0xd8, 0x0f, 0xbf, 0x1d,
	0x70, 0xd1, 0x03, 0xf9, 0x02, 0xf7, 0x07, 0x6f, 0x83, 0x3c, 0xb5, 0xaa, 0x18, 0xff, 0xc7, 0x5b,
	0x3a, 0x05, 0x69, 0x34, 0xe8, 0x3d, 0xf2, 0x19, 0xce, 0xae, 0xb8, 0xee, 0x7d, 0x09, 0xf2, 0x64,
	0xd8, 0xa2, 0xc1, 0xa7, 0xdc, 0xee, 0xff, 0x0d, 0x68, 0x89, 0x43, 0x0f, 0x21, 0xaf, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x29, 0xf5, 0x75, 0xe7, 0x82, 0x04, 0x00, 0x00,
}
