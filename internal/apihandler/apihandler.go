package apihandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func createHandler() *Handler {
	handler := Handler{}
	return &handler
}

func (handler *Handler) ping(context *gin.Context) {
	response := gin.H{
		"message": "pong",
	}
	context.JSON(http.StatusOK, response)
}

func CreateEngine() *gin.Engine {
	engine := gin.Default()

	handler := createHandler()
	engine.GET("/ping", handler.ping)

	return engine
}
