package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-core/gocrafter/api/service"
	"github.com/game-core/gocrafter/api/presentation/parameter"
	"github.com/game-core/gocrafter/api/presentation/output"
	"github.com/game-core/gocrafter/api/presentation/response"
)

type ExampleController interface {
	GetExample() echo.HandlerFunc
}

type exampleController struct {
	exampleService service.ExampleService
}

func NewExampleController(exampleService service.ExampleService) ExampleController {
    return &exampleController{
        exampleService: exampleService,
    }
}

// @tags        Example
// @Summary     確認用
// @Accept      json
// @Produce     json
// @Param       example_key path string true "example_key" maxlength(20)
// @Router      /example/{example_key}/get_example [get]
// @Success     200  {object} response.Success{items=output.Example}
// @Failure     500  {array}  output.Error
func (e *exampleController) GetExample() echo.HandlerFunc {
	return func(c echo.Context) error {
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
