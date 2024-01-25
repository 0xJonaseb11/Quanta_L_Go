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

// ProductHistory represents the history of a product
type ProductHistory struct {
	Timestamp uint64        `json:"timestamp"`
	Action    string        `json:"action"`
	Location  string        `json:"location"`
	State     ProductState `json:"state"`
}

// ProductState represents the state of a product
type ProductState int

const (
	PRODUCT_REGISTERED ProductState = iota
	QUALITY_ASSURANCE
	PRODUCT_TRANSIT
	PRODUCT_IN_INVENTORY
	PRODUCT_SOLD
	PRODUCT_RECALLED
	CONSUMPTION
	PENDING
	VALIDATING
	PUBLISHING
)