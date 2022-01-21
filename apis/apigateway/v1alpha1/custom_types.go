/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

type CustomAPIKeyParameters struct{}

type CustomAuthorizerParameters struct {
	// RestAPIID is the ID for the API.
	// +immutable
	RestAPIID *string `json:"restAPIID,omitempty"`

	// RestAPIIDRef is a reference to an API used to set
	// the restAPIID.
	// +optional
	RestAPIIDRef *xpv1.Reference `json:"restAPIIDRef,omitempty"`

	// RestAPIIDSelector is a reference to an API used to set
	// to set the RestAPIID.
	// +optional
	RestAPIIDSelector *xpv1.Selector `json:"Selector,omitempty"`
}

type CustomBasePathMappingParameters struct {
	// RestAPIID is the ID for the API.
	// +immutable
	RestAPIID *string `json:"restAPIID,omitempty"`

	// RestAPIIDRef is a reference to an API used to set
	// the restAPIID.
	// +optional
	RestAPIIDRef *xpv1.Reference `json:"restAPIIDRef,omitempty"`

	// RestAPIIDSelector is a reference to an API used to set
	// to set the RestAPIID.
	// +optional
	RestAPIIDSelector *xpv1.Selector `json:"restAPIIDSelector,omitempty"`
}

type CustomDeploymentParameters struct {
	// RestAPIID is the ID for the API.
	// +immutable
	RestAPIID *string `json:"restAPIID,omitempty"`

	// RestAPIIDRef is a reference to an API used to set
	// the restAPIID.
	// +optional
	RestAPIIDRef *xpv1.Reference `json:"restAPIIDRef,omitempty"`

	// RestAPIIDSelector is a reference to an API used to set
	// to set the RestAPIID.
	// +optional
	RestAPIIDSelector *xpv1.Selector `json:"restAPIIDSelector,omitempty"`

	// StageName is a name for the Stage
	// +optional
	// +crossplane:generate:reference:type=Stage
	StageName *string `json:"stageName,omitempty"`

	// StageNameRef is a reference to an Stage used to set
	// StageName
	// +optional
	StageNameRef *xpv1.Reference `json:"StageNameRef,omitempty"`

	// StageNameSelector is a reference to an Stage used to set
	// StageName
	// +optional
	StageNameSelector *xpv1.Selector `json:"StageNameSelector,omitempty"`
}

type CustomDocumentationPartParameters struct {
	// RestAPIID is the ID for the API.
	// +immutable
	RestAPIID *string `json:"restAPIID,omitempty"`

	// RestAPIIDRef is a reference to an API used to set
	// the restAPIID.
	// +optional
	RestAPIIDRef *xpv1.Reference `json:"restAPIIDRef,omitempty"`

	// RestAPIIDSelector is a reference to an API used to set
	// to set the RestAPIID.
	// +optional
	RestAPIIDSelector *xpv1.Selector `json:"restAPIIDSelector,omitempty"`
}

type CustomDocumentationVersionParameters struct {
	// RestAPIID is the ID for the API.
	// +immutable
	RestAPIID *string `json:"restAPIID,omitempty"`

	// RestAPIIDRef is a reference to an API used to set
	// the restAPIID.
	// +optional
	RestAPIIDRef *xpv1.Reference `json:"restAPIIDRef,omitempty"`

	// RestAPIIDSelector is a reference to an API used to set
	// to set the RestAPIID.
	// +optional
	RestAPIIDSelector *xpv1.Selector `json:"restAPIIDSelector,omitempty"`

	// StageName is a name for the Stage
	// +optional
	StageName *string `json:"stageName,omitempty"`

	// StageNameRef is a reference to an Stage used to set
	// StageName
	// +optional
	StageNameRef *xpv1.Reference `json:"stageNameRef,omitempty"`

	// StageNameSelector is a reference to an Stage used to set
	// StageName
	// +optional
	StageNameSelector *xpv1.Selector `json:"stageNameSelector,omitempty"`
}

type CustomDomainNameParameters struct{}

type CustomModelParameters struct {
	// RestAPIID is the ID for the API.
	// +immutable
	RestAPIID *string `json:"restAPIID,omitempty"`

	// RestAPIIDRef is a reference to an API used to set
	// the restAPIID.
	// +optional
	RestAPIIDRef *xpv1.Reference `json:"restAPIIDRef,omitempty"`

	// RestAPIIDSelector is a reference to an API used to set
	// to set the RestAPIID.
	// +optional
	RestAPIIDSelector *xpv1.Selector `json:"restAPIIDSelector,omitempty"`
}

