package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	config "github.com/slowhigh/umm/internal/conf"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", cfg.Server.Auth.Port))
}
