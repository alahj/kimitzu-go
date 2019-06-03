// Code generated by protoc-gen-go. DO NOT EDIT.
// source: posts.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Post_PostType int32

const (
	Post_POST    Post_PostType = 0
	Post_COMMENT Post_PostType = 1
	Post_REPOST  Post_PostType = 2
)

var Post_PostType_name = map[int32]string{
	0: "POST",
	1: "COMMENT",
	2: "REPOST",
}

var Post_PostType_value = map[string]int32{
	"POST":    0,
	"COMMENT": 1,
	"REPOST":  2,
}

func (x Post_PostType) String() string {
	return proto.EnumName(Post_PostType_name, int32(x))
}

func (Post_PostType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b14bd1586479c33d, []int{0, 0}
}

type Post struct {
	Slug                 string               `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	VendorID             *ID                  `protobuf:"bytes,2,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
	Status               string               `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	LongForm             string               `protobuf:"bytes,4,opt,name=longForm,proto3" json:"longForm,omitempty"`
	Images               []*Post_Image        `protobuf:"bytes,5,rep,name=images,proto3" json:"images,omitempty"`
	Tags                 []string             `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Channels             []string             `protobuf:"bytes,7,rep,name=channels,proto3" json:"channels,omitempty"`
	PostType             Post_PostType        `protobuf:"varint,8,opt,name=postType,proto3,enum=Post_PostType" json:"postType,omitempty"`
	Reference            string               `protobuf:"bytes,9,opt,name=reference,proto3" json:"reference,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_b14bd1586479c33d, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Post) GetVendorID() *ID {
	if m != nil {
		return m.VendorID
	}
	return nil
}

func (m *Post) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Post) GetLongForm() string {
	if m != nil {
		return m.LongForm
	}
	return ""
}

func (m *Post) GetImages() []*Post_Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Post) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Post) GetChannels() []string {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *Post) GetPostType() Post_PostType {
	if m != nil {
		return m.PostType
	}
	return Post_POST
}

func (m *Post) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *Post) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Post_Image struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Original             string   `protobuf:"bytes,2,opt,name=original,proto3" json:"original,omitempty"`
	Large                string   `protobuf:"bytes,3,opt,name=large,proto3" json:"large,omitempty"`
	Medium               string   `protobuf:"bytes,4,opt,name=medium,proto3" json:"medium,omitempty"`
	Small                string   `protobuf:"bytes,5,opt,name=small,proto3" json:"small,omitempty"`
	Tiny                 string   `protobuf:"bytes,6,opt,name=tiny,proto3" json:"tiny,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post_Image) Reset()         { *m = Post_Image{} }
func (m *Post_Image) String() string { return proto.CompactTextString(m) }
func (*Post_Image) ProtoMessage()    {}
func (*Post_Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_b14bd1586479c33d, []int{0, 0}
}

func (m *Post_Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post_Image.Unmarshal(m, b)
}
func (m *Post_Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post_Image.Marshal(b, m, deterministic)
}
func (m *Post_Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post_Image.Merge(m, src)
}
func (m *Post_Image) XXX_Size() int {
	return xxx_messageInfo_Post_Image.Size(m)
}
func (m *Post_Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Post_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Post_Image proto.InternalMessageInfo

func (m *Post_Image) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Post_Image) GetOriginal() string {
	if m != nil {
		return m.Original
	}
	return ""
}

func (m *Post_Image) GetLarge() string {
	if m != nil {
		return m.Large
	}
	return ""
}

func (m *Post_Image) GetMedium() string {
	if m != nil {
		return m.Medium
	}
	return ""
}

func (m *Post_Image) GetSmall() string {
	if m != nil {
		return m.Small
	}
	return ""
}

func (m *Post_Image) GetTiny() string {
	if m != nil {
		return m.Tiny
	}
	return ""
}

type SignedPost struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedPost) Reset()         { *m = SignedPost{} }
func (m *SignedPost) String() string { return proto.CompactTextString(m) }
func (*SignedPost) ProtoMessage()    {}
func (*SignedPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_b14bd1586479c33d, []int{1}
}

func (m *SignedPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedPost.Unmarshal(m, b)
}
func (m *SignedPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedPost.Marshal(b, m, deterministic)
}
func (m *SignedPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedPost.Merge(m, src)
}
func (m *SignedPost) XXX_Size() int {
	return xxx_messageInfo_SignedPost.Size(m)
}
func (m *SignedPost) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedPost.DiscardUnknown(m)
}

var xxx_messageInfo_SignedPost proto.InternalMessageInfo

func (m *SignedPost) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *SignedPost) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *SignedPost) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("Post_PostType", Post_PostType_name, Post_PostType_value)
	proto.RegisterType((*Post)(nil), "Post")
	proto.RegisterType((*Post_Image)(nil), "Post.Image")
	proto.RegisterType((*SignedPost)(nil), "SignedPost")
}

func init() { proto.RegisterFile("posts.proto", fileDescriptor_b14bd1586479c33d) }

