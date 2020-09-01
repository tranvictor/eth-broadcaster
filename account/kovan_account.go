package account

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tranvictor/ethutils/account/ledgereum"
	"github.com/tranvictor/ethutils/account/trezoreum"
	"github.com/tranvictor/ethutils/broadcaster"
	"github.com/tranvictor/ethutils/reader"
)

func NewKovanAccountFromKeystore(file string, password string) (*Account, error) {
	_, key, err := PrivateKeyFromKeystore(file, password)
	if err != nil {
		return nil, err
	}
	return &Account{
		NewKeySigner(key),
		reader.NewKovanReader(),
		broadcaster.NewKovanBroadcaster(),
		crypto.PubkeyToAddress(key.PublicKey),
	}, nil
}

func NewKovanAccountFromPrivateKey(hex string) (*Account, error) {
	_, key, err := PrivateKeyFromHex(hex)
	if err != nil {
		return nil, err
	}
	return &Account{
		NewKeySigner(key),
		reader.NewKovanReader(),
		broadcaster.NewKovanBroadcaster(),
		crypto.PubkeyToAddress(key.PublicKey),
	}, nil
}

func NewKovanAccountFromPrivateKeyFile(file string) (*Account, error) {
	_, key, err := PrivateKeyFromFile(file)
	if err != nil {
		return nil, err
	}
	return &Account{
		NewKeySigner(key),
		reader.NewKovanReader(),
		broadcaster.NewKovanBroadcaster(),
		crypto.PubkeyToAddress(key.PublicKey),
	}, nil
}

func NewKovanTrezorAccount(path string, address string) (*Account, error) {
	signer, err := trezoreum.NewTrezorSigner(path, address)
	if err != nil {
		return nil, err
	}
	return &Account{
		signer,
		reader.NewKovanReader(),
		broadcaster.NewKovanBroadcaster(),
		common.HexToAddress(address),
	}, nil
}

func NewKovanLedgerAccount(path string, address string) (*Account, error) {
	signer, err := ledgereum.NewLedgerSigner(path, address)
	if err != nil {
		return nil, err
	}
	return &Account{
		signer,
		reader.NewKovanReader(),
		broadcaster.NewKovanBroadcaster(),
		common.HexToAddress(address),
	}, nil
}
