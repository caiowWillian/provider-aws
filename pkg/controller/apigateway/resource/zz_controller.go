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

package resource

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigateway"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigateway/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an Resource resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Resource in AWS"
	errUpdate        = "cannot update Resource in AWS"
	errDescribe      = "failed to describe Resource"
	errDelete        = "failed to delete Resource"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Resource)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.Resource)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetResourceInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetResourceWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateResource(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.Resource)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateResourceInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateResourceWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.Id != nil {
		cr.Status.AtProvider.ID = resp.Id
	} else {
		cr.Status.AtProvider.ID = nil
	}
	if resp.ParentId != nil {
		cr.Spec.ForProvider.ParentID = resp.ParentId
	} else {
		cr.Spec.ForProvider.ParentID = nil
	}
	if resp.Path != nil {
		cr.Status.AtProvider.Path = resp.Path
	} else {
		cr.Status.AtProvider.Path = nil
	}
	if resp.PathPart != nil {
		cr.Spec.ForProvider.PathPart = resp.PathPart
	} else {
		cr.Spec.ForProvider.PathPart = nil
	}
	if resp.ResourceMethods != nil {
		f4 := map[string]*svcapitypes.Method{}
		for f4key, f4valiter := range resp.ResourceMethods {
			f4val := &svcapitypes.Method{}
			if f4valiter.ApiKeyRequired != nil {
				f4val.APIKeyRequired = f4valiter.ApiKeyRequired
			}
			if f4valiter.AuthorizationScopes != nil {
				f4valf1 := []*string{}
				for _, f4valf1iter := range f4valiter.AuthorizationScopes {
					var f4valf1elem string
					f4valf1elem = *f4valf1iter
					f4valf1 = append(f4valf1, &f4valf1elem)
				}
				f4val.AuthorizationScopes = f4valf1
			}
			if f4valiter.AuthorizationType != nil {
				f4val.AuthorizationType = f4valiter.AuthorizationType
			}
			if f4valiter.AuthorizerId != nil {
				f4val.AuthorizerID = f4valiter.AuthorizerId
			}
			if f4valiter.HttpMethod != nil {
				f4val.HTTPMethod = f4valiter.HttpMethod
			}
			if f4valiter.MethodIntegration != nil {
				f4valf5 := &svcapitypes.Integration{}
				if f4valiter.MethodIntegration.CacheKeyParameters != nil {
					f4valf5f0 := []*string{}
					for _, f4valf5f0iter := range f4valiter.MethodIntegration.CacheKeyParameters {
						var f4valf5f0elem string
						f4valf5f0elem = *f4valf5f0iter
						f4valf5f0 = append(f4valf5f0, &f4valf5f0elem)
					}
					f4valf5.CacheKeyParameters = f4valf5f0
				}
				if f4valiter.MethodIntegration.CacheNamespace != nil {
					f4valf5.CacheNamespace = f4valiter.MethodIntegration.CacheNamespace
				}
				if f4valiter.MethodIntegration.ConnectionId != nil {
					f4valf5.ConnectionID = f4valiter.MethodIntegration.ConnectionId
				}
				if f4valiter.MethodIntegration.ConnectionType != nil {
					f4valf5.ConnectionType = f4valiter.MethodIntegration.ConnectionType
				}
				if f4valiter.MethodIntegration.ContentHandling != nil {
					f4valf5.ContentHandling = f4valiter.MethodIntegration.ContentHandling
				}
				if f4valiter.MethodIntegration.Credentials != nil {
					f4valf5.Credentials = f4valiter.MethodIntegration.Credentials
				}
				if f4valiter.MethodIntegration.HttpMethod != nil {
					f4valf5.HTTPMethod = f4valiter.MethodIntegration.HttpMethod
				}
				if f4valiter.MethodIntegration.IntegrationResponses != nil {
					f4valf5f7 := map[string]*svcapitypes.IntegrationResponse{}
					for f4valf5f7key, f4valf5f7valiter := range f4valiter.MethodIntegration.IntegrationResponses {
						f4valf5f7val := &svcapitypes.IntegrationResponse{}
						if f4valf5f7valiter.ContentHandling != nil {
							f4valf5f7val.ContentHandling = f4valf5f7valiter.ContentHandling
						}
						if f4valf5f7valiter.ResponseParameters != nil {
							f4valf5f7valf1 := map[string]*string{}
							for f4valf5f7valf1key, f4valf5f7valf1valiter := range f4valf5f7valiter.ResponseParameters {
								var f4valf5f7valf1val string
								f4valf5f7valf1val = *f4valf5f7valf1valiter
								f4valf5f7valf1[f4valf5f7valf1key] = &f4valf5f7valf1val
							}
							f4valf5f7val.ResponseParameters = f4valf5f7valf1
						}
						if f4valf5f7valiter.ResponseTemplates != nil {
							f4valf5f7valf2 := map[string]*string{}
							for f4valf5f7valf2key, f4valf5f7valf2valiter := range f4valf5f7valiter.ResponseTemplates {
								var f4valf5f7valf2val string
								f4valf5f7valf2val = *f4valf5f7valf2valiter
								f4valf5f7valf2[f4valf5f7valf2key] = &f4valf5f7valf2val
							}
							f4valf5f7val.ResponseTemplates = f4valf5f7valf2
						}
						if f4valf5f7valiter.SelectionPattern != nil {
							f4valf5f7val.SelectionPattern = f4valf5f7valiter.SelectionPattern
						}
						if f4valf5f7valiter.StatusCode != nil {
							f4valf5f7val.StatusCode = f4valf5f7valiter.StatusCode
						}
						f4valf5f7[f4valf5f7key] = f4valf5f7val
					}
					f4valf5.IntegrationResponses = f4valf5f7
				}
				if f4valiter.MethodIntegration.PassthroughBehavior != nil {
					f4valf5.PassthroughBehavior = f4valiter.MethodIntegration.PassthroughBehavior
				}
				if f4valiter.MethodIntegration.RequestParameters != nil {
					f4valf5f9 := map[string]*string{}
					for f4valf5f9key, f4valf5f9valiter := range f4valiter.MethodIntegration.RequestParameters {
						var f4valf5f9val string
						f4valf5f9val = *f4valf5f9valiter
						f4valf5f9[f4valf5f9key] = &f4valf5f9val
					}
					f4valf5.RequestParameters = f4valf5f9
				}
				if f4valiter.MethodIntegration.RequestTemplates != nil {
					f4valf5f10 := map[string]*string{}
					for f4valf5f10key, f4valf5f10valiter := range f4valiter.MethodIntegration.RequestTemplates {
						var f4valf5f10val string
						f4valf5f10val = *f4valf5f10valiter
						f4valf5f10[f4valf5f10key] = &f4valf5f10val
					}
					f4valf5.RequestTemplates = f4valf5f10
				}
				if f4valiter.MethodIntegration.TimeoutInMillis != nil {
					f4valf5.TimeoutInMillis = f4valiter.MethodIntegration.TimeoutInMillis
				}
				if f4valiter.MethodIntegration.TlsConfig != nil {
					f4valf5f12 := &svcapitypes.TLSConfig{}
					if f4valiter.MethodIntegration.TlsConfig.InsecureSkipVerification != nil {
						f4valf5f12.InsecureSkipVerification = f4valiter.MethodIntegration.TlsConfig.InsecureSkipVerification
					}
					f4valf5.TLSConfig = f4valf5f12
				}
				if f4valiter.MethodIntegration.Type != nil {
					f4valf5.Type = f4valiter.MethodIntegration.Type
				}
				if f4valiter.MethodIntegration.Uri != nil {
					f4valf5.URI = f4valiter.MethodIntegration.Uri
				}
				f4val.MethodIntegration = f4valf5
			}
			if f4valiter.MethodResponses != nil {
				f4valf6 := map[string]*svcapitypes.MethodResponse{}
				for f4valf6key, f4valf6valiter := range f4valiter.MethodResponses {
					f4valf6val := &svcapitypes.MethodResponse{}
					if f4valf6valiter.ResponseModels != nil {
						f4valf6valf0 := map[string]*string{}
						for f4valf6valf0key, f4valf6valf0valiter := range f4valf6valiter.ResponseModels {
							var f4valf6valf0val string
							f4valf6valf0val = *f4valf6valf0valiter
							f4valf6valf0[f4valf6valf0key] = &f4valf6valf0val
						}
						f4valf6val.ResponseModels = f4valf6valf0
					}
					if f4valf6valiter.ResponseParameters != nil {
						f4valf6valf1 := map[string]*bool{}
						for f4valf6valf1key, f4valf6valf1valiter := range f4valf6valiter.ResponseParameters {
							var f4valf6valf1val bool
							f4valf6valf1val = *f4valf6valf1valiter
							f4valf6valf1[f4valf6valf1key] = &f4valf6valf1val
						}
						f4valf6val.ResponseParameters = f4valf6valf1
					}
					if f4valf6valiter.StatusCode != nil {
						f4valf6val.StatusCode = f4valf6valiter.StatusCode
					}
					f4valf6[f4valf6key] = f4valf6val
				}
				f4val.MethodResponses = f4valf6
			}
			if f4valiter.OperationName != nil {
				f4val.OperationName = f4valiter.OperationName
			}
			if f4valiter.RequestModels != nil {
				f4valf8 := map[string]*string{}
				for f4valf8key, f4valf8valiter := range f4valiter.RequestModels {
					var f4valf8val string
					f4valf8val = *f4valf8valiter
					f4valf8[f4valf8key] = &f4valf8val
				}
				f4val.RequestModels = f4valf8
			}
			if f4valiter.RequestParameters != nil {
				f4valf9 := map[string]*bool{}
				for f4valf9key, f4valf9valiter := range f4valiter.RequestParameters {
					var f4valf9val bool
					f4valf9val = *f4valf9valiter
					f4valf9[f4valf9key] = &f4valf9val
				}
				f4val.RequestParameters = f4valf9
			}
			if f4valiter.RequestValidatorId != nil {
				f4val.RequestValidatorID = f4valiter.RequestValidatorId
			}
			f4[f4key] = f4val
		}
		cr.Status.AtProvider.ResourceMethods = f4
	} else {
		cr.Status.AtProvider.ResourceMethods = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.Resource)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateResourceInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateResourceWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.Resource)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteResourceInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteResourceWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.APIGatewayAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.APIGatewayAPI
	preObserve     func(context.Context, *svcapitypes.Resource, *svcsdk.GetResourceInput) error
	postObserve    func(context.Context, *svcapitypes.Resource, *svcsdk.Resource, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.ResourceParameters, *svcsdk.Resource) error
	isUpToDate     func(*svcapitypes.Resource, *svcsdk.Resource) (bool, error)
	preCreate      func(context.Context, *svcapitypes.Resource, *svcsdk.CreateResourceInput) error
	postCreate     func(context.Context, *svcapitypes.Resource, *svcsdk.Resource, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Resource, *svcsdk.DeleteResourceInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Resource, *svcsdk.DeleteResourceOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.Resource, *svcsdk.UpdateResourceInput) error
	postUpdate     func(context.Context, *svcapitypes.Resource, *svcsdk.Resource, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Resource, *svcsdk.GetResourceInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.Resource, _ *svcsdk.Resource, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.ResourceParameters, *svcsdk.Resource) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.Resource, *svcsdk.Resource) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.Resource, *svcsdk.CreateResourceInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Resource, _ *svcsdk.Resource, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Resource, *svcsdk.DeleteResourceInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Resource, _ *svcsdk.DeleteResourceOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.Resource, *svcsdk.UpdateResourceInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.Resource, _ *svcsdk.Resource, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
