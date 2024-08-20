package customecho

import (
	"context"
	"fmt"
	"log"

	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/weather_service/pkg/http/server/custom_echo/configs"
	"github.com/labstack/echo/v4"
)

type echoServer struct {
	echo   *echo.Echo
	config *configs.EchoOptions
}

type EchoServer interface {
	Start()
	Shutdown(ctx context.Context) error
	GetEchoInstance() *echo.Echo
	ApplyVersioningFromHeader()
}

func NewEchoServer(config *configs.EchoOptions) EchoServer {
	e := echo.New()
	e.HideBanner = false

	return &echoServer{
		echo:   e,
		config: config,
	}
}

func (e *echoServer) Start() {
	e.echo.Logger.Fatal(e.echo.Start(e.config.Port))
}

func (e *echoServer) Shutdown(ctx context.Context) error {
	err := e.echo.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (e *echoServer) GetEchoInstance() *echo.Echo {
	return e.echo
}

func (e *echoServer) ApplyVersioningFromHeader() {
	e.echo.Pre(apiVersion)
}

// APIVersion Header Based Versioning
func apiVersion(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		headers := req.Header

		apiVersion := headers.Get("version")

		if apiVersion == "" && req.URL.Path != "/" {
			apiVersion = "v1"
		}

		req.URL.Path = fmt.Sprintf("/%s%s", apiVersion, req.URL.Path)

		log.Println(req.URL.Path)

		return next(c)
	}
}
