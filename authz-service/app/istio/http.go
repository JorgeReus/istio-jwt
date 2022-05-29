package istio

import (
	"authorization/core/validator"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
)

const (
	resultHeader   = "x-ext-authz-check-result"
	receivedHeader = "x-ext-authz-check-received"
	overrideHeader = "x-ext-authz-additional-header-override"
	resultAllowed  = "allowed"
	resultDenied   = "denied"
)

var (
	denyBody = fmt.Sprintf("Your role doesn't have the required permissions")
)

type extHttpAuthzServer struct {
	httpServer *http.Server
	port       uint16
}

// ServeHTTP implements the HTTP check request.
func (s *extHttpAuthzServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("[HTTP] read body failed: %v", err)
	}
	authorizationHeaders := request.Header["Authorization"]
	allowed := false
	for _, header := range authorizationHeaders {
		base64Token := strings.Split(header, "Bearer ")[1]
		// Only parse do not validate, that is taken care by istio itself
		if validator.IsJWTAuthorized(&base64Token) {
			allowed = true
			break
		}
	}
	l := fmt.Sprintf("%s %s%s, headers: %v, body: [%s]\n", request.Method, request.Host, request.URL, request.Header, body)
	if allowed {
		log.Printf("[HTTP][allowed]: %s", l)
		response.Header().Set(resultHeader, resultAllowed)
		response.Header().Set(overrideHeader, request.Header.Get(overrideHeader))
		response.Header().Set(receivedHeader, l)
		response.WriteHeader(http.StatusOK)
	} else {
		log.Printf("[HTTP][denied]: %s", l)
		response.Header().Set(resultHeader, resultDenied)
		response.Header().Set(overrideHeader, request.Header.Get(overrideHeader))
		response.Header().Set(receivedHeader, l)
		response.WriteHeader(http.StatusForbidden)
		_, _ = response.Write([]byte(denyBody))
	}
}

func (s *extHttpAuthzServer) Start(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		log.Printf("Stopped HTTP server")
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.port))
	if err != nil {
		log.Fatalf("Failed to create HTTP server: %v", err)
	}
	s.httpServer = &http.Server{Handler: s}

	log.Printf("Starting HTTP server at %s", listener.Addr())
	if err := s.httpServer.Serve(listener); err != http.ErrServerClosed {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func (s *extHttpAuthzServer) Stop() {
	log.Printf("HTTP server stopped: %v", s.httpServer.Close())
}

func NewHttpAuthorizer(port uint16) *extHttpAuthzServer {
	return &extHttpAuthzServer{
		port: port,
	}
}
