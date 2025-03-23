package main

import (
	"backend/utils"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
	"strings"
)

// AccountManager handles Ethereum account operations.
type AccountManager struct {
	logger         *zap.Logger
	env            utils.Envs
	PrivateKeyPath string
	privateKey     *ecdsa.PrivateKey
	publicKey      common.Address
}

func NewAccountManagerSecure(privateKeyPath string, logger *zap.Logger) (*AccountManager, error) {
	logger.Info("NewAccountManagerSecure")

	var privateKey *ecdsa.PrivateKey

	// Check if the private key file exists
	if privateKeyBytes, err := os.ReadFile(privateKeyPath); err == nil {
		// Parse the private key from the file
		privateKey, err = crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse private key")
		}
	} else {
		// Generate a new private key if the file does not exist
		privateKey, err = crypto.GenerateKey()
		if err != nil {
			return nil, errors.Wrapf(err, "failed to generate private key")
		}

		newPrivateKeyBytes := crypto.FromECDSA(privateKey)
		if err := os.WriteFile(privateKeyPath, newPrivateKeyBytes, 0600); err != nil {
			return nil, errors.Wrapf(err, "failed to write private key to file")
		}
	}

	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)
	return &AccountManager{
		PrivateKeyPath: privateKeyPath,
		privateKey:     privateKey,
		publicKey:      publicKey,
	}, nil
}

func NewAccountManager(privateKeyPath string, env utils.Envs, logger *zap.Logger) (*AccountManager, error) {
	logger.Info("NewAccountManager")

	key := strings.TrimPrefix(env.PRIVATE_KEY, "0x")
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	return &AccountManager{
		PrivateKeyPath: privateKeyPath,
		privateKey:     privateKey,
		publicKey:      publicKey,
	}, nil
}

func (am *AccountManager) GetPublicKey() string {
	return am.publicKey.Hex()
}

func (am *AccountManager) WritePrivateKeyToFile() error {
	privateKeyBytes := crypto.FromECDSA(am.privateKey)
	return os.WriteFile(am.PrivateKeyPath, privateKeyBytes, 0600)
}
