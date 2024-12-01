package handlers

import (
	"casbin/pkg"
	"casbin/repositories"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
)

type (
	UserHandler struct {
		r repositories.IUserRepository
		v *validator.Validate
	}

	userRequest struct {
		Name string                `json:"name"`
		Type repositories.UserType `json:"type"`
	}
)

func NewUserHandler(r repositories.IUserRepository) *UserHandler {
	return &UserHandler{r: r, v: validator.New()}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.r.GetUsers()
	render.JSON(w, r, users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctxv := r.Context().Value(pkg.ContextKey)
	if ctxv == nil {
		pkg.ErrorResp(w, r, http.StatusBadRequest, pkg.ErrContextValueNotFound)
		return
	}

	uid, err := strconv.Atoi(ctxv.(string))
	if err != nil {
		pkg.ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	user := h.r.GetUser(uid)
	pkg.SuccessResp(w, r, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := &userRequest{}
	if err := render.DecodeJSON(r.Body, req); err != nil {
		pkg.ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	if err := h.v.Struct(req); err != nil {
		pkg.ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	if err := h.r.CreateUser(repositories.User{Name: req.Name, Type: req.Type}); err != nil {
		pkg.ErrorResp(w, r, http.StatusInternalServerError, err)
		return
	}

	pkg.SuccessResp(w, r, map[string]string{"message": "user created"})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	if uid == "" {
		pkg.ErrorResp(w, r, http.StatusBadRequest, pkg.ErrPathParameterNotFound)
		return
	}

	id, err := strconv.Atoi(uid)
	if err != nil {
		pkg.ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	if err := h.r.DeleteUser(id); err != nil {
		pkg.ErrorResp(w, r, http.StatusInternalServerError, err)
		return
	}

	pkg.SuccessResp(w, r, map[string]string{"message": "user deleted"})
}
