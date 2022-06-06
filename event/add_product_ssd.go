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

var addProductSsdEndpoint = os.Getenv("SERVICE_HOST") + "/api/users/ssds/products"

type ProductAddedToSSD struct {
}

type ProductAddedToSSDRequest struct {
	SsdId    int    `json:"ssd_id"`
	Barcode  string `json:"barcode"`
	Quantity int    `json:"n_products"`
}

func (passd *ProductAddedToSSD) Handle() {
	request := ProductAddedToSSDRequest{
		Quantity: rand.Intn(5),
	}

	/**
	get random product
	*/
	res, err := http.Get(os.Getenv("SERVICE_HOST") + "/api/products/random")
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

	/**
	get random ssd
	*/
	res, err = http.Get(os.Getenv("SERVICE_HOST") + "/api/users/ssds/random")
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

	/**
	send to service
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
		log.Println("event failed")
		return
	}

	var resBody map[string]any
	json.NewDecoder(res.Body).Decode(&resBody)
	log.Println(resBody)
}

func (passd *ProductAddedToSSD) Type() string {
	return "product-added-to-ssd"
}
