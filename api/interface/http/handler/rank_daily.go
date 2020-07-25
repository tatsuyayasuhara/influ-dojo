package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeDailyRankHandler(
	follower client.Follower,
	user repository.User,
	work repository.DailyWork,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.DailyRank{
			FollowerClient: follower,
			UserRepo:       user,
			DailyWorkRepo:  work,
		}

		out, err := in.GetDailyRank()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, out)
	}
}
