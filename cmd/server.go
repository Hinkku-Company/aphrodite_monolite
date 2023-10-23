package main

import (
	"net"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/graphql/generated"
	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/graphql/resolvers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	config config.Config
	echo   *echo.Echo
}

func NewAPIServer(config config.Config) *server {
	return &server{
		config: config,
		echo:   echo.New(),
	}
}

func (s *server) Config() *server {
	s.echo.HideBanner = true
	s.echo.HidePort = true

	return s.cors().
		bodyLimit().
		timeout().
		rateLimiter().
		logger()
}

func (s *server) Run() {
	url := net.JoinHostPort("0.0.0.0", s.config.APPPort)
	logger.Log().Info("Starting server", "host", "http://"+url, "mode", s.config.AppENV)
	s.echo.Logger.Fatal(
		s.echo.Start(url),
	)
}

func (s *server) StartGraphql(resolver *resolvers.Resolver) *echo.Group {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: resolver,
		}),
	)

	group := s.echo.Group("/graphql")
	s.basicAuth(group)
	group.GET("/play", echo.WrapHandler(
		playground.Handler("GraphQL playground", "/graphql/query")),
	)
	group.POST("/query", echo.WrapHandler(srv))
	return group
}

func (s *server) StartRest() *echo.Group {
	group := s.echo.Group("/api")
	group.GET("/", func(c echo.Context) error {
		return c.JSON(200, "Rest API")
	})
	return group
}

func (s *server) cors() *server {
	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderAuthorization},
		AllowMethods:     []string{echo.GET, echo.POST},
		AllowCredentials: true,
	}))
	return s
}

func (s *server) basicAuth(group *echo.Group) {
	group.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/graphql/play"
		},
		Validator: func(username, password string, c echo.Context) (bool, error) {
			if username == "admin" && password == s.config.AdminPassword {
				return true, nil
			}
			return false, nil
		},
	}))
}

func (s *server) bodyLimit() *server {
	s.echo.Use(middleware.BodyLimit("2M"))
	return s
}

func (s *server) timeout() *server {
	s.echo.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
	return s
}

func (s *server) rateLimiter() *server {
	s.echo.Use(
		middleware.RateLimiter(
			middleware.NewRateLimiterMemoryStore(20)),
	)
	return s
}

func (s *server) logger() *server {
	log := logger.Log()
	s.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.Info("echo",
				"URI", values.URI,
				"status", values.Status,
				"start_time", values.StartTime,
				"latency", values.Latency,
				"host", values.Host,
				"method", values.Method,
				"remote_ip", values.RemoteIP,
				"user_agent", values.UserAgent)

			return nil
		},
	}))
	return s
}
