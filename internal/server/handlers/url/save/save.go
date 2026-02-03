package save

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gasuhwbab/url-shortener/internal/lib/api/response"
	"github.com/gasuhwbab/url-shortener/internal/lib/random"
	"github.com/gasuhwbab/url-shortener/internal/logger"
	"github.com/gasuhwbab/url-shortener/internal/storage"
	"github.com/go-chi/render"

	"github.com/go-playground/validator/v10"
)

type Request struct {
	URL   string `json:"url" validate:"required,url`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	Resp  response.Response
	Alias string
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) error
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.url.save.new"

		log = log.With(slog.String("op", op))

		var req Request

		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("failed to decode request body", logger.ErrorAttr(err))
			render.JSON(w, r, Response{Resp: response.Error("failed to decode request body")})
			return
		}

		log.Info("request body successfully decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", logger.ErrorAttr(err))
			render.JSON(w, r, Response{Resp: response.Error("invalid request")})
			return
		}
		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomStrinng(6)
		}
		if err := urlSaver.SaveURL(req.URL, alias); err != nil {
			if errors.Is(err, storage.ErrURLExists) {
				log.Info("url already exists", slog.String("url", req.URL))
				render.JSON(w, r, Response{Resp: response.Error("url already exists")})
				return
			}
			log.Error("failed to save url", logger.ErrorAttr(err))
			render.JSON(w, r, Response{Resp: response.Error("failed to decode save url")})
			return
		}
		log.Info("url successfully saved")
		render.JSON(w, r, Response{Resp: response.OK(), Alias: alias})
	}
}
