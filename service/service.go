package service

import (
	"api_crud/adapters"
	"api_crud/app"
	"api_crud/app/command"
	"api_crud/app/query"
	"api_crud/core/setting"
	"fmt"
	"github.com/sirupsen/logrus"
)

func NewApplication() *app.Application {
	return newApplication()
}

func newApplication() *app.Application {
	databaseSetting := setting.NewDatabaseSetting("config/app.env")
	rabbitMQSetting := setting.NewRabbitMQSetting("config/app.env")

	rb, err := adapters.NewRabbitConnection(*rabbitMQSetting)
	if err != nil {
		panic(fmt.Sprintf("cannot init rabbitmq connection: %v", err))
	}

	callRabbit := adapters.NewCallRabbitMQRepository(rb)
	//fmt.Println(callRabbit.Conn)

	db, err := adapters.NewMysqlConnection(*databaseSetting)
	if err != nil {
		panic(fmt.Sprintf("cannot init database connection: %v", err))
	}

	gormMigrator := adapters.NewGORMMigrator(db)

	callRepo := adapters.NewCallMysqlRepository(db, gormMigrator)

	err = gormMigrator.MakeMigrations()
	if err != nil {
		panic(fmt.Sprintf("something wrong when makemigration GORM: %v", err))
	}
	logger := logrus.NewEntry(logrus.StandardLogger())

	return &app.Application{
		Commands: app.Commands{
			AddCall:    command.NewAddCallHandle(callRepo, logger, callRabbit),
			UpdateCall: command.NewUpdateCallHandler(callRepo, logger),
			DeleteCall: command.NewDeleteCallHandler(callRepo, logger),
		},
		Queries: app.Queries{
			GetCalls:       query.NewGetCallsHandler(callRepo, logger),
			GetCallCreated: query.NewGetCallCreatedHandle(callRabbit, callRepo),
		},
	}
}
