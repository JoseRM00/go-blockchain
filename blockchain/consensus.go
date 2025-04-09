package blockchain

import (
	"crypto/rand"
	"log"
	"math/big"
	//"math/rand"
	//"time"
)

//realDavePartner@gmail.com
/*
1. PoSvalidator: represents a validator in the PoS system with a public key
2. ProofOfStake(): selects a validator based on their stake.
The more stake a validator has the higher the probability of being chosen to validate a block
*/

// PoSvalidator: represents a validator in the PoS system with a public key
type PoSValidator struct {
	PublicKey []byte
	Stake     int
}

/*
ProofOfStake(): selects a validator based on their stake.
The more stake a validator has the higher the probability of being chosen to validate a block
*/
func ProofOfStake(validators map[string]*PoSValidator) string {
	totalStake := 0
	for _, validator := range validators {
		totalStake += validator.Stake
	}

	//select a validator randomly based on their stake
	randomBig, err := rand.Int(rand.Reader, big.NewInt(int64(totalStake)))
	if err != nil {
		log.Panic(err)
	}
	random := randomBig.Int64()

	//rand.Seed(time.Now().UnixNano())
	//random := rand.Intn(totalstake)

	//select a validator based on their stake
	for _, validator := range validators {
		random -= int64(validator.Stake)
		if random <= 0 {
			return string(validator.PublicKey)
		}
	}

	log.Panic("Unable to find a validator")
	return ""
}
