package account

import (
	"github.com/labstack/echo/v4"

	request "github.com/game-core/gocrafter/auth/presentation/request/account"
	_ "github.com/game-core/gocrafter/auth/presentation/response/account"
	errorResponse "github.com/game-core/gocrafter/auth/presentation/response/error"
	accountService "github.com/game-core/gocrafter/domain/service/auth/account"
)

type AccountController interface {
	RegisterAccount() echo.HandlerFunc
}

type accountController struct {
	accountService accountService.AccountService
}

func NewAccountController(
	accountService accountService.AccountService,
) AccountController {
	return &accountController{
		accountService: accountService,
	}
}

// @tags    Account
// @Summary アカウント登録
// @Accept  json
// @Produce json
// @Param   body body request.RegisterAccount true "アカウント登録"
// @Router  /account/register_account [post]
// @Success 200 {object} account.RegisterAccount
// @Failure 500 {object} errorResponse.Error
func (a *accountController) RegisterAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &request.RegisterAccount{}
		c.Bind(req)

		res, err := a.accountService.RegisterAccount(req)
		if err != nil {
			return c.JSON(500, &errorResponse.Error{
				Status:       500,
				ErrorMessage: err.Error(),
			})
		}

		return c.JSON(200, res)
	}
}
