package main_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/quii/go-specs-greet/adapters"
	"github.com/quii/go-specs-greet/adapters/httpserver"
	"github.com/quii/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port = "8080"
	)

	container := adapters.StartDockerServer(t, port, "httpserver")
	host, err := container.Host(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	mappedPort, err := container.MappedPort(context.Background(), port+"/tcp")
	if err != nil {
		t.Fatal(err)
	}

	driver := httpserver.Driver{
		BaseURL: fmt.Sprintf("http://%s:%s", host, mappedPort.Port()),
		Client:  &http.Client{Timeout: 1 * time.Second},
	}
	specifications.GreetSpecification(t, &driver)
}
