package app

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/configs/environment"
	customecho "github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/http/server/custom_echo"
	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/http/server/custom_echo/configs"
	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/services/weather-service/internal/zip_code/contracts/params"
	endpointsCity "github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/services/weather-service/internal/zip_code/features/getting_city_by_zip_code/v1/endpoints"
	endpointsWeather "github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/services/weather-service/internal/zip_code/features/getting_weather_by_zip_code/v1/endpoints"

	"github.com/labstack/echo/v4"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	env := environment.ProvideAppEnv()

	cfg, err := configs.ProvideConfig(string(env))
	if err != nil {
		panic(err)
	}

	e := customecho.NewEchoServer(cfg)
	e.ApplyVersioningFromHeader()

	e.GetEchoInstance().GET("//", func(c echo.Context) error {
		return c.String(http.StatusOK, "zip code service is running")
	})

	// routes params
	g := e.GetEchoInstance().Group("")
	v := validator.New()

	// endpoint city
	pc := params.NewZipCodeRouteParams(g, v)
	epc := endpointsCity.NewGetCityByZipCodeEndpoint(pc)
	epc.MapEndpoint()

	// endpoint weather
	pw := params.NewWeatherRouteParams(g, v)
	epw := endpointsWeather.NewGetWeatherByZipCodeEndpoint(pw)
	epw.MapEndpoint()

	// routes in json
	// routes, err := json.MarshalIndent(e.GetEchoInstance().Routes(), "", "  ")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// os.WriteFile("routes.json", routes, 0644)

	e.Start()
}
