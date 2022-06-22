package event

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var (
	addProductSsdEndpoint = os.Getenv("USERS_SERVICE") + "/ssds/products"
	randonSsdEndpoint     = os.Getenv("USERS_SERVICE") + "/ssds/random"

	randomProductEndpoint = os.Getenv("ORDERS_SERVICE") + "/random"
)

type ProductAddedToSSD struct{}

type ProductAddedToSSDRequest struct {
	SsdId    int    `json:"ssd_id"`
	Barcode  string `json:"barcode"`
	Quantity int    `json:"n_products"`
}

func (passd *ProductAddedToSSD) Handle() {
	request := ProductAddedToSSDRequest{
		Quantity: rand.Intn(5),
	}

	/*
		getting a random product
	*/
	res, err := http.Get(randomProductEndpoint)
	if err != nil {
		panic(err.Error())
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("random product error: %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal(body, &request); err != nil {
		panic(err.Error())
	}

	/*
		getting a random device
	*/
	res, err = http.Get(randonSsdEndpoint)
	if err != nil {
		panic(err.Error())
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("random ssd error: %v", res.Status)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal(body, &request); err != nil {
		panic(err.Error())
	}

	/*
		send request to users service
	*/
	jsonData, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	res, err = http.Post(
		addProductSsdEndpoint,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Println("event", passd.Type(), "failed")
		return
	}
}

func (passd *ProductAddedToSSD) Type() string {
	return "product-added-to-ssd"
}
