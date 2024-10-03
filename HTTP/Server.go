package HTTP

import (
	"crypto/subtle"
	"dankey/Config"
	"dankey/DTO"
	"dankey/Storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"sync"
)

type Server struct {
	Provider Storage.Provider
	Echo     *echo.Echo
	conf     *Config.Config
}

func NewServer(provider Storage.Provider, config *Config.Config) *Server {
	return &Server{
		Provider: provider,
		Echo:     echo.New(),
		conf:     config,
	}
}

func (s *Server) Start() {
	wg := sync.WaitGroup{}
	configureLogger()
	s.configureEcho()
	s.setRoutes()
	s.startWithWaitGroup(&wg)
	testDankeyServer()
	wg.Wait()
}

func (s *Server) startWithWaitGroup(wg *sync.WaitGroup) {
	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		defer wg.Done()
		if err := s.Echo.Start(":6969"); err != nil {
			s.Echo.Logger.Error("Error starting Echo server: ", err)
		}
	}(wg)
}

func testDankeyServer() {
	_, err := http.Get("http://localhost:6969")
	if err != nil {
		log.Err(err).Msg("")
		log.Fatal().Msg("Dankey server failed to start up")
	}
	log.Info().Msg("Dankey server started up successfully")
}

func (s *Server) setRoutes() {
	s.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Dankey!")
	})
	s.Echo.GET("/routes", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, s.Echo.Routes(), "  ")
	})

	basicAuthGroup := s.Echo.Group("")
	basicAuthGroup.Use(middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(s.conf.Username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(s.conf.Password)) == 1 {
			return true, nil
		}
		return false, nil
	}))

	basicAuthGroup.PUT("/put", genHandlerFunc[DTO.PutRequestDTO, DTO.PutResponseDTO](s.Provider.Put))
	basicAuthGroup.GET("/get", genHandlerFunc[DTO.GetRequestDTO, DTO.GetResponseDTO](s.Provider.Get))
	basicAuthGroup.DELETE("/delete", genHandlerFunc[DTO.DeleteRequestDTO, DTO.DeleteResponseDTO](s.Provider.Delete))
	basicAuthGroup.POST("/increment", genHandlerFunc[DTO.IncrementRequestDTO, DTO.IncrementResponseDTO](s.Provider.Increment))
	basicAuthGroup.POST("/decrement", genHandlerFunc[DTO.DecrementRequestDTO, DTO.DecrementResponseDTO](s.Provider.Decrement))
	basicAuthGroup.POST("/saveToFile", genHandlerFunc[DTO.SaveToFileRequestDTO, DTO.SaveToFileResponseDTO](s.Provider.SaveToFile))
	basicAuthGroup.POST("/retrieveFromFile", genHandlerFunc[DTO.RetrieveFromFileRequestDTO, DTO.RetrieveFromFileResponseDTO](s.Provider.RetrieveFromFile))
}

func (s *Server) configureEcho() {
	s.Echo.HideBanner = true
	s.Echo.HidePort = true
}

func configureLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: zerolog.TimeFieldFormat,
	})
}

func genHandlerFunc[ReqT DTO.RequestDTOType, ResT DTO.ResponseDTOType](f func(ReqT) ResT) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ReqT
		err := c.Bind(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, DTO.ResponseDTO{
				Success: false,
				Message: err.Error(),
			})
		}

		res := f(req)
		return c.JSON(http.StatusOK, res)
	}
}
