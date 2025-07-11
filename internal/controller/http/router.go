package http

import (
	"net/http"

	"github.com/myrat012/to-do/internal/usecase"
)

func registerRouter(u *usecase.UseCases) *http.ServeMux {
	r := http.NewServeMux()
	userRegister(r, u.UserUseCase)
	return r
}
