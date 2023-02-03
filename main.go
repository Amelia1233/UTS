package main

import (
	"uts/connection"
	"uts/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := connection.ConnectToDb()
	rh := router.HandlerEndPoint{
		DB: db,
		R:  r,
	}

	rh.Routers()

	r.Run()
}