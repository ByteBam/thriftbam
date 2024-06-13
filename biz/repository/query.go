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
	GetGiteeIDByUserID(ctx context.Context, userId string) (int64, error)
	GetBranchInfoByID(ctx context.Context, branchId string) (*gen.Branch, error)
	UpdateBranchStatusByID(ctx context.Context, branchId string, status int) error
	UpdateBranchInterfaceNumByID(ctx context.Context, branchId string, nums int) error
	GetStoreLinkByID(ctx context.Context, id string) (string, error)
	CreateModule(ctx context.Context, m *gen.ModuleInfo) error
	CreateInterface(ctx context.Context, m *gen.InterfaceInfo) error
}

func (q *queryRepository) GetGiteeIDByUserID(ctx context.Context, userId string) (int64, error) {
	thirdUser, err := query.Use(q.db).LinkedThirdUser.WithContext(ctx).Where(query.Use(q.db).LinkedThirdUser.UserID.Eq(userId)).First()
	if err != nil {
		return 0, err
	}
	return thirdUser.GiteeID, nil
}

func (q *queryRepository) GetBranchInfoByID(ctx context.Context, branchId string) (*gen.Branch, error) {
	branch, err := query.Use(q.db).Branch.WithContext(ctx).Where(query.Use(q.db).Branch.ID.Eq(branchId)).First()
	if err != nil {
		return nil, err
	}
	return branch, nil
}

func (q *queryRepository) UpdateBranchStatusByID(ctx context.Context, branchId string, status int) error {
	raw, err := query.Use(q.db).Branch.WithContext(ctx).Where(query.Use(q.db).Branch.ID.Eq(branchId)).Update(query.Use(q.db).Branch.Status, status)
	if err != nil {
		return err
	}
	if raw.RowsAffected != 0 {
		return raw.Error
	}
	return nil
}

func (q *queryRepository) UpdateBranchInterfaceNumByID(ctx context.Context, branchId string, nums int) error {
	raw, err := query.Use(q.db).Branch.WithContext(ctx).Where(query.Use(q.db).Branch.ID.Eq(branchId)).Update(query.Use(q.db).Branch.InterfaceNum, nums)
	if err != nil {
		return err
	}
	if raw.RowsAffected != 0 {
		return raw.Error
	}
	return nil
}

func (q *queryRepository) GetStoreLinkByID(ctx context.Context, storeId string) (string, error) {
	storeInfo, err := query.Use(q.db).Store.WithContext(ctx).Where(query.Use(q.db).Store.ID.Eq(storeId)).First()
	if err != nil {
		return "", err
	}
	return storeInfo.StoreLink, nil
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

func (q *queryRepository) GetTokenByRDS(ctx context.Context, key string) (string, error) {
	result, err := q.rdb.HGet(ctx, key, "accessToken").Result()
	if err != nil {
		return "", err
	}
	return result, err
}
