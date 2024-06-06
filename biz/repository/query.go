package repository

import (
	"context"
	"github.com/ByteBam/thirftbam/biz/model/gen"
	"github.com/ByteBam/thirftbam/biz/repository/query"
)

type CaptchaRepository interface {
	GetTokenByRDS(ctx context.Context, key string) (string, error)
}

type QueryRepository interface {
	GetBranchByID(ctx context.Context, id string) (string, error)
	CreateModule(ctx context.Context, m *gen.ModuleInfo) error
	CreateInterface(ctx context.Context, m *gen.InterfaceInfo) error
}

func (q *queryRepository) GetBranchByID(ctx context.Context, id string) (string, error) {
	branchInfo, err := query.Use(q.db).Branch.WithContext(ctx).Where(query.Use(q.db).Branch.ID.Eq(id)).First()
	if err != nil {
		return "", nil
	}

	return branchInfo.BranchName, nil
}

func (q *queryRepository) CreateModule(ctx context.Context, m *gen.ModuleInfo) error {
	if err := query.Use(q.db).ModuleInfo.WithContext(ctx).Create(m); err != nil {
		return err
	}
	return nil
}

func (q *queryRepository) CreateInterface(ctx context.Context, m *gen.InterfaceInfo) error {
	if err := query.Use(q.db).InterfaceInfo.WithContext(ctx).Create(m); err != nil {
		return err
	}
	return nil
}

type queryRepository struct {
	*Repository
}

func NewQueryRepository(
	r *Repository,
) QueryRepository {
	return &queryRepository{
		Repository: r,
	}
}

func NewCaptchaRepository(
	r *Repository,
) CaptchaRepository {
	return &queryRepository{
		Repository: r,
	}
}

func (q *queryRepository) GetTokenByRDS(ctx context.Context, query string) (string, error) {
	result, err := q.rdb.Get(ctx, query).Result()
	if err != nil {
		return "", err
	}
	return result, err
}
