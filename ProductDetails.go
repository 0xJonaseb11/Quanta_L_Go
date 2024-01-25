package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ProductDetailsContract struct {
	contractapi.Contract
}

// Product represents the product details
type Product struct {
	ID              uint64 `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ManufactureDate uint64 `json:"manufactureDate"`
	BatchNumber     string `json:"batchNumber"`
}