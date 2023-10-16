package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slowhigh/umm/configs"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
