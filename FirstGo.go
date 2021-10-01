package main

import (


	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

)
type Target struct {
	Address string `json:"SOURCE_ADDRESS"`
	EndPoint    string     `json:"ETH_NODE_ADDRESS"`
	key string;
}

func MakeTransaction(){
	MyTarget:=Target{}
	MyLink:=MyTarget.Address

	client, err := ethclient.Dial("MyLink") //connect to a client

	if err != nil {
		log.Fatal(err)
	}


	privateKey:= Target.key //getting private key

	publicKey := privateKey.Public() //public address of of the account we're sending from

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // установить тип нашей переменной publicKey и присвоить ее publicKeyECDSA
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	from:= crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), from) //getting nonce
	if err != nil {
		log.Fatal(err)
	}


	value := big.NewInt(1000000000000000000)
	gasLimit := uint64(21000)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	to:=MyTarget.EndPoint
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, nil) //generate unsigned transaction

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("Transaction was made successfully")

}

func main() {
	d := time.NewTicker(3 * time.Second)
	MyChannel := make(chan bool)

	go func() {
		time.Sleep(9 * time.Second)
		MyChannel <- true
	}()

	for {
		select {
		case <-MyChannel:
			fmt.Println("Completed!")
			return
		case tm := <-d.C:
			MakeTransaction()
			fmt.Println("The Current time is: ", tm)
		}
	}
}