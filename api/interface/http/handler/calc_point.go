package handler

import (
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"log"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
)

func MakeCalcPointHandler(
	work repository.Work,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := &input.CalcPoint{
			Path:     path.Base(c.Path()),
			WorkRepo: work,
		}

		if err := in.CalcPoint(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func CalcPoint(
	work repository.Work,
	path string,
) {
	in := &input.CalcPoint{
		Path:     path,
		WorkRepo: work,
	}

	if err := in.CalcPoint(); err != nil {
		log.Printf("%+v", err)
	}
}
