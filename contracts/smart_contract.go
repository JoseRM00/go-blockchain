package contracts

import (
	"errors"
	"fmt"
	"time"
)

//SmartContract represents a smart contract with ID, code, state and timestamp
//NewSmartContract initializes a new smart contract
//Execute simulates executing the contract by logging the input and updating the state
//Validate checks the contract

// SmartContract represents a smart contract with ID, code, state and timestamp
type SmartContract struct {
	ID        string
	Code      string
	State     map[string]interface{}
	CreatedAt time.Time
}

// Execute simulates executing the contract by logging the input and updating the state
func (sc *SmartContract) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println("Executing contract with input:", input)
	sc.State["lastExecution"] = input
	return sc.State, nil
}

// Validate checks the contract
func (sc *SmartContract) Validate() error {
	if sc.Code == "" {
		return errors.New("Contract code is empty")
	}
	return nil
}
