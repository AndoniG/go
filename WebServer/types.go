package main

import (
	"net/http"
	"encoding/json"
)

type Server struct {
	port string
	router *Router
}

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func(u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

type MetaData interface{}