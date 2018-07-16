// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/types.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// LINT.IfChange
type DataType int32

const (
	// Not a legal value for DataType.  Used to indicate a DataType field
	// has not been set.
	DataType_DT_INVALID DataType = 0
	// Data types that all computation devices are expected to be
	// capable to support.
	DataType_DT_FLOAT      DataType = 1
	DataType_DT_DOUBLE     DataType = 2
	DataType_DT_INT32      DataType = 3
	DataType_DT_UINT8      DataType = 4
	DataType_DT_INT16      DataType = 5
	DataType_DT_INT8       DataType = 6
	DataType_DT_STRING     DataType = 7
	DataType_DT_COMPLEX64  DataType = 8
	DataType_DT_INT64      DataType = 9
	DataType_DT_BOOL       DataType = 10
	DataType_DT_QINT8      DataType = 11
	DataType_DT_QUINT8     DataType = 12
	DataType_DT_QINT32     DataType = 13
	DataType_DT_BFLOAT16   DataType = 14
	DataType_DT_QINT16     DataType = 15
	DataType_DT_QUINT16    DataType = 16
	DataType_DT_UINT16     DataType = 17
	DataType_DT_COMPLEX128 DataType = 18
	DataType_DT_HALF       DataType = 19
	DataType_DT_RESOURCE   DataType = 20
	// Do not use!  These are only for parameters.  Every enum above
	// should have a corresponding value below (verified by types_test).
	DataType_DT_FLOAT_REF      DataType = 101
	DataType_DT_DOUBLE_REF     DataType = 102
	DataType_DT_INT32_REF      DataType = 103
	DataType_DT_UINT8_REF      DataType = 104
	DataType_DT_INT16_REF      DataType = 105
	DataType_DT_INT8_REF       DataType = 106
	DataType_DT_STRING_REF     DataType = 107
	DataType_DT_COMPLEX64_REF  DataType = 108
	DataType_DT_INT64_REF      DataType = 109
	DataType_DT_BOOL_REF       DataType = 110
	DataType_DT_QINT8_REF      DataType = 111
	DataType_DT_QUINT8_REF     DataType = 112
	DataType_DT_QINT32_REF     DataType = 113
	DataType_DT_BFLOAT16_REF   DataType = 114
	DataType_DT_QINT16_REF     DataType = 115
	DataType_DT_QUINT16_REF    DataType = 116
	DataType_DT_UINT16_REF     DataType = 117
	DataType_DT_COMPLEX128_REF DataType = 118
	DataType_DT_HALF_REF       DataType = 119
	DataType_DT_RESOURCE_REF   DataType = 120
)

