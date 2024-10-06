package HTTP

import (
	"dankey/Config"
	"dankey/Storage"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
)

type Server struct {
	Provider Storage.Provider
	Echo     *echo.Echo
	conf     *Config.Config
}

// NewServer
//
//	@title						Dankey API
//	@version					1.0
//	@description				This is a simple key-value store that allows you to store any JSON-serializable data under a key.
//	@license.name				MIT
//	@license.url				http://opensource.org/licenses/MIT
//	@securityDefinitions.basic	BasicAuth
func NewServer(provider Storage.Provider, config *Config.Config) *Server {
	return &Server{
		Provider: provider,
		Echo:     echo.New(),
		conf:     config,
	}
}

func (s *Server) Start() {
	log.Info().Msg("Starting Dankey server")
	wg := sync.WaitGroup{}
	s.configureEcho()
	s.setRoutes()
	s.startWithWaitGroup(&wg)
	s.testDankeyServer()
	wg.Wait()
}

func (s *Server) configureEcho() {
	s.Echo.HideBanner = true
	s.Echo.HidePort = true
}

func (s *Server) startWithWaitGroup(wg *sync.WaitGroup) {
	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		defer wg.Done()
		if err := s.Echo.Start(":" + s.conf.Port); err != nil {
			log.Err(err).Msg("")
			log.Fatal().Msg("Dankey server failed to start up")
		}
	}(wg)
}

func (s *Server) testDankeyServer() {
	_, err := http.Get("http://localhost:" + s.conf.Port)
	if err != nil {
		log.Err(err).Msg("")
		log.Fatal().Msg("Dankey server failed to start up")
	}
	log.Info().Msgf("Dankey server started successfully on :%s", s.conf.Port)
}
