package server

import (
	"fmt"

	"github.com/erfanshekari/url-shortener/config"
	"github.com/erfanshekari/url-shortener/server/middlewares/throttle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e *echo.Echo
}

func (s *Server) preInit() {

	s.e = echo.New()

	s.e.HideBanner = true

	s.e.Use(middleware.Recover())

	conf := config.GetConfigInstance()

	if conf.Debug {
		s.e.Use(middleware.Logger())
	}

	if conf.Server.Throttle != (config.Throttle{}) {
		s.e.Use(throttle.GetThrottleMiddleware())
	}

	if conf.Server.GZip != (config.GZip{}) {
		s.e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: conf.Server.GZip.Level,
		}))
	}
}

func (s *Server) Init() {

	s.preInit()

	conf := config.GetConfigInstance()

	if conf.PublicDir != "" {
		s.e.Static("/*", conf.PublicDir)
	}

	s.e.POST("/add", addLink)

	s.e.GET("/l:slug", redirectHandler)
}

func (s *Server) Listen() error {
	conf := config.GetConfigInstance().Server

	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}

	if conf.Port == 0 {
		conf.Port = 5000
	}

	return s.e.Start(conf.Host + ":" + fmt.Sprint(conf.Port))
}
