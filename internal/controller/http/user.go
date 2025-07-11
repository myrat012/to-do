package http

import (
	"fmt"
	"net/http"

	"github.com/myrat012/to-do/internal/usecase"
)

type userRoutes struct {
	u usecase.User
}

func userRegister(r *http.ServeMux, u usecase.User) {
	d := &userRoutes{u}

	r.HandleFunc("/api/login", d.login)

}

func (u *userRoutes) login(w http.ResponseWriter, r *http.Request) {
	h := "UserLogin"
	fmt.Fprint(w, h)
}
