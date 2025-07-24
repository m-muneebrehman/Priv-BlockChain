package wallet

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "./tmp/wallets_%s.data"

type Wallets struct {
	Wallets map[string]*Wallet
}

// SerializableWallets represents wallets that can be encoded/decoded with gob
type SerializableWallets struct {
	Wallets map[string]*SerializableWallet
}

// ToSerializable converts Wallets to SerializableWallets
func (ws *Wallets) ToSerializable() *SerializableWallets {
	serializableWallets := &SerializableWallets{
		Wallets: make(map[string]*SerializableWallet),
	}

	for address, wallet := range ws.Wallets {
		serializableWallets.Wallets[address] = wallet.ToSerializable()
	}

	return serializableWallets
}

// FromSerializable converts SerializableWallets back to Wallets
func (sws *SerializableWallets) ToWallets() *Wallets {
	wallets := &Wallets{
		Wallets: make(map[string]*Wallet),
	}

	for address, serializableWallet := range sws.Wallets {
		wallets.Wallets[address] = serializableWallet.ToWallet()
	}

	return wallets
}

func CreateWallets(nodeId string) (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFile(nodeId)

	return &wallets, err
}

func (ws *Wallets) AddWallet() string {
	wallet := MakeWallet()
	address := fmt.Sprintf("%s", wallet.Address())

	ws.Wallets[address] = wallet

	return address
}

func (ws *Wallets) GetAllAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) LoadFile(nodeId string) error {
	walletFile := fmt.Sprintf(walletFile, nodeId)
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		return err
	}

	var serializableWallets SerializableWallets
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&serializableWallets)
	if err != nil {
		return err
	}

	// Convert back to regular wallets
	loadedWallets := serializableWallets.ToWallets()
	ws.Wallets = loadedWallets.Wallets

	return nil
}

func (ws *Wallets) SaveFile(nodeId string) {
	var content bytes.Buffer
	walletFile := fmt.Sprintf(walletFile, nodeId)

	// Convert to serializable format
	serializableWallets := ws.ToSerializable()

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(serializableWallets)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
