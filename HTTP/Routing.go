package HTTP

import (
	"crypto/subtle"
	"dankey/DTO"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"net/http"

	_ "dankey/docs"
)

func (s *Server) setRoutes() {

	// Welcome message
	//
	//	@Summary		Welcome message
	//	@Description	Welcome to Dankey!
	//	@ID				welcome
	//	@Produce		html
	//	@Success		200	{string}	string
	//	@Router			/ [get]
	s.Echo.GET("/", func(c echo.Context) error {
		// return index.html file from the public directory
		return c.File("public/index.html")
	})

	s.Echo.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	s.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	basicAuthGroup := s.Echo.Group("")
	basicAuthGroup.Use(middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(s.conf.Username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(s.conf.Password)) == 1 {
			return true, nil
		}
		return false, nil
	}))

	basicAuthGroup.PUT("/put", s.put)
	basicAuthGroup.GET("/get", s.get)
	basicAuthGroup.DELETE("/delete", s.delete)
	basicAuthGroup.POST("/increment", s.increment)
	basicAuthGroup.POST("/decrement", s.decrement)
	basicAuthGroup.POST("/saveToFile", s.saveToFile)
	basicAuthGroup.POST("/retrieveFromFile", s.retrieveFromFile)
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

func generalHandlerFunc[ReqT DTO.RequestDTOType, ResT DTO.ResponseDTOType](c echo.Context, f func(ReqT) ResT) error {
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

// Put a key-value pair
//
//	@Summary		Put a key-value pair
//	@Description	Put any JSON-serializable data under a key
//	@ID				put
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DTO.PutRequestDTO	true	"PutRequestDTO"
//	@Success		200		{object}	DTO.PutResponseDTO
//	@Failure		401		{object}	DTO.ResponseDTO
//	@Router			/put [put]
func (s *Server) put(c echo.Context) error {
	return generalHandlerFunc(c, s.Provider.Put)
}

// Get a key-value pair
//
//	@Summary		Get a key-value pair
//	@Description	Retrieve the JSON Onject stored under a key
//	@ID				get
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DTO.GetRequestDTO	true	"GetRequestDTO"
//	@Success		200		{object}	DTO.GetResponseDTO
//	@Router			/get [get]
func (s *Server) get(c echo.Context) error {
	return generalHandlerFunc(c, s.Provider.Get)
}

// Delete a key-value pair
//
//	@Summary		Delete a key-value pair
//	@Description	Delete a key-value pair
//	@ID				delete
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DTO.DeleteRequestDTO	true	"DeleteRequestDTO"
//	@Success		200		{object}	DTO.DeleteResponseDTO
//	@Router			/delete [delete]
func (s *Server) delete(c echo.Context) error {
	return generalHandlerFunc(c, s.Provider.Delete)
}

// Increment a key-value pair
//
//	@Summary		Increment a key-value pair
//	@Description	Increment the value of a key by 1. If they value is not an integer, an error will be returned.
//	@ID				increment
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DTO.IncrementRequestDTO	true	"IncrementRequestDTO"
//	@Success		200		{object}	DTO.IncrementResponseDTO
//	@Router			/increment [post]
func (s *Server) increment(c echo.Context) error {
	return generalHandlerFunc(c, s.Provider.Increment)
}

// Decrement a key-value pair
//
//	@Summary		Decrement a key-value pair
//	@Description	Decrement the value of a key by 1. If they value is not an integer, an error will be returned.
//	@ID				decrement
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DTO.DecrementRequestDTO	true	"DecrementRequestDTO"
//	@Success		200		{object}	DTO.DecrementResponseDTO
//	@Router			/decrement [post]
func (s *Server) decrement(c echo.Context) error {
	return generalHandlerFunc(c, s.Provider.Decrement)
}

// Save to file
//
//	@Summary		Save to file
//	@Description	Save the current state of the database to a file
//	@ID				saveToFile
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DTO.SaveToFileRequestDTO	true	"SaveToFileRequestDTO"
//	@Success		200		{object}	DTO.SaveToFileResponseDTO
//	@Router			/saveToFile [post]
func (s *Server) saveToFile(c echo.Context) error {
	return generalHandlerFunc(c, s.Provider.SaveToFile)
}

// Retrieve from file
//
//	@Summary		Retrieve from file
//	@Description	Retrieve the state of the database from a file
//	@ID				retrieveFromFile
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DTO.RetrieveFromFileRequestDTO	true	"RetrieveFromFileRequestDTO"
//	@Success		200		{object}	DTO.RetrieveFromFileResponseDTO
//	@Router			/retrieveFromFile [post]
func (s *Server) retrieveFromFile(c echo.Context) error {
	return generalHandlerFunc(c, s.Provider.RetrieveFromFile)
}
