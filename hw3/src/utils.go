package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/jackc/pgconn"
	"golang.org/x/xerrors"
)

func IsAlreadyExistedErr(err error) bool {
	var pgErr *pgconn.PgError
	if xerrors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}

const (
	APIStatusOK    = "ok"
	APIStatusError = "error"
)

type APIResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

func Response(statusCode int, w http.ResponseWriter, v interface{}) error {
	respBytes, err := json.Marshal(v)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err = w.Write(respBytes)

	return err
}

func APIOKWithStatus(w http.ResponseWriter, code int, data interface{}) error {
	resp := &APIResponse{
		Status: APIStatusOK,
		Code:   code,
		Data:   data,
	}

	return Response(code, w, resp)
}

func APIErrorWithData(w http.ResponseWriter, code int, data interface{}) error {
	resp := &APIResponse{
		Status: APIStatusError,
		Code:   code,
		Data:   data,
	}

	return Response(code, w, resp)
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		if offset < 0 {
			offset = 0
		}

		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		return db.Offset(offset).Limit(limit)
	}
}
