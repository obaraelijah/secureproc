// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: jobmanager.proto

// The jobmanager.v1 package includes the API for the first major
// version of the JobManager API.  This API can be extended with
// non-breaking changes. If breaking changes are required, then
// we can introduce new versions of the API alongside this one to
// maintain backward compatibility.

package jobmanagerv1

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

// The OutputStream enumeration captures the set of output stream
// the JobManager can stream from the process.
type OutputStream int32

const (
	// The unset value
	OutputStream_UNSET OutputStream = 0
	// Stream the process’ standard output stream
	OutputStream_STDOUT OutputStream = 1
	// Stream the process’ standard error stream
	OutputStream_STDERR OutputStream = 2
)

// Enum value maps for OutputStream.
var (
	OutputStream_name = map[int32]string{
		0: "UNSET",
		1: "STDOUT",
		2: "STDERR",
	}
	OutputStream_value = map[string]int32{
		"UNSET":  0,
		"STDOUT": 1,
		"STDERR": 2,
	}
)

func (x OutputStream) Enum() *OutputStream {
	p := new(OutputStream)
	*p = x
	return p
}

func (x OutputStream) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OutputStream) Descriptor() protoreflect.EnumDescriptor {
	return file_jobmanager_proto_enumTypes[0].Descriptor()
}

func (OutputStream) Type() protoreflect.EnumType {
	return &file_jobmanager_proto_enumTypes[0]
}

func (x OutputStream) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OutputStream.Descriptor instead.
func (OutputStream) EnumDescriptor() ([]byte, []int) {
	return file_jobmanager_proto_rawDescGZIP(), []int{0}
}

// A JobCreationRequest is a message that clients use to request
// the service to create a new Job.  Possible extensions to this
// would enable clients to include additional metadata (e.g., labels)
// to be associated with the newly-created jobs.
type JobCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A client-specified name for the job.  The user cannot have
	// any other current job with the same name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The path of the program to run
	ProgramPath string `protobuf:"bytes,2,opt,name=programPath,proto3" json:"programPath,omitempty"`
	// Arguments to pass to the the program
	Arguments []string `protobuf:"bytes,3,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *JobCreationRequest) Reset() {
	*x = JobCreationRequest{}
	mi := &file_jobmanager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobCreationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobCreationRequest) ProtoMessage() {}

func (x *JobCreationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobmanager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobCreationRequest.ProtoReflect.Descriptor instead.
func (*JobCreationRequest) Descriptor() ([]byte, []int) {
	return file_jobmanager_proto_rawDescGZIP(), []int{0}
}

func (x *JobCreationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobCreationRequest) GetProgramPath() string {
	if x != nil {
		return x.ProgramPath
	}
	return ""
}

func (x *JobCreationRequest) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

// A JobID is a message that client use to uniquely identify a job
// managed by the JobManager.
type JobID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A server-assigned UUIDv4 for the job
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *JobID) Reset() {
	*x = JobID{}
	mi := &file_jobmanager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobID) ProtoMessage() {}

func (x *JobID) ProtoReflect() protoreflect.Message {
	mi := &file_jobmanager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobID.ProtoReflect.Descriptor instead.
func (*JobID) Descriptor() ([]byte, []int) {
	return file_jobmanager_proto_rawDescGZIP(), []int{1}
}

func (x *JobID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// A Job is a message that uniquely identifies one of the Jobs managed
// by the JobManager service.
type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The server-assigned ID
	Id *JobID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The client-specified name for the job
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	mi := &file_jobmanager_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_jobmanager_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_jobmanager_proto_rawDescGZIP(), []int{2}
}

func (x *Job) GetId() *JobID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The JobStatus message is used to communicate the status of a job.
type JobStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Job with which this status is associated
	Job *Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	// The user who started the job
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// Is the job running?
	IsRunning bool `protobuf:"varint,3,opt,name=isRunning,proto3" json:"isRunning,omitempty"`
	// The process ID of the job
	Pid int32 `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
	// If the job is not running and was terminated by a signal, which
	// signal?  If the process is terminated as a result of the Stop
	// API, then this will indicate the KILL signal (9).  If the
	// process was terminated by an extraneous signal, then this field
	// will indicate the signal number that was received.
	SignalNumber int32 `protobuf:"varint,5,opt,name=signalNumber,proto3" json:"signalNumber,omitempty"`
	// If the job is not running, what was its exit code?
	ExitCode int32 `protobuf:"varint,6,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	// If a job failed to start, what was the cause of the failure?
	ErrorMessage string `protobuf:"bytes,7,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *JobStatus) Reset() {
	*x = JobStatus{}
	mi := &file_jobmanager_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatus) ProtoMessage() {}

