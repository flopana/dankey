package HTTP

import (
	"dankey/Config"
	"dankey/Storage"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Server struct {
	Provider     Storage.Provider
	Echo         *echo.Echo
	conf         *Config.Config
	requestCount *atomic.Uint64
}

// NewServer
//
//	@title						Dankey API
//	@version					1.0
//	@description				This is a simple key-value store that allows you to store any JSON-serializable data under a key.
//	@externalDocs.description	Visit the GitHub repository
//	@externalDocs.url			https://github.com/flopana/dankey
//	@license.name				MPL
//	@license.url				https://www.mozilla.org/en-US/MPL/2.0/
//	@securityDefinitions.basic	BasicAuth
func NewServer(provider Storage.Provider, config *Config.Config) *Server {
	requestCount := atomic.Uint64{}
	requestCount.Store(0)
	return &Server{
		Provider:     provider,
		Echo:         echo.New(),
		conf:         config,
		requestCount: &requestCount,
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
	const maxRetries = 5
	const delayBetweenRetries = 10 * time.Millisecond

	err := error(nil)

	for i := 0; i < maxRetries; i++ {
		err = s.sendTestRequestToDankey()
		if err == nil {
			log.Info().Msgf("Dankey server started successfully on :%s", s.conf.Port)
			log.Info().Msgf("Visit http://localhost:%s for the index page", s.conf.Port)
			return
		}
		time.Sleep(delayBetweenRetries)
	}

	log.Err(err).Msg("")
	log.Fatal().Msg("Dankey server failed to start up")
}

func (s *Server) sendTestRequestToDankey() error {
	_, err := http.Get("http://localhost:" + s.conf.Port)
	return err
}
