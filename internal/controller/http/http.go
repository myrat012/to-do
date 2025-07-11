package http

import (
	"context"
	"fmt"

	"net"
	"net/http"

	"github.com/myrat012/to-do/internal/model"
	"github.com/myrat012/to-do/internal/usecase"
	"github.com/myrat012/to-do/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type userAuthHttpWithLog func(ctx context.Context, zLog *zerolog.Logger, w http.ResponseWriter, r *http.Request, us *userSessionInfo)

type userSessionInfo struct {
	UserId   int            `json:"user_id"`
	Username string         `json:"username"`
	Role     model.UserRole `json:"role"`
}

// HTTP struct holds all the dependencies required for starting HTTP server
type HTTP struct {
	Server *http.Server
}

func NewService(cfg *config.Config, usecases *usecase.UseCases) *HTTP {
	return &HTTP{
		Server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
			Handler: registerRouter(usecases),
		},
	}

}

func requirementsForUser(handleName string, w http.ResponseWriter, r *http.Request, role model.UserRole) (bool, error) {
	ctx := r.Context()
	zerolog.Ctx(ctx).With().
		Str("remote-addr", GetRemoteAddres(r)).
		Str("uri", r.RequestURI).
		Str("method", r.Method).
		Str("handle", handleName)
	return false, nil
}

func GetRemoteAddres(r *http.Request) string {
	if val := r.Header.Get("X-Forwarded-For"); val != "" {
		return val
	} else if val := r.Header.Get("X-Real-IP"); val != "" {
		return val
	} else {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Error().Err(err).
				Str("remote-addr", r.RemoteAddr).
				Msg("error parsing remote addres")
			return "0.0.0.0"
		}
		return host
	}
}
