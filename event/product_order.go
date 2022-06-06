package event

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type ProductOrder struct {
}

func (pr *ProductOrder) Handle() {
	res, err := http.Get(os.Getenv("SERVICE_HOST") + "/api/products/randomProductFromUserSSD")
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Println("event failed")
	}

	var jsonRes map[string]any
	json.NewDecoder(res.Body).Decode(&jsonRes)
	log.Println(jsonRes)

	/**
	send the order to the API
	*/
}

func (pr *ProductOrder) Type() string {
	return "product-order"
}
