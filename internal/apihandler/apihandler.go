package apihandler

import (
	"net/http"
	"server/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Handler struct {
	conn *pgx.Conn
}

func createHandler(dbConfig *db.DBConfig) *Handler {
	conn := db.CreateDBClient(dbConfig)
	handler := Handler{conn: conn}
	return &handler
}

func (handler *Handler) ping(context *gin.Context) {
	response := gin.H{
		"message": "pong",
	}
	context.JSON(http.StatusOK, response)
}

func (handler *Handler) createUser(context *gin.Context) {
}

func (handler *Handler) getUser(context *gin.Context) {
}

func CreateEngine(dbConfig *db.DBConfig) *gin.Engine {
	engine := gin.Default()

	handler := createHandler(dbConfig)
	engine.GET("/ping", handler.ping)

	return engine
}
