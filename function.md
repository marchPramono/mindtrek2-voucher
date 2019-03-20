


func addInvoiceItem(c echo.Context) error {

	u := new(InvoiceItem)
	if err := c.Bind(u); err != nil {
		return err
	}

	itemsID := u.ItemID
	itemsAmount := u.ItemAmount

	for i := 0; i < u.ItemAmount; i++ {

		sqlStatement := `INSERT INTO mind_voucher(product_id, voucher_code, nominal, duration_month, expired_at)
		VALUES($1, random_string(8), $2 , $3, now() + '1 month'::interval * 120)
		RETURNING voucher_code`

		voucherCode := ""
		res, err := db.Query(sqlStatement, u.ProductID, u.Nominal, u.DurationMonth).Scan(&voucherCode)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("voucher inserted", i, voucherCode)

		sqlStatement = `INSERT INTO mind_partner_voucher(invoice_id, partner_id, voucher_code, purchase_value)
		VALUES($1, $2, $3, $4)`

		res, err := db.Query(sqlStatement, u.ProductID, u.Nominal, u.DurationMonth).Scan(&voucherCode)
		if err != nil {
			fmt.Println(err)
		}

	}

	return c.String(http.StatusOK, "ok")

	/*
		for item, amount := range itemsID {
			for i := 0; i < u.ItemAmount; i++ {
				return itemsAmount
			}
			return itemsID
		}

		sqlStatement := `INSERT INTO mind_invoice_item(item_id, item_amount, item_price, item_discount)
		VALUES($1, $2, $3, #4)`

		res, err := db.Query(sqlStatement, u.ItemID, u.ItemAmount, u.ItemPrice, u.ItemDiscount)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, "ok")
	*/
}

func getAllVouchers(c echo.Context) error {
	sqlStatement := `SELECT invoice_id, voucher_code, partner_id, purchase_value 
	FROM mind_partner_voucher ORDER BY invoice_id`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := PartnerVouchers{}

	for rows.Next() {

		voucher := PartnerVoucher{}

		err2 := rows.Scan(&voucher.InvoiceID, &voucher.PartnerID, &voucher.VoucherCode, &voucher.PurchaseValue)
		if err2 != nil {
			return err2
		}
		result.PartnerVouchers = append(result.PartnerVouchers, voucher)
	}
	return c.JSON(http.StatusAccepted, result)
}
