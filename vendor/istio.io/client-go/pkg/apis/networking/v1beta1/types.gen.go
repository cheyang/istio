// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "istio.io/api/meta/v1alpha1"
	networkingv1beta1 "istio.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DestinationRule defines policies that apply to traffic intended for a service
// after routing has occurred.
//
// <!-- crd generation tags
// +cue-gen:DestinationRule:groupName:networking.istio.io
// +cue-gen:DestinationRule:version:v1beta1
// +cue-gen:DestinationRule:annotations:helm.sh/resource-policy=keep
// +cue-gen:DestinationRule:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:DestinationRule:subresource:status
// +cue-gen:DestinationRule:scope:Namespaced
// +cue-gen:DestinationRule:resource:categories=istio-io,networking-istio-io,shortNames=dr
// +cue-gen:DestinationRule:printerColumn:name=Host,type=string,JSONPath=.spec.host,description="The name of a service from the service registry"
// +cue-gen:DestinationRule:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:DestinationRule:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// <!-- istio code generation tags
// +istio.io/sync-from:networking/v1alpha3/destination_rule.proto
// -->
type DestinationRule struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec networkingv1beta1.DestinationRule `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DestinationRuleList is a collection of DestinationRules.
type DestinationRuleList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []DestinationRule `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Gateway describes a load balancer operating at the edge of the mesh
// receiving incoming or outgoing HTTP/TCP connections.
//
// <!-- crd generation tags
// +cue-gen:Gateway:groupName:networking.istio.io
// +cue-gen:Gateway:version:v1beta1
// +cue-gen:Gateway:annotations:helm.sh/resource-policy=keep
// +cue-gen:Gateway:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:Gateway:subresource:status
// +cue-gen:Gateway:scope:Namespaced
// +cue-gen:Gateway:resource:categories=istio-io,networking-istio-io,shortNames=gw
// +cue-gen:Gateway:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// <!-- istio code generation tags
// +istio.io/sync-from:networking/v1alpha3/gateway.proto
// -->
type Gateway struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec networkingv1beta1.Gateway `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GatewayList is a collection of Gateways.
type GatewayList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Gateway `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// please upgrade the proto package
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// `ProxyConfig` exposes proxy level configuration options.
//
// <!-- crd generation tags
// +cue-gen:ProxyConfig:groupName:networking.istio.io
// +cue-gen:ProxyConfig:version:v1beta1
// +cue-gen:ProxyConfig:storageVersion
// +cue-gen:ProxyConfig:annotations:helm.sh/resource-policy=keep
// +cue-gen:ProxyConfig:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:ProxyConfig:subresource:status
// +cue-gen:ProxyConfig:scope:Namespaced
// +cue-gen:ProxyConfig:resource:categories=istio-io,networking-istio-io,plural=proxyconfigs
// +cue-gen:ProxyConfig:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type ProxyConfig struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec networkingv1beta1.ProxyConfig `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProxyConfigList is a collection of ProxyConfigs.
type ProxyConfigList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []ProxyConfig `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceEntry enables adding additional entries into Istio's internal
// service registry.
//
// <!-- crd generation tags
// +cue-gen:ServiceEntry:groupName:networking.istio.io
// +cue-gen:ServiceEntry:version:v1beta1
// +cue-gen:ServiceEntry:annotations:helm.sh/resource-policy=keep
// +cue-gen:ServiceEntry:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:ServiceEntry:subresource:status
// +cue-gen:ServiceEntry:scope:Namespaced
// +cue-gen:ServiceEntry:resource:categories=istio-io,networking-istio-io,shortNames=se,plural=serviceentries
// +cue-gen:ServiceEntry:printerColumn:name=Hosts,type=string,JSONPath=.spec.hosts,description="The hosts associated with the ServiceEntry"
// +cue-gen:ServiceEntry:printerColumn:name=Location,type=string,JSONPath=.spec.location,description="Whether the service is external to the
// mesh or part of the mesh (MESH_EXTERNAL or MESH_INTERNAL)"
// +cue-gen:ServiceEntry:printerColumn:name=Resolution,type=string,JSONPath=.spec.resolution,description="Service discovery mode for the hosts
// (NONE, STATIC, or DNS)"
// +cue-gen:ServiceEntry:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:ServiceEntry:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// <!-- istio code generation tags
// +istio.io/sync-from:networking/v1alpha3/service_entry.proto
// -->
type ServiceEntry struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec networkingv1beta1.ServiceEntry `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceEntryList is a collection of ServiceEntries.
type ServiceEntryList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []ServiceEntry `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// `Sidecar` describes the configuration of the sidecar proxy that mediates
// inbound and outbound communication of the workload instance to which it is
// attached.
//
// <!-- crd generation tags
// +cue-gen:Sidecar:groupName:networking.istio.io
// +cue-gen:Sidecar:version:v1beta1
// +cue-gen:Sidecar:annotations:helm.sh/resource-policy=keep
// +cue-gen:Sidecar:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:Sidecar:subresource:status
// +cue-gen:Sidecar:scope:Namespaced
// +cue-gen:Sidecar:resource:categories=istio-io,networking-istio-io
// +cue-gen:Sidecar:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// <!-- istio code generation tags
// +istio.io/sync-from:networking/v1alpha3/sidecar.proto
// -->
type Sidecar struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec networkingv1beta1.Sidecar `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SidecarList is a collection of Sidecars.
type SidecarList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Sidecar `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Configuration affecting traffic routing.
//
// <!-- crd generation tags
// +cue-gen:VirtualService:groupName:networking.istio.io
// +cue-gen:VirtualService:version:v1beta1
// +cue-gen:VirtualService:annotations:helm.sh/resource-policy=keep
// +cue-gen:VirtualService:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:VirtualService:subresource:status
// +cue-gen:VirtualService:scope:Namespaced
// +cue-gen:VirtualService:resource:categories=istio-io,networking-istio-io,shortNames=vs
// +cue-gen:VirtualService:printerColumn:name=Gateways,type=string,JSONPath=.spec.gateways,description="The names of gateways and sidecars
// that should apply these routes"
// +cue-gen:VirtualService:printerColumn:name=Hosts,type=string,JSONPath=.spec.hosts,description="The destination hosts to which traffic is being sent"
// +cue-gen:VirtualService:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:VirtualService:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// <!-- istio code generation tags
// +istio.io/sync-from:networking/v1alpha3/virtual_service.proto
// -->
type VirtualService struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec networkingv1beta1.VirtualService `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualServiceList is a collection of VirtualServices.
type VirtualServiceList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []VirtualService `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// please upgrade the proto package
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkloadEntry enables specifying the properties of a single non-Kubernetes workload such a VM or a bare metal services that can be referred to by service entries.
//
// <!-- crd generation tags
// +cue-gen:WorkloadEntry:groupName:networking.istio.io
// +cue-gen:WorkloadEntry:version:v1beta1
// +cue-gen:WorkloadEntry:annotations:helm.sh/resource-policy=keep
// +cue-gen:WorkloadEntry:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:WorkloadEntry:subresource:status
// +cue-gen:WorkloadEntry:scope:Namespaced
// +cue-gen:WorkloadEntry:resource:categories=istio-io,networking-istio-io,shortNames=we,plural=workloadentries
// +cue-gen:WorkloadEntry:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:WorkloadEntry:printerColumn:name=Address,type=string,JSONPath=.spec.address,description="Address associated with the network endpoint."
// +cue-gen:WorkloadEntry:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// <!-- istio code generation tags
// +istio.io/sync-from:networking/v1alpha3/workload_entry.proto
// -->
type WorkloadEntry struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec networkingv1beta1.WorkloadEntry `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkloadEntryList is a collection of WorkloadEntries.
type WorkloadEntryList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []WorkloadEntry `json:"items" protobuf:"bytes,2,rep,name=items"`
}
