package main

import (
	"github.com/gin-gonic/gin"
	"heaven-door/service"
	"net/http"
)

func main() {
	r := gin.Default()

	Router(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

// Router
func Router(e *gin.Engine) {
	e.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
	srv := service.NewHeavenDoorService()
	e.GET("/dir", srv.ListDir)
	e.GET("/content", srv.Content)
}
