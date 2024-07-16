package port

import (
	"api_crud/app"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type HttpServer struct {
	app *app.Application
}

func NewHttpServer(app *app.Application) *HttpServer {
	return &HttpServer{
		app: app,
	}
}

func (httpServer *HttpServer) GetListCall(ctx *gin.Context) {
	var requestCall ParamCallRequest

	if err := ctx.ShouldBindQuery(&requestCall); err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := httpServer.app.Queries.GetCalls.Handle(ctx, requestCall.extractToQuery())
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})

	return
}

func (httpServer *HttpServer) AddCall(ctx *gin.Context) {
	var callAdd AddCallRequest

	if err := ctx.ShouldBind(&callAdd); err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	callAdded, err := httpServer.app.Commands.AddCall.Handle(ctx, callAdd.extractToModelCreate())
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": callAdded,
	})

	return
}

func (httpServer *HttpServer) UpdateCall(ctx *gin.Context) {
	var callUpdate UpdateCallRequest
	if err := ctx.ShouldBind(&callUpdate); err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	idParam, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  false,
			"error": err.Error(),
		})
		return
	}

	result, err := httpServer.app.Commands.UpdateCall.Handle(ctx, callUpdate.extractToModelUpdate(idParam))

	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  result,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})

	return

}

func (httpServer *HttpServer) DeleteCall(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  false,
			"error": err.Error(),
		})
		return
	}

	result, err := httpServer.app.Commands.DeleteCall.Handle(ctx, id)
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  result,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})

	return

}
