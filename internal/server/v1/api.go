package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/martinpaz/restfulapi/internal/data"
)

// New returns the API V1 Handler with configuration.
func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())

	pr := &PostRouter{
		Repository: &data.PostRepository{
			Data: data.New(),
		},
	}

	r.Mount("/posts", pr.Routes())

	rr := &RoleRouter{
		Repository: &data.RoleRepository{
			Data: data.New(),
		},
	}

	r.Mount("/role", rr.Routes())

	return r
}
