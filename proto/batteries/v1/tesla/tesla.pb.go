// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/batteries/v1/tesla/tesla.proto

package tesla

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

type TeslaBatteryState struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	SiteId             string                 `protobuf:"bytes,1,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
	BackupReserve      int32                  `protobuf:"varint,2,opt,name=backup_reserve,json=backupReserve,proto3" json:"backup_reserve,omitempty"`
	BatteryCharging    bool                   `protobuf:"varint,3,opt,name=battery_charging,json=batteryCharging,proto3" json:"battery_charging,omitempty"`
	BatteryDischarging bool                   `protobuf:"varint,4,opt,name=battery_discharging,json=batteryDischarging,proto3" json:"battery_discharging,omitempty"`
	ConsumptionW       int32                  `protobuf:"varint,5,opt,name=consumption_w,json=consumptionW,proto3" json:"consumption_w,omitempty"`
	GridImportW        int32                  `protobuf:"varint,6,opt,name=grid_import_w,json=gridImportW,proto3" json:"grid_import_w,omitempty"`
	GridExportW        int32                  `protobuf:"varint,7,opt,name=grid_export_w,json=gridExportW,proto3" json:"grid_export_w,omitempty"`
	ProductionW        int32                  `protobuf:"varint,8,opt,name=production_w,json=productionW,proto3" json:"production_w,omitempty"`
	StorageLevel       float64                `protobuf:"fixed64,9,opt,name=storage_level,json=storageLevel,proto3" json:"storage_level,omitempty"`
	SystemStatus       string                 `protobuf:"bytes,10,opt,name=system_status,json=systemStatus,proto3" json:"system_status,omitempty"`
	StormModeEnabled   bool                   `protobuf:"varint,11,opt,name=storm_mode_enabled,json=stormModeEnabled,proto3" json:"storm_mode_enabled,omitempty"`
	OperationMode      string                 `protobuf:"bytes,12,opt,name=operation_mode,json=operationMode,proto3" json:"operation_mode,omitempty"`
	LastUpdated        string                 `protobuf:"bytes,13,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *TeslaBatteryState) Reset() {
	*x = TeslaBatteryState{}
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeslaBatteryState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeslaBatteryState) ProtoMessage() {}

func (x *TeslaBatteryState) ProtoReflect() protoreflect.Message {
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeslaBatteryState.ProtoReflect.Descriptor instead.
func (*TeslaBatteryState) Descriptor() ([]byte, []int) {
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP(), []int{0}
}

func (x *TeslaBatteryState) GetSiteId() string {
	if x != nil {
		return x.SiteId
	}
	return ""
}

func (x *TeslaBatteryState) GetBackupReserve() int32 {
	if x != nil {
		return x.BackupReserve
	}
	return 0
}

func (x *TeslaBatteryState) GetBatteryCharging() bool {
	if x != nil {
		return x.BatteryCharging
	}
	return false
}

func (x *TeslaBatteryState) GetBatteryDischarging() bool {
	if x != nil {
		return x.BatteryDischarging
	}
	return false
}

func (x *TeslaBatteryState) GetConsumptionW() int32 {
	if x != nil {
		return x.ConsumptionW
	}
	return 0
}

func (x *TeslaBatteryState) GetGridImportW() int32 {
	if x != nil {
		return x.GridImportW
	}
	return 0
}

func (x *TeslaBatteryState) GetGridExportW() int32 {
	if x != nil {
		return x.GridExportW
	}
	return 0
}

func (x *TeslaBatteryState) GetProductionW() int32 {
	if x != nil {
		return x.ProductionW
	}
	return 0
}

func (x *TeslaBatteryState) GetStorageLevel() float64 {
	if x != nil {
		return x.StorageLevel
	}
	return 0
}

func (x *TeslaBatteryState) GetSystemStatus() string {
	if x != nil {
		return x.SystemStatus
	}
	return ""
}

func (x *TeslaBatteryState) GetStormModeEnabled() bool {
	if x != nil {
		return x.StormModeEnabled
	}
	return false
}

func (x *TeslaBatteryState) GetOperationMode() string {
	if x != nil {
		return x.OperationMode
	}
	return ""
}

func (x *TeslaBatteryState) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

type GetTeslaBatteryStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTeslaBatteryStatusRequest) Reset() {
	*x = GetTeslaBatteryStatusRequest{}
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTeslaBatteryStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeslaBatteryStatusRequest) ProtoMessage() {}

func (x *GetTeslaBatteryStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeslaBatteryStatusRequest.ProtoReflect.Descriptor instead.
func (*GetTeslaBatteryStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP(), []int{1}
}

type UpdateBackupReserveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Percentage    int32                  `protobuf:"varint,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBackupReserveRequest) Reset() {
	*x = UpdateBackupReserveRequest{}
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBackupReserveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBackupReserveRequest) ProtoMessage() {}

func (x *UpdateBackupReserveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBackupReserveRequest.ProtoReflect.Descriptor instead.
func (*UpdateBackupReserveRequest) Descriptor() ([]byte, []int) {
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBackupReserveRequest) GetPercentage() int32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

type UpdateBackupReserveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Percentage    int32                  `protobuf:"varint,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBackupReserveResponse) Reset() {
	*x = UpdateBackupReserveResponse{}
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBackupReserveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBackupReserveResponse) ProtoMessage() {}

