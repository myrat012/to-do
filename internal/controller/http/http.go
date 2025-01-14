package http

import (
	"fmt"
	"net/http"

	"github.com/myrat012/to-do/pkg/config"
)

// HTTP struct holds all the dependencies required for starting HTTP server
type HTTP struct {
	Server *http.Server
}

func NewService(cfg *config.Config) *HTTP {
	return &HTTP{
		Server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
			Handler: registerRouter(),
		},
	}

}
