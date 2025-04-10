// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: calculator/v1/calculator.proto

package calculatorv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OperateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number1       float64                `protobuf:"fixed64,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2       float64                `protobuf:"fixed64,2,opt,name=number2,proto3" json:"number2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperateRequest) Reset() {
	*x = OperateRequest{}
	mi := &file_calculator_v1_calculator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateRequest) ProtoMessage() {}

func (x *OperateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_v1_calculator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateRequest.ProtoReflect.Descriptor instead.
func (*OperateRequest) Descriptor() ([]byte, []int) {
	return file_calculator_v1_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *OperateRequest) GetNumber1() float64 {
	if x != nil {
		return x.Number1
	}
	return 0
}

func (x *OperateRequest) GetNumber2() float64 {
	if x != nil {
		return x.Number2
	}
	return 0
}

type OperateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        float64                `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperateResponse) Reset() {
	*x = OperateResponse{}
	mi := &file_calculator_v1_calculator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateResponse) ProtoMessage() {}

func (x *OperateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_v1_calculator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateResponse.ProtoReflect.Descriptor instead.
func (*OperateResponse) Descriptor() ([]byte, []int) {
	return file_calculator_v1_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *OperateResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_v1_calculator_proto protoreflect.FileDescriptor

const file_calculator_v1_calculator_proto_rawDesc = "" +
	"\n" +
	"\x1ecalculator/v1/calculator.proto\x12\rcalculator.v1\"D\n" +
	"\x0eOperateRequest\x12\x18\n" +
	"\anumber1\x18\x01 \x01(\x01R\anumber1\x12\x18\n" +
	"\anumber2\x18\x02 \x01(\x01R\anumber2\")\n" +
	"\x0fOperateResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\x01R\x06result2\xb8\x02\n" +
	"\x11CalculatorService\x12D\n" +
	"\x03Add\x12\x1d.calculator.v1.OperateRequest\x1a\x1e.calculator.v1.OperateResponse\x12I\n" +
	"\bSubtract\x12\x1d.calculator.v1.OperateRequest\x1a\x1e.calculator.v1.OperateResponse\x12I\n" +
	"\bMultiply\x12\x1d.calculator.v1.OperateRequest\x1a\x1e.calculator.v1.OperateResponse\x12G\n" +
	"\x06Divide\x12\x1d.calculator.v1.OperateRequest\x1a\x1e.calculator.v1.OperateResponseB0Z.ConnectRPC-Demo/gen/calculator/v1;calculatorv1b\x06proto3"

var (
	file_calculator_v1_calculator_proto_rawDescOnce sync.Once
	file_calculator_v1_calculator_proto_rawDescData []byte
)

func file_calculator_v1_calculator_proto_rawDescGZIP() []byte {
	file_calculator_v1_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_v1_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_calculator_v1_calculator_proto_rawDesc), len(file_calculator_v1_calculator_proto_rawDesc)))
	})
	return file_calculator_v1_calculator_proto_rawDescData
}

var file_calculator_v1_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculator_v1_calculator_proto_goTypes = []any{
	(*OperateRequest)(nil),  // 0: calculator.v1.OperateRequest
	(*OperateResponse)(nil), // 1: calculator.v1.OperateResponse
}
var file_calculator_v1_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.v1.CalculatorService.Add:input_type -> calculator.v1.OperateRequest
	0, // 1: calculator.v1.CalculatorService.Subtract:input_type -> calculator.v1.OperateRequest
	0, // 2: calculator.v1.CalculatorService.Multiply:input_type -> calculator.v1.OperateRequest
	0, // 3: calculator.v1.CalculatorService.Divide:input_type -> calculator.v1.OperateRequest
	1, // 4: calculator.v1.CalculatorService.Add:output_type -> calculator.v1.OperateResponse
	1, // 5: calculator.v1.CalculatorService.Subtract:output_type -> calculator.v1.OperateResponse
	1, // 6: calculator.v1.CalculatorService.Multiply:output_type -> calculator.v1.OperateResponse
	1, // 7: calculator.v1.CalculatorService.Divide:output_type -> calculator.v1.OperateResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_v1_calculator_proto_init() }
func file_calculator_v1_calculator_proto_init() {
	if File_calculator_v1_calculator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_calculator_v1_calculator_proto_rawDesc), len(file_calculator_v1_calculator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_v1_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_v1_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_v1_calculator_proto_msgTypes,
	}.Build()
	File_calculator_v1_calculator_proto = out.File
	file_calculator_v1_calculator_proto_goTypes = nil
	file_calculator_v1_calculator_proto_depIdxs = nil
}
