package istio

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	authv2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	authv3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type V2GrpcReqBuilder func() *authv2.CheckRequest
type V3GrpcReqBuilder func() *authv3.CheckRequest

var grpcTests = []struct {
	name               string
	expectedStatusCode codes.Code
	v2reqBuilder       V2GrpcReqBuilder
	v3reqBuilder       V3GrpcReqBuilder
}{
	{"Test Valid GRPC request", codes.OK, func() *authv2.CheckRequest {
		return &authv2.CheckRequest{
			Attributes: &authv2.AttributeContext{
				Request: &authv2.AttributeContext_Request{
					Http: &authv2.AttributeContext_HttpRequest{
						Headers: map[string]string{
							// Headers are lowercased
							"authorization": fmt.Sprintf("Bearer %s", validJwt),
						},
					},
				},
			},
		}
	}, func() *authv3.CheckRequest {
		return &authv3.CheckRequest{
			Attributes: &authv3.AttributeContext{
				Request: &authv3.AttributeContext_Request{
					Http: &authv3.AttributeContext_HttpRequest{
						Headers: map[string]string{
							// Headers are lowercased
							"authorization": fmt.Sprintf("Bearer %s", validJwt),
						},
					},
				},
			},
		}
	}},
	{"Test inValid GRPC request", codes.PermissionDenied, func() *authv2.CheckRequest {
		return &authv2.CheckRequest{
			Attributes: &authv2.AttributeContext{
				Request: &authv2.AttributeContext_Request{
					Http: &authv2.AttributeContext_HttpRequest{
						Headers: map[string]string{
							// Headers are lowercased
							"authorization": fmt.Sprintf("Bearer %s", invalidJwt),
						},
					},
				},
			},
		}
	}, func() *authv3.CheckRequest {
		return &authv3.CheckRequest{
			Attributes: &authv3.AttributeContext{
				Request: &authv3.AttributeContext_Request{
					Http: &authv3.AttributeContext_HttpRequest{
						Headers: map[string]string{
							// Headers are lowercased
							"authorization": fmt.Sprintf("Bearer %s", invalidJwt),
						},
					},
				},
			},
		}
	}},
}

func TestGrpcAuthz(t *testing.T) {
	// Start the service in a user defined port, could be better :)
	var port uint16 = 9090
	grpcAuthz := NewGrpcAuthorizer(port)
	var wg sync.WaitGroup
	wg.Add(1)
	go grpcAuthz.Start(&wg)
	time.Sleep(500 * time.Millisecond)
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}
	v2Client := authv2.NewAuthorizationClient(conn)
	v3Client := authv3.NewAuthorizationClient(conn)
	defer conn.Close()
	for _, tt := range grpcTests {
		t.Run(tt.name, func(t *testing.T) {
			respv2, err := v2Client.Check(context.Background(), tt.v2reqBuilder())
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			respv3, err := v3Client.Check(context.Background(), tt.v3reqBuilder())
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			// Verify status code
			if respv2.Status.Code != int32(tt.expectedStatusCode) {
				t.Errorf("expected status code %v got %v", tt.expectedStatusCode, respv2.Status.Code)
			}
			if respv3.Status.Code != int32(tt.expectedStatusCode) {
				t.Errorf("expected status code %v got %v", tt.expectedStatusCode, respv3.Status.Code)
			}
		})
	}
	grpcAuthz.Stop()
	wg.Wait()
}
