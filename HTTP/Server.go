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
	s.Echo.Use(middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte(s.conf.Username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(s.conf.Password)) == 1 {
			return true, nil
		}
		return false, nil
	}))

	s.Echo.PUT("/put", s.Put)
	s.Echo.GET("/get", s.Get)
	//e.DELETE("/delete", s.Delete)
	//e.GET("/increment", s.Increment)
	//e.GET("/decrement", s.Decrement)
	s.Echo.POST("/saveToFile", s.SaveToFile)
	s.Echo.POST("/retrieveFromFile", s.RetrieveFromFile)
	s.Echo.Logger.Fatal(s.Echo.Start(":6969"))
}

func (s *Server) Put(c echo.Context) error {
	var putRequestDTO DTO.PutRequestDTO
	err := c.Bind(&putRequestDTO)

	if err != nil {
		return c.JSON(http.StatusBadRequest, DTO.ResponseDTO{
			Success: false,
			Message: err.Error(),
		})
	}

	putResponseDTO := s.Provider.Put(putRequestDTO)
	return c.JSON(http.StatusOK, putResponseDTO)
}

func (s *Server) Get(c echo.Context) error {
	var getRequestDTO DTO.GetRequestDTO
	err := c.Bind(&getRequestDTO)

	if err != nil {
		return c.JSON(http.StatusBadRequest, DTO.ResponseDTO{
			Success: false,
			Message: err.Error(),
		})
	}

	getResponseDTO := s.Provider.Get(getRequestDTO)
	return c.JSON(http.StatusOK, getResponseDTO)
}

func (s *Server) SaveToFile(c echo.Context) error {
	var saveToFileRequestDTO DTO.SaveToFileRequestDTO
	err := c.Bind(&saveToFileRequestDTO)

	if err != nil {
		return c.JSON(http.StatusBadRequest, DTO.ResponseDTO{
			Success: false,
			Message: err.Error(),
		})
	}

	saveToFileResponseDTO := s.Provider.SaveToFile(saveToFileRequestDTO)
	return c.JSON(http.StatusOK, saveToFileResponseDTO)
}

func (s *Server) RetrieveFromFile(c echo.Context) error {
	var retrieveFromFileRequestDTO DTO.RetrieveFromFileRequestDTO
	err := c.Bind(&retrieveFromFileRequestDTO)

	if err != nil {
		return c.JSON(http.StatusBadRequest, DTO.ResponseDTO{
			Success: false,
			Message: err.Error(),
		})
	}

	retrieveFromFileResponseDTO := s.Provider.RetrieveFromFile(retrieveFromFileRequestDTO)
	return c.JSON(http.StatusOK, retrieveFromFileResponseDTO)
}
