/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	//	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// Case describes basic details of what makes up a Case
type Case struct {
	ID            string `json:"id"`
	Rev           string `json:"rev"`
	CaseId        string `json:"caseId"`
	EvidenceId    string `json:"evidenceId"`
	Removalowner  string `json:"removalowner"`
	Removalreason string `json:"removalreason"`
	StateId       string `json:"stateId"`
	Timestamp     string `json:"timestamp"`
	Version       string `json:"version"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	ID     string `json:"Key"`
	Record *Case
}

// InitLedger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	cases := []Case{
		Case{ID: "evidence1", Rev: "1-8a4a6d4ba723a6675c5c906c48aa532a", CaseId: "case1", EvidenceId: "evidence1", Removalowner: "Tomoko", Removalreason: "DISPOSED", StateId: "REMOVED", Timestamp: "1674526440", Version: "CgMBUWA="},
		Case{ID: "evidence2", Rev: "2-8a4a6d4ba723a6675c5c906c48aa532a", CaseId: "case2", EvidenceId: "evidence2", Removalowner: "Tudo", Removalreason: "DISPOSED", StateId: "REMOVED", Timestamp: "2674526440", Version: "CgMBUWA="},
		Case{ID: "evidence3", Rev: "3-8a4a6d4ba723a6675c5c906c48aa532a", CaseId: "case3", EvidenceId: "evidence3", Removalowner: "Signa", Removalreason: "DISPOSED", StateId: "REMOVED", Timestamp: "3674526440", Version: "CgMBUWA="},
		Case{ID: "evidence4", Rev: "4-8a4a6d4ba723a6675c5c906c48aa532a", CaseId: "case4", EvidenceId: "evidence4", Removalowner: "Bula", Removalreason: "DISPOSED", StateId: "REMOVED", Timestamp: "4674526440", Version: "CgMBUWA="},
		Case{ID: "evidence5", Rev: "5-8a4a6d4ba723a6675c5c906c48aa532a", CaseId: "case5", EvidenceId: "evidence5", Removalowner: "Quafg", Removalreason: "DISPOSED", StateId: "REMOVED", Timestamp: "5674526440", Version: "CgMBUWA="},
		Case{ID: "evidence6", Rev: "6-8a4a6d4ba723a6675c5c906c48aa532a", CaseId: "case6", EvidenceId: "evidence6", Removalowner: "Zagt", Removalreason: "DISPOSED", StateId: "REMOVED", Timestamp: "6674526440", Version: "CgMBUWA="},
	}

	for _, Case := range cases {
		caseAsBytes, _ := json.Marshal(Case)
		err := ctx.GetStub().PutState(Case.ID, caseAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// ADD

func (s *SmartContract) CreateCase(ctx contractapi.TransactionContextInterface, id string, rev string, caseId string, evidenceId string, removalowner string, removalreason string, stateId string, timestamp string, version string) error {
	Case := Case{
		ID:            id,
		Rev:           rev,
		CaseId:        caseId,
		EvidenceId:    evidenceId,
		Removalowner:  removalowner,
		Removalreason: removalreason,
		StateId:       stateId,
		Timestamp:     timestamp,
		Version:       version,
	}

	caseAsBytes, _ := json.Marshal(Case)

	return ctx.GetStub().PutState(id, caseAsBytes)
}

// checkout
func (s *SmartContract) Checkout(ctx contractapi.TransactionContextInterface, id string, rev string, timestamp string) error {
	Case, err := s.QueryCase(ctx, id)

	if err != nil {
		return err
	}

	Case.ID = id
	Case.Rev = rev
	Case.Timestamp = timestamp
	Case.StateId = "CHECKEDOUT"

	caseAsBytes, _ := json.Marshal(Case)

	return ctx.GetStub().PutState(id, caseAsBytes)
}

// CheckIn
func (s *SmartContract) Checkin(ctx contractapi.TransactionContextInterface, id string, rev string, timestamp string) error {
	Case, err := s.QueryCase(ctx, id)

	if err != nil {
		return err
	}

	Case.ID = id
	Case.Rev = rev
	Case.Timestamp = timestamp
	Case.StateId = "CHECKEDIN"

	caseAsBytes, _ := json.Marshal(Case)

	return ctx.GetStub().PutState(id, caseAsBytes)
}

// Remove
func (s *SmartContract) Remove(ctx contractapi.TransactionContextInterface, id string, rev string, timestamp string) error {
	Case, err := s.QueryCase(ctx, id)

	if err != nil {
		return err
	}

	Case.ID = id
	Case.Rev = rev
	Case.Timestamp = timestamp
	Case.StateId = "DISPOSED"

	caseAsBytes, _ := json.Marshal(Case)

	return ctx.GetStub().PutState(id, caseAsBytes)
}

// queryCase
func (s *SmartContract) QueryCase(ctx contractapi.TransactionContextInterface, id string) (*Case, error) {
	caseAsBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if caseAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", id)
	}

	Case := new(Case)
	_ = json.Unmarshal(caseAsBytes, Case)

	return Case, nil
}

// queryallCase
func (s *SmartContract) QueryAllCases(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		Case := new(Case)
		_ = json.Unmarshal(queryResponse.Value, Case)

		queryResult := QueryResult{ID: queryResponse.Key, Record: Case}
		results = append(results, queryResult)
	}

	return results, nil
}

// DeleteCase
func (s *SmartContract) DeleteCase(ctx contractapi.TransactionContextInterface, id string) error {
	caseAsBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if caseAsBytes == nil {
		return fmt.Errorf("%s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create case chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting case chaincode: %s", err.Error())
	}
}
