// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newRegionSecurityPoliciesClientHook clientHook

// RegionSecurityPoliciesCallOptions contains the retry settings for each method of RegionSecurityPoliciesClient.
type RegionSecurityPoliciesCallOptions struct {
	AddRule    []gax.CallOption
	Delete     []gax.CallOption
	Get        []gax.CallOption
	GetRule    []gax.CallOption
	Insert     []gax.CallOption
	List       []gax.CallOption
	Patch      []gax.CallOption
	PatchRule  []gax.CallOption
	RemoveRule []gax.CallOption
}

func defaultRegionSecurityPoliciesRESTCallOptions() *RegionSecurityPoliciesCallOptions {
	return &RegionSecurityPoliciesCallOptions{
		AddRule: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		Delete: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		Get: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		GetRule: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		Insert: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		List: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		Patch: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		PatchRule: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		RemoveRule: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
	}
}

// internalRegionSecurityPoliciesClient is an interface that defines the methods available from Google Compute Engine API.
type internalRegionSecurityPoliciesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddRule(context.Context, *computepb.AddRuleRegionSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
	Delete(context.Context, *computepb.DeleteRegionSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetRegionSecurityPolicyRequest, ...gax.CallOption) (*computepb.SecurityPolicy, error)
	GetRule(context.Context, *computepb.GetRuleRegionSecurityPolicyRequest, ...gax.CallOption) (*computepb.SecurityPolicyRule, error)
	Insert(context.Context, *computepb.InsertRegionSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListRegionSecurityPoliciesRequest, ...gax.CallOption) *SecurityPolicyIterator
	Patch(context.Context, *computepb.PatchRegionSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
	PatchRule(context.Context, *computepb.PatchRuleRegionSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
	RemoveRule(context.Context, *computepb.RemoveRuleRegionSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
}

// RegionSecurityPoliciesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The RegionSecurityPolicies API.
type RegionSecurityPoliciesClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionSecurityPoliciesClient

	// The call options for this service.
	CallOptions *RegionSecurityPoliciesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionSecurityPoliciesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionSecurityPoliciesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RegionSecurityPoliciesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddRule inserts a rule into a security policy.
func (c *RegionSecurityPoliciesClient) AddRule(ctx context.Context, req *computepb.AddRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.AddRule(ctx, req, opts...)
}

// Delete deletes the specified policy.
func (c *RegionSecurityPoliciesClient) Delete(ctx context.Context, req *computepb.DeleteRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get list all of the ordered rules present in a single specified policy.
func (c *RegionSecurityPoliciesClient) Get(ctx context.Context, req *computepb.GetRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetRule gets a rule at the specified priority.
func (c *RegionSecurityPoliciesClient) GetRule(ctx context.Context, req *computepb.GetRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicyRule, error) {
	return c.internalClient.GetRule(ctx, req, opts...)
}

// Insert creates a new policy in the specified project using the data included in the request.
func (c *RegionSecurityPoliciesClient) Insert(ctx context.Context, req *computepb.InsertRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List list all the policies that have been configured for the specified project and region.
func (c *RegionSecurityPoliciesClient) List(ctx context.Context, req *computepb.ListRegionSecurityPoliciesRequest, opts ...gax.CallOption) *SecurityPolicyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified policy with the data included in the request. To clear fields in the policy, leave the fields empty and specify them in the updateMask. This cannot be used to be update the rules in the policy. Please use the per rule methods like addRule, patchRule, and removeRule instead.
func (c *RegionSecurityPoliciesClient) Patch(ctx context.Context, req *computepb.PatchRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// PatchRule patches a rule at the specified priority. To clear fields in the rule, leave the fields empty and specify them in the updateMask.
func (c *RegionSecurityPoliciesClient) PatchRule(ctx context.Context, req *computepb.PatchRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.PatchRule(ctx, req, opts...)
}

// RemoveRule deletes a rule at the specified priority.
func (c *RegionSecurityPoliciesClient) RemoveRule(ctx context.Context, req *computepb.RemoveRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.RemoveRule(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionSecurityPoliciesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// operationClient is used to call the operation-specific management service.
	operationClient *RegionOperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing RegionSecurityPoliciesClient
	CallOptions **RegionSecurityPoliciesCallOptions
}

// NewRegionSecurityPoliciesRESTClient creates a new region security policies rest client.
//
// The RegionSecurityPolicies API.
func NewRegionSecurityPoliciesRESTClient(ctx context.Context, opts ...option.ClientOption) (*RegionSecurityPoliciesClient, error) {
	clientOpts := append(defaultRegionSecurityPoliciesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRegionSecurityPoliciesRESTCallOptions()
	c := &regionSecurityPoliciesRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	o := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opC, err := NewRegionOperationsRESTClient(ctx, o...)
	if err != nil {
		return nil, err
	}
	c.operationClient = opC

	return &RegionSecurityPoliciesClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRegionSecurityPoliciesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://compute.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionSecurityPoliciesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionSecurityPoliciesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	if err := c.operationClient.Close(); err != nil {
		return err
	}
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *regionSecurityPoliciesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AddRule inserts a rule into a security policy.
func (c *regionSecurityPoliciesRESTClient) AddRule(ctx context.Context, req *computepb.AddRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSecurityPolicyRuleResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/securityPolicies/%v/addRule", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.ValidateOnly != nil {
		params.Add("validateOnly", fmt.Sprintf("%v", req.GetValidateOnly()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).AddRule[0:len((*c.CallOptions).AddRule):len((*c.CallOptions).AddRule)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// Delete deletes the specified policy.
func (c *regionSecurityPoliciesRESTClient) Delete(ctx context.Context, req *computepb.DeleteRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/securityPolicies/%v", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// Get list all of the ordered rules present in a single specified policy.
func (c *regionSecurityPoliciesRESTClient) Get(ctx context.Context, req *computepb.GetRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicy, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/securityPolicies/%v", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.SecurityPolicy{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetRule gets a rule at the specified priority.
func (c *regionSecurityPoliciesRESTClient) GetRule(ctx context.Context, req *computepb.GetRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicyRule, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/securityPolicies/%v/getRule", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.Priority != nil {
		params.Add("priority", fmt.Sprintf("%v", req.GetPriority()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetRule[0:len((*c.CallOptions).GetRule):len((*c.CallOptions).GetRule)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.SecurityPolicyRule{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// Insert creates a new policy in the specified project using the data included in the request.
func (c *regionSecurityPoliciesRESTClient) Insert(ctx context.Context, req *computepb.InsertRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSecurityPolicyResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/securityPolicies", req.GetProject(), req.GetRegion())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}
	if req != nil && req.ValidateOnly != nil {
		params.Add("validateOnly", fmt.Sprintf("%v", req.GetValidateOnly()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// List list all the policies that have been configured for the specified project and region.
func (c *regionSecurityPoliciesRESTClient) List(ctx context.Context, req *computepb.ListRegionSecurityPoliciesRequest, opts ...gax.CallOption) *SecurityPolicyIterator {
	it := &SecurityPolicyIterator{}
	req = proto.Clone(req).(*computepb.ListRegionSecurityPoliciesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.SecurityPolicy, string, error) {
		resp := &computepb.SecurityPolicyList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/securityPolicies", req.GetProject(), req.GetRegion())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Patch patches the specified policy with the data included in the request. To clear fields in the policy, leave the fields empty and specify them in the updateMask. This cannot be used to be update the rules in the policy. Please use the per rule methods like addRule, patchRule, and removeRule instead.
func (c *regionSecurityPoliciesRESTClient) Patch(ctx context.Context, req *computepb.PatchRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSecurityPolicyResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/securityPolicies/%v", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}
	if req != nil && req.UpdateMask != nil {
		params.Add("updateMask", fmt.Sprintf("%v", req.GetUpdateMask()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// PatchRule patches a rule at the specified priority. To clear fields in the rule, leave the fields empty and specify them in the updateMask.
func (c *regionSecurityPoliciesRESTClient) PatchRule(ctx context.Context, req *computepb.PatchRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSecurityPolicyRuleResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/securityPolicies/%v/patchRule", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.Priority != nil {
		params.Add("priority", fmt.Sprintf("%v", req.GetPriority()))
	}
	if req != nil && req.UpdateMask != nil {
		params.Add("updateMask", fmt.Sprintf("%v", req.GetUpdateMask()))
	}
	if req != nil && req.ValidateOnly != nil {
		params.Add("validateOnly", fmt.Sprintf("%v", req.GetValidateOnly()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).PatchRule[0:len((*c.CallOptions).PatchRule):len((*c.CallOptions).PatchRule)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// RemoveRule deletes a rule at the specified priority.
func (c *regionSecurityPoliciesRESTClient) RemoveRule(ctx context.Context, req *computepb.RemoveRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/securityPolicies/%v/removeRule", req.GetProject(), req.GetRegion(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.Priority != nil {
		params.Add("priority", fmt.Sprintf("%v", req.GetPriority()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).RemoveRule[0:len((*c.CallOptions).RemoveRule):len((*c.CallOptions).RemoveRule)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}
