package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//Transaction contains input (senders's public key and signature) and output (recipient's public key and value)
//NewTransaction creates a new transaction, signs it with ESDSA, and assigns an ID
//HashTransaction hashes the transaction datato create a unique ID
//Serializae and Deserialize for db storage and retrieval

type Transaction struct {
	ID     []byte
	Input  []TxInput
	Output []TxOutput
}

type TxInput struct {
	Signature []byte
	PublicKey []byte
}

type TxOutput struct {
	Value     int
	PublicKey []byte
}

// NewTransaction creates a new transaction, signs it with ESDSA, and assigns an ID
func NewTransaction(privateKey ecdsa.PrivateKey, recipient []byte, amount int) *Transaction {
	txIn := TxInput{}
	txOut := TxOutput{Value: amount, PublicKey: recipient}

	tx := Transaction{
		Input:  []TxInput{txIn},
		Output: []TxOutput{txOut},
	}
	tx.ID = tx.hashTransaction()

	//sign the transaction with the sender's
	r, s, err := ecdsa.Sign(rand.Reader, &privateKey, tx.ID)
	//check for errors
	if err != nil {
		log.Panic(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)
	txIn.Signature = signature
	return &tx
}

// HashTransaction hashes the transaction datato create a unique ID
func (tx *Transaction) hashTransaction() []byte {
	var hash [32]byte
	hash = sha256.Sum256(bytes.Join([][]byte{
		tx.Input[0].PublicKey,
		tx.Output[0].PublicKey,
		[]byte(string(tx.Output[0].Value)),
	}, []byte{}))

	return hash[:]
}

// Serializae and Deserialize for db storage and retrieval

func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return encoded.Bytes()
}

func DeserializeTransaction(data []byte) *Transaction {
	var transaction Transaction
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&transaction)
	if err != nil {
		log.Panic(err)
	}
	return &transaction
}
