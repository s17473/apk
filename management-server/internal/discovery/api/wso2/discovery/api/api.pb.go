//
//  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: api/proto/wso2/discovery/api/api.proto

package api

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

type API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid           string      `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Provider       string      `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Version        string      `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Name           string      `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Context        string      `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
	Type           string      `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	OrganizationId string      `protobuf:"bytes,7,opt,name=organizationId,proto3" json:"organizationId,omitempty"`
	CreatedBy      string      `protobuf:"bytes,8,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	CreatedTime    string      `protobuf:"bytes,9,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedBy      string      `protobuf:"bytes,10,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	UpdatedTime    string      `protobuf:"bytes,11,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	Definition     string      `protobuf:"bytes,12,opt,name=definition,proto3" json:"definition,omitempty"`
	Transports     []string    `protobuf:"bytes,13,rep,name=transports,proto3" json:"transports,omitempty"`
	Resources      []*Resource `protobuf:"bytes,16,rep,name=resources,proto3" json:"resources,omitempty"`
	CorsConfig     *CorsConfig `protobuf:"bytes,17,opt,name=corsConfig,proto3" json:"corsConfig,omitempty"`
}

func (x *API) Reset() {
	*x = API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*API) ProtoMessage() {}

func (x *API) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use API.ProtoReflect.Descriptor instead.
func (*API) Descriptor() ([]byte, []int) {
	return file_api_proto_wso2_discovery_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *API) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *API) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *API) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *API) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *API) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *API) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *API) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *API) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *API) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *API) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *API) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *API) GetDefinition() string {
	if x != nil {
		return x.Definition
	}
	return ""
}

func (x *API) GetTransports() []string {
	if x != nil {
		return x.Transports
	}
	return nil
}

func (x *API) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *API) GetCorsConfig() *CorsConfig {
	if x != nil {
		return x.CorsConfig
	}
	return nil
}

type CorsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorsConfigurationEnabled      bool     `protobuf:"varint,1,opt,name=corsConfigurationEnabled,proto3" json:"corsConfigurationEnabled,omitempty"`
	AccessControlAllowOrigins     []string `protobuf:"bytes,2,rep,name=accessControlAllowOrigins,proto3" json:"accessControlAllowOrigins,omitempty"`
	AccessControlAllowCredentials bool     `protobuf:"varint,3,opt,name=accessControlAllowCredentials,proto3" json:"accessControlAllowCredentials,omitempty"`
	AccessControlAllowHeaders     []string `protobuf:"bytes,4,rep,name=accessControlAllowHeaders,proto3" json:"accessControlAllowHeaders,omitempty"`
	AccessControlAllowMethods     []string `protobuf:"bytes,5,rep,name=accessControlAllowMethods,proto3" json:"accessControlAllowMethods,omitempty"`
}

func (x *CorsConfig) Reset() {
	*x = CorsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorsConfig) ProtoMessage() {}

func (x *CorsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorsConfig.ProtoReflect.Descriptor instead.
func (*CorsConfig) Descriptor() ([]byte, []int) {
	return file_api_proto_wso2_discovery_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *CorsConfig) GetCorsConfigurationEnabled() bool {
	if x != nil {
		return x.CorsConfigurationEnabled
	}
	return false
}

func (x *CorsConfig) GetAccessControlAllowOrigins() []string {
	if x != nil {
		return x.AccessControlAllowOrigins
	}
	return nil
}

func (x *CorsConfig) GetAccessControlAllowCredentials() bool {
	if x != nil {
		return x.AccessControlAllowCredentials
	}
	return false
}

func (x *CorsConfig) GetAccessControlAllowHeaders() []string {
	if x != nil {
		return x.AccessControlAllowHeaders
	}
	return nil
}

func (x *CorsConfig) GetAccessControlAllowMethods() []string {
	if x != nil {
		return x.AccessControlAllowMethods
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path              string             `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Verb              string             `protobuf:"bytes,2,opt,name=verb,proto3" json:"verb,omitempty"`
	Authentications   []*Authentication  `protobuf:"bytes,3,rep,name=authentications,proto3" json:"authentications,omitempty"`
	Scopes            []*Scope           `protobuf:"bytes,4,rep,name=scopes,proto3" json:"scopes,omitempty"`
	OperationPolicies *OperationPolicies `protobuf:"bytes,6,opt,name=operationPolicies,proto3" json:"operationPolicies,omitempty"`
	QueryParams       []*QueryParam      `protobuf:"bytes,7,rep,name=queryParams,proto3" json:"queryParams,omitempty"`
	Hostname          []string           `protobuf:"bytes,8,rep,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_api_proto_wso2_discovery_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *Resource) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Resource) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *Resource) GetAuthentications() []*Authentication {
	if x != nil {
		return x.Authentications
	}
	return nil
}

func (x *Resource) GetScopes() []*Scope {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *Resource) GetOperationPolicies() *OperationPolicies {
	if x != nil {
		return x.OperationPolicies
	}
	return nil
}

func (x *Resource) GetQueryParams() []*QueryParam {
	if x != nil {
		return x.QueryParams
	}
	return nil
}

func (x *Resource) GetHostname() []string {
	if x != nil {
		return x.Hostname
	}
	return nil
}

type OperationPolicies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OperationPolicies) Reset() {
	*x = OperationPolicies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationPolicies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationPolicies) ProtoMessage() {}

func (x *OperationPolicies) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationPolicies.ProtoReflect.Descriptor instead.
func (*OperationPolicies) Descriptor() ([]byte, []int) {
	return file_api_proto_wso2_discovery_api_api_proto_rawDescGZIP(), []int{3}
}

type QueryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *QueryParam) Reset() {
	*x = QueryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParam) ProtoMessage() {}

func (x *QueryParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParam.ProtoReflect.Descriptor instead.
func (*QueryParam) Descriptor() ([]byte, []int) {
	return file_api_proto_wso2_discovery_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *QueryParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryParam) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Scope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName string   `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Bindings    []string `protobuf:"bytes,4,rep,name=bindings,proto3" json:"bindings,omitempty"`
}

