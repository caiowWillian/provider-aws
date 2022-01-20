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

package apikey

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigateway/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetApiKeyInput returns input for read
// operation.
func GenerateGetApiKeyInput(cr *svcapitypes.APIKey) *svcsdk.GetApiKeyInput {
	res := &svcsdk.GetApiKeyInput{}

	return res
}

// GenerateAPIKey returns the current state in the form of *svcapitypes.APIKey.
func GenerateAPIKey(resp *svcsdk.ApiKey) *svcapitypes.APIKey {
	cr := &svcapitypes.APIKey{}

	if resp.CreatedDate != nil {
		cr.Status.AtProvider.CreatedDate = &metav1.Time{*resp.CreatedDate}
	} else {
		cr.Status.AtProvider.CreatedDate = nil
	}
	if resp.CustomerId != nil {
		cr.Spec.ForProvider.CustomerID = resp.CustomerId
	} else {
		cr.Spec.ForProvider.CustomerID = nil
	}
	if resp.Description != nil {
		cr.Spec.ForProvider.Description = resp.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.Enabled != nil {
		cr.Spec.ForProvider.Enabled = resp.Enabled
	} else {
		cr.Spec.ForProvider.Enabled = nil
	}
	if resp.Id != nil {
		cr.Status.AtProvider.ID = resp.Id
	} else {
		cr.Status.AtProvider.ID = nil
	}
	if resp.LastUpdatedDate != nil {
		cr.Status.AtProvider.LastUpdatedDate = &metav1.Time{*resp.LastUpdatedDate}
	} else {
		cr.Status.AtProvider.LastUpdatedDate = nil
	}
	if resp.Name != nil {
		cr.Spec.ForProvider.Name = resp.Name
	} else {
		cr.Spec.ForProvider.Name = nil
	}
	if resp.StageKeys != nil {
		f7 := []*svcapitypes.StageKey{}
		for _, f7iter := range resp.StageKeys {
			f7elem := &svcapitypes.StageKey{}
			f7elem = *f7iter
			f7 = append(f7, f7elem)
		}
		cr.Spec.ForProvider.StageKeys = f7
	} else {
		cr.Spec.ForProvider.StageKeys = nil
	}
	if resp.Tags != nil {
		f8 := map[string]*string{}
		for f8key, f8valiter := range resp.Tags {
			var f8val string
			f8val = *f8valiter
			f8[f8key] = &f8val
		}
		cr.Spec.ForProvider.Tags = f8
	} else {
		cr.Spec.ForProvider.Tags = nil
	}
	if resp.Value != nil {
		cr.Spec.ForProvider.Value = resp.Value
	} else {
		cr.Spec.ForProvider.Value = nil
	}

	return cr
}

// GenerateCreateApiKeyInput returns a create input.
func GenerateCreateApiKeyInput(cr *svcapitypes.APIKey) *svcsdk.CreateApiKeyInput {
	res := &svcsdk.CreateApiKeyInput{}

	if cr.Spec.ForProvider.CustomerID != nil {
		res.SetCustomerId(*cr.Spec.ForProvider.CustomerID)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.Enabled != nil {
		res.SetEnabled(*cr.Spec.ForProvider.Enabled)
	}
	if cr.Spec.ForProvider.GenerateDistinctID != nil {
		res.SetGenerateDistinctId(*cr.Spec.ForProvider.GenerateDistinctID)
	}
	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}
	if cr.Spec.ForProvider.StageKeys != nil {
		f5 := []*svcsdk.StageKey{}
		for _, f5iter := range cr.Spec.ForProvider.StageKeys {
			f5elem := &svcsdk.StageKey{}
			if f5iter.RestAPIID != nil {
				f5elem.SetRestApiId(*f5iter.RestAPIID)
			}
			if f5iter.StageName != nil {
				f5elem.SetStageName(*f5iter.StageName)
			}
			f5 = append(f5, f5elem)
		}
		res.SetStageKeys(f5)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f6 := map[string]*string{}
		for f6key, f6valiter := range cr.Spec.ForProvider.Tags {
			var f6val string
			f6val = *f6valiter
			f6[f6key] = &f6val
		}
		res.SetTags(f6)
	}
	if cr.Spec.ForProvider.Value != nil {
		res.SetValue(*cr.Spec.ForProvider.Value)
	}

	return res
}

// GenerateUpdateApiKeyInput returns an update input.
func GenerateUpdateApiKeyInput(cr *svcapitypes.APIKey) *svcsdk.UpdateApiKeyInput {
	res := &svcsdk.UpdateApiKeyInput{}

	return res
}

// GenerateDeleteApiKeyInput returns a deletion input.
func GenerateDeleteApiKeyInput(cr *svcapitypes.APIKey) *svcsdk.DeleteApiKeyInput {
	res := &svcsdk.DeleteApiKeyInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