var DataType_name = map[int32]string{
	0:   "DT_INVALID",
	1:   "DT_FLOAT",
	2:   "DT_DOUBLE",
	3:   "DT_INT32",
	4:   "DT_UINT8",
	5:   "DT_INT16",
	6:   "DT_INT8",
	7:   "DT_STRING",
	8:   "DT_COMPLEX64",
	9:   "DT_INT64",
	10:  "DT_BOOL",
	11:  "DT_QINT8",
	12:  "DT_QUINT8",
	13:  "DT_QINT32",
	14:  "DT_BFLOAT16",
	15:  "DT_QINT16",
	16:  "DT_QUINT16",
	17:  "DT_UINT16",
	18:  "DT_COMPLEX128",
	19:  "DT_HALF",
	20:  "DT_RESOURCE",
	101: "DT_FLOAT_REF",
	102: "DT_DOUBLE_REF",
	103: "DT_INT32_REF",
	104: "DT_UINT8_REF",
	105: "DT_INT16_REF",
	106: "DT_INT8_REF",
	107: "DT_STRING_REF",
	108: "DT_COMPLEX64_REF",
	109: "DT_INT64_REF",
	110: "DT_BOOL_REF",
	111: "DT_QINT8_REF",
	112: "DT_QUINT8_REF",
	113: "DT_QINT32_REF",
	114: "DT_BFLOAT16_REF",
	115: "DT_QINT16_REF",
	116: "DT_QUINT16_REF",
	117: "DT_UINT16_REF",
	118: "DT_COMPLEX128_REF",
	119: "DT_HALF_REF",
	120: "DT_RESOURCE_REF",
}
var DataType_value = map[string]int32{
	"DT_INVALID":        0,
	"DT_FLOAT":          1,
	"DT_DOUBLE":         2,
	"DT_INT32":          3,
	"DT_UINT8":          4,
	"DT_INT16":          5,
	"DT_INT8":           6,
	"DT_STRING":         7,
	"DT_COMPLEX64":      8,
	"DT_INT64":          9,
	"DT_BOOL":           10,
	"DT_QINT8":          11,
	"DT_QUINT8":         12,
	"DT_QINT32":         13,
	"DT_BFLOAT16":       14,
	"DT_QINT16":         15,
	"DT_QUINT16":        16,
	"DT_UINT16":         17,
	"DT_COMPLEX128":     18,
	"DT_HALF":           19,
	"DT_RESOURCE":       20,
	"DT_FLOAT_REF":      101,
	"DT_DOUBLE_REF":     102,
	"DT_INT32_REF":      103,
	"DT_UINT8_REF":      104,
	"DT_INT16_REF":      105,
	"DT_INT8_REF":       106,
	"DT_STRING_REF":     107,
	"DT_COMPLEX64_REF":  108,
	"DT_INT64_REF":      109,
	"DT_BOOL_REF":       110,
	"DT_QINT8_REF":      111,
	"DT_QUINT8_REF":     112,
	"DT_QINT32_REF":     113,
	"DT_BFLOAT16_REF":   114,
	"DT_QINT16_REF":     115,
	"DT_QUINT16_REF":    116,
	"DT_UINT16_REF":     117,
	"DT_COMPLEX128_REF": 118,
	"DT_HALF_REF":       119,
	"DT_RESOURCE_REF":   120,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}