func (x *JobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_jobmanager_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatus.ProtoReflect.Descriptor instead.
func (*JobStatus) Descriptor() ([]byte, []int) {
	return file_jobmanager_proto_rawDescGZIP(), []int{3}
}

func (x *JobStatus) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *JobStatus) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *JobStatus) GetIsRunning() bool {
	if x != nil {
		return x.IsRunning
	}
	return false
}

func (x *JobStatus) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *JobStatus) GetSignalNumber() int32 {
	if x != nil {
		return x.SignalNumber
	}
	return 0
}

func (x *JobStatus) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *JobStatus) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// The JobOutput message is used to stream the output of the command.
// This message can be enhanced in the future to include information
// about the byte offset into the output if this information would
// be useful to clients.
type JobOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of bytes corresponding to the next “chunk” of command
	// output (either stdout or stderr).
	Output []byte `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *JobOutput) Reset() {
	*x = JobOutput{}
	mi := &file_jobmanager_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOutput) ProtoMessage() {}

func (x *JobOutput) ProtoReflect() protoreflect.Message {
	mi := &file_jobmanager_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOutput.ProtoReflect.Descriptor instead.
func (*JobOutput) Descriptor() ([]byte, []int) {
	return file_jobmanager_proto_rawDescGZIP(), []int{4}
}

func (x *JobOutput) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

// The JobStatusList message is used to communicate the list of jobs
// managed by the JobManager and their status.
type JobStatusList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobStatusList []*JobStatus `protobuf:"bytes,1,rep,name=jobStatusList,proto3" json:"jobStatusList,omitempty"`
}

func (x *JobStatusList) Reset() {
	*x = JobStatusList{}
	mi := &file_jobmanager_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobStatusList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusList) ProtoMessage() {}

func (x *JobStatusList) ProtoReflect() protoreflect.Message {
	mi := &file_jobmanager_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusList.ProtoReflect.Descriptor instead.
func (*JobStatusList) Descriptor() ([]byte, []int) {
	return file_jobmanager_proto_rawDescGZIP(), []int{5}
}

func (x *JobStatusList) GetJobStatusList() []*JobStatus {
	if x != nil {
		return x.JobStatusList
	}
	return nil
}

// The StreamOutputRequest message is used to request that the service
// stream the output of the job.
type StreamOutputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The server-assigned ID
	JobID *JobID `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	// The output stream to return
	OutputStream OutputStream `protobuf:"varint,2,opt,name=outputStream,proto3,enum=jobmanager.v1.OutputStream" json:"outputStream,omitempty"`
}

func (x *StreamOutputRequest) Reset() {
	*x = StreamOutputRequest{}
	mi := &file_jobmanager_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamOutputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamOutputRequest) ProtoMessage() {}

func (x *StreamOutputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jobmanager_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamOutputRequest.ProtoReflect.Descriptor instead.
func (*StreamOutputRequest) Descriptor() ([]byte, []int) {
	return file_jobmanager_proto_rawDescGZIP(), []int{6}
}

func (x *StreamOutputRequest) GetJobID() *JobID {
	if x != nil {
		return x.JobID
	}
	return nil
}

func (x *StreamOutputRequest) GetOutputStream() OutputStream {
	if x != nil {
		return x.OutputStream
	}
	return OutputStream_UNSET
}

// The NilMessage message is used when no other message is needed.
type NilMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NilMessage) Reset() {
	*x = NilMessage{}
	mi := &file_jobmanager_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NilMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NilMessage) ProtoMessage() {}

