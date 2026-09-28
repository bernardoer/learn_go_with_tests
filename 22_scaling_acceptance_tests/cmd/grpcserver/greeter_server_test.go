package main_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/quii/go-specs-greet/adapters"
	"github.com/quii/go-specs-greet/adapters/grpcserver"
	"github.com/quii/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port       = "50051"
		binToBuild = "grpcserver"
	)

	container := adapters.StartDockerServer(t, port, binToBuild)
	host, err := container.Host(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	mappedPort, err := container.MappedPort(context.Background(), port+"/tcp")
	if err != nil {
		t.Fatal(err)
	}

	driver := grpcserver.Driver{Addr: fmt.Sprintf("%s:%s", host, mappedPort.Port())}
	specifications.GreetSpecification(t, &driver)
}
