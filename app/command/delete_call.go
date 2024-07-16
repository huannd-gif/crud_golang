package command

import (
	"api_crud/common/decorator"
	"context"
	"github.com/sirupsen/logrus"
)

type DeleteCallHandler decorator.QueryHandler[int, bool]

type DeleteToDB interface {
	DeleteCall(ctx context.Context, id *int) error
}

type deleteCallHandler struct {
	repoCall DeleteToDB
}

func (d deleteCallHandler) Handle(ctx context.Context, id int) (bool, error) {
	err := d.repoCall.DeleteCall(ctx, &id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewDeleteCallHandler(
	repoCall DeleteToDB,
	logger *logrus.Entry,
) DeleteCallHandler {
	return decorator.ApplyQueryDecorators[int, bool](
		deleteCallHandler{repoCall: repoCall},
		logger,
	)
}
