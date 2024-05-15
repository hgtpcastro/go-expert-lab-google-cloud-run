package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/web/routes"
	zipcodeapi "github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/services/weather-service/internal/service/zip_code_api"
	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/services/weather-service/internal/zip_code/contracts/params"
)

type getCityByZipCodeEndpoint struct {
	params.ZipCodeRouteParams
}

func NewGetCityByZipCodeEndpoint(params params.ZipCodeRouteParams) routes.Endpoint {
	return &getCityByZipCodeEndpoint{ZipCodeRouteParams: params}
}

func (ep *getCityByZipCodeEndpoint) MapEndpoint() {
	ep.RoutesGroup.Group("/v1/zip-code").GET("/:zip_code", ep.handler())
}

func (ep *getCityByZipCodeEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		zipCode := c.Param("zip_code")

		if zipCode == "" {
			return c.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
		}

		ep.Validator.SetTagName("zip_code")
		err := ep.Validator.Var(zipCode, "required,len=8,number")
		if err != nil {
			// fmt.Println(err)
			return c.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
		}

		data, err := zipcodeapi.GetCityByZipCode(zipCode)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if data.Erro {
			return c.JSON(http.StatusNotFound, "can not find zipcode")
		}

		return c.JSON(http.StatusOK, data)
	}
}
