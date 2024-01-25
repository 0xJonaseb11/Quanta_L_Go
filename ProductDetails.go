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

// Init initializes the chaincode
func (c *ProductDetailsContract) Init(ctx contractapi.TransactionContextInterface) error {
	// Initialization logic goes here
	return nil
}

// AddProduct adds a new product
func (c *ProductDetailsContract) AddProduct(ctx contractapi.TransactionContextInterface, name string, description string, manufacturedDate uint64, batchNumber string) error {
	nextProductID, err := c.generateNextProductID(ctx)
	if err != nil {
		return err
	}

	product := Product{
		ID:              nextProductID,
		Name:            name,
		Description:     description,
		ManufactureDate: manufacturedDate,
		BatchNumber:     batchNumber,
	}

	err = ctx.GetStub().PutState(fmt.Sprintf("PRODUCT-%d", nextProductID), []byte(product))
	if err != nil {
		return fmt.Errorf("failed to put product on the ledger: %v", err)
	}

	return nil

}

// RetrieveProductDetails retrieves the details of a product
func (c *ProductDetailsContract) RetrieveProductDetails(ctx contractapi.TransactionContextInterface, productID uint64) (*Product, error) {
	productBytes, err := ctx.GetStub().GetState(fmt.Sprintf("PRODUCT-%d", productID))
	if err != nil {
		return nil, fmt.Errorf("failed to read product from the ledger: %v", err)
	}
	if productBytes == nil {
		return nil, fmt.Errorf("product with ID %d does not exist", productID)
	}

	product := new(Product)
	err = json.Unmarshal(productBytes, product)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal product JSON: %v", err)
	}

	return product, nil
}

// UpdateProductState updates the state of a product
func (c *ProductDetailsContract) UpdateProductState(ctx contractapi.TransactionContextInterface, productID uint64, currentState ProductState) error {
	product, err := c.RetrieveProductDetails(ctx, productID)
	if err != nil {
		return err
	}

	// check for valid state transitions
	if product.State == PRODUCT_REGISTERED && currentState != PRODUCT_TRANSIT {
		return fmt.Errorf("invalid state transition")
	}

	product.State = currentState
	productBytes, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to marshal product JSON: %v", err)
	}

}

