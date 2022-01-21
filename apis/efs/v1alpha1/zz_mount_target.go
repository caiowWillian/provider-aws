/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// MountTargetParameters defines the desired state of MountTarget
type MountTargetParameters struct {
	// Region is which region the MountTarget will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Valid IPv4 address within the address range of the specified subnet.
	 IPAddress *string `json:"ipAddress,omitempty"` 
	CustomMountTargetParameters `json:",inline"`
}

// MountTargetSpec defines the desired state of MountTarget
type MountTargetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider MountTargetParameters `json:"forProvider"`
}

// MountTargetObservation defines the observed state of MountTarget
type MountTargetObservation struct {
	// The unique and consistent identifier of the Availability Zone (AZ) that the
// mount target resides in. For example, use1-az1 is an AZ ID for the us-east-1
// Region and it has the same location in every AWS account.
	AvailabilityZoneID *string `json:"availabilityZoneID,omitempty"`
	// The name of the Availability Zone (AZ) that the mount target resides in.
// AZs are independently mapped to names for each AWS account. For example,
// the Availability Zone us-east-1a for your AWS account might not be the same
// location as us-east-1a for another AWS account.
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty"`
	// The ID of the file system for which the mount target is intended.
	FileSystemID *string `json:"fileSystemID,omitempty"`
	// Lifecycle state of the mount target.
	LifeCycleState *string `json:"lifeCycleState,omitempty"`
	// System-assigned mount target ID.
	MountTargetID *string `json:"mountTargetID,omitempty"`
	// The ID of the network interface that Amazon EFS created when it created the
// mount target.
	NetworkInterfaceID *string `json:"networkInterfaceID,omitempty"`
	// AWS account ID that owns the resource.
	OwnerID *string `json:"ownerID,omitempty"`
	// The ID of the mount target's subnet.
	SubnetID *string `json:"subnetID,omitempty"`
	// The Virtual Private Cloud (VPC) ID that the mount target is configured in.
	VPCID *string `json:"vpcID,omitempty"`
}

// MountTargetStatus defines the observed state of MountTarget.
type MountTargetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider MountTargetObservation `json:"atProvider,omitempty"`
}


// +kubebuilder:object:root=true

// MountTarget is the Schema for the MountTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MountTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   MountTargetSpec   `json:"spec"`
	Status MountTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MountTargetList contains a list of MountTargets
type MountTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []MountTarget `json:"items"`
}

// Repository type metadata.
var (
	MountTargetKind             = "MountTarget"
	MountTargetGroupKind        = schema.GroupKind{Group: Group, Kind: MountTargetKind}.String()
	MountTargetKindAPIVersion   = MountTargetKind + "." + GroupVersion.String()
	MountTargetGroupVersionKind = GroupVersion.WithKind(MountTargetKind)
)

func init() {
	SchemeBuilder.Register(&MountTarget{}, &MountTargetList{})
}

