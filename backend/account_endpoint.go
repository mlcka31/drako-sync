package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"os"
)

// AccountManager handles Ethereum account operations.
type AccountManager struct {
	privateKey *ecdsa.PrivateKey
	publicKey  common.Address
}

// NewAccountManager creates a new Ethereum account and saves the private key to file.
func NewAccountManager(privateKeyPath string) (*AccountManager, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	err = ioutil.WriteFile(privateKeyPath, privateKeyBytes, 0600)
	if err != nil {
		return nil, err
	}

	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)
	return &AccountManager{privateKey, publicKey}, nil
}

// GetPublicKey returns the Ethereum public address.
func (am *AccountManager) GetPublicKey() string {
	return am.publicKey.Hex()
}

func (am *AccountManager) WritePrivateKeyToFile(path string) error {
	privateKeyBytes := crypto.FromECDSA(am.privateKey)
	return os.WriteFile(path, privateKeyBytes, 0600)
}
