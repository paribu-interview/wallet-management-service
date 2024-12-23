package wallet

import (
	"context"
	"github.com/safayildirim/wallet-management-service/internal/wallet/entity"
	"github.com/safayildirim/wallet-management-service/internal/wallet/request"
)

type Service interface {
	CreateWallet(ctx context.Context, request *request.CreateWalletRequest) (*entity.Wallet, error)
	GetWallet(ctx context.Context, id uint) (*entity.Wallet, error)
	DeleteWallet(ctx context.Context, id uint) error
}

type service struct {
	walletRepository Repository
}

func NewService(walletRepository Repository) Service {
	return &service{walletRepository: walletRepository}
}

// CreateWallet creates a new wallet with the provided details.
//
// Parameters:
//   - ctx: Context for managing request lifecycle and cancellation.
//   - request: Request object containing the wallet address and network details.
//
// Returns:
//   - The created wallet entity.
//   - An error if wallet creation fails.
func (s *service) CreateWallet(ctx context.Context, request *request.CreateWalletRequest) (*entity.Wallet, error) {
	// Map the request data to the Wallet entity
	item := entity.Wallet{
		Address: request.Address,
		Network: request.Network,
	}
	// Delegate wallet creation to the repository
	return s.walletRepository.CreateWallet(ctx, &item)
}

// GetWallet retrieves a wallet by its unique ID.
// Parameters:
//   - ctx: Context for managing request lifecycle and cancellation.
//   - id: The unique identifier of the wallet.
//
// Returns:
//   - The wallet entity if found.
//   - An error if the wallet does not exist or retrieval fails.
func (s *service) GetWallet(ctx context.Context, id uint) (*entity.Wallet, error) {
	return s.walletRepository.GetWallet(ctx, id)
}

// DeleteWallet deletes a wallet by its unique ID.
//
// Parameters:
//   - ctx: Context for managing request lifecycle and cancellation.
//   - id: The unique identifier of the wallet to delete.
//
// Returns:
//   - An error if the wallet does not exist or deletion fails.
func (s *service) DeleteWallet(ctx context.Context, id uint) error {
	return s.walletRepository.DeleteWallet(ctx, id)
}
