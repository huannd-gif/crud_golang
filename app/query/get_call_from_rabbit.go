package query

import (
	"api_crud/domain"
	"context"
	"fmt"
)

type GetCallCreatedHandler interface {
	Handle()
}

type GetCallFromRabbit interface {
	GetCallAfterCreated(handleUpdate func(call domain.Call))
}

type UpdateCallToDB interface {
	UpdateCall(ctx context.Context, call *domain.Call) error
}

type getCallRabbitMQHandler struct {
	callRabbit GetCallFromRabbit
	callRepo   UpdateCallToDB
}

func (g getCallRabbitMQHandler) HandleUpdateCallRabbit(call domain.Call) {
	call.SetResult("SUCCESS")
	err := g.callRepo.UpdateCall(nil, &call)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("update success")
}

func (g getCallRabbitMQHandler) Handle() {
	g.callRabbit.GetCallAfterCreated(g.HandleUpdateCallRabbit)
}

func NewGetCallCreatedHandle(
	callRabbit GetCallFromRabbit,
	callRepo UpdateCallToDB,
) GetCallCreatedHandler {
	return getCallRabbitMQHandler{
		callRabbit: callRabbit,
		callRepo:   callRepo,
	}
}
