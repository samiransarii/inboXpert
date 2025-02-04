package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

var user