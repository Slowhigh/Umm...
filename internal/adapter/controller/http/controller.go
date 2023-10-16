package http

import (
	"github.com/gin-gonic/gin"
	"github.com/slowhigh/umm/internal/adapter/controller/http/handler"
)

type Controller struct {
	router *gin.Engine
	memoHandler handler.MemoHandler
}

func NewController() {
	
}