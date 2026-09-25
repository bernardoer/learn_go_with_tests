package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/bernardoer/learn_go_with_tests/go-specs-greet/adapters"
	httpserver "github.com/bernardoer/learn_go_with_tests/go-specs-greet/adapters"
	"github.com/bernardoer/learn_go_with_tests/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	// ctx := context.Background()

	// req := testcontainers.ContainerRequest{
	// 	FromDockerfile: testcontainers.FromDockerfile{
	// 		Context:    "../../.",
	// 		Dockerfile: "./cmd/httpserver/Dockerfile",
	// 		// set to false if you want less spam, but this is helpful if you're having troubles
	// 		PrintBuildLog: true,
	// 	},
	// 	ExposedPorts: []string{"8080/tcp"},
	// 	WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	// }
	// container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
	// 	ContainerRequest: req,
	// 	Started:          true,
	// })
	// assert.NoError(t, err)
	// t.Cleanup(func() {
	// 	assert.NoError(t, container.Terminate(ctx))
	// })

	// client := http.Client{
	// 	Timeout: 1 * time.Second,
	// }

	// mappedPort, err := container.MappedPort(ctx, "8080/tcp")
	// assert.NoError(t, err)
	// baseURL := fmt.Sprintf("http://localhost:%s", mappedPort.Port())

	// driver := httpserver.Driver{BaseURL: baseURL, Client: &client}
	// specifications.GreetSpecification(t, driver)

	var (
		port           = "8080"
		dockerFilePath = "./cmd/httpserver/Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}
