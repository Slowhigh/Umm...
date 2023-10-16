package usecase

import (
	"context"
	"time"

	"github.com/slowhigh/umm/internal/domain/entity"
)

type memoUsecase struct {
	memoRepo entity.MemoRepository
	timeout  time.Duration
}

type MemoUsecase interface {
	Create(ctx context.Context, content string) (*entity.Memo, error)
	FindOneByID(ctx context.Context, id string) (*entity.Memo, error)
	FindAllByContentKeyword(ctx context.Context, keyword string) (*[]entity.Memo, error)
	Update(ctx context.Context, id, content string) (*entity.Memo, error)
	Delete(ctx context.Context, id string) error
}

func NewMemoUsecase(userRepo entity.MemoRepository, timeout time.Duration) MemoUsecase {
	return &memoUsecase{
		memoRepo: userRepo,
		timeout:  timeout,
	}
}

func (mu memoUsecase) Create(ctx context.Context, content string) (*entity.Memo, error) {
	panic("not implemented") // TODO: Implement
}

func (mu memoUsecase) FindOneByID(ctx context.Context, id string) (*entity.Memo, error) {
	panic("not implemented") // TODO: Implement
}

func (mu memoUsecase) FindAllByContentKeyword(ctx context.Context, keyword string) (*[]entity.Memo, error) {
	panic("not implemented") // TODO: Implement
}

func (mu memoUsecase) Update(ctx context.Context, id string, content string) (*entity.Memo, error) {
	panic("not implemented") // TODO: Implement
}

func (mu memoUsecase) Delete(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
