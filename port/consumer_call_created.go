package port

import "api_crud/app"

type ConsumerReceiverCallCreated struct {
	app *app.Application
}

func NewConsumerReceiverCallCreated(app *app.Application) *ConsumerReceiverCallCreated {
	return &ConsumerReceiverCallCreated{
		app: app,
	}
}

func (c *ConsumerReceiverCallCreated) GetCallCreated() {
	c.app.Queries.GetCallCreated.Handle()
}