func (x *UpdateBackupReserveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBackupReserveResponse.ProtoReflect.Descriptor instead.
func (*UpdateBackupReserveResponse) Descriptor() ([]byte, []int) {
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBackupReserveResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateBackupReserveResponse) GetPercentage() int32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

type ToggleStormModeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enable        bool                   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToggleStormModeRequest) Reset() {
	*x = ToggleStormModeRequest{}
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToggleStormModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleStormModeRequest) ProtoMessage() {}

func (x *ToggleStormModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleStormModeRequest.ProtoReflect.Descriptor instead.
func (*ToggleStormModeRequest) Descriptor() ([]byte, []int) {
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP(), []int{4}
}

func (x *ToggleStormModeRequest) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

type ToggleStormModeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Enabled       bool                   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToggleStormModeResponse) Reset() {
	*x = ToggleStormModeResponse{}
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToggleStormModeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleStormModeResponse) ProtoMessage() {}

func (x *ToggleStormModeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleStormModeResponse.ProtoReflect.Descriptor instead.
func (*ToggleStormModeResponse) Descriptor() ([]byte, []int) {
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP(), []int{5}
}

func (x *ToggleStormModeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ToggleStormModeResponse) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type SetOperationModeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mode          string                 `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetOperationModeRequest) Reset() {
	*x = SetOperationModeRequest{}
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetOperationModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOperationModeRequest) ProtoMessage() {}

func (x *SetOperationModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOperationModeRequest.ProtoReflect.Descriptor instead.
func (*SetOperationModeRequest) Descriptor() ([]byte, []int) {
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP(), []int{6}
}

func (x *SetOperationModeRequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type SetOperationModeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Mode          string                 `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetOperationModeResponse) Reset() {
	*x = SetOperationModeResponse{}
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetOperationModeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOperationModeResponse) ProtoMessage() {}

func (x *SetOperationModeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOperationModeResponse.ProtoReflect.Descriptor instead.
func (*SetOperationModeResponse) Descriptor() ([]byte, []int) {
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP(), []int{7}
}

func (x *SetOperationModeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SetOperationModeResponse) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type GetPowerUsageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPowerUsageRequest) Reset() {
	*x = GetPowerUsageRequest{}
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPowerUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerUsageRequest) ProtoMessage() {}

