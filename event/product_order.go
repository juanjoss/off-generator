package event

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var (
	randomProductFromUserSsdEndpoint = os.Getenv("ORDERS_SERVICE") + "/randomProductFromUserSSD"

	ordersEndpoint = os.Getenv("ORDERS_SERVICE") + "/orders"
)

type ProductOrder struct{}

type ProductOrderRequest struct {
	SsdId    int    `json:"ssd_id"`
	Barcode  string `json:"barcode"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"`
}

func (pr *ProductOrder) Handle() {
	/*
		get a device (id) and a associated product (barcode)
	*/
	res, err := http.Get(randomProductFromUserSsdEndpoint)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Println("error getting: ", randomProductFromUserSsdEndpoint)
	}

	var request ProductOrderRequest
	err = json.NewDecoder(res.Body).Decode(&request)
	if err != nil {
		panic(err.Error())
	}
	request.Quantity = rand.Intn(5)
	request.Status = ""

	/*
		send request to orders service
	*/
	jsonData, err := json.Marshal(request)
	if err != nil {
		panic(err.Error())
	}

	res, err = http.Post(
		ordersEndpoint,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		panic(err.Error())
	}

	if res.StatusCode != http.StatusOK {
		log.Println("event", pr.Type(), "failed")
		return
	}
}

func (pr *ProductOrder) Type() string {
	return "product-order"
}
