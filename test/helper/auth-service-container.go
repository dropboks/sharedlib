package helper

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

type AuthServiceContainer struct {
	Container testcontainers.Container
}

func StartAuthServiceContainer(ctx context.Context, sharedNetwork, version string) (*FileServiceContainer, error) {
	image := fmt.Sprintf("auth:%s", version)
	req := testcontainers.ContainerRequest{
		Name:         "file_service",
		Image:        image,
		ExposedPorts: []string{"50052:50052/tcp"},
		Env:          map[string]string{"ENV": "test-dependence"},
		Networks:     []string{sharedNetwork},
		Cmd:          []string{"/file_service"},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return &FileServiceContainer{
		Container: container,
	}, nil
}

func (a *AuthServiceContainer) Terminate(ctx context.Context) error {
	if a.Container != nil {
		return a.Container.Terminate(ctx)
	}
	return nil
}
