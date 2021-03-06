package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeRecordHandler(
	follower client.Follower,
	user repository.User,
	work repository.Work,
	result repository.Result,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.Record{
			FollowerClient: follower,
			UserRepo:       user,
			WorkRepo:       work,
			ResultRepo:     result,
		}

		if err := in.Record(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func Record(
	follower client.Follower,
	user repository.User,
	work repository.Work,
	result repository.Result,
) {
	in := input.Record{
		FollowerClient: follower,
		UserRepo:       user,
		WorkRepo:       work,
		ResultRepo:     result,
	}

	if err := in.Record(); err != nil {
		log.Printf("%+v", err)
	}
}
