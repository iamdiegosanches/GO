package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) handleBook(c *gin.Context) error {
	if c.Request.Method == "GET" {
		return s.handleGetBook(c)
	}
	if c.Request.Method == "POST" {
		return s.handleCreateBook(c)
	}
	if c.Request.Method == "DELETE" {
		return s.handleDeleteBook(c)
	}

	return fmt.Errorf("method not allowed %s", c.Request.Method)
}

func (s *APIServer) handleGetBook(c *gin.Context) error {
	defaultBook := Book{
		ID:              uuid.New(),
		Name:            "My Book",
		Author:          "John Doe",
		PublicationDate: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		Publisher:       "Acme Publishing",
	}
	c.JSON(http.StatusOK, defaultBook)
	return nil
}

func (s *APIServer) handleCreateBook(c *gin.Context) error {
	return nil
}

func (s *APIServer) handleDeleteBook(c *gin.Context) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

// This allows to define functions that handle HTTP requests and return erros
// and then convert them into functions that handle HTTP request and handle errors internally
// This way, you don't have to repeat the error handling logic for every function that handles
// HTTP requests
func makeHTTPHandleFunc(f func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := f(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
