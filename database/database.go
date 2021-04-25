package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

/**
  "id": "ambp16",
  "name": "Apple",
  "description": "Apple MacBook Pro 16",
  "inventory": "13"
*/

type (
	product struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Inventory   string `json:"inventory"`
	}
)

var products []product

func init() {
	wd, err := os.Getwd()
	if err != nil {
		er(err)
	}

	file, err := os.Open(fmt.Sprintf("%v/database/data.json", wd))
	if err != nil {
		er(err)
	}

	byteFile, err := ioutil.ReadAll(file)
	if err != nil {
		er(err)
	}

	err = json.Unmarshal(byteFile, &products)
	if err != nil {
		er(err)
	}
}

func GetProduct(id string) (*product, error) {
	for _, v := range products {
		if id == v.Id {
			return &v, nil
		}
	}
	return nil, errors.New("No product found")
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
