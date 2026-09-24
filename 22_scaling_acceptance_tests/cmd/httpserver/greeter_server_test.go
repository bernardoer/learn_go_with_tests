package main_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	go_specs_greet "github.com/bernardoer/learn_go_with_tests/go-specs-greet"
	"github.com/bernardoer/learn_go_with_tests/go-specs-greet/specifications"
	"github.com/quii/go-graceful-shutdown/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../.",
			Dockerfile: "./cmd/httpserver/Dockerfile",
			// set to false if you want less spam, but this is helpful if you're having troubles
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	mappedPort, err := container.MappedPort(ctx, "8080/tcp")
	assert.NoError(t, err)
	baseURL := fmt.Sprintf("http://localhost:%s", mappedPort.Port())

	driver := go_specs_greet.Driver{BaseURL: baseURL, Client: &client}
	specifications.GreetSpecification(t, driver)
}
