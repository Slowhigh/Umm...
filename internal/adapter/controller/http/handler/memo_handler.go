package handler

import "github.com/slowhigh/umm/internal/application/usecase"

type MemoHandler struct {
	memoUsecase usecase.MemoUsecase
}

func NewMemoHandler()