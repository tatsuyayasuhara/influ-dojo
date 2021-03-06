package handler

import (
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeRankingHandler(
	work repository.Work,
	result repository.Result,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.Ranking{
			WorkRepo:   work,
			ResultRepo: result,
		}

		if err := in.Rank(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func Ranking(
	work repository.Work,
	result repository.Result,
) {
	in := input.Ranking{
		WorkRepo:   work,
		ResultRepo: result,
	}

	if err := in.Rank(); err != nil {
		log.Printf("%+v", err)
	}
}
