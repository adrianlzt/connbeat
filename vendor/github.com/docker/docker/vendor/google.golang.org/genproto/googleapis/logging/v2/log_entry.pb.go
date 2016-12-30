// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/logging/v2/log_entry.proto
// DO NOT EDIT!

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/logging/v2/log_entry.proto
	google.golang.org/genproto/googleapis/logging/v2/logging.proto
	google.golang.org/genproto/googleapis/logging/v2/logging_config.proto
	google.golang.org/genproto/googleapis/logging/v2/logging_metrics.proto

It has these top-level messages:
	LogEntry
	LogEntryOperation
	DeleteLogRequest
	WriteLogEntriesRequest
	WriteLogEntriesResponse
	ListLogEntriesRequest
	ListLogEntriesResponse
	ListMonitoredResourceDescriptorsRequest
	ListMonitoredResourceDescriptorsResponse
	LogSink
	ListSinksRequest
	ListSinksResponse
	GetSinkRequest
	CreateSinkRequest
	UpdateSinkRequest
	DeleteSinkRequest
	LogMetric
	ListLogMetricsRequest
	ListLogMetricsResponse
	GetLogMetricRequest
	CreateLogMetricRequest
	UpdateLogMetricRequest
	DeleteLogMetricRequest
*/
package v2 // import "google.golang.org/genproto/googleapis/logging/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_api3 "google.golang.org/genproto/googleapis/api/monitoredres"
import google_logging_type "google.golang.org/genproto/googleapis/logging/type"
import google_logging_type1 "google.golang.org/genproto/googleapis/logging/type"
import google_protobuf2 "github.com/golang/protobuf/ptypes/any"
import google_protobuf3 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf4 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An individual entry in a log.
type LogEntry struct {
	// Required. The resource name of the log to which this log entry
	// belongs. The format of the name is
	// `"projects/<project-id>/logs/<log-id>"`.  Examples:
	// `"projects/my-projectid/logs/syslog"`,
	// `"projects/my-projectid/logs/library.googleapis.com%2Fbook_log"`.
	//
	// The log ID part of resource name must be less than 512 characters
	// long and can only include the following characters: upper and
	// lower case alphanumeric characters: [A-Za-z0-9]; and punctuation
	// characters: forward-slash, underscore, hyphen, and period.
	// Forward-slash (`/`) characters in the log ID must be URL-encoded.
	LogName string `protobuf:"bytes,12,opt,name=log_name,json=logName" json:"log_name,omitempty"`
	// Required. The monitored resource associated with this log entry.
	// Example: a log entry that reports a database error would be
	// associated with the monitored resource designating the particular
	// database that reported the error.
	Resource *google_api3.MonitoredResource `protobuf:"bytes,8,opt,name=resource" json:"resource,omitempty"`
	// Optional. The log entry payload, which can be one of multiple types.
	//
	// Types that are valid to be assigned to Payload:
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_JsonPayload
	Payload isLogEntry_Payload `protobuf_oneof:"payload"`
	// Optional. The time the event described by the log entry occurred.  If
	// omitted, Stackdriver Logging will use the time the log entry is received.
	Timestamp *google_protobuf4.Timestamp `protobuf:"bytes,9,opt,name=timestamp" json:"timestamp,omitempty"`
	// Optional. The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity google_logging_type1.LogSeverity `protobuf:"varint,10,opt,name=severity,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// Optional. A unique ID for the log entry. If you provide this
	// field, the logging service considers other log entries in the
	// same project with the same ID as duplicates which can be removed.  If
	// omitted, Stackdriver Logging will generate a unique ID for this
	// log entry.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId" json:"insert_id,omitempty"`
	// Optional. Information about the HTTP request associated with this
	// log entry, if applicable.
	HttpRequest *google_logging_type.HttpRequest `protobuf:"bytes,7,opt,name=http_request,json=httpRequest" json:"http_request,omitempty"`
	// Optional. A set of user-defined (key, value) data that provides additional
	// information about the log entry.
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Optional. Information about an operation associated with the log entry, if
	// applicable.
	Operation *LogEntryOperation `protobuf:"bytes,15,opt,name=operation" json:"operation,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isLogEntry_Payload interface {
	isLogEntry_Payload()
}

type LogEntry_ProtoPayload struct {
	ProtoPayload *google_protobuf2.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,oneof"`
}
type LogEntry_TextPayload struct {
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,oneof"`
}
type LogEntry_JsonPayload struct {
	JsonPayload *google_protobuf3.Struct `protobuf:"bytes,6,opt,name=json_payload,json=jsonPayload,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload() {}
func (*LogEntry_TextPayload) isLogEntry_Payload()  {}
func (*LogEntry_JsonPayload) isLogEntry_Payload()  {}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *LogEntry) GetResource() *google_api3.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *LogEntry) GetProtoPayload() *google_protobuf2.Any {
	if x, ok := m.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (m *LogEntry) GetTextPayload() string {
	if x, ok := m.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (m *LogEntry) GetJsonPayload() *google_protobuf3.Struct {
	if x, ok := m.GetPayload().(*LogEntry_JsonPayload); ok {
		return x.JsonPayload
	}
	return nil
}

func (m *LogEntry) GetTimestamp() *google_protobuf4.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetHttpRequest() *google_logging_type.HttpRequest {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

func (m *LogEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LogEntry) GetOperation() *LogEntryOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogEntry_OneofMarshaler, _LogEntry_OneofUnmarshaler, _LogEntry_OneofSizer, []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_JsonPayload)(nil),
	}
}

func _LogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProtoPayload); err != nil {
			return err
		}
	case *LogEntry_TextPayload:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TextPayload)
	case *LogEntry_JsonPayload:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JsonPayload); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LogEntry.Payload has unexpected type %T", x)
	}
	return nil
}

func _LogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEntry)
	switch tag {
	case 2: // payload.proto_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf2.Any)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_ProtoPayload{msg}
		return true, err
	case 3: // payload.text_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &LogEntry_TextPayload{x}
		return true, err
	case 6: // payload.json_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf3.Struct)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_JsonPayload{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		s := proto.Size(x.ProtoPayload)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LogEntry_TextPayload:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TextPayload)))
		n += len(x.TextPayload)
	case *LogEntry_JsonPayload:
		s := proto.Size(x.JsonPayload)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Additional information about a potentially long-running operation with which
// a log entry is associated.
type LogEntryOperation struct {
	// Optional. An arbitrary operation identifier. Log entries with the
	// same identifier are assumed to be part of the same operation.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Optional. An arbitrary producer identifier. The combination of
	// `id` and `producer` must be globally unique.  Examples for `producer`:
	// `"MyDivision.MyBigCompany.com"`, `"github.com/MyProject/MyApplication"`.
	Producer string `protobuf:"bytes,2,opt,name=producer" json:"producer,omitempty"`
	// Optional. Set this to True if this is the first log entry in the operation.
	First bool `protobuf:"varint,3,opt,name=first" json:"first,omitempty"`
	// Optional. Set this to True if this is the last log entry in the operation.
	Last bool `protobuf:"varint,4,opt,name=last" json:"last,omitempty"`
}

func (m *LogEntryOperation) Reset()                    { *m = LogEntryOperation{} }
func (m *LogEntryOperation) String() string            { return proto.CompactTextString(m) }
func (*LogEntryOperation) ProtoMessage()               {}
func (*LogEntryOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*LogEntry)(nil), "google.logging.v2.LogEntry")
	proto.RegisterType((*LogEntryOperation)(nil), "google.logging.v2.LogEntryOperation")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/logging/v2/log_entry.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xfd, 0x65, 0xff, 0x9a, 0xb8, 0xdd, 0x7e, 0xcc, 0x1a, 0x22, 0x04, 0x21, 0x55, 0x03, 0x09,
	0x9e, 0x1c, 0x54, 0x84, 0xb4, 0x31, 0x24, 0xa0, 0x68, 0xd2, 0x26, 0x0d, 0x36, 0x79, 0x3c, 0x21,
	0xa1, 0x2a, 0x4d, 0x5d, 0xd7, 0x90, 0xda, 0xc1, 0x71, 0x2a, 0xf2, 0x81, 0xf8, 0x8e, 0x3c, 0x62,
	0x3b, 0x4e, 0x3a, 0x51, 0xc4, 0x36, 0xf1, 0xd0, 0xf6, 0x5e, 0xfb, 0x9c, 0x73, 0x7d, 0xcf, 0xb5,
	0x0b, 0xde, 0x50, 0x21, 0x68, 0x46, 0x10, 0x15, 0x59, 0xc2, 0x29, 0x12, 0x92, 0xc6, 0x94, 0xf0,
	0x5c, 0x0a, 0x25, 0xe2, 0x7a, 0x2b, 0xc9, 0x59, 0x11, 0x67, 0x82, 0x52, 0xc6, 0x69, 0xbc, 0x18,
	0x98, 0x70, 0x44, 0xb8, 0x92, 0x15, 0xb2, 0x28, 0xb8, 0xeb, 0x14, 0x1c, 0x04, 0x2d, 0x06, 0xd1,
	0xe9, 0xcd, 0x44, 0xf5, 0x57, 0x5c, 0x10, 0xb9, 0x60, 0x29, 0x49, 0x05, 0x9f, 0x32, 0x1a, 0x27,
	0x9c, 0x0b, 0x95, 0x28, 0x26, 0x78, 0x51, 0xab, 0x47, 0xe7, 0x37, 0x97, 0x9a, 0x0b, 0xce, 0x94,
	0x90, 0x64, 0x22, 0x49, 0xb1, 0x4c, 0x46, 0x3a, 0x13, 0xa5, 0x4c, 0x89, 0x13, 0x3c, 0xbe, 0x5d,
	0xc3, 0xaa, 0xca, 0x49, 0x3c, 0x53, 0x2a, 0xd7, 0x3a, 0xdf, 0x4a, 0x52, 0xa8, 0x7f, 0x90, 0x31,
	0xce, 0x15, 0x64, 0x41, 0x24, 0x53, 0xce, 0xbc, 0x28, 0xa6, 0x4c, 0xcd, 0xca, 0x31, 0x4a, 0xc5,
	0x3c, 0xae, 0xa5, 0x62, 0xbb, 0x31, 0x2e, 0xa7, 0x71, 0x6e, 0x48, 0xba, 0x35, 0x5e, 0x99, 0x8f,
	0x23, 0xbc, 0xb8, 0x9e, 0x50, 0x28, 0x59, 0xa6, 0xca, 0xfd, 0x38, 0xda, 0xd1, 0xf5, 0x34, 0xc5,
	0xe6, 0xba, 0xbd, 0x64, 0x9e, 0x2f, 0xa3, 0x9a, 0xbc, 0xff, 0x63, 0x13, 0xf8, 0x67, 0x82, 0x1e,
	0x9b, 0xa1, 0xc3, 0xfb, 0xc0, 0x37, 0x7d, 0xf0, 0x64, 0x4e, 0xc2, 0x5e, 0xdf, 0x7b, 0x1a, 0xe0,
	0x8e, 0xce, 0x3f, 0xe8, 0x14, 0x1e, 0x02, 0xbf, 0x31, 0x3b, 0xf4, 0xf5, 0x56, 0x77, 0xf0, 0x10,
	0x39, 0x9b, 0xb4, 0x19, 0xe8, 0x7d, 0x33, 0x12, 0xec, 0x40, 0xb8, 0x85, 0xc3, 0x23, 0xb0, 0x6d,
	0x6b, 0x8d, 0xf2, 0xa4, 0xca, 0x44, 0x32, 0x09, 0xd7, 0x2c, 0x7f, 0xaf, 0xe1, 0x37, 0x87, 0x45,
	0x6f, 0x79, 0x75, 0xf2, 0x1f, 0xee, 0xd9, 0xfc, 0xa2, 0xc6, 0xc2, 0x47, 0xa0, 0xa7, 0xc8, 0x77,
	0xd5, 0x72, 0xd7, 0xcd, 0xb1, 0x34, 0xaa, 0x6b, 0x56, 0x1b, 0xd0, 0x2b, 0xd0, 0xfb, 0x52, 0x08,
	0xde, 0x82, 0xb6, 0x6c, 0x81, 0x7b, 0x2b, 0x05, 0x2e, 0xad, 0x6d, 0x86, 0x6d, 0xe0, 0x0d, 0xfb,
	0x00, 0x04, 0xad, 0x2b, 0x61, 0x60, 0xa9, 0xd1, 0x0a, 0xf5, 0x63, 0x83, 0xc0, 0x4b, 0xb0, 0xae,
	0xeb, 0x37, 0x33, 0x0f, 0x81, 0x26, 0xee, 0x0c, 0xfa, 0xe8, 0xb7, 0x17, 0x63, 0xfc, 0x47, 0xda,
	0xe0, 0x4b, 0x87, 0xc3, 0x2d, 0x03, 0x3e, 0x00, 0x01, 0xe3, 0xfa, 0x8d, 0xa8, 0x11, 0x9b, 0x84,
	0x1b, 0xd6, 0x6e, 0xbf, 0x5e, 0x38, 0x9d, 0xc0, 0x77, 0xa0, 0x77, 0xf5, 0x66, 0x86, 0x1d, 0x7b,
	0xae, 0x3f, 0xcb, 0x9f, 0x68, 0x20, 0xae, 0x71, 0xb8, 0x3b, 0x5b, 0x26, 0xf0, 0x35, 0xd8, 0xca,
	0x92, 0x31, 0xc9, 0x8a, 0xb0, 0xdb, 0x5f, 0xd7, 0xf4, 0x27, 0x68, 0xe5, 0x3d, 0xa3, 0x66, 0xf8,
	0xe8, 0xcc, 0x22, 0x6d, 0x8c, 0x1d, 0x0d, 0x0e, 0x41, 0x20, 0x72, 0x22, 0xed, 0xab, 0x0d, 0xff,
	0xb7, 0x47, 0x78, 0xfc, 0x17, 0x8d, 0xf3, 0x06, 0x8b, 0x97, 0xb4, 0xe8, 0x10, 0x74, 0xaf, 0x48,
	0xc3, 0x3b, 0x60, 0xfd, 0x2b, 0xa9, 0x42, 0xcf, 0xf6, 0x6b, 0x42, 0xb8, 0x07, 0x36, 0x17, 0x49,
	0x56, 0x12, 0x7b, 0x2f, 0x02, 0x5c, 0x27, 0x2f, 0xd7, 0x0e, 0xbc, 0x61, 0x00, 0x3a, 0x6e, 0xa4,
	0xfb, 0x0c, 0xec, 0xae, 0x54, 0x81, 0x3b, 0x60, 0x4d, 0x5b, 0x57, 0x4b, 0xe9, 0x08, 0x46, 0xc0,
	0xd7, 0x03, 0x9b, 0x94, 0x29, 0x91, 0x4e, 0xac, 0xcd, 0x4d, 0x95, 0x29, 0x93, 0xda, 0x49, 0x73,
	0x83, 0x7c, 0x5c, 0x27, 0x10, 0x82, 0x8d, 0x2c, 0xd1, 0x8b, 0x1b, 0x76, 0xd1, 0xc6, 0xc3, 0xcf,
	0xe0, 0xae, 0x7e, 0x4a, 0xab, 0x6d, 0x0e, 0xb7, 0x9b, 0x13, 0x5c, 0xd8, 0x1b, 0xea, 0x7d, 0x7a,
	0x76, 0xdb, 0x3f, 0xd8, 0x9f, 0x9e, 0x37, 0xde, 0xb2, 0xfb, 0xcf, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x1e, 0x4b, 0x59, 0x17, 0x9e, 0x05, 0x00, 0x00,
}
