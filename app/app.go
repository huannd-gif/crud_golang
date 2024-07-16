package app

import (
	"api_crud/app/command"
	"api_crud/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AddCall    command.AddCallHandler
	UpdateCall command.UpdateCallHandler
	DeleteCall command.DeleteCallHandler
}

type Queries struct {
	GetCalls       query.GetCallsHandler
	GetCallCreated query.GetCallCreatedHandler
}
