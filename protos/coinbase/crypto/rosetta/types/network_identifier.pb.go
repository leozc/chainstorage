// (-- api-linter: core::0140::abbreviations=disabled
//     aip.dev/not-precedent: Matching the naming convention of rosetta --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: coinbase/crypto/rosetta/types/network_identifier.proto

// The stable release for rosetta types

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// NetworkIdentifier specifies which network a particular object is associated with.
type NetworkIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Blockchain the network is associated with
	Blockchain string `protobuf:"bytes,1,opt,name=blockchain,proto3" json:"blockchain,omitempty"`
	// If a blockchain has a specific chain-id or network identifier, it should go in
	// this field. It is up to the client to determine which network-specific identifier
	// is mainnet or testnet.
	Network string `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	// In blockchains with sharded state, the SubNetworkIdentifier is required to query
	// some object on a specific shard. This identifier is optional for all non-sharded
	// blockchains.
	SubNetworkIdentifier *SubNetworkIdentifier `protobuf:"bytes,3,opt,name=sub_network_identifier,json=subNetworkIdentifier,proto3" json:"sub_network_identifier,omitempty"`
}

func (x *NetworkIdentifier) Reset() {
	*x = NetworkIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_network_identifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkIdentifier) ProtoMessage() {}

func (x *NetworkIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_network_identifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkIdentifier.ProtoReflect.Descriptor instead.
func (*NetworkIdentifier) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkIdentifier) GetBlockchain() string {
	if x != nil {
		return x.Blockchain
	}
	return ""
}

func (x *NetworkIdentifier) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *NetworkIdentifier) GetSubNetworkIdentifier() *SubNetworkIdentifier {
	if x != nil {
		return x.SubNetworkIdentifier
	}
	return nil
}

// In blockchains with sharded state, the SubNetworkIdentifier is required to
// query some object on a specific shard. This identifier is optional for
// all non-sharded blockchains.
type SubNetworkIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subnetwork identifier.
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	// Metadata info for the subnetwork.
	Metadata map[string]*anypb.Any `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SubNetworkIdentifier) Reset() {
	*x = SubNetworkIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_network_identifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubNetworkIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubNetworkIdentifier) ProtoMessage() {}

func (x *SubNetworkIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_network_identifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubNetworkIdentifier.ProtoReflect.Descriptor instead.
func (*SubNetworkIdentifier) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDescGZIP(), []int{1}
}

func (x *SubNetworkIdentifier) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *SubNetworkIdentifier) GetMetadata() map[string]*anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_coinbase_crypto_rosetta_types_network_identifier_proto protoreflect.FileDescriptor

var file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74,
	0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x69, 0x0a, 0x16, 0x73, 0x75, 0x62, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x14, 0x73, 0x75, 0x62, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xe2, 0x01,
	0x0a, 0x14, 0x53, 0x75, 0x62, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x5d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x51, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f,
	0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDescOnce sync.Once
	file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDescData = file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDesc
)

func file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDescGZIP() []byte {
	file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDescOnce.Do(func() {
		file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDescData)
	})
	return file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDescData
}

var file_coinbase_crypto_rosetta_types_network_identifier_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coinbase_crypto_rosetta_types_network_identifier_proto_goTypes = []interface{}{
	(*NetworkIdentifier)(nil),    // 0: coinbase.crypto.rosetta.types.NetworkIdentifier
	(*SubNetworkIdentifier)(nil), // 1: coinbase.crypto.rosetta.types.SubNetworkIdentifier
	nil,                          // 2: coinbase.crypto.rosetta.types.SubNetworkIdentifier.MetadataEntry
	(*anypb.Any)(nil),            // 3: google.protobuf.Any
}
var file_coinbase_crypto_rosetta_types_network_identifier_proto_depIdxs = []int32{
	1, // 0: coinbase.crypto.rosetta.types.NetworkIdentifier.sub_network_identifier:type_name -> coinbase.crypto.rosetta.types.SubNetworkIdentifier
	2, // 1: coinbase.crypto.rosetta.types.SubNetworkIdentifier.metadata:type_name -> coinbase.crypto.rosetta.types.SubNetworkIdentifier.MetadataEntry
	3, // 2: coinbase.crypto.rosetta.types.SubNetworkIdentifier.MetadataEntry.value:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_coinbase_crypto_rosetta_types_network_identifier_proto_init() }
func file_coinbase_crypto_rosetta_types_network_identifier_proto_init() {
	if File_coinbase_crypto_rosetta_types_network_identifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_crypto_rosetta_types_network_identifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkIdentifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinbase_crypto_rosetta_types_network_identifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubNetworkIdentifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_crypto_rosetta_types_network_identifier_proto_goTypes,
		DependencyIndexes: file_coinbase_crypto_rosetta_types_network_identifier_proto_depIdxs,
		MessageInfos:      file_coinbase_crypto_rosetta_types_network_identifier_proto_msgTypes,
	}.Build()
	File_coinbase_crypto_rosetta_types_network_identifier_proto = out.File
	file_coinbase_crypto_rosetta_types_network_identifier_proto_rawDesc = nil
	file_coinbase_crypto_rosetta_types_network_identifier_proto_goTypes = nil
	file_coinbase_crypto_rosetta_types_network_identifier_proto_depIdxs = nil
}
