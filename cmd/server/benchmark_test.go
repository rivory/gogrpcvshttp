package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/rivory/gogrpcvshttp/handler"
	service "github.com/rivory/gogrpcvshttp/pkg/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// SETUP
// Importantly you need to call Run() once you've done what you need
func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

var test bool = false

func init() {
	go func() {
		serverHTTP := serverHTTP{
			handler: handler.ProvideHTTPHandler(),
		}
		http.HandleFunc("/hello", serverHTTP.handler.Handle)
		fmt.Printf("Starting http server at port 8080\n")
		log.Fatal(http.ListenAndServe(":8080", nil))
		test = true
	}()
}

func BenchmarkHTTP(b *testing.B) {
	client := &http.Client{}
	ts := httptest.NewServer(http.HandlerFunc(handler.ProvideHTTPHandler().Handle))

	body, err := json.Marshal(map[string]interface{}{
		"message": "tototototototototototest",
	})
	if err != nil {
		b.Fatalf("Expected no error, got %v", err)
	}

	req, err := http.NewRequest("POST", ts.URL+"/hello", bytes.NewBuffer(body))
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		resp, err := client.Do(req)
		if err != nil {
			b.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()
		bodyRes, _ := ioutil.ReadAll(resp.Body)
		var bodyArr (map[string]interface{})
		json.Unmarshal(bodyRes, &bodyArr)
		assert.Equal(b, "TOTOTOTOTOTOTOTOTOTOTEST", bodyArr["message"])
	}
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	service.RegisterServiceServer(server, handler.ProvideGrpcHandler())

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func BenchmarkGrpc(b *testing.B) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < b.N; i++ {
		r, err := service.NewServiceClient(conn).Handle(context.Background(), &service.HelloWorld{Message: "tototototototototototest"})
		if err != nil {
			b.Error("error: received", err)
		}
		assert.Equal(b, "TOTOTOTOTOTOTOTOTOTOTEST", r.GetMessage())
	}
}
