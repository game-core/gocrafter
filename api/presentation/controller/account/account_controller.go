package account

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-core/gocrafter/api/presentation/controller"
	request "github.com/game-core/gocrafter/api/presentation/request/account"
	response "github.com/game-core/gocrafter/api/presentation/response/account"
	errorResponse "github.com/game-core/gocrafter/api/presentation/response/error"
	accountService "github.com/game-core/gocrafter/domain/service/account"
)

type AccountController interface {
	RegisterAccount() echo.HandlerFunc
}

type exampleController struct {
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
// @Param       body body parameter.RegisterAccount true "アカウント登録"
// @Router      /account/register [post]
// @Success     200  {object} response.RegisterAccount
// @Failure     500  {array}  output.Error
func (a *accountController) RegisterAccount() echo.HandlerFunc {
	return func(c echo.controller.Context) error {
		param := &request.RegisterAccount{}
		c.controller.Bind(param)

		response, err := a.accountService.RegisterAccount(param)
		if err != nil {
			return c.JSON(500, errorResponse.ErrorResponse())
		}

		return c.JSON(200, response)
	}
}
