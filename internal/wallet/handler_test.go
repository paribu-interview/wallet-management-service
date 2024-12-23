package wallet

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/safayildirim/wallet-management-service/internal/wallet/entity"
	walletmock "github.com/safayildirim/wallet-management-service/internal/wallet/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandler_GetWallet(t *testing.T) {
	e := echo.New()

	tests := []struct {
		name                 string
		walletID             string
		mockService          bool
		mockReturnData       *entity.Wallet
		mockReturnErr        error
		expectedStatus       int
		expectErr            bool
		expectedErrorMessage string
	}{
		{
			name:        "when valid wallet id is provided then should return wallet",
			walletID:    "1",
			mockService: true,
			mockReturnData: &entity.Wallet{
				ID: 1, Address: "address1", Network: "network1",
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "when wallet not found then should return not found",
			walletID:       "1",
			mockService:    true,
			mockReturnData: &entity.Wallet{},
			mockReturnErr:  nil,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:                 "when invalid wallet id is provided then should return bad request",
			walletID:             "not-integer",
			mockService:          false,
			mockReturnData:       nil,
			mockReturnErr:        nil,
			expectedStatus:       http.StatusBadRequest,
			expectErr:            true,
			expectedErrorMessage: "invalid syntax",
		},
		{
			name:                 "when service returns error then should return internal server error",
			walletID:             "1",
			mockService:          true,
			mockReturnData:       nil,
			mockReturnErr:        errors.New("service error"),
			expectedStatus:       http.StatusInternalServerError,
			expectErr:            true,
			expectedErrorMessage: "service error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := walletmock.NewMockWalletService(t)
			handler := NewHandler(mockService)

			if tt.mockService {
				mockService.EXPECT().GetWallet(mock.Anything, mock.Anything).
					Return(tt.mockReturnData, tt.mockReturnErr).Once()
			}

			req := httptest.NewRequest(http.MethodGet, "/wallets/:id", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/wallets/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(fmt.Sprint(tt.walletID))

			err := handler.GetWallet(ctx)
			if tt.expectErr {
				assert.Error(t, err)
				httpErr := err.(*echo.HTTPError)
				assert.Contains(t, httpErr.Message, tt.expectedErrorMessage)
			} else {
				assert.NoError(t, err)

			}
		})
	}
}

func TestHandler_CreateWallet(t *testing.T) {
	e := echo.New()

	tests := []struct {
		name                 string
		body                 string
		mockService          bool
		mockReturn           *entity.Wallet
		mockError            error
		expectedStatus       int
		expectErr            bool
		expectedErrorMessage string
	}{
		{
			name:        "when valid request body is provided then should create wallet",
			body:        `{"address":"address1","network":"network1"}`,
			mockService: true,
			mockReturn: &entity.Wallet{
				ID:        1,
				CreatedAt: time.Time{},
				UpdatedAt: null.Time{},
				DeletedAt: null.Time{},
				Address:   "address1",
				Network:   "network1",
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
		},
		{
			name:                 "when invalid request body is provided then should return bad request",
			body:                 `{"address":123}`,
			mockReturn:           nil,
			mockError:            nil,
			expectedStatus:       http.StatusBadRequest,
			expectErr:            true,
			expectedErrorMessage: "Unmarshal type error",
		},
		{
			name:                 "when empty request body is provided then should return bad request",
			body:                 `{"address":""}`,
			mockReturn:           nil,
			mockError:            nil,
			expectedStatus:       http.StatusBadRequest,
			expectErr:            true,
			expectedErrorMessage: "address: cannot be blank",
		},
		{
			name:                 "when wallet already exists then should return conflict",
			body:                 `{"address":"address1","network":"network1"}`,
			mockService:          true,
			mockReturn:           nil,
			mockError:            ErrDuplicateWallet,
			expectedStatus:       http.StatusConflict,
			expectErr:            true,
			expectedErrorMessage: "wallet already exist",
		},
		{
			name:                 "when service returns error then should return internal server error",
			body:                 `{"address":"address1","network":"network1"}`,
			mockService:          true,
			mockReturn:           nil,
			mockError:            errors.New("internal server error"),
			expectedStatus:       http.StatusInternalServerError,
			expectErr:            true,
			expectedErrorMessage: "internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := walletmock.NewMockWalletService(t)
			handler := NewHandler(mockService)

			if tt.mockService {
				mockService.EXPECT().CreateWallet(mock.Anything, mock.Anything).
					Return(tt.mockReturn, tt.mockError).Once()
			}

			req := httptest.NewRequest(http.MethodPost, "/wallets", strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := handler.CreateWallet(ctx)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMessage)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestHandler_DeleteWallet(t *testing.T) {
	e := echo.New()

	tests := []struct {
		name                 string
		walletID             string
		mockService          bool
		mockReturnErr        error
		expectedStatus       int
		expectErr            bool
		expectedErrorMessage string
	}{
		{
			name:           "when valid request body is provided then should deposit asset",
			walletID:       "1",
			mockService:    true,
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:                 "when invalid wallet id is provided then should return bad request",
			walletID:             "not-integer",
			mockReturnErr:        nil,
			expectedStatus:       http.StatusBadRequest,
			expectErr:            true,
			expectedErrorMessage: "invalid syntax",
		},
		{
			name:                 "when service returns error then should return internal server error",
			walletID:             "1",
			mockService:          true,
			mockReturnErr:        errors.New("database error"),
			expectedStatus:       http.StatusInternalServerError,
			expectErr:            true,
			expectedErrorMessage: "database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := walletmock.NewMockWalletService(t)
			handler := NewHandler(mockService)

			if tt.mockService {
				mockService.EXPECT().DeleteWallet(mock.Anything, mock.Anything).
					Return(tt.mockReturnErr).Once()
			}

			req := httptest.NewRequest(http.MethodDelete, "/wallets/:id", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/wallets/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(fmt.Sprint(tt.walletID))

			err := handler.DeleteWallet(ctx)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMessage)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
