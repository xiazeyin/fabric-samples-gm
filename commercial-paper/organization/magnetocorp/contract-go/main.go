/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"

	"github.com/xiazeyin/fabric-contract-api-go-gm/contractapi"
	commercialpaper "github.com/xiazeyin/fabric-samples-gm/commercial-paper/organization/magnetocorp/contract-go/commercial-paper"
)

func main() {

	contract := new(commercialpaper.Contract)
	contract.TransactionContextHandler = new(commercialpaper.TransactionContext)
	contract.Name = "org.papernet.commercialpaper"
	contract.Info.Version = "0.0.1"

	chaincode, err := contractapi.NewChaincode(contract)

	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode. %s", err.Error()))
	}

	chaincode.Info.Title = "CommercialPaperChaincode"
	chaincode.Info.Version = "0.0.1"

	err = chaincode.Start()

	if err != nil {
		panic(fmt.Sprintf("Error starting chaincode. %s", err.Error()))
	}
}
