package params

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ZipCodeRouteParams struct {
	RoutesGroup *echo.Group
	Validator   *validator.Validate
}

func NewZipCodeRouteParams(routesGroup *echo.Group, validator *validator.Validate) ZipCodeRouteParams {
	return ZipCodeRouteParams{
		RoutesGroup: routesGroup,
		Validator:   validator,
	}
}
