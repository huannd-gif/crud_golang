package query

import (
	"api_crud/common/decorator"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
)

type GetCallsHandler decorator.QueryHandler[CallRequest, ListCallsPaginated]

type CallRequest struct {
	PhoneNumber          string
	MetadataDisplayField string
	PageNum              int
	PageSize             int
}

type GetCallsReadModel interface {
	GetCalls(ctx context.Context, callRequest CallRequest) (ListCallsPaginated, error)
}

type getCallsHandler struct {
	callRepo GetCallsReadModel
}

func (g getCallsHandler) Handle(ctx context.Context, re CallRequest) (ListCallsPaginated, error) {
	calls, err := g.callRepo.GetCalls(ctx, re)

	if err != nil {
		return ListCallsPaginated{}, errors.New(err.Error())
	}
	return calls, nil
}

func NewGetCallsHandler(
	callRepo GetCallsReadModel,
	logger *logrus.Entry,
) GetCallsHandler {
	return decorator.ApplyQueryDecorators[CallRequest, ListCallsPaginated](
		getCallsHandler{callRepo: callRepo},
		logger,
	)
}
