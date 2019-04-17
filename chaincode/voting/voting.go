package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Vote struct {
	Participant   string `json:"participant"`
	Time  string `json:"club"`
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
	} else if function == "getScore" {
		return s.getScore(APIstub)
	} else if function == "castVote" {
		return s.castVote(APIstub,args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	voteCount := make(map[string]int)
	votes := []Vote{}

	voteCount["FC Barcelona"] = 5
	voteCount["Real Madrid"] = 5
	voteCount["Manchester United"] = 5
	voteCount["AC Milan"] = 5
	voteCount["Bayern Munich"] = 5

	voteCountAsBytes, _ := json.Marshal( voteCount )
	votesAsBytes, _ := json.Marshal( votes )

	APIstub.PutState( "VoteCount" , voteCountAsBytes )
	APIstub.PutState( "Votes" , votesAsBytes )

	return shim.Success(nil)
}

func (s *SmartContract) getScore(APIstub shim.ChaincodeStubInterface) sc.Response{

	voteCountAsBytes, _ := APIstub.GetState("VoteCount")

	return shim.Success(voteCountAsBytes)
}

func (s *SmartContract) castVote(APIstub shim.ChaincodeStubInterface , args []string) sc.Response{

	// var vote = Vote{ Participant : args[1] , Time : args[2] }

	// votesAsBytes, _ := APIstub.GetState("Votes")
	voteCountAsBytes, _ := APIstub.GetState("VoteCount")

	// var Votes []Vote
	// err := json.Unmarshal( votesAsBytes , &Votes )

	// if  err != nil {
	// 	fmt.Println(err)
	// }

	// Votes = append( Votes , vote )

	// updatesVotesAsBytes, err1 := json.Marshal( Votes )
	// APIstub.PutState( "Votes" , updatesVotesAsBytes )

	// if err1 != nil {
	// 	fmt.Println(err1)
	// }

	VoteCount := make(map[string]int)
	err := json.Unmarshal( voteCountAsBytes , &VoteCount )

	if err != nil {
		fmt.Println(err)
	}

	count := VoteCount[args[0]]
	VoteCount[args[0]] = count + 1

	updatedVoteCountAsBytes, err2 := json.Marshal( VoteCount )
	APIstub.PutState( "VoteCount" , updatedVoteCountAsBytes )

	if err2 != nil {
		fmt.Println(err)
	}

	err = APIstub.SetEvent("voteCasted", []byte(args[0]))

	return shim.Success( nil )
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}