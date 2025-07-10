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
	image := fmt.Sprintf("auth_service:%s", version)
	req := testcontainers.ContainerRequest{
		Name:         "auth_service",
		Image:        image,
		ExposedPorts: []string{"8081:8081/tcp"},
		Env:          map[string]string{"ENV": "test"},
		Networks:     []string{sharedNetwork},
		Cmd:          []string{"/auth_service"},
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
