package main

import (
	"api_crud/domain"
	"api_crud/port"
	"api_crud/service"
	"bytes"
	"encoding/json"
	"gorm.io/gorm/utils/tests"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ResponseCreate struct {
	Data domain.Call `json:"data"`
}

type ResponseDelete struct {
	Data bool `json:"data"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	httpServer := port.NewHttpServer(service.NewApplication())
	v1 := router.Group("/call/")
	{
		v1.GET("/", httpServer.GetListCall)
		v1.POST("/", httpServer.AddCall)
		v1.PUT("/:id", httpServer.UpdateCall)
		v1.DELETE("/:id", httpServer.DeleteCall)
	}

	return router
}

func TestGetListCall(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/call/?page_num=1&page_size=15&phone_number=22222&metadata_display_field=name", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestAddCall(t *testing.T) {
	//GIVEN
	router := setupRouter()

	newItem := port.AddCallRequest{
		PhoneNumber: "02929292",
		Result:      "INIT",
		CallAt:      tests.Now(),
		EndAt:       tests.Now(),
		CallPress:   tests.Now(),
		ReceiverAt:  tests.Now(),
		Metadata: map[string]interface{}{
			"nghe": "da nghe",
		},
	}
	jsonValue, _ := json.Marshal(newItem)
	req, _ := http.NewRequest("POST", "/call/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	//WHEN
	router.ServeHTTP(w, req)

	//THEN

	assert.Equal(t, http.StatusOK, w.Code)

	var response ResponseCreate
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, newItem.PhoneNumber, response.Data.PhoneNumber)
	assert.Equal(t, newItem.Metadata, response.Data.Metadata)
	assert.Equal(t, string(newItem.Result), string(response.Data.Result))
	assert.Equal(t, newItem.CallAt.Truncate(time.Second), response.Data.CallAt.Truncate(time.Second))
}

func TestAddCallFailResult(t *testing.T) {
	//GIVEN
	router := setupRouter()

	newItem := port.AddCallRequest{
		PhoneNumber: "02929292",
		Result:      "koko",
		CallAt:      tests.Now(),
		EndAt:       tests.Now(),
		CallPress:   tests.Now(),
		ReceiverAt:  tests.Now(),
		Metadata: map[string]interface{}{
			"nghe": "da nghe",
		},
	}
	jsonValue, _ := json.Marshal(newItem)
	req, _ := http.NewRequest("POST", "/call/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	//WHEN
	router.ServeHTTP(w, req)

	//THEN

	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestUpdateCall(t *testing.T) {
	router := setupRouter()

	updatedItem := port.UpdateCallRequest{
		PhoneNumber: "02929292",
		Result:      "INIT",
		CallAt:      tests.Now(),
		EndAt:       tests.Now(),
		CallPress:   tests.Now(),
		ReceiverAt:  tests.Now(),
		Metadata: map[string]interface{}{
			"nghe": "update",
		},
	}
	jsonValue, _ := json.Marshal(updatedItem)
	req, _ := http.NewRequest("PUT", "/call/43", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response ResponseDelete
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, true, response.Data)
}

func TestUpdateCallFailValidateResult(t *testing.T) {
	router := setupRouter()

	updatedItem := port.UpdateCallRequest{
		PhoneNumber: "02929292",
		Result:      "cdcdcd",
		CallAt:      tests.Now(),
		EndAt:       tests.Now(),
		CallPress:   tests.Now(),
		ReceiverAt:  tests.Now(),
		Metadata: map[string]interface{}{
			"nghe": 2,
		},
	}
	jsonValue, _ := json.Marshal(updatedItem)
	req, _ := http.NewRequest("PUT", "/call/44", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ResponseDelete
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, false, response.Data)
}

func TestDeleteCall(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/call/42", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response ResponseDelete
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, true, response.Data)
}
