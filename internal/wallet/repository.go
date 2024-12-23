package wallet

import (
	"context"
	"github.com/safayildirim/wallet-management-service/internal/wallet/entity"
	"gorm.io/gorm"
	"strings"
)

type Repository interface {
	CreateWallet(ctx context.Context, entity *entity.Wallet) (*entity.Wallet, error)
	GetWallet(ctx context.Context, id uint) (*entity.Wallet, error)
	DeleteWallet(ctx context.Context, id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateWallet(ctx context.Context, entity *entity.Wallet) (*entity.Wallet, error) {
	err := r.db.WithContext(ctx).Create(entity).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, ErrDuplicateWallet
		}

		return nil, err
	}

	return entity, nil
}

func (r *repository) GetWallet(ctx context.Context, id uint) (*entity.Wallet, error) {
	var item entity.Wallet
	err := r.db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *repository) DeleteWallet(ctx context.Context, id uint) error {
	var item entity.Wallet
	err := r.db.WithContext(ctx).Delete(&item, id).Error
	if err != nil {
		return err
	}

	return nil
}
