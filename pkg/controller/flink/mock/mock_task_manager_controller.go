package mock

import (
	"context"

	"github.com/lyft/flinkk8soperator/pkg/apis/app/v1beta1"
)

type CreateIfNotExistFunc func(ctx context.Context, application *v1beta1.FlinkApplication) (bool, error)

type TaskManagerController struct {
	CreateIfNotExistFunc CreateIfNotExistFunc
}

func (m *TaskManagerController) CreateIfNotExist(
	ctx context.Context, application *v1beta1.FlinkApplication) (bool, error) {
	if m.CreateIfNotExistFunc != nil {
		return m.CreateIfNotExistFunc(ctx, application)
	}
	return false, nil
}
