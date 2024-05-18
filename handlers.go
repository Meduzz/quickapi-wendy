package quickapiwendy

import (
	"log/slog"

	"github.com/Meduzz/quickapi-wendy/api"
	"github.com/Meduzz/wendy"
)

type (
	handler struct {
		logger  *slog.Logger
		storage *storage
	}

	ErrorDTO struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

const (
	codeBadJson = "BAD_JSON"
)

func newHandler(storage *storage) *handler {
	logger := slog.With("logger", "handler")
	return &handler{logger, storage}
}

func (h *handler) Create(r *wendy.Request) *wendy.Response {
	log := h.logger.With("method", "create")

	def := &api.Create{}
	err := r.Body.AsJson(def)

	if err != nil {
		log.Error("parsing json threw error", "error", err)
		return wendy.BadRequest(createError(codeBadJson, err))
	}

	res, err := h.storage.Create(def)

	if err != nil {
		log.Error("creating entity threw error", "error", err)
		return wendy.Error(createError("TODO", err))
	}

	return wendy.Ok(wendy.Json(res))
}

func (h *handler) Read(r *wendy.Request) *wendy.Response {
	log := h.logger.With("method", "read")

	def := &api.Read{}
	err := r.Body.AsJson(def)

	if err != nil {
		log.Error("parsing json threw error", "error", err)
		return wendy.BadRequest(createError(codeBadJson, err))
	}

	res, err := h.storage.Read(def)

	if err != nil {
		log.Error("creating entity threw error", "error", err)
		return wendy.Error(createError("TODO", err))
	}

	return wendy.Ok(wendy.Json(res))
}

func (h *handler) Update(r *wendy.Request) *wendy.Response {
	log := h.logger.With("method", "update")

	def := &api.Update{}
	err := r.Body.AsJson(def)

	if err != nil {
		log.Error("parsing json threw error", "error", err)
		return wendy.BadRequest(createError(codeBadJson, err))
	}

	res, err := h.storage.Update(def)

	if err != nil {
		log.Error("updating entity threw error", "error", err)
		return wendy.Error(createError("TODO", err))
	}

	return wendy.Ok(wendy.Json(res))
}

func (h *handler) Delete(r *wendy.Request) *wendy.Response {
	log := h.logger.With("method", "delete")

	def := &api.Delete{}
	err := r.Body.AsJson(def)

	if err != nil {
		log.Error("parsing json threw error", "error", err)
		return wendy.BadRequest(createError(codeBadJson, err))
	}

	err = h.storage.Delete(def)

	if err != nil {
		log.Error("deleting entity threw error", "error", err)
		return wendy.Error(createError("TODO", err))
	}

	return wendy.Ok(nil)
}

func (h *handler) Search(r *wendy.Request) *wendy.Response {
	log := h.logger.With("method", "search")

	def := &api.Search{}
	err := r.Body.AsJson(def)

	if err != nil {
		log.Error("parsing json threw error", "error", err)
		return wendy.BadRequest(createError(codeBadJson, err))
	}

	res, err := h.storage.Search(def)

	if err != nil {
		log.Error("searching for entities threw error", "error", err)
		return wendy.Error(createError("TODO", err))
	}

	return wendy.Ok(wendy.Json(res))
}

func (h *handler) Patch(r *wendy.Request) *wendy.Response {
	log := h.logger.With("method", "patch")

	def := &api.Patch{}
	err := r.Body.AsJson(def)

	if err != nil {
		log.Error("parsing json threw error", "error", err)
		return wendy.BadRequest(createError(codeBadJson, err))
	}

	res, err := h.storage.Patch(def)

	if err != nil {
		log.Error("patching entity threw error", "error", err)
		return wendy.Error(createError("TODO", err))
	}

	return wendy.Ok(wendy.Json(res))
}

func createError(code string, err error) *wendy.Body {
	return wendy.Json(&ErrorDTO{
		Code:    code,
		Message: err.Error(),
	})
}
