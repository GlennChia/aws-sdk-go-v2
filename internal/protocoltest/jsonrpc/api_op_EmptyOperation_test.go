// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/smithy-go/middleware"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_EmptyOperation_awsAwsjson11Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *EmptyOperationInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Sends requests to /
		"sends_requests_to_slash": {
			Params:        &EmptyOperationInput{},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
		},
		// Includes X-Amz-Target header and Content-Type
		"includes_x_amz_target_and_content_type": {
			Params:        &EmptyOperationInput{},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.EmptyOperation"},
			},
		},
		// Clients must always send an empty JSON object payload for operations with no
		// input (that is, {}). While AWS service implementations support requests with no
		// payload or requests that send {}, always sending {} from the client is preferred
		// for forward compatibility in case input is ever added to an operation.
		"json_1_1_client_sends_empty_payload_for_no_input_shape": {
			Params:        &EmptyOperationInput{},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{}`))
			},
		},
		// Service implementations must support no payload or an empty object payload for
		// operations that define no input. However, despite the lack of a payload, a
		// Content-Type header is still required in order for the service to properly
		// detect the protocol.
		"json_1_1_service_supports_empty_payload_for_no_input_shape": {
			Params:        &EmptyOperationInput{},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if name == "json_1_1_service_supports_empty_payload_for_no_input_shape" {
				t.Skip("disabled test aws.protocoltests.json#JsonProtocol aws.protocoltests.json#EmptyOperation")
			}

			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			url := server.URL
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient: awshttp.NewBuildableClient(),
				Region:     "us-west-2",
			})
			result, err := client.EmptyOperation(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_EmptyOperation_awsAwsjson11Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *EmptyOperationOutput
	}{
		// When no output is defined, the service is expected to return an empty payload,
		// however, client must ignore a JSON payload if one is returned. This ensures that
		// if output is added later, then it will not break the client.
		"handles_empty_output_shape": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{}`),
			ExpectResult:  &EmptyOperationOutput{},
		},
		// This client-only test builds on handles_empty_output_shape, by including
		// unexpected fields in the JSON. A client needs to ignore JSON output that is
		// empty or that contains JSON object data.
		"handles_unexpected_json_output": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "foo": true
			}`),
			ExpectResult: &EmptyOperationOutput{},
		},
		// When no output is defined, the service is expected to return an empty payload.
		// Despite the lack of a payload, the service is expected to always send a
		// Content-Type header. Clients must handle cases where a service returns a JSON
		// object and where a service returns no JSON at all.
		"json_1_1_service_responds_with_no_payload": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			Body:         []byte(``),
			ExpectResult: &EmptyOperationOutput{},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2",
			})
			var params EmptyOperationInput
			result, err := client.EmptyOperation(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
