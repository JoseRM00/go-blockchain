package api

import (
	"net/http"

	"github.com/JoseRM00/go-blockchain/blockchain"
	"github.com/labstack/echo/v4"
)

/*
1. StartServer initializes de Echo Framework server and sets up routs for handling transactions,
retrieving blocks and deploying/executing smart contracts
2. HandleTransaction accepts transactions via POST and processes them
3. HandleGetBlock fetchs a block by its hash
4. HandleDeployContract deploys a smart contract to the blockchain
5. HandleExecuteContract executes a smart contract on the blockchain
*/

// Blockchain and storage references (global for simplicity)
var bc *blockchain.Blockchain
var db *blockchain.BlockchainDB

// StartServer initializes de Echo Framework server
func StartServer(blockchain *blockchain.Blockchain, database *blockchain.BlockchainDB) {
	bc = blockchain
	db = database

	e := echo.New()

	//Define API routes
	e.POST("/transaction", handleTransaction)
	e.GET("/block/:hash", handleGetBlock)
	e.POST("/contract", handleDeployContract)
	e.POST("/contract/execute", handleExecuteContract)

	e.Logger.Fatal(e.Start(":1323")) //8080
}

// HandleTransaction accepts transactions via POST and processes them
func handleTransaction(c echo.Context) error {
	from := c.QueryParam("from")
	to := c.QueryParam("to")
	amount := c.QueryParam("amount")

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Transaction submitted",
		"from":    from,
		"to":      to,
		"amount":  amount,
	})
}

// HandleGetBlock fetchs a block by its hash
func handleGetBlock(c echo.Context) error {
	hash := c.Param("hash")
	block, err := bc.GetBlock(hash)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Block not found",
		})
	}
	return c.JSON(http.StatusOK, block)
}

// HandleDeployContract deploys a smart contract to the blockchain
func handleDeployContract(c echo.Context) error {
	id := c.QueryParam("id")
	code := c.QueryParam("code")

	contract := contracts.NewSmartContract(id, code)
	err := contract.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Contract deployed",
		"id":      id,
	})
}

// HandleExecuteContract executes a smart contract on the blockchain
func handleExecuteContract(c echo.Context) error {
	id := c.QueryParam("id")
	input := map[string]interface{}{}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Contract executed",
		"id":      id,
		"input":   input,
	})
}
