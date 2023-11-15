package account

import (
	"github.com/labstack/echo/v4"
	
	_ "github.com/game-core/gocrafter/api/presentation/response/account"
	request "github.com/game-core/gocrafter/api/presentation/request/account"
	errorResponse "github.com/game-core/gocrafter/api/presentation/response/error"
	accountService "github.com/game-core/gocrafter/domain/service/account"
)

type AccountController interface {
	RegisterAccount() echo.HandlerFunc
	LoginAccount() echo.HandlerFunc
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

// @tags        Account
// @Summary     アカウント登録
// @Accept      json
// @Produce     json
// @Param       body body request.RegisterAccount true "アカウント登録"
// @Router      /account/register_account [post]
// @Success     200  {object} account.RegisterAccount
// @Failure     500  {object}  errorResponse.Error
func (a *accountController) RegisterAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := &request.RegisterAccount{}
		c.Bind(request)

		response, err := a.accountService.RegisterAccount(request)
		if err != nil {
			return c.JSON(500, &errorResponse.Error{
				Status: 500,
				ErrorMessage: "",
			})
		}

		return c.JSON(200, response)
	}
}

// @tags        Account
// @Summary     アカウントログイン
// @Accept      json
// @Produce     json
// @Param       body body request.LoginAccount true "アカウントログイン"
// @Router      /account/login_account [post]
// @Success     200  {object} account.LoginAccount
// @Failure     500  {object}  errorResponse.Error
func (a *accountController) LoginAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := &request.LoginAccount{}
		c.Bind(request)

		response, err := a.accountService.LoginAccount(request)
		if err != nil {
			return c.JSON(500, &errorResponse.Error{
				Status: 500,
				ErrorMessage: "",
			})
		}

		return c.JSON(200, response)
	}
}
