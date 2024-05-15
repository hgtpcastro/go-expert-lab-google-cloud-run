package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	httpClient "github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/http/client"
	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/web/routes"
	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/services/weather-service/internal/service/weather/converter"
	zipcodeapi "github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/services/weather-service/internal/service/zip_code_api"
	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/services/weather-service/internal/zip_code/contracts/params"
	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/services/weather-service/internal/zip_code/features/getting_weather_by_zip_code/v1/dtos"

	"github.com/labstack/echo/v4"
)

type getWeatherByZipCodeEndpoint struct {
	params.WeatherRouteParams
}

func NewGetWeatherByZipCodeEndpoint(params params.WeatherRouteParams) routes.Endpoint {
	return &getWeatherByZipCodeEndpoint{WeatherRouteParams: params}
}

func (ep *getWeatherByZipCodeEndpoint) MapEndpoint() {
	ep.RoutesGroup.Group("/v1/weather").GET("/:zip_code", ep.handler())
}

func (ep *getWeatherByZipCodeEndpoint) handler() echo.HandlerFunc {
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

		dataCity, err := zipcodeapi.GetCityByZipCode(zipCode)
		if err != nil {
			return err
		}

		if dataCity.Erro {
			return c.JSON(http.StatusNotFound, "can not find zipcode")
		}

		dataWeather, err := getWeatherByZipCode(dataCity.Localidade)
		if err != nil {
			return err
		}

		tempC := dataWeather.Current.TempC
		tempF := converter.NewConverter().CelsiusToFahrenheit(tempC)
		tempK := converter.NewConverter().CelsiusToKelvin(tempC)

		response := map[string]float64{
			"temp_C": tempC,
			"temp_F": tempF,
			"temp_K": tempK,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func getWeatherByZipCode(city string) (dtos.GetWeatherByZipCodeEndpointDTO, error) {
	params := url.Values{}
	params.Add("q", strings.ToLower(city))
	api := os.Getenv("WEATHER_API_URL")
	url := fmt.Sprintf(api, params.Encode())
	// log.Println(url)

	client := httpClient.NewHttpClient()
	resp, err := client.R().Get(url)
	if err != nil {
		return dtos.GetWeatherByZipCodeEndpointDTO{}, err
	}

	var data dtos.GetWeatherByZipCodeEndpointDTO
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return dtos.GetWeatherByZipCodeEndpointDTO{}, err
	}

	return data, nil
}
