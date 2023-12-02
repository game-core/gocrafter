package example

import (
	"github.com/labstack/echo/v4"

	request "github.com/game-core/gocrafter/admin/presentation/request/example"
	errorResponse "github.com/game-core/gocrafter/admin/presentation/response/error"
	_ "github.com/game-core/gocrafter/admin/presentation/response/example"
	exampleService "github.com/game-core/gocrafter/domain/service/admin/example"
)

type ExampleController interface {
	GetExample() echo.HandlerFunc
}

type exampleController struct {
	exampleService exampleService.ExampleService
}

func NewExampleController(
	exampleService exampleService.ExampleService,
) ExampleController {
	return &exampleController{
		exampleService: exampleService,
	}
}

// @tags    Example
// @Summary アカウント確認
// @Accept  json
// @Produce json
// @Param   body body request.CheckExample true "アカウント確認"
// @Router  /example/get_example [post]
// @Success 200 {object} example.CheckExample
// @Failure 500 {object} errorResponse.Error
func (a *exampleController) GetExample() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &request.GetExample{}
		c.Bind(req)

		res, err := a.exampleService.GetExample(req)
		if err != nil {
			return c.JSON(500, &errorResponse.Error{
				Status:       500,
				ErrorMessage: err.Error(),
			})
		}

		return c.JSON(200, res)
	}
}
