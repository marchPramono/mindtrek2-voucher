package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func init() {

	connStr := "host=localhost port=5432 user=postgres password=nav123 dbname=mind1_voucher sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB connected...")
	}

}

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server connected!")
	})

	e.POST("/voucherType", addVoucherType)
	e.POST("/partner", addPartner)
	e.POST("/voucher", addVoucher)
	e.GET("/voucherType/{id}", getVoucherType)
	e.POST("/invoiceItem", addInvoiceItem)

	// e.POST("/invoice")

	// e.GET("/voucher", getAllVouchers)

	e.Logger.Fatal(e.Start(":2005"))
}