func (x *NilMessage) ProtoReflect() protoreflect.Message {
	mi := &file_jobmanager_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NilMessage.ProtoReflect.Descriptor instead.
func (*NilMessage) Descriptor() ([]byte, []int) {
	return file_jobmanager_proto_rawDescGZIP(), []int{7}
}

var File_jobmanager_proto protoreflect.FileDescriptor

var file_jobmanager_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x22, 0x68, 0x0a, 0x12, 0x4a, 0x6f, 0x62, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x17, 0x0a, 0x05, 0x4a,
	0x6f, 0x62, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x24, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xdb, 0x01, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4f, 0x0a, 0x0d, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x6a, 0x6f, 0x62,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x6a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x13, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x3f, 0x0a,
	0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x0c,
	0x0a, 0x0a, 0x4e, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x31, 0x0a, 0x0c,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x09, 0x0a, 0x05,
	0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x44, 0x4f, 0x55,
	0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x44, 0x45, 0x52, 0x52, 0x10, 0x02, 0x32,
	0xd9, 0x02, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x40,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x21, 0x2e, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6a, 0x6f, 0x62,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x2e, 0x6a, 0x6f, 0x62, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x1a, 0x19,
	0x2e, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x2e, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x6a, 0x6f, 0x62,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19,
	0x2e, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x2e, 0x6a, 0x6f, 0x62, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0c, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x22, 0x2e, 0x6a, 0x6f, 0x62, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f,
	0x62, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x43, 0x5a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x62, 0x61, 0x72, 0x61, 0x65,
	0x6c, 0x69, 0x6a, 0x61, 0x68, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2f, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jobmanager_proto_rawDescOnce sync.Once
	file_jobmanager_proto_rawDescData = file_jobmanager_proto_rawDesc
)

func file_jobmanager_proto_rawDescGZIP() []byte {
	file_jobmanager_proto_rawDescOnce.Do(func() {
		file_jobmanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_jobmanager_proto_rawDescData)
	})
	return file_jobmanager_proto_rawDescData
}

var file_jobmanager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_jobmanager_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_jobmanager_proto_goTypes = []any{
	(OutputStream)(0),           // 0: jobmanager.v1.OutputStream
	(*JobCreationRequest)(nil),  // 1: jobmanager.v1.JobCreationRequest
	(*JobID)(nil),               // 2: jobmanager.v1.JobID
	(*Job)(nil),                 // 3: jobmanager.v1.Job
	(*JobStatus)(nil),           // 4: jobmanager.v1.JobStatus
	(*JobOutput)(nil),           // 5: jobmanager.v1.JobOutput
	(*JobStatusList)(nil),       // 6: jobmanager.v1.JobStatusList
	(*StreamOutputRequest)(nil), // 7: jobmanager.v1.StreamOutputRequest
	(*NilMessage)(nil),          // 8: jobmanager.v1.NilMessage
}
var file_jobmanager_proto_depIdxs = []int32{
	2,  // 0: jobmanager.v1.Job.id:type_name -> jobmanager.v1.JobID
	3,  // 1: jobmanager.v1.JobStatus.job:type_name -> jobmanager.v1.Job
	4,  // 2: jobmanager.v1.JobStatusList.jobStatusList:type_name -> jobmanager.v1.JobStatus
	2,  // 3: jobmanager.v1.StreamOutputRequest.jobID:type_name -> jobmanager.v1.JobID
	0,  // 4: jobmanager.v1.StreamOutputRequest.outputStream:type_name -> jobmanager.v1.OutputStream
	1,  // 5: jobmanager.v1.JobManager.Start:input_type -> jobmanager.v1.JobCreationRequest
	2,  // 6: jobmanager.v1.JobManager.Stop:input_type -> jobmanager.v1.JobID
	2,  // 7: jobmanager.v1.JobManager.Query:input_type -> jobmanager.v1.JobID
	8,  // 8: jobmanager.v1.JobManager.List:input_type -> jobmanager.v1.NilMessage
	7,  // 9: jobmanager.v1.JobManager.StreamOutput:input_type -> jobmanager.v1.StreamOutputRequest
	3,  // 10: jobmanager.v1.JobManager.Start:output_type -> jobmanager.v1.Job
	8,  // 11: jobmanager.v1.JobManager.Stop:output_type -> jobmanager.v1.NilMessage
	4,  // 12: jobmanager.v1.JobManager.Query:output_type -> jobmanager.v1.JobStatus
	6,  // 13: jobmanager.v1.JobManager.List:output_type -> jobmanager.v1.JobStatusList
	5,  // 14: jobmanager.v1.JobManager.StreamOutput:output_type -> jobmanager.v1.JobOutput
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_jobmanager_proto_init() }
func file_jobmanager_proto_init() {
	if File_jobmanager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_jobmanager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jobmanager_proto_goTypes,
		DependencyIndexes: file_jobmanager_proto_depIdxs,
		EnumInfos:         file_jobmanager_proto_enumTypes,
		MessageInfos:      file_jobmanager_proto_msgTypes,
	}.Build()
	File_jobmanager_proto = out.File
	file_jobmanager_proto_rawDesc = nil
	file_jobmanager_proto_goTypes = nil
	file_jobmanager_proto_depIdxs = nil
}