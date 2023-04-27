// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: types/v1/filter.proto

package pbscribe

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractAddress *NullableString `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	ChainId         uint32          `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	BlockNumber     *NullableUint64 `protobuf:"bytes,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	TxHash          *NullableString `protobuf:"bytes,4,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	TxIndex         *NullableUint64 `protobuf:"bytes,5,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	BlockHash       *NullableString `protobuf:"bytes,6,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	Index           *NullableUint64 `protobuf:"bytes,7,opt,name=index,proto3" json:"index,omitempty"`
	Confirmed       *NullableBool   `protobuf:"bytes,8,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
}

func (x *LogFilter) Reset() {
	*x = LogFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogFilter) ProtoMessage() {}

func (x *LogFilter) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogFilter.ProtoReflect.Descriptor instead.
func (*LogFilter) Descriptor() ([]byte, []int) {
	return file_types_v1_filter_proto_rawDescGZIP(), []int{0}
}

func (x *LogFilter) GetContractAddress() *NullableString {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

func (x *LogFilter) GetChainId() uint32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *LogFilter) GetBlockNumber() *NullableUint64 {
	if x != nil {
		return x.BlockNumber
	}
	return nil
}

func (x *LogFilter) GetTxHash() *NullableString {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *LogFilter) GetTxIndex() *NullableUint64 {
	if x != nil {
		return x.TxIndex
	}
	return nil
}

func (x *LogFilter) GetBlockHash() *NullableString {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *LogFilter) GetIndex() *NullableUint64 {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *LogFilter) GetConfirmed() *NullableBool {
	if x != nil {
		return x.Confirmed
	}
	return nil
}

var File_types_v1_filter_proto protoreflect.FileDescriptor

var file_types_v1_filter_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x03, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x55,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x31, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x74,
	0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x52, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x37, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x75,
	0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x34, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x63,
	0x6e, 0x73, 0x2f, 0x73, 0x61, 0x6e, 0x67, 0x75, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x70, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_v1_filter_proto_rawDescOnce sync.Once
	file_types_v1_filter_proto_rawDescData = file_types_v1_filter_proto_rawDesc
)

func file_types_v1_filter_proto_rawDescGZIP() []byte {
	file_types_v1_filter_proto_rawDescOnce.Do(func() {
		file_types_v1_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_v1_filter_proto_rawDescData)
	})
	return file_types_v1_filter_proto_rawDescData
}

var file_types_v1_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_v1_filter_proto_goTypes = []interface{}{
	(*LogFilter)(nil),      // 0: types.v1.LogFilter
	(*NullableString)(nil), // 1: types.v1.NullableString
	(*NullableUint64)(nil), // 2: types.v1.NullableUint64
	(*NullableBool)(nil),   // 3: types.v1.NullableBool
}
var file_types_v1_filter_proto_depIdxs = []int32{
	1, // 0: types.v1.LogFilter.contract_address:type_name -> types.v1.NullableString
	2, // 1: types.v1.LogFilter.block_number:type_name -> types.v1.NullableUint64
	1, // 2: types.v1.LogFilter.tx_hash:type_name -> types.v1.NullableString
	2, // 3: types.v1.LogFilter.tx_index:type_name -> types.v1.NullableUint64
	1, // 4: types.v1.LogFilter.block_hash:type_name -> types.v1.NullableString
	2, // 5: types.v1.LogFilter.index:type_name -> types.v1.NullableUint64
	3, // 6: types.v1.LogFilter.confirmed:type_name -> types.v1.NullableBool
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_types_v1_filter_proto_init() }
func file_types_v1_filter_proto_init() {
	if File_types_v1_filter_proto != nil {
		return
	}
	file_types_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_types_v1_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogFilter); i {
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
			RawDescriptor: file_types_v1_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_filter_proto_goTypes,
		DependencyIndexes: file_types_v1_filter_proto_depIdxs,
		MessageInfos:      file_types_v1_filter_proto_msgTypes,
	}.Build()
	File_types_v1_filter_proto = out.File
	file_types_v1_filter_proto_rawDesc = nil
	file_types_v1_filter_proto_goTypes = nil
	file_types_v1_filter_proto_depIdxs = nil
}