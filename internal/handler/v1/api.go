package v1

import (
	"context"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rigmas/joybox/internal/core/port"
	"github.com/rigmas/joybox/pkg/handler"
)

type ApiHandler struct {
	Host    string
	Echo    *echo.Echo
	BookSrv port.BookService
}

func NewApiHandler(host string,
	bookSrv port.BookService) handler.Handler {

	return &ApiHandler{
		Host:    host,
		BookSrv: bookSrv,
	}
}

func (a *ApiHandler) route() {
	route := a.Echo.Group("/api")
	a.Echo.Use(middleware.CORS())
	Route(route, a.BookSrv)
}

func (a *ApiHandler) startServer() {
	go func() {
		a.Echo.Start(a.Host)
	}()
}

func (a *ApiHandler) Stop() {
	a.Echo.Shutdown(context.TODO())
}

func (a *ApiHandler) Start() error {
	a.Echo = echo.New()
	a.route()
	a.startServer()

	return nil
}
