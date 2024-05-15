package params

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type WeatherRouteParams struct {
	RoutesGroup *echo.Group
	Validator   *validator.Validate
}

func NewWeatherRouteParams(routesGroup *echo.Group, validator *validator.Validate) WeatherRouteParams {
	return WeatherRouteParams{
		RoutesGroup: routesGroup,
		Validator:   validator,
	}
}
