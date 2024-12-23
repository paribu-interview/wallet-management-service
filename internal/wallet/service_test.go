package wallet

import (
	"context"
	"github.com/pkg/errors"
	"github.com/safayildirim/wallet-management-service/internal/wallet/entity"
	walletmock "github.com/safayildirim/wallet-management-service/internal/wallet/mock"
	"github.com/safayildirim/wallet-management-service/internal/wallet/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestService_CreateWallet(t *testing.T) {
	tests := []struct {
		name           string
		request        *request.CreateWalletRequest
		mockRepository bool
		mockReturn     *entity.Wallet
		mockError      error
		expectedResult *entity.Wallet
		expectedError  error
	}{
		{
			name: "when request is valid then should return wallet",
			request: &request.CreateWalletRequest{
				Address: "1A2B3C",
				Network: "Bitcoin",
			},
			mockRepository: true,
			mockReturn: &entity.Wallet{
				ID:      1,
				Address: "1A2B3C",
				Network: "Bitcoin",
			},
			mockError: nil,
			expectedResult: &entity.Wallet{
				ID:      1,
				Address: "1A2B3C",
				Network: "Bitcoin",
			},
			expectedError: nil,
		},
		{
			name: "when repository returns an error then should return error",
			request: &request.CreateWalletRequest{
				Address: "1A2B3C",
				Network: "Bitcoin",
			},
			mockRepository: true,
			mockReturn:     nil,
			mockError:      errors.New("repository error"),
			expectedResult: nil,
			expectedError:  errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := walletmock.NewMockWalletRepository(t)
			s := NewService(mockRepository)

			if tt.mockRepository {
				mockRepository.EXPECT().CreateWallet(mock.Anything, &entity.Wallet{
					Address: tt.request.Address,
					Network: tt.request.Network,
				}).Return(tt.mockReturn, tt.mockError).Once()
			}

			result, err := s.CreateWallet(context.Background(), tt.request)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

func TestService_GetWallet(t *testing.T) {
	tests := []struct {
		name           string
		walletID       uint
		mockRepository bool
		mockReturn     *entity.Wallet
		mockError      error
		expectedResult *entity.Wallet
		expectedError  error
	}{
		{
			name:           "when wallet is found then should return wallet",
			walletID:       1,
			mockRepository: true,
			mockReturn: &entity.Wallet{
				ID:      1,
				Address: "1A2B3C",
				Network: "Bitcoin",
			},
			mockError: nil,
			expectedResult: &entity.Wallet{
				ID:      1,
				Address: "1A2B3C",
				Network: "Bitcoin",
			},
			expectedError: nil,
		},
		{
			name:           "when wallet is not found then should return error",
			walletID:       2,
			mockRepository: true,
			mockReturn:     nil,
			mockError:      errors.New("wallet not found"),
			expectedResult: nil,
			expectedError:  errors.New("wallet not found"),
		},
		{
			name:           "when repository returns an error then should return error",
			walletID:       3,
			mockRepository: true,
			mockReturn:     nil,
			mockError:      errors.New("repository error"),
			expectedResult: nil,
			expectedError:  errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := walletmock.NewMockWalletRepository(t)
			s := NewService(mockRepository)

			if tt.mockRepository {
				mockRepository.EXPECT().GetWallet(mock.Anything, tt.walletID).Return(tt.mockReturn, tt.mockError).Once()
			}

			result, err := s.GetWallet(context.Background(), tt.walletID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

func TestService_DeleteWallet(t *testing.T) {
	tests := []struct {
		name           string
		walletID       uint
		mockRepository bool
		mockError      error
		expectedResult *entity.Wallet
		expectedError  error
	}{
		{
			name:           "when wallet is exists then should delete wallet",
			walletID:       1,
			mockRepository: true,
			mockError:      nil,
			expectedError:  nil,
		},
		{
			name:           "when repository returns an error then should return error",
			walletID:       3,
			mockRepository: true,
			mockError:      errors.New("repository error"),
			expectedResult: nil,
			expectedError:  errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := walletmock.NewMockWalletRepository(t)
			s := NewService(mockRepository)

			if tt.mockRepository {
				mockRepository.EXPECT().DeleteWallet(mock.Anything, tt.walletID).Return(tt.mockError).Once()
			}

			err := s.DeleteWallet(context.Background(), tt.walletID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
