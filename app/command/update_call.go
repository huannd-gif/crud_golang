package command

import (
	"api_crud/common/decorator"
	"api_crud/domain"
	"context"
	"github.com/sirupsen/logrus"
)

type UpdateCallHandler decorator.QueryHandler[domain.Call, bool]

type UpdateCallToDB interface {
	UpdateCall(ctx context.Context, call *domain.Call) error
}

type updateCallHandler struct {
	callRepo UpdateCallToDB
}

func (u updateCallHandler) Handle(ctx context.Context, call domain.Call) (bool, error) {
	err := u.callRepo.UpdateCall(ctx, &call)

	if err != nil {
		return false, err
	}
	return true, nil
}

func NewUpdateCallHandler(
	callRepo UpdateCallToDB,
	logger *logrus.Entry,
) UpdateCallHandler {
	return decorator.ApplyQueryDecorators[domain.Call, bool](
		updateCallHandler{callRepo: callRepo},
		logger,
	)
}
