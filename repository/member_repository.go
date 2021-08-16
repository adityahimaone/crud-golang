package repository

import (
	"context"
	"onepage/entity"
)

type MemberRepository interface {
	FindAll(ctx context.Context) ([]entity.Fellow, error)
	Insert(ctx context.Context, member entity.Fellow) (entity.Fellow, error)
}
