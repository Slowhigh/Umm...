package repository

import (
	"context"

	"github.com/slowhigh/umm/internal/domain/entity"
)

type MemoRepository struct {
	db *[]entity.Memo
}

func NewMemoRepository(db *[]entity.Memo) entity.MemoRepository {
	return &MemoRepository{
		db: db,
	}
}

func (mr MemoRepository) Create(ctx context.Context, content string) (*entity.Memo, error) {
	panic("not implemented") // TODO: Implement
}

func (mr MemoRepository) FindOne(ctx context.Context, conditions entity.Memo) (*entity.Memo, error) {
	panic("not implemented") // TODO: Implement
}

func (mr MemoRepository) FindAll(ctx context.Context, conditions entity.Memo) (*[]entity.Memo, error) {
	panic("not implemented") // TODO: Implement
}

func (mr MemoRepository) Update(ctx context.Context, id entity.Memo, string entity.Memo, memo entity.Memo) (*entity.Memo, error) {
	panic("not implemented") // TODO: Implement
}

func (mr MemoRepository) Delete(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