var fileDescriptor_b14bd1586479c33d = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x92, 0x4f, 0x8b, 0xdb, 0x30,
	0x10, 0xc5, 0xeb, 0xc4, 0xf6, 0xda, 0xe3, 0xb2, 0x5d, 0x44, 0x29, 0x6a, 0x28, 0xac, 0x49, 0x2f,
	0xa1, 0x50, 0x2f, 0xa4, 0x97, 0x9e, 0xdb, 0xdd, 0x42, 0x0e, 0xdb, 0x0d, 0xde, 0x5c, 0xda, 0x9b,
	0x92, 0x28, 0x8a, 0x40, 0x96, 0x8c, 0x24, 0x17, 0xf6, 0x63, 0xf4, 0xd4, 0xaf, 0x5b, 0x34, 0xfe,
	0x93, 0xdb, 0xbc, 0x37, 0x23, 0x69, 0xde, 0xcf, 0x86, 0xa2, 0x35, 0xce, 0xbb, 0xaa, 0xb5, 0xc6,
	0x9b, 0xc5, 0xad, 0x30, 0x46, 0x28, 0x7e, 0x87, 0x6a, 0xdf, 0x9d, 0xee, 0xbc, 0x6c, 0xb8, 0xf3,
	0xac, 0x69, 0x87, 0x81, 0x37, 0x07, 0xa3, 0xbd, 0x65, 0x87, 0xf1, 0xc4, 0xf2, 0x6f, 0x0c, 0xf1,
	0xd6, 0x38, 0x4f, 0x08, 0xc4, 0x4e, 0x75, 0x82, 0x46, 0x65, 0xb4, 0xca, 0x6b, 0xac, 0xc9, 0x2d,
	0x64, 0x7f, 0xb8, 0x3e, 0x1a, 0xbb, 0xb9, 0xa7, 0xb3, 0x32, 0x5a, 0x15, 0xeb, 0x79, 0xb5, 0xb9,
	0xaf, 0x27, 0x93, 0xbc, 0x83, 0xd4, 0x79, 0xe6, 0x3b, 0x47, 0xe7, 0x78, 0x6c, 0x50, 0x64, 0x01,
	0x99, 0x32, 0x5a, 0xfc, 0x30, 0xb6, 0xa1, 0x31, 0x76, 0x26, 0x4d, 0x3e, 0x42, 0x2a, 0x1b, 0x26,
	0xb8, 0xa3, 0x49, 0x39, 0x5f, 0x15, 0xeb, 0xa2, 0x0a, 0xef, 0x57, 0x9b, 0xe0, 0xd5, 0x43, 0x2b,
	0x6c, 0xe3, 0x99, 0x70, 0x34, 0x2d, 0xe7, 0x61, 0x9b, 0x50, 0x87, 0x4b, 0x0f, 0x67, 0xa6, 0x35,
	0x57, 0x8e, 0x5e, 0xa1, 0x3f, 0x69, 0xf2, 0x09, 0xb2, 0xc0, 0x61, 0xf7, 0xd2, 0x72, 0x9a, 0x95,
	0xd1, 0xea, 0x7a, 0x7d, 0xdd, 0x5f, 0xbb, 0x1d, 0xdc, 0x7a, 0xea, 0x93, 0x0f, 0x90, 0x5b, 0x7e,
	0xe2, 0x96, 0xeb, 0x03, 0xa7, 0x39, 0x6e, 0x77, 0x31, 0xc8, 0x57, 0xc8, 0x27, 0x68, 0x14, 0x30,
	0xf4, 0xa2, 0xea, 0xb1, 0x56, 0x23, 0xd6, 0x6a, 0x37, 0x4e, 0xd4, 0x97, 0xe1, 0xc5, 0xbf, 0x08,
	0x12, 0x4c, 0x11, 0x36, 0x3d, 0x49, 0xc5, 0x35, 0x6b, 0xf8, 0xc0, 0x73, 0xd2, 0xa1, 0x67, 0xac,
	0x14, 0x52, 0x33, 0x85, 0x4c, 0xf3, 0x7a, 0xd2, 0xe4, 0x2d, 0x24, 0x8a, 0x59, 0xc1, 0x07, 0x9a,
	0xbd, 0x08, 0x90, 0x1b, 0x7e, 0x94, 0xdd, 0x88, 0x72, 0x50, 0x61, 0xda, 0x35, 0x4c, 0x29, 0x9a,
	0xf4, 0xd3, 0x28, 0x90, 0x9c, 0xd4, 0x2f, 0x34, 0xed, 0xbf, 0x63, 0xa8, 0x97, 0x9f, 0x21, 0x1b,
	0x39, 0x90, 0x0c, 0xe2, 0xed, 0xd3, 0xf3, 0xee, 0xe6, 0x15, 0x29, 0xe0, 0xea, 0xfb, 0xd3, 0xe3,
	0xe3, 0xc3, 0xcf, 0xdd, 0x4d, 0x44, 0x00, 0xd2, 0xfa, 0x01, 0x1b, 0xb3, 0xe5, 0x2f, 0x80, 0x67,
	0x29, 0x34, 0x3f, 0xe2, 0x8f, 0xf1, 0x1e, 0xe2, 0x80, 0x0e, 0x83, 0x14, 0xeb, 0x04, 0x89, 0xd6,
	0x68, 0x85, 0xb7, 0xce, 0xcc, 0x9d, 0x87, 0x1c, 0x58, 0x07, 0xba, 0x4e, 0x0a, 0xcd, 0x7c, 0x67,
	0xfb, 0x1c, 0xaf, 0xeb, 0x8b, 0xf1, 0x2d, 0xfe, 0x3d, 0x6b, 0xf7, 0xfb, 0x14, 0x41, 0x7e, 0xf9,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x2a, 0x4a, 0xa9, 0xbc, 0x02, 0x00, 0x00,
}
