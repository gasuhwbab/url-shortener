package delete

import (
	"log/slog"
	"net/http"

	resp "github.com/gasuhwbab/url-shortener/internal/lib/api/response"
	"github.com/gasuhwbab/url-shortener/internal/logger"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Alias string `json:"alias" validate:"required"`
}

type Response struct {
	Resp resp.Response
}

type URLDeleter interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.url.delete.new"

		log = log.With(slog.String("op", op))

		var req Request
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("failed to decode request", logger.ErrorAttr(err))
			render.JSON(w, r, Response{resp.Error("failed to decode request")})
			return
		}

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", logger.ErrorAttr(err))
			render.JSON(w, r, Response{resp.Error("invalid request")})
			return
		}

		if err := urlDeleter.DeleteURL(req.Alias); err != nil {
			log.Error("failed to delete alias", logger.ErrorAttr(err))
			render.JSON(w, r, Response{resp.Error("failed to delete alilas")})
			return
		}

		render.JSON(w, r, Response{resp.OK()})
	}
}
