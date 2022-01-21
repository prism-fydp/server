package apihandler

import (
	"context"
	"fmt"
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

func (handler *Handler) ping(c *gin.Context) {
	// A simple JSON response
	response := gin.H{
		"message": "pong",
	}
	c.JSON(http.StatusOK, response)
}

func (handler *Handler) createUser(c *gin.Context) {
	// Get the name field from query parameters (ex: /users?name=jeffrey)
	name := c.PostForm("name")

	// Insert a new record into the users table and get the resulting ID
	var id int
	query := fmt.Sprintf("INSERT INTO users (name) VALUES ('%s') RETURNING id", name)
	err := handler.conn.QueryRow(context.Background(), query).Scan(&id)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Failed to insert user")
		return
	}

	// Return the ID back to the requester as JSON
	response := gin.H{
		"id": id,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *Handler) getUser(c *gin.Context) {
	// Get the id field from the URI (ex: /users/1)
	id := c.Param("id")

	// Query the database to find the name associated with this user ID
	var name string
	query := fmt.Sprintf("SELECT name FROM users WHERE id='%s'", id)
	err := handler.conn.QueryRow(context.Background(), query).Scan(&name)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Failed to fetch user")
		return
	}

	// Return the ID to the request if found
	response := gin.H{
		"name": name,
	}
	c.JSON(http.StatusOK, response)
}

func CreateEngine(dbConfig *db.DBConfig) *gin.Engine {
	engine := gin.Default()

	handler := createHandler(dbConfig)
	engine.GET("/ping", handler.ping)
	engine.POST("/users", handler.createUser)
	engine.GET("/users/:id", handler.getUser)

	return engine
}