func (x *Scope) Reset() {
	*x = Scope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scope) ProtoMessage() {}

func (x *Scope) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scope.ProtoReflect.Descriptor instead.
func (*Scope) Descriptor() ([]byte, []int) {
	return file_api_proto_wso2_discovery_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *Scope) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scope) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Scope) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Scope) GetBindings() []string {
	if x != nil {
		return x.Bindings
	}
	return nil
}

type Authentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           string        `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Iss            string        `protobuf:"bytes,2,opt,name=iss,proto3" json:"iss,omitempty"`
	Aud            string        `protobuf:"bytes,3,opt,name=aud,proto3" json:"aud,omitempty"`
	JwksUri        string        `protobuf:"bytes,4,opt,name=jwksUri,proto3" json:"jwksUri,omitempty"`
	CredentialList []*Credential `protobuf:"bytes,5,rep,name=credentialList,proto3" json:"credentialList,omitempty"`
}

func (x *Authentication) Reset() {
	*x = Authentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authentication) ProtoMessage() {}

func (x *Authentication) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authentication.ProtoReflect.Descriptor instead.
func (*Authentication) Descriptor() ([]byte, []int) {
	return file_api_proto_wso2_discovery_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *Authentication) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Authentication) GetIss() string {
	if x != nil {
		return x.Iss
	}
	return ""
}

func (x *Authentication) GetAud() string {
	if x != nil {
		return x.Aud
	}
	return ""
}

func (x *Authentication) GetJwksUri() string {
	if x != nil {
		return x.JwksUri
	}
	return ""
}

func (x *Authentication) GetCredentialList() []*Credential {
	if x != nil {
		return x.CredentialList
	}
	return nil
}

type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_api_proto_wso2_discovery_api_api_proto_rawDescGZIP(), []int{7}
}

func (x *Credential) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Credential) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wso2_discovery_api_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_api_proto_wso2_discovery_api_api_proto_rawDescGZIP(), []int{8}
}

func (x *Response) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_api_proto_wso2_discovery_api_api_proto protoreflect.FileDescriptor

var file_api_proto_wso2_discovery_api_api_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x73, 0x6f, 0x32,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x22, 0xf5, 0x03, 0x0a,
	0x03, 0x41, 0x50, 0x49, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x73,
	0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f,
	0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x63, 0x6f, 0x72, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0xc8, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a, 0x18, 0x63, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x63, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x3c, 0x0a, 0x19, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x19, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x44, 0x0a,
	0x1d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x19, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x19, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x3c, 0x0a, 0x19, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x19, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x22,
	0xe6, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x76, 0x65, 0x72, 0x62, 0x12, 0x4c, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x06, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x53, 0x0a, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52,
	0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x22, 0x36, 0x0a,
	0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7b, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x75, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x75, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6a, 0x77, 0x6b, 0x73, 0x55, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6a, 0x77, 0x6b, 0x73, 0x55, 0x72, 0x69, 0x12, 0x46, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x0e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x44, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xd8, 0x01, 0x0a, 0x0a, 0x41, 0x50,
	0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x50, 0x49, 0x12, 0x17, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x50, 0x49, 0x1a, 0x1c,
	0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x12, 0x17, 0x2e, 0x77, 0x73, 0x6f, 0x32,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x50, 0x49, 0x1a, 0x1c, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x50, 0x49, 0x12, 0x17, 0x2e,
	0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x50, 0x49, 0x1a, 0x1c, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x61, 0x70, 0x6b, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x77, 0x73,
	0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_wso2_discovery_api_api_proto_rawDescOnce sync.Once
	file_api_proto_wso2_discovery_api_api_proto_rawDescData = file_api_proto_wso2_discovery_api_api_proto_rawDesc
)

func file_api_proto_wso2_discovery_api_api_proto_rawDescGZIP() []byte {
	file_api_proto_wso2_discovery_api_api_proto_rawDescOnce.Do(func() {
		file_api_proto_wso2_discovery_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_wso2_discovery_api_api_proto_rawDescData)
	})
	return file_api_proto_wso2_discovery_api_api_proto_rawDescData
}

var file_api_proto_wso2_discovery_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_wso2_discovery_api_api_proto_goTypes = []interface{}{
	(*API)(nil),               // 0: wso2.discovery.api.API
	(*CorsConfig)(nil),        // 1: wso2.discovery.api.CorsConfig
	(*Resource)(nil),          // 2: wso2.discovery.api.Resource
	(*OperationPolicies)(nil), // 3: wso2.discovery.api.OperationPolicies
	(*QueryParam)(nil),        // 4: wso2.discovery.api.QueryParam
	(*Scope)(nil),             // 5: wso2.discovery.api.Scope
	(*Authentication)(nil),    // 6: wso2.discovery.api.Authentication
	(*Credential)(nil),        // 7: wso2.discovery.api.Credential
	(*Response)(nil),          // 8: wso2.discovery.api.Response
}
var file_api_proto_wso2_discovery_api_api_proto_depIdxs = []int32{
	2,  // 0: wso2.discovery.api.API.resources:type_name -> wso2.discovery.api.Resource
	1,  // 1: wso2.discovery.api.API.corsConfig:type_name -> wso2.discovery.api.CorsConfig
	6,  // 2: wso2.discovery.api.Resource.authentications:type_name -> wso2.discovery.api.Authentication
	5,  // 3: wso2.discovery.api.Resource.scopes:type_name -> wso2.discovery.api.Scope
	3,  // 4: wso2.discovery.api.Resource.operationPolicies:type_name -> wso2.discovery.api.OperationPolicies
	4,  // 5: wso2.discovery.api.Resource.queryParams:type_name -> wso2.discovery.api.QueryParam
	7,  // 6: wso2.discovery.api.Authentication.credentialList:type_name -> wso2.discovery.api.Credential
	0,  // 7: wso2.discovery.api.APIService.createAPI:input_type -> wso2.discovery.api.API
	0,  // 8: wso2.discovery.api.APIService.updateAPI:input_type -> wso2.discovery.api.API
	0,  // 9: wso2.discovery.api.APIService.deleteAPI:input_type -> wso2.discovery.api.API
	8,  // 10: wso2.discovery.api.APIService.createAPI:output_type -> wso2.discovery.api.Response
	8,  // 11: wso2.discovery.api.APIService.updateAPI:output_type -> wso2.discovery.api.Response
	8,  // 12: wso2.discovery.api.APIService.deleteAPI:output_type -> wso2.discovery.api.Response
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_proto_wso2_discovery_api_api_proto_init() }
func file_api_proto_wso2_discovery_api_api_proto_init() {
	if File_api_proto_wso2_discovery_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_wso2_discovery_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*API); i {
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
		file_api_proto_wso2_discovery_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CorsConfig); i {
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
		file_api_proto_wso2_discovery_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_api_proto_wso2_discovery_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationPolicies); i {
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
		file_api_proto_wso2_discovery_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParam); i {
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
		file_api_proto_wso2_discovery_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scope); i {
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
		file_api_proto_wso2_discovery_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authentication); i {
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
		file_api_proto_wso2_discovery_api_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
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
		file_api_proto_wso2_discovery_api_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_api_proto_wso2_discovery_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_wso2_discovery_api_api_proto_goTypes,
		DependencyIndexes: file_api_proto_wso2_discovery_api_api_proto_depIdxs,
		MessageInfos:      file_api_proto_wso2_discovery_api_api_proto_msgTypes,
	}.Build()
	File_api_proto_wso2_discovery_api_api_proto = out.File
	file_api_proto_wso2_discovery_api_api_proto_rawDesc = nil
	file_api_proto_wso2_discovery_api_api_proto_goTypes = nil
	file_api_proto_wso2_discovery_api_api_proto_depIdxs = nil
}
