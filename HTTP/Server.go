package HTTP

import (
	"crypto/subtle"
	"dankey/Config"
	"dankey/DTO"
	"dankey/Storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
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
	s.Echo.HideBanner = true
	s.Echo.HidePort = true

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
	s.Echo.Logger.Fatal(s.Echo.Start(":6969"))
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
