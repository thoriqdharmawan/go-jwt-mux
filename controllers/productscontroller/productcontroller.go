package productscontroller

import (
	"net/http"

	"github.com/thoriqdharmawan/go-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	// dumy
	data := []map[string]any{
		{
			"id":           1,
			"product_name": "Kemeja",
			"stock":        100,
		},
		{
			"id":           2,
			"product_name": "Topi",
			"stock":        1100,
		},
		{
			"id":           3,
			"product_name": "Sepatu",
			"stock":        80,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}