func (DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

func init() {
	proto.RegisterEnum("tensorflow.DataType", DataType_name, DataType_value)
}

func init() { proto.RegisterFile("tensorflow/core/framework/types.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0xdd, 0x6e, 0xd3, 0x40,
	0x10, 0x85, 0x09, 0xd0, 0x36, 0x1d, 0xe7, 0x67, 0xb2, 0x2d, 0x12, 0x2f, 0xc0, 0x0d, 0x42, 0x89,
	0xdc, 0x56, 0x56, 0x6e, 0xe3, 0xda, 0x01, 0x4b, 0xc6, 0x1b, 0xbb, 0x1b, 0xc4, 0x9d, 0x15, 0x90,
	0x53, 0xa0, 0x3f, 0x6b, 0x6c, 0x43, 0xc8, 0xeb, 0xf1, 0x54, 0x5c, 0x22, 0x8f, 0xc7, 0x59, 0xf7,
	0x72, 0xbf, 0x99, 0x39, 0x73, 0x7c, 0x34, 0x86, 0x37, 0x55, 0xf6, 0x58, 0xea, 0x62, 0x7b, 0xaf,
	0x77, 0xb3, 0xaf, 0xba, 0xc8, 0x66, 0xdb, 0x62, 0xf3, 0x90, 0xed, 0x74, 0x71, 0x37, 0xab, 0xf6,
	0x79, 0x56, 0x4e, 0xf3, 0x42, 0x57, 0x5a, 0x80, 0x69, 0x7b, 0xfb, 0xf7, 0x08, 0xfa, 0xde, 0xa6,
	0xda, 0xa8, 0x7d, 0x9e, 0x89, 0x11, 0x80, 0xa7, 0xd2, 0x20, 0xfa, 0xb4, 0x08, 0x03, 0x0f, 0x9f,
	0x89, 0x01, 0xf4, 0x3d, 0x95, 0x2e, 0x43, 0xb9, 0x50, 0xd8, 0x13, 0x43, 0x38, 0xf5, 0x54, 0xea,
	0xc9, 0xb5, 0x1b, 0xfa, 0xf8, 0x9c, 0x8b, 0x41, 0xa4, 0x2e, 0x2f, 0xf0, 0x05, 0xbf, 0xd6, 0x41,
	0xa4, 0xe6, 0xf8, 0xd2, 0xd4, 0x6c, 0x07, 0x8f, 0x84, 0x05, 0x27, 0xcd, 0x6b, 0x8e, 0xc7, 0xac,
	0x72, 0xa3, 0x92, 0x20, 0x7a, 0x8f, 0x27, 0x02, 0x61, 0xe0, 0xa9, 0xf4, 0x5a, 0x7e, 0x5c, 0x85,
	0xfe, 0x67, 0xe7, 0x0a, 0xfb, 0x66, 0xd6, 0xb9, 0xc2, 0x53, 0x9e, 0x75, 0xa5, 0x0c, 0x11, 0xb8,
	0x14, 0x93, 0x92, 0xc5, 0x4a, 0x71, 0xb3, 0x73, 0xd0, 0x3e, 0x1b, 0x43, 0x43, 0x31, 0x06, 0xab,
	0x1e, 0x24, 0xf3, 0xb6, 0x83, 0xa3, 0x4e, 0xdd, 0x76, 0x70, 0xcc, 0xdf, 0x4a, 0xd3, 0xb6, 0x83,
	0xc8, 0x65, 0x7e, 0x4e, 0xc4, 0x04, 0x86, 0xc6, 0x97, 0x7d, 0x31, 0x47, 0xc1, 0x56, 0x3e, 0x2c,
	0xc2, 0x25, 0x9e, 0xb1, 0x7c, 0xe2, 0xdf, 0xc8, 0x75, 0x72, 0xed, 0xe3, 0x39, 0x7f, 0x08, 0xad,
	0x4b, 0x13, 0x7f, 0x89, 0x19, 0x4b, 0x34, 0x79, 0x11, 0xda, 0x72, 0x13, 0x59, 0x24, 0x72, 0xcb,
	0x84, 0xbe, 0x81, 0xc8, 0x37, 0xd3, 0x63, 0x3b, 0x44, 0xbe, 0xf3, 0xae, 0x43, 0xcb, 0x0f, 0x56,
	0x6e, 0x32, 0x24, 0x74, 0x27, 0xce, 0x01, 0xbb, 0x39, 0x12, 0xbd, 0x37, 0x5a, 0x4c, 0x1e, 0xda,
	0x58, 0xa4, 0x0c, 0x09, 0x3c, 0x72, 0x4b, 0x7c, 0x50, 0xd7, 0xac, 0x1e, 0x1b, 0x4f, 0x79, 0x8b,
	0x8c, 0xf1, 0x9f, 0xe2, 0x0c, 0xc6, 0x9d, 0x7c, 0x09, 0x16, 0x9d, 0x3e, 0x46, 0xa5, 0x10, 0x30,
	0x32, 0x39, 0x13, 0xab, 0xb8, 0xad, 0x83, 0x7e, 0x89, 0x57, 0x30, 0x79, 0x92, 0x37, 0xe1, 0xdf,
	0x6c, 0xb7, 0xce, 0x9c, 0xc0, 0x8e, 0xd7, 0xb6, 0xb9, 0x13, 0xfc, 0xe3, 0xbe, 0x83, 0xd7, 0xba,
	0xb8, 0x9d, 0x9a, 0xb3, 0x9e, 0x1e, 0x0e, 0xdf, 0xb5, 0xea, 0xcb, 0x2e, 0x57, 0xf5, 0xe1, 0x97,
	0xab, 0xde, 0xbf, 0x5e, 0xef, 0xcb, 0x31, 0xfd, 0x05, 0x97, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x14, 0x03, 0xf4, 0x32, 0x2e, 0x03, 0x00, 0x00,
}
