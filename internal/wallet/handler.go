package wallet

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/safayildirim/wallet-management-service/internal/common"
	"github.com/safayildirim/wallet-management-service/internal/wallet/request"
	"net/http"
)

type Handler struct {
	walletService Service
}

func NewHandler(walletService Service) *Handler {
	return &Handler{walletService: walletService}
}

func (h Handler) RegisterRoutes(e *echo.Group) {
	e.POST("/wallets", h.CreateWallet)
	e.GET("/wallets/:id", h.GetWallet)
	e.DELETE("/wallets/:id", h.DeleteWallet)
}

// CreateWallet handles the creation of a new wallet.
//
// Parameters:
//   - ctx: The Echo context containing the HTTP request and response.
//
// Returns:
//   - 200 OK with the created wallet on success.
//   - 400 Bad Request if the request payload is invalid.
//   - 409 Conflict if a duplicate wallet is detected.
//   - 500 Internal Server Error for unexpected issues.
func (h Handler) CreateWallet(ctx echo.Context) error {
	var req request.CreateWalletRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := req.Validate(); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	wallet, err := h.walletService.CreateWallet(ctx.Request().Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, ErrDuplicateWallet):
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, Response{Data: wallet})
}

// GetWallet retrieves a wallet by its unique ID.
//
// Parameters:
//   - ctx: The Echo context containing the HTTP request and response.
//
// Returns:
//   - 200 OK with the wallet on success.
//   - 400 Bad Request if the ID is invalid.
//   - 500 Internal Server Error for unexpected issues.
func (h Handler) GetWallet(ctx echo.Context) error {
	id, err := common.ParseIntFromString[uint](ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	wallet, err := h.walletService.GetWallet(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, wallet)
}

// DeleteWallet deletes a wallet by its unique ID.
//
// Parameters:
//   - ctx: The Echo context containing the HTTP request and response.
//
// Returns:
//   - 204 No Content on success.
//   - 400 Bad Request if the ID is invalid.
//   - 500 Internal Server Error for unexpected issues.
func (h Handler) DeleteWallet(ctx echo.Context) error {
	id, err := common.ParseIntFromString[uint](ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.walletService.DeleteWallet(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
