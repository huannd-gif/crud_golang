package command

import (
	"api_crud/common/decorator"
	"api_crud/domain"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type AddCallHandler decorator.QueryHandler[domain.Call, domain.Call]

type AddCallToDB interface {
	AddCall(ctx context.Context, call *domain.Call) error
}

type SendCallToRabbit interface {
	SendCall(ctx context.Context, call *domain.Call) error
}

type addCallHandler struct {
	callRepo   AddCallToDB
	callRabbit SendCallToRabbit
}

func (a addCallHandler) Handle(ctx context.Context, call domain.Call) (domain.Call, error) {
	err := a.callRepo.AddCall(ctx, &call)

	if err != nil {
		return *domain.NewCallNoArgument(), err
	}

	err = a.callRabbit.SendCall(ctx, &call)

	if err != nil {
		fmt.Println(err)
	}

	return call, nil

}

func NewAddCallHandle(
	callRepo AddCallToDB,
	logger *logrus.Entry,
	rabbitMq SendCallToRabbit, // cho nay la con tro ma de interface cung duoc ha
) AddCallHandler {
	return decorator.ApplyQueryDecorators[domain.Call, domain.Call](
		addCallHandler{callRepo: callRepo, callRabbit: rabbitMq},
		logger,
	)
}
