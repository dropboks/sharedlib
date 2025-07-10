package helper

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

type FileServiceContainer struct {
	Container testcontainers.Container
}

func StartFileServiceContainer(ctx context.Context, sharedNetwork, version string) (*FileServiceContainer, error) {
	image := fmt.Sprintf("file_service:%s", version)
	req := testcontainers.ContainerRequest{
		Name:         "file_service",
		Image:        image,
		ExposedPorts: []string{"50052:50051/tcp"},
		Env:          map[string]string{"ENV": "test"},
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

func (f *FileServiceContainer) Terminate(ctx context.Context) error {
	if f.Container != nil {
		return f.Container.Terminate(ctx)
	}
	return nil
}
