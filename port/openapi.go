package port

import (
	"api_crud/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

//func (g *GetOneCommand) Handle(ctx context.Context) error {
//
//}
//
//type (
//	GetOneCommand struct {
//	}
//	Application struct {
//		Query struct {
//			GetOneHandler GetOneCommand
//		}
//		Command struct {
//		}
//	}
//	ServerImpl struct {
//		app Application
//	}
//
//	Prototype interface {
//		CalculateArea() int
//	}
//
//	Rectangle struct {
//		height int
//		width  int
//	}
//
//	Triangle struct {
//		height int
//		width  int
//	}
//
//	CalcualteAreaHandler struct {
//		CalcuatePrototype Prototype
//	}
//)
//
//func (r *Rectangle) CalculateArea() int {
//	return r.height * r.width
//}
//
//func (t *Triangle) CalculateArea() int {
//	return t.height * t.width
//}
//
//func NewApplication() *Application {
//	return &Application{}
//}

func HttpApp() {
	r := gin.Default()
	httpServer := NewHttpServer(service.NewApplication())

	v1 := r.Group("/call")
	{
		v1.GET("/", httpServer.GetListCall)
		v1.POST("/", httpServer.AddCall)
		v1.PUT("/:id", httpServer.UpdateCall)
		v1.DELETE("/:id", httpServer.DeleteCall)
	}

	err := r.Run()
	if err != nil {

		panic(fmt.Sprintf("cannot init database connection: %v", err))
	}

}

func RabbitMQApp() {
	consumerCallCreated := NewConsumerReceiverCallCreated(service.NewApplication())
	consumerCallCreated.GetCallCreated()
}
