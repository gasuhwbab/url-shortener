package get

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gasuhwbab/url-shortener/internal/lib/api/response"
	"github.com/gasuhwbab/url-shortener/internal/logger"
	"github.com/gasuhwbab/url-shortener/internal/storage"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Alias string `json:"alias" validate:"required"`
}

type Response struct {
	Resp response.Response
	URL  string
}

type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.url.get.new"

		log = log.With(slog.String("op", op))

		var req Request
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("failed to decode request", logger.ErrorAttr(err))
			render.JSON(w, r, Response{Resp: response.Error("failed to decode request")})
			return
		}

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", logger.ErrorAttr(err))
			render.JSON(w, r, Response{Resp: response.Error("invalid request")})
			return
		}

		url, err := urlGetter.GetURL(req.Alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				log.Error("url not found", logger.ErrorAttr(err))
				render.JSON(w, r, Response{Resp: response.Error("url not found")})
				return
			}
			log.Error("failed to get url", logger.ErrorAttr(err))
			render.JSON(w, r, Response{Resp: response.Error("failed to get url")})
			return
		}
		render.JSON(w, r, Response{Resp: response.OK(), URL: url})
	}
}
