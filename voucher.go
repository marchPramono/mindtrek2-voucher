package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// VoucherType hold data for voucher_type table
type VoucherType struct {
	VoucherTypeID string `json:"voucher_type_id"`
	VoucherName   string `json:"voucher_name"`
	Nominal       string `json:"nominal"`
	DurationMonth string `json:"duration_month"`
	CreatedAt     string `json:"created_at"`
}

// VouchersType hold vouchers_type slice
type VouchersType struct {
	VouchersType []VoucherType `json:"vouchers_type"`
}

// Voucher hold data for voucher table
type Voucher struct {
	TypeID      string `json:"type_id"`
	VoucherCode string `json:"voucher_code"`
	CreatedAt   string `json:"created_at"`
	ExpiredAt   string `json:"expired_at"`
	ActivatedAt string `json:"activated_at"`
}

// Vouchers hold vouchers slice
type Vouchers struct {
	Vouchers []Voucher `json:"vouchers"`
}

// Partner hold data for partner table
type Partner struct {
	PartnerID     string `json:"partner_id"`
	PartnerName   string `json:"partner_name"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	OtherContacts string `json:"other_contacts"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

// Partners hold partner slice
type Partners struct {
	Partners []Partner `jason:"partners"`
}

// Invoice hold data for invoice table
type Invoice struct {
	InvoiceID       string `json:"invoice_id"`
	PartnerID       string `json:"partner_id"`
	PaymentMethodID string `json:"payment_method_id"`
	CreatedAt       string `json:"created_at"`
}

// Invoices hold invoice struct
type Invoices struct {
	Invoices []Invoice `json:"invoices"`
}

// InvoiceItem hold data for invoice_item table
type InvoiceItem struct {
	InvoiceItemID string `json:"invoice_item_id"`
	ItemID        string `json:"item_id"`
	ItemAmount    int    `json:"item_amount"`
	ItemPrice     string `json:"item_price"`
	ItemDiscount  string `json:"item_discount"`
}

// InvoiceItems Generate Invoices per Item
type InvoiceItems struct {
	InvoiceItems []InvoiceItem `json:"invoice_items"`
}

// PartnerVoucher hold data for partner_voucher table
type PartnerVoucher struct {
	InvoiceID     string `json:"invoice_id"`
	PartnerID     string `json:"partner_id"`
	VoucherCode   string `json:"voucher_code"`
	PurchaseValue string `json:"purchase_value"`
}

// PartnerVouchers generate vouchers for partner purchase
type PartnerVouchers struct {
	PartnerVouchers []PartnerVoucher `json:"partner_vouchers"`
}

func addVoucherType(c echo.Context) error {

	u := new(VoucherType)
	if err := c.Bind(u); err != nil {
		return err
	}

	sqlStatement := `INSERT INTO voucher_type(voucher_name, nominal, duration_month)
		VALUES($1, $2 , $3)`

	res, err := db.Query(sqlStatement, u.VoucherName, u.Nominal, u.DurationMonth)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")

}

func getVoucherType(c echo.Context) error {

	voucher := VoucherType{}
	id := c.Param("voucher_type_id")

	sqlStatement := `SELECT FROM voucher_type WHERE voucher_type_id=$1`

	res, err := db.Query(sqlStatement, voucher)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, voucher)
	}
	return c.String(http.StatusOK, id+"Selected")
}
func addPartner(c echo.Context) error {

	u := new(Partner)
	if err := c.Bind(u); err != nil {
		return err
	}

	sqlStatement := `INSERT INTO partner(partner_name, address, phone, email, other_contacts)
	VALUES($1, $2, $3, $4, $5)`

	res, err := db.Query(sqlStatement, u.PartnerName, u.Address, u.Phone, u.Email, u.OtherContacts)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")
}

func addVoucher(c echo.Context) error {

	u := new(Voucher)
	if err := c.Bind(u); err != nil {
		return err
	}

	sqlStatement := `INSERT INTO voucher(type_id, voucher_code, activated_at, expired_at)
		VALUES($1, random_string(8), $2, now() + '1 month'::interval * 6)
		RETURNING voucher_code`

	res, err := db.Query(sqlStatement, u.TypeID, u.ActivatedAt, u.ExpiredAt)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")

}
