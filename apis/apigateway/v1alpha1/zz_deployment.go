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

// DeploymentParameters defines the desired state of Deployment
type DeploymentParameters struct {
	// Region is which region the Deployment will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Enables a cache cluster for the Stage resource specified in the input.
	CacheClusterEnabled *bool `json:"cacheClusterEnabled,omitempty"`
	// Specifies the cache cluster size for the Stage resource specified in the
	// input, if a cache cluster is enabled.
	CacheClusterSize *string `json:"cacheClusterSize,omitempty"`
	// The input configuration for the canary deployment when the deployment is
	// a canary release deployment.
	CanarySettings *DeploymentCanarySettings `json:"canarySettings,omitempty"`
	// The description for the Deployment resource to create.
	Description *string `json:"description,omitempty"`
	// [Required] The string identifier of the associated RestApi.
	// +kubebuilder:validation:Required
	RestAPIID *string `json:"restAPIID"`
	// The description of the Stage resource for the Deployment resource to create.
	StageDescription *string `json:"stageDescription,omitempty"`
	// The name of the Stage resource for the Deployment resource to create.
	StageName *string `json:"stageName,omitempty"`
	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled *bool `json:"tracingEnabled,omitempty"`
	// A map that defines the stage variables for the Stage resource that is associated
	// with the new deployment. Variable names can have alphanumeric and underscore
	// characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.
	Variables                  map[string]*string `json:"variables,omitempty"`
	CustomDeploymentParameters `json:",inline"`
}

// DeploymentSpec defines the desired state of Deployment
type DeploymentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DeploymentParameters `json:"forProvider"`
}

// DeploymentObservation defines the observed state of Deployment
type DeploymentObservation struct {
	// The date and time that the deployment resource was created.
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`
	// The identifier for the deployment resource.
	ID *string `json:"id,omitempty"`
}

// DeploymentStatus defines the observed state of Deployment.
type DeploymentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Deployment is the Schema for the Deployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentSpec   `json:"spec"`
	Status            DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployments
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

// Repository type metadata.
var (
	DeploymentKind             = "Deployment"
	DeploymentGroupKind        = schema.GroupKind{Group: Group, Kind: DeploymentKind}.String()
	DeploymentKindAPIVersion   = DeploymentKind + "." + GroupVersion.String()
	DeploymentGroupVersionKind = GroupVersion.WithKind(DeploymentKind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
