/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/xiazeyin/fabric-contract-api-go-gm/contractapi"
	auction "github.com/xiazeyin/fabric-samples-gm/auction/chaincode-go/smart-contract"
)

func main() {
	auctionSmartContract, err := contractapi.NewChaincode(&auction.SmartContract{})
	if err != nil {
		log.Panicf("Error creating auction chaincode: %v", err)
	}

	if err := auctionSmartContract.Start(); err != nil {
		log.Panicf("Error starting auction chaincode: %v", err)
	}
}
