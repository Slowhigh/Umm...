package entity

import "context"

type Memo struct {
	ID      string
	Content string
}

type MemoRepository interface {
	Create(ctx context.Context, content string) (*Memo, error)
	FindOne(ctx context.Context, conditions Memo) (*Memo, error)
	FindAll(ctx context.Context, conditions Memo) (*[]Memo, error)
	Update(ctx context.Context, id, string, memo Memo) (*Memo, error)
	Delete(ctx context.Context, id string) error
}
