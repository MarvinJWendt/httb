package service

import (
	"github.com/labstack/echo/v4"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"github.com/marvinjwendt/httb/internal/pkg/random"
	"net/http"
)

func (s Service) GetJsonRandomLog(ctx echo.Context, params api.GetJsonRandomLogParams) error {
	if params.Count == nil {
		params.Count = new(int)
		*params.Count = 1
	}

	if params.LogLevels == nil {
		params.LogLevels = new([]string)
		*params.LogLevels = []string{"debug", "info", "warn", "error"}
	}

	if params.LogLevelWeights == nil {
		params.LogLevelWeights = new([]float32)
		*params.LogLevelWeights = []float32{1, 5, 2, 1}
	}

	if len(*params.LogLevels) != len(*params.LogLevelWeights) {
		return echo.NewHTTPError(http.StatusBadRequest, "logLevels and logLevelWeights must have the same length")
	}

	logLevels := make(map[string]float32)
	for i, level := range *params.LogLevels {
		logLevels[level] = (*params.LogLevelWeights)[i]
	}

	return ctx.String(http.StatusOK, random.NewLog(*params.Count, logLevels).String())
}

func (s Service) GetJsonRandomAddress(ctx echo.Context, params api.GetJsonRandomAddressParams) error {
	return ctx.JSON(http.StatusOK, random.Address())
}

func (s Service) GetJsonRandomContact(ctx echo.Context, params api.GetJsonRandomContactParams) error {
	return ctx.JSON(http.StatusOK, random.Contact())
}

func (s Service) GetJsonRandom(c echo.Context, params api.GetJsonRandomParams) error {
	panic("not implemented")
}

func (s Service) GetJsonRandomUser(ctx echo.Context, params api.GetJsonRandomUserParams) error {
	return ctx.JSON(http.StatusOK, random.User())
}
