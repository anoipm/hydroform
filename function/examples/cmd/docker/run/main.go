package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/client"
	"github.com/kyma-project/hydroform/function/pkg/docker"
	"github.com/kyma-project/hydroform/function/pkg/docker/runtimes"
	"github.com/kyma-project/hydroform/function/pkg/resources/types"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	runOpts := docker.RunOpts{
		Ports:         map[string]string{"8080": "8080"},
		Envs:          runtimes.ContainerEnvs(types.Nodejs12, false),
		ContainerName: "test123",
		Image:         runtimes.ContainerImage(types.Nodejs12),
		Commands:      runtimes.ContainerCommands(types.Nodejs12, false, false),
		User:          runtimes.ContainerUser,
	}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ctx := context.Background()
	_, err = docker.RunContainer(ctx, cli, runOpts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
