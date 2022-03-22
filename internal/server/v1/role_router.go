package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/martinpaz/restfulapi/pkg/response"
	"github.com/martinpaz/restfulapi/pkg/role"
)

// UserRouter is the router of the users.
type RoleRouter struct {
	Repository role.Repository
}

// Routes returns user router with each endpoint.
func (rr *RoleRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/{id}", rr.GetOneHandler)

	return r
}

func (rr *RoleRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	role, err := rr.Repository.GetRole(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	/*storedRoles, err := rr.Repository.GetAll(ctx)

	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	var roles_string []string
	for r := range role {
		roles_string = append(roles_string, storedRoles[r.role_id].Name)
	}*/

	response.JSON(w, r, http.StatusOK, response.Map{"role": role})
}
