package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Player struct {
	ObjectType string `json:"docType"`
	Name   string `json:"name"`
	Club  string `json:"club"`
	Country string `json:"country"`
	KitNo  uint `json:"kit_no"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "queryAllPlayers" {
		return s.queryAllPlayers(APIstub , args)
	} else if function == "queryPlayerCount" {
		return s.queryPlayerCount(APIstub)
	} else if function == "queryPlayers" {
		return s.queryPlayers(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	players := []Player{
		Player{ObjectType: "Player", Name: "Marc Andre Ter Stegen", Club: "FC Barcelona", Country: "Germany", KitNo: 1},
		Player{ObjectType: "Player", Name: "Nelson Semedo", Club: "FC Barcelona", Country: "Portugal", KitNo: 2},
		Player{ObjectType: "Player", Name: "Gerard Pique", Club: "FC Barcelona", Country: "Spain", KitNo: 3},
		Player{ObjectType: "Player", Name: "Ivan Rakitic", Club: "FC Barcelona", Country: "Croatia", KitNo: 4},
		Player{ObjectType: "Player", Name: "Sergio Busquets", Club: "FC Barcelona", Country: "Spain", KitNo: 5},
		Player{ObjectType: "Player", Name: "Denis Suarez", Club: "FC Barcelona", Country: "Spain", KitNo: 6},
		Player{ObjectType: "Player", Name: "Philip Coutinho", Club: "FC Barcelona", Country: "Brazil", KitNo: 7},
		Player{ObjectType: "Player", Name: "Arthur Melo", Club: "FC Barcelona", Country: "Brazil", KitNo: 8},
		Player{ObjectType: "Player", Name: "Luis Suarez", Club: "FC Barcelona", Country: "Uruguay", KitNo: 9},
		Player{ObjectType: "Player", Name: "Lionel Messi", Club: "FC Barcelona", Country: "Argentina", KitNo: 10},
		Player{ObjectType: "Player", Name: "Osmane Dembele", Club: "FC Barcelona", Country: "France", KitNo: 11},
		Player{ObjectType: "Player", Name: "Rafinha Alacantara", Club: "FC Barcelona", Country: "Brazil", KitNo: 12},
		Player{ObjectType: "Player", Name: "Jasper Cillessen", Club: "FC Barcelona", Country: "Netherlands", KitNo: 13},
		Player{ObjectType: "Player", Name: "Malcom", Club: "FC Barcelona", Country: "Brazil", KitNo: 14},
		Player{ObjectType: "Player", Name: "Clement Lenglet", Club: "FC Barcelona", Country: "France", KitNo: 15},
		Player{ObjectType: "Player", Name: "Paco Alcacer", Club: "FC Barcelona", Country: "Spain", KitNo: 17},
		Player{ObjectType: "Player", Name: "Jordi Alba", Club: "FC Barcelona", Country: "Spain", KitNo: 18},
		Player{ObjectType: "Player", Name: "Munir Al Haddadi", Club: "FC Barcelona", Country: "Spain", KitNo: 19},
		Player{ObjectType: "Player", Name: "Sergio Roberto", Club: "FC Barcelona", Country: "Spain", KitNo: 20},
		Player{ObjectType: "Player", Name: "Arturo Vidal", Club: "FC Barcelona", Country: "Chile", KitNo: 22},
		Player{ObjectType: "Player", Name: "Samuel Umtiti", Club: "FC Barcelona", Country: "France", KitNo: 23},
		Player{ObjectType: "Player", Name: "Thomas Vermaelen", Club: "FC Barcelona", Country: "Belgium", KitNo: 25},
		Player{ObjectType: "Player", Name: "Claudio Bravo", Club: "Manchester City", Country: "Chile", KitNo: 1},
		Player{ObjectType: "Player", Name: "Kyle Walker", Club: "Manchester City", Country: "England", KitNo: 2},
		Player{ObjectType: "Player", Name: "Danilo", Club: "Manchester City", Country: "Brazil", KitNo: 3},
		Player{ObjectType: "Player", Name: "Vincent Kompany", Club: "Manchester City", Country: "Belgium", KitNo: 4},
		Player{ObjectType: "Player", Name: "John Stones", Club: "Manchester City", Country: "England", KitNo: 5},
		Player{ObjectType: "Player", Name: "Raheem Sterling", Club: "Manchester City", Country: "England", KitNo: 7},
		Player{ObjectType: "Player", Name: "Ikkay Gundogon", Club: "Manchester City", Country: "Germany", KitNo: 8},
		Player{ObjectType: "Player", Name: "Sergio Aguero", Club: "Manchester City", Country: "Argentina", KitNo: 10},
		Player{ObjectType: "Player", Name: "Aymeric Laporte", Club: "Manchester City", Country: "France", KitNo: 14},
		Player{ObjectType: "Player", Name: "Eliaquim Mangala", Club: "Manchester City", Country: "France", KitNo: 15},
		Player{ObjectType: "Player", Name: "Kevin De Bruyne", Club: "Manchester City", Country: "Belgium", KitNo: 17},
		Player{ObjectType: "Player", Name: "Fabian Delph", Club: "Manchester City", Country: "England", KitNo: 18},
		Player{ObjectType: "Player", Name: "Leroy Sane", Club: "Manchester City", Country: "Germany", KitNo: 19},
		Player{ObjectType: "Player", Name: "Bernardo Silva", Club: "Manchester City", Country: "Portugal", KitNo: 20},
		Player{ObjectType: "Player", Name: "Davil Silva", Club: "Manchester City", Country: "Spain", KitNo: 21},
		Player{ObjectType: "Player", Name: "Benjamin Mendy", Club: "Manchester City", Country: "France", KitNo: 22},
		Player{ObjectType: "Player", Name: "Ferdandinho", Club: "Manchester City", Country: "Brazil", KitNo: 25},
		Player{ObjectType: "Player", Name: "Riyad Mahrez", Club: "Manchester City", Country: "Algeria", KitNo: 26},
		Player{ObjectType: "Player", Name: "Nicolas Otamendi", Club: "Manchester City", Country: "Argentina", KitNo: 30},
		Player{ObjectType: "Player", Name: "Ederson", Club: "Manchester City", Country: "Brazil", KitNo: 31},
		Player{ObjectType: "Player", Name: "Gabriel Jesus", Club: "Manchester City", Country: "Brazil", KitNo: 33},
	}

	APIstub.PutState("PLAYER_COUNT", []byte(strconv.Itoa(len(players))))

	i := 0
	for i < len(players) {
		fmt.Println("i is ", i)
		playerAsBytes, _ := json.Marshal(players[i])
		APIstub.PutState("PLAYER"+strconv.Itoa(i), playerAsBytes)
		fmt.Println("Added", players[i])
		i = i + 1
	}

	return shim.Success(nil)
}

// func (s *SmartContract) queryAllPlayers(APIstub shim.ChaincodeStubInterface , args []string) sc.Response {

// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 2")
// 	}

// 	offset, err1 := strconv.Atoi( args[0] )
// 	limit, err2 := strconv.Atoi( args[1] )

// 	if err1 != nil || err2 != nil {
// 		return shim.Error("Invalid arguments")
// 	}

// 	var keys []string

// 	i := 0
// 	for i < limit {
// 		keys[i] = "PLAYER"+strconv.Itoa( offset + i )
// 		i = i + 1
// 	}

// 	// startKey := "PLAYER"+strconv.Itoa(offset)
// 	// endKey := "PLAYER"+strconv.Itoa(offset+limit)

// 	// resultsIterator, err := APIstub.GetStateByPartialCompositeKey("player", keys)
// 	// if err != nil {
// 	// 	return shim.Error(err.Error())
// 	// }
// 	// defer resultsIterator.Close()

// 	// buffer is a JSON array containing QueryResults
// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")

// 	bArrayMemberAlreadyWritten := false
// 	i = 0
// 	for i < limit {

// 		playerAsBytes, err := APIstub.GetState(keys[i])
// 		// var player Player
// 		// err1 := json.Unmarshal( playerAsBytes , &player )

// 		if err != nil {
// 			return shim.Error(err.Error())
// 		}
// 		// Add a comma before array members, suppress it for the first array member
// 		if bArrayMemberAlreadyWritten == true {
// 			buffer.WriteString(",")
// 		}
// 		buffer.WriteString("{\"Key\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(keys[i])
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"Record\":")
// 		// Record is a JSON object, so we write as-is
// 		buffer.Write(playerAsBytes)
// 		buffer.WriteString("}")
// 		bArrayMemberAlreadyWritten = true

// 		i = i + 1
// 	}
// 	buffer.WriteString("]")

// 	fmt.Printf("- queryAllPlayers:\n%s\n", buffer.String())

// 	return shim.Success(buffer.Bytes())
// }


func (s *SmartContract) queryAllPlayers(APIstub shim.ChaincodeStubInterface , args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	offset, err1 := strconv.Atoi( args[0] )
	limit, err2 := strconv.Atoi( args[1] )

	if err1 != nil || err2 != nil {
		return shim.Error("Invalid arguments")
	}

	var players []Player

	var i = 0
	for i < limit {
		
		index := offset + i
		var key string = "PLAYER"+strconv.Itoa(index)

		playerAsBytes, err := APIstub.GetState(key)

		var player Player

		err = json.Unmarshal(playerAsBytes , &player  )
		players = append(players , player)

		if err != nil {
			return shim.Error("Error at KEY - "+key)
		}

		i = i+1

	}

	bytestowrite , err3 := json.Marshal(players)

	if  err3 != nil {
		fmt.Println(err3)
	}

	return shim.Success(bytestowrite)
}

func (s *SmartContract) queryPlayerCount(APIstub shim.ChaincodeStubInterface ) sc.Response {
	
	PlayerCountBytes, err := APIstub.GetState("PLAYER_COUNT")
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for Player Count\"}"
		return shim.Error(jsonResp)
	}
	
	return shim.Success(PlayerCountBytes)
}

func (t *SmartContract) queryPlayers(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}