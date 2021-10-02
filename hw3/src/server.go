package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func SetDBMiddleware(next http.Handler, db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second)
		ctx := context.WithValue(r.Context(), "DB", db.WithContext(timeoutContext))
		next.ServeHTTP(w, r.WithContext(ctx))
		cancel()
	})
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func PostSaveOrder(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	db, ok := ctx.Value("DB").(*gorm.DB)
	if !ok {
		return APIErrorWithData(w, http.StatusInternalServerError, "no db")

	}
	order := Order{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&order); err != nil {
		return APIErrorWithData(w, http.StatusBadRequest, err)

	}

	tx := db.Create(&order)
	if tx.Error == nil {
		return APIOKWithStatus(w, http.StatusCreated, order)

	}
	if IsAlreadyExistedErr(tx.Error) {
		return APIOKWithStatus(w, http.StatusOK, "key already existed")

	}
	return APIErrorWithData(w, http.StatusInternalServerError, tx.Error)
}

func GetPaginateOrders(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	db, ok := ctx.Value("DB").(*gorm.DB)
	if !ok {
		return APIErrorWithData(w, http.StatusInternalServerError, "no db")
	}

	var orders []Order
	var tx *gorm.DB
	noPaginate := r.URL.Query().Get("no_pages") == "1"
	if noPaginate {
		tx = db.Find(&orders)
	} else {
		tx = db.Scopes(Paginate(r)).Find(&orders)
	}
	if tx.Error != nil {
		return APIErrorWithData(w, http.StatusInternalServerError, tx.Error)
	}

	response := struct {
		Count    int     `json:"count"`
		Orders   []Order `json:"orders"`
		Paginate bool    `json:"paginate"`
	}{len(orders), orders, !noPaginate}

	return APIOKWithStatus(w, http.StatusOK, response)
}

func GetOrderByExternalID(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	db, ok := ctx.Value("DB").(*gorm.DB)
	if !ok {
		return APIErrorWithData(w, http.StatusInternalServerError, "no db")
	}

	var order Order
	tx := db.Find(&order, &Order{ExternalID: chi.URLParam(r, "externalID")})

	if tx.Error != nil {
		return APIErrorWithData(w, http.StatusInternalServerError, tx.Error)
	}
	if tx.RowsAffected == 0 {
		return APIErrorWithData(w, http.StatusNotFound, "no such external id")
	}
	return APIOKWithStatus(w, http.StatusOK, order)
}

func DeleteOrderByExternalID(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	db, ok := ctx.Value("DB").(*gorm.DB)
	if !ok {
		return APIErrorWithData(w, http.StatusInternalServerError, "no db")
	}

	var order Order
	tx := db.Delete(&order, &Order{ExternalID: chi.URLParam(r, "externalID")})

	if tx.Error != nil {
		return APIErrorWithData(w, http.StatusInternalServerError, tx.Error)
	}
	if tx.RowsAffected == 0 {
		return APIErrorWithData(w, http.StatusNotFound, "no such external id")
	}
	return APIOKWithStatus(w, http.StatusOK, chi.URLParam(r, "externalID"))
}