func (x *GetPowerUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerUsageRequest.ProtoReflect.Descriptor instead.
func (*GetPowerUsageRequest) Descriptor() ([]byte, []int) {
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP(), []int{8}
}

type GetPowerUsageResponse struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	GridImportW         int32                  `protobuf:"varint,1,opt,name=grid_import_w,json=gridImportW,proto3" json:"grid_import_w,omitempty"`
	GridExportW         int32                  `protobuf:"varint,2,opt,name=grid_export_w,json=gridExportW,proto3" json:"grid_export_w,omitempty"`
	BatteryStorageLevel float64                `protobuf:"fixed64,3,opt,name=battery_storage_level,json=batteryStorageLevel,proto3" json:"battery_storage_level,omitempty"`
	Timestamp           string                 `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *GetPowerUsageResponse) Reset() {
	*x = GetPowerUsageResponse{}
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPowerUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerUsageResponse) ProtoMessage() {}

func (x *GetPowerUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_batteries_v1_tesla_tesla_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerUsageResponse.ProtoReflect.Descriptor instead.
func (*GetPowerUsageResponse) Descriptor() ([]byte, []int) {
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP(), []int{9}
}

func (x *GetPowerUsageResponse) GetGridImportW() int32 {
	if x != nil {
		return x.GridImportW
	}
	return 0
}

func (x *GetPowerUsageResponse) GetGridExportW() int32 {
	if x != nil {
		return x.GridExportW
	}
	return 0
}

func (x *GetPowerUsageResponse) GetBatteryStorageLevel() float64 {
	if x != nil {
		return x.BatteryStorageLevel
	}
	return 0
}

func (x *GetPowerUsageResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_proto_batteries_v1_tesla_tesla_proto protoreflect.FileDescriptor

var file_proto_batteries_v1_tesla_tesla_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x6c, 0x61, 0x2f, 0x74, 0x65, 0x73, 0x6c, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0x81, 0x04, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x6c, 0x61, 0x42, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x43, 0x68, 0x61,
	0x72, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x13, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x5f, 0x64, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x44, 0x69, 0x73, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x12, 0x22, 0x0a, 0x0d, 0x67,
	0x72, 0x69, 0x64, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x77, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x67, 0x72, 0x69, 0x64, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x12,
	0x22, 0x0a, 0x0d, 0x67, 0x72, 0x69, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x77,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x72, 0x69, 0x64, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x57, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x6d, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x74,
	0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x73, 0x6c, 0x61, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x22,
	0x30, 0x0a, 0x16, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x6d, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x22, 0x4d, 0x0a, 0x17, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x6d,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x22, 0x2d, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22,
	0x48, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xb1, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x67,
	0x72, 0x69, 0x64, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x67, 0x72, 0x69, 0x64, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x12,
	0x22, 0x0a, 0x0d, 0x67, 0x72, 0x69, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x77,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x72, 0x69, 0x64, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x57, 0x12, 0x32, 0x0a, 0x15, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x13, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x8e, 0x04, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x6c, 0x61, 0x42,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x6c, 0x61, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x2e, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x6c, 0x61, 0x42, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x73, 0x6c, 0x61, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x12, 0x28, 0x2e, 0x62,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0f, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x2e, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x67, 0x67,
	0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x2e, 0x62, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x62, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x62, 0x74, 0x72, 0x61, 0x64, 0x2f, 0x67, 0x6f, 0x5f, 0x62,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x6c,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_batteries_v1_tesla_tesla_proto_rawDescOnce sync.Once
	file_proto_batteries_v1_tesla_tesla_proto_rawDescData []byte
)

func file_proto_batteries_v1_tesla_tesla_proto_rawDescGZIP() []byte {
	file_proto_batteries_v1_tesla_tesla_proto_rawDescOnce.Do(func() {
		file_proto_batteries_v1_tesla_tesla_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_batteries_v1_tesla_tesla_proto_rawDesc), len(file_proto_batteries_v1_tesla_tesla_proto_rawDesc)))
	})
	return file_proto_batteries_v1_tesla_tesla_proto_rawDescData
}

var file_proto_batteries_v1_tesla_tesla_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_batteries_v1_tesla_tesla_proto_goTypes = []any{
	(*TeslaBatteryState)(nil),            // 0: batteries.v1.TeslaBatteryState
	(*GetTeslaBatteryStatusRequest)(nil), // 1: batteries.v1.GetTeslaBatteryStatusRequest
	(*UpdateBackupReserveRequest)(nil),   // 2: batteries.v1.UpdateBackupReserveRequest
	(*UpdateBackupReserveResponse)(nil),  // 3: batteries.v1.UpdateBackupReserveResponse
	(*ToggleStormModeRequest)(nil),       // 4: batteries.v1.ToggleStormModeRequest
	(*ToggleStormModeResponse)(nil),      // 5: batteries.v1.ToggleStormModeResponse
	(*SetOperationModeRequest)(nil),      // 6: batteries.v1.SetOperationModeRequest
	(*SetOperationModeResponse)(nil),     // 7: batteries.v1.SetOperationModeResponse
	(*GetPowerUsageRequest)(nil),         // 8: batteries.v1.GetPowerUsageRequest
	(*GetPowerUsageResponse)(nil),        // 9: batteries.v1.GetPowerUsageResponse
}
var file_proto_batteries_v1_tesla_tesla_proto_depIdxs = []int32{
	1, // 0: batteries.v1.TeslaBatteryService.GetTeslaBatteryStatus:input_type -> batteries.v1.GetTeslaBatteryStatusRequest
	2, // 1: batteries.v1.TeslaBatteryService.UpdateBackupReserve:input_type -> batteries.v1.UpdateBackupReserveRequest
	4, // 2: batteries.v1.TeslaBatteryService.ToggleStormMode:input_type -> batteries.v1.ToggleStormModeRequest
	6, // 3: batteries.v1.TeslaBatteryService.SetOperationMode:input_type -> batteries.v1.SetOperationModeRequest
	8, // 4: batteries.v1.TeslaBatteryService.GetPowerUsage:input_type -> batteries.v1.GetPowerUsageRequest
	0, // 5: batteries.v1.TeslaBatteryService.GetTeslaBatteryStatus:output_type -> batteries.v1.TeslaBatteryState
	3, // 6: batteries.v1.TeslaBatteryService.UpdateBackupReserve:output_type -> batteries.v1.UpdateBackupReserveResponse
	5, // 7: batteries.v1.TeslaBatteryService.ToggleStormMode:output_type -> batteries.v1.ToggleStormModeResponse
	7, // 8: batteries.v1.TeslaBatteryService.SetOperationMode:output_type -> batteries.v1.SetOperationModeResponse
	9, // 9: batteries.v1.TeslaBatteryService.GetPowerUsage:output_type -> batteries.v1.GetPowerUsageResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_batteries_v1_tesla_tesla_proto_init() }
func file_proto_batteries_v1_tesla_tesla_proto_init() {
	if File_proto_batteries_v1_tesla_tesla_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_batteries_v1_tesla_tesla_proto_rawDesc), len(file_proto_batteries_v1_tesla_tesla_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_batteries_v1_tesla_tesla_proto_goTypes,
		DependencyIndexes: file_proto_batteries_v1_tesla_tesla_proto_depIdxs,
		MessageInfos:      file_proto_batteries_v1_tesla_tesla_proto_msgTypes,
	}.Build()
	File_proto_batteries_v1_tesla_tesla_proto = out.File
	file_proto_batteries_v1_tesla_tesla_proto_goTypes = nil
	file_proto_batteries_v1_tesla_tesla_proto_depIdxs = nil
}
