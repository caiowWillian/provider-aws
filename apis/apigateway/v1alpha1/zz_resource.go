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

// ResourceParameters defines the desired state of Resource
type ResourceParameters struct {
	// Region is which region the Resource will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// [Required] The parent resource's identifier.
	// +kubebuilder:validation:Required
	ParentID *string `json:"parentID"`
	// The last path segment for this resource.
	// +kubebuilder:validation:Required
	PathPart *string `json:"pathPart"`
	// [Required] The string identifier of the associated RestApi.
	// +kubebuilder:validation:Required
	RestAPIID                *string `json:"restAPIID"`
	CustomResourceParameters `json:",inline"`
}

// ResourceSpec defines the desired state of Resource
type ResourceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ResourceParameters `json:"forProvider"`
}

// ResourceObservation defines the observed state of Resource
type ResourceObservation struct {
	// The resource's identifier.
	ID *string `json:"id,omitempty"`
	// The full path for this resource.
	Path *string `json:"path,omitempty"`
	// Gets an API resource's method of a given HTTP verb.
	//
	// The resource methods are a map of methods indexed by methods' HTTP verbs
	// enabled on the resource. This method map is included in the 200 OK response
	// of the GET /restapis/{restapi_id}/resources/{resource_id} or GET /restapis/{restapi_id}/resources/{resource_id}?embed=methods
	// request.
	//
	// Example: Get the GET method of an API resource
	//
	// Request
	//  GET /restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET HTTP/1.1 Content-Type:
	//  application/json Host: apigateway.us-east-1.amazonaws.com X-Amz-Date: 20170223T031827Z
	//  Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request,
	//  SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
	// Response
	//  { "_links": { "curies": [ { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-{rel}.html",
	//  "name": "integration", "templated": true }, { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-response-{rel}.html",
	//  "name": "integrationresponse", "templated": true }, { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-{rel}.html",
	//  "name": "method", "templated": true }, { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-response-{rel}.html",
	//  "name": "methodresponse", "templated": true } ], "self": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET",
	//  "name": "GET", "title": "GET" }, "integration:put": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
	//  }, "method:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET"
	//  }, "method:integration": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
	//  }, "method:responses": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200",
	//  "name": "200", "title": "200" }, "method:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET"
	//  }, "methodresponse:put": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/{status_code}",
	//  "templated": true } }, "apiKeyRequired": false, "authorizationType": "NONE",
	//  "httpMethod": "GET", "_embedded": { "method:integration": { "_links": {
	//  "self": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
	//  }, "integration:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
	//  }, "integration:responses": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
	//  "name": "200", "title": "200" }, "integration:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
	//  }, "integrationresponse:put": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/{status_code}",
	//  "templated": true } }, "cacheKeyParameters": [], "cacheNamespace": "3kzxbg5sa2",
	//  "credentials": "arn:aws:iam::123456789012:role/apigAwsProxyRole", "httpMethod":
	//  "POST", "passthroughBehavior": "WHEN_NO_MATCH", "requestParameters": { "integration.request.header.Content-Type":
	//  "'application/x-amz-json-1.1'" }, "requestTemplates": { "application/json":
	//  "{\n}" }, "type": "AWS", "uri": "arn:aws:apigateway:us-east-1:kinesis:action/ListStreams",
	//  "_embedded": { "integration:responses": { "_links": { "self": { "href":
	//  "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
	//  "name": "200", "title": "200" }, "integrationresponse:delete": { "href":
	//  "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
	//  }, "integrationresponse:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
	//  } }, "responseParameters": { "method.response.header.Content-Type": "'application/xml'"
	//  }, "responseTemplates": { "application/json": "$util.urlDecode(\"%3CkinesisStreams%3E#foreach($stream
	//  in $input.path('$.StreamNames'))%3Cstream%3E%3Cname%3E$stream%3C/name%3E%3C/stream%3E#end%3C/kinesisStreams%3E\")\n"
	//  }, "statusCode": "200" } } }, "method:responses": { "_links": { "self":
	//  { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200",
	//  "name": "200", "title": "200" }, "methodresponse:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200"
	//  }, "methodresponse:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200"
	//  } }, "responseModels": { "application/json": "Empty" }, "responseParameters":
	//  { "method.response.header.Content-Type": false }, "statusCode": "200" }
	//  } }
	// If the OPTIONS is enabled on the resource, you can follow the example here
	// to get that method. Just replace the GET of the last path segment in the
	// request URL with OPTIONS.
	ResourceMethods map[string]*Method `json:"resourceMethods,omitempty"`
}

// ResourceStatus defines the observed state of Resource.
type ResourceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Resource is the Schema for the Resources API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Resource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceSpec   `json:"spec"`
	Status            ResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceList contains a list of Resources
type ResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Resource `json:"items"`
}

// Repository type metadata.
var (
	ResourceKind             = "Resource"
	ResourceGroupKind        = schema.GroupKind{Group: Group, Kind: ResourceKind}.String()
	ResourceKindAPIVersion   = ResourceKind + "." + GroupVersion.String()
	ResourceGroupVersionKind = GroupVersion.WithKind(ResourceKind)
)

func init() {
	SchemeBuilder.Register(&Resource{}, &ResourceList{})
}
