package main

import (
	"database/sql/driver"
	"fmt"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"

	"github.com/shopspring/decimal"
)

type OrderStatus string

const (
	statusNew      OrderStatus = "new"
	statusPending  OrderStatus = "pending"
	statusDone     OrderStatus = "done"
	statusRejected OrderStatus = "rejected"
)

var (
	empty    struct{}
	statuses = map[OrderStatus]struct{}{
		statusNew:      empty,
		statusPending:  empty,
		statusDone:     empty,
		statusRejected: empty,
	}
)

func (p *OrderStatus) Scan(value interface{}) error {
	*p = OrderStatus(value.(string))
	return nil
}

func (p OrderStatus) Value() (driver.Value, error) {
	return string(p), nil
}

func (p OrderStatus) Validate() bool {
	_, ok := statuses[p]
	return ok
}

type Order struct {
	gorm.Model
	ExternalID string          `gorm:"uniqueIndex" json:"external_id"`
	ClientID   string          `json:"client_id"`
	Amount     decimal.Decimal `json:"amount"`
	ContractID string          `json:"contract_id"`
	Status     OrderStatus     `sql:"type:order_status" json:"status"`
}

func main() {
	dbDSN := os.Getenv("DBDSN")
	fmt.Println(dbDSN)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbDSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{Logger: glog.Default.LogMode(glog.Info)})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Order{})
	if err != nil {
		fmt.Println(err)
		panic("cannot migrate db")
	}

	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		Debug:            true,
	}))
	r.Use(func(h http.Handler) http.Handler {
		return SetDBMiddleware(h, db)
	})

	r.Method("POST", "/order", Handler(PostSaveOrder))
	r.Method("GET", "/order", Handler(GetPaginateOrders))
	r.Method("GET", "/order/{externalID}", Handler(GetOrderByExternalID))
	r.Method("DELETE", "/order/{externalID}", Handler(DeleteOrderByExternalID))
	// /order/123123
	// /order/2414

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println(err)
		return
	}
}
