package loginReward

import (
	"github.com/labstack/echo/v4"
	"time"

	request "github.com/game-core/gocrafter/api/presentation/request/loginReward"
	errorResponse "github.com/game-core/gocrafter/api/presentation/response/error"
	_ "github.com/game-core/gocrafter/api/presentation/response/loginReward"
	loginRewardService "github.com/game-core/gocrafter/domain/service/loginReward"
)

type LoginRewardController interface {
	ReceiveLoginReward() echo.HandlerFunc
}

type loginRewardController struct {
	loginRewardService loginRewardService.LoginRewardService
}

func NewLoginRewardController(
	loginRewardService loginRewardService.LoginRewardService,
) LoginRewardController {
	return &loginRewardController{
		loginRewardService: loginRewardService,
	}
}

// @tags    LoginReward
// @Summary ログイン報酬受け取り
// @Accept  json
// @Produce json
// @Param   body body request.ReceiveLoginReward true "ログイン報酬受け取り"
// @Router  /loginReward/receive_loginReward [post]
// @Success 200 {object} loginReward.ReceiveLoginReward
// @Failure 500 {object} errorResponse.Error
func (a *loginRewardController) ReceiveLoginReward() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &request.ReceiveLoginReward{}
		c.Bind(req)

		res, err := a.loginRewardService.ReceiveLoginReward(req, time.Now())
		if err != nil {
			return c.JSON(500, &errorResponse.Error{
				Status:       500,
				ErrorMessage: err.Error(),
			})
		}

		return c.JSON(200, res)
	}
}
