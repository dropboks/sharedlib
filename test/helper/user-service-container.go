package helper

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

type UserServiceContainer struct {
	Container testcontainers.Container
}

func StartUserServiceContainer(ctx context.Context, sharedNetwork, version string) (*UserServiceContainer, error) {
	image := fmt.Sprintf("user_service:%s", version)
	req := testcontainers.ContainerRequest{
		Name:         "user_service",
		Image:        image,
		ExposedPorts: []string{"50051:50051/tcp", "8182:8182/tcp"},
		Env:          map[string]string{"ENV": "test-dependence"},
		Networks:     []string{sharedNetwork},
		Cmd:          []string{"/user_service"},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return &UserServiceContainer{
		Container: container,
	}, nil
}

func (u *UserServiceContainer) Terminate(ctx context.Context) error {
	if u.Container != nil {
		return u.Container.Terminate(ctx)
	}
	return nil
}
