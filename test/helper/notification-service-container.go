package helper

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

type NotificationServiceContainer struct {
	Container testcontainers.Container
}

func StartNotificationServiceContainer(ctx context.Context, sharedNetwork, version string) (*NotificationServiceContainer, error) {
	image := fmt.Sprintf("notification_service:%s", version)
	req := testcontainers.ContainerRequest{
		Name:     "notification_service",
		Image:    image,
		Env:      map[string]string{"ENV": "test-dependence"},
		Networks: []string{sharedNetwork},
		Cmd:      []string{"/notification_service"},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return &NotificationServiceContainer{
		Container: container,
	}, nil
}

func (f *NotificationServiceContainer) Terminate(ctx context.Context) error {
	if f.Container != nil {
		return f.Container.Terminate(ctx)
	}
	return nil
}
