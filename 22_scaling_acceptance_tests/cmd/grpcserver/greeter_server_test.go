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
		port           = "50051"
		dockerFilePath = "./cmd/grpcserver/Dockerfile"
	)

	container := adapters.StartDockerServer(t, port, dockerFilePath)
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