type CustomRequestValidatorParameters struct {
	// RestAPIID is the ID for the API.
	// +immutable
	RestAPIID *string `json:"restAPIID,omitempty"`

	// RestAPIIDRef is a reference to an API used to set
	// the restAPIID.
	// +optional
	RestAPIIDRef *xpv1.Reference `json:"restAPIIDRef,omitempty"`

	// RestAPIIDSelector is a reference to an API used to set
	// to set the RestAPIID.
	// +optional
	RestAPIIDSelector *xpv1.Selector `json:"restAPIIDSelector,omitempty"`
}

type CustomResourceParameters struct {
	// RestAPIID is the ID for the API.
	// +immutable
	RestAPIID *string `json:"restAPIID,omitempty"`

	// RestAPIIDRef is a reference to an API used to set
	// the restAPIID.
	// +optional
	RestAPIIDRef *xpv1.Reference `json:"restAPIIDRef,omitempty"`

	// RestAPIIDSelector is a reference to an API used to set
	// to set the RestAPIID.
	// +optional
	RestAPIIDSelector *xpv1.Selector `json:"restAPIIDSelector,omitempty"`

	// ParentID is the ID for the Parent.
	// +immutable
	// +crossplane:generate:reference:type=Resource
	ParentID *string `json:"parentID"`

	// ParentIdRef is a reference to an API used to set
	// the parentID.
	// +optional
	ParentIDRef *xpv1.Reference `json:"parentIdRef,omitempty"`

	// ParentIdSelector is a reference to an API used to set
	// the parentID.
	// +optional
	ParentIDSelector *xpv1.Selector `json:"ParentIdSelector,omitempty"`
}

type CustomRestAPIParameters struct {
	// RestAPIID is the ID for the API.
	// +immutable
	RestAPIID *string `json:"restAPIID,omitempty"`

	// RestAPIIDRef is a reference to an API used to set
	// the restAPIID.
	// +optional
	RestAPIIDRef *xpv1.Reference `json:"restAPIIDRef,omitempty"`

	// RestAPIIDSelector is a reference to an API used to set
	// to set the RestAPIID.
	// +optional
	RestAPIIDSelector *xpv1.Selector `json:"restAPIIDSelector,omitempty"`
}

type CustomStageParameters struct {
	// RestAPIID is the ID for the API.
	// +immutable
	RestAPIID *string `json:"restAPIID,omitempty"`

	// RestAPIIDRef is a reference to an API used to set
	// the restAPIID.
	// +optional
	RestAPIIDRef *xpv1.Reference `json:"restAPIIDRef,omitempty"`

	// RestAPIIDSelector is a reference to an API used to set
	// to set the RestAPIID.
	// +optional
	RestAPIIDSelector *xpv1.Selector `json:"restAPIIDSelector,omitempty"`

	// DeploymentID is the ID for the Deployment.
	// +immutable
	// +crossplane:generate:reference:type=Deployment
	DeploymentID *string `json:"deploymentID"`

	// DeploymentIDRef is a reference to an API used to set
	// the DeploymentID.
	// +optional
	DeploymentIDRef *xpv1.Reference `json:"deploymentIDRef,omitempty"`

	// DeploymentIDSelector is a reference to an API used to set
	// the DeploymentID.
	// +optional
	DeploymentIDSelector *xpv1.Selector `json:"deploymentIDSelector,omitempty"`
}

type CustomUsagePlanKeyParameters struct{}

type CustomUsagePlanParameters struct{}
type CustomVPCLinkParameters struct {

	// TargetARNs is the arns for the target.
	// +immutable
	// +crossplane:generate:reference:type=github.com/crossplane/provider-aws/apis/elbv2/v1alpha1.LoadBalancer
	TargetARNs []*string `json:"targetARNs"`

	// TargetARNsRef is a reference to an API used to set
	// the TargetARNs.
	// +optional
	TargetARNsRefs []xpv1.Reference `json:"targetARNsRefs,omitempty"`

	// TargetARNsSelector is a reference to an API used to set
	//	TargetARNs
	// +optional
	TargetARNsSelector *xpv1.Selector `json:"targetARNsSelector,omitempty"`
}
