package account

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-core/gocrafter/domain/service"
	"github.com/game-core/gocrafter/api/presentation/parameter"
	"github.com/game-core/gocrafter/api/presentation/output"
	"github.com/game-core/gocrafter/api/presentation/response"
	"github.com/game-core/gocrafter/api/presentation/controller"
)

type AccountController interface {
	Register() echo.HandlerFunc
}

type exampleController struct {
	accountService service.AccountService
}

func NewAccountController(accountService service.AccountService) AccountController {
    return &accountController{
        accountService: accountService,
    }
}

// @tags        Account
// @Summary     確認用
// @Accept      json
// @Produce     json
// @Param       example_key path string true "example_key" maxlength(20)
// @Router      /example/{example_key}/get_example [get]
// @Success     200  {object} response.Success{items=output.Example}
// @Failure     500  {array}  output.Error
func (a *accountController) GetExample() echo.HandlerFunc {
	return func(c echo.controller.Context) error {
		exampleKey := &parameter.ExampleKey{
			ExampleKey: c.Param("exampleKey"),
		}

		result, err := e.exampleService.FindByExampleKey(exampleKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("get_example", 500, err)

			return c.JSON(500, response)
		}

		out := output.ToExample(result)
		response := response.SuccessWith("get_example", 200, out)

		return c.JSON(200, response)
	}
}
