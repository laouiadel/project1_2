// Look the official documentation documentation, fabric-sample/chaincode, and the fabric source code for deeper informations.
package main

import (
		//"encoding/base64"
		"encoding/hex"
		"encoding/json"
		"fmt"
		"github.com/hyperledger/fabric/core/chaincode/shim"
		pb "github.com/hyperledger/fabric/protos/peer"
		// "strconv"
		//"time"
		// /fabric/vendor/github.com/op/go-logging
		"time"
		"strconv"
		"github.com/cbergoon/merkletree"
		"log"
		"crypto/sha256"
			)

const this_chaincode = "test_cc2.go" // This is the name of the file that contains the chaincode, used for error messages
//const layout = time.RFC3339                 //or:ANSIC, UnixDate::: this is used with time.parse() to convert time.Time to string
var chainlogger = shim.NewLogger("FirstLogger")
// var chainLogLevel shim.logging.Level
// var chainLogger *shim.ChaincodeLogger
const users = 12
type DataStruct struct { // This is a basic struct
		SM string    `json:"SmartMeter"`
		Time string    `json:"Timestamp"`
		Data1 [users] string    `json:"Consommations"` //the data data: could be anything
		Data2 string   `json:"MerklerootDATA"`
		//Time time.Time `json:"time"`
				}

//const users = 12
type VoteStruct struct { // This is a basic struct
		VOTES string    `json:"Votes"`
		LIST_VOTES [users+1] string    `json:"List Votes"`

				}


type RepStruct struct { // This is a basic struct
		ID string    `json:"Data"`
		//Data1 [10]string    `json:"Consommations"` // the data data: could be anything
		Data string   `json:"MerklerootReputations"`
		//Time time.Time `json:"time"`
					}

func (t *DataStruct) Init(stub shim.ChaincodeStubInterface) pb.Response {
// This function must be implimented, it is executed when instantiating or upgrading the chaincode, i don't think that we need to specify any actions yet (maybe in the future [-_-] )
// chainLogger = chainlogger.Info(loggername)

	chainlogger.Info("Init function::: test_cc2.go")
	
	return shim.Success(nil) // this return value is required
						}



///////////////////////////////////////////////////// merkle root functions ////////////////////


//TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	  x string
  }

  //CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() ([]byte, error) {
	    h := sha256.New()
	      if _, err := h.Write([]byte(t.x)); err != nil {
		          return nil, err
			    }

			      return h.Sum(nil), nil
		      }

		      //Equals tests for equality of two Contents
func (t TestContent) Equals(other merkletree.Content) (bool, error) {
			        return t.x == other.(TestContent).x, nil
			}








///////////////////////////////////////////////////// END of merkle root functions ////////////////////



func (t *DataStruct) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
// This function must be implimented, it is executed when consulting the chaincode or the blockchain(when read/write/update...)
// I think that we need the function '''response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)'''
// Look "https://github.com/hyperledger/fabric/blob/release-1.1/examples/chaincode/go/chaincode_example05/chaincode_example05.go#L94"

		chainlogger.Info("Invoke function::: test_cc2.go")
		function, args := stub.GetFunctionAndParameters() // function: the name of the function to execute

		chainlogger.Info("Executing function " + function + " ...")
		switch function {
			case "NewData":
				return NewData(stub, args)
			case "GetData":
				return GetData(stub, args)
                        case "Votes":
                                return Votes(stub, args)
                        case "GetVotes":
                                return GetVotes(stub, args)
                        case "NewRep":
                                return NewRep(stub, args)
                        case "GetRep":
                                return GetRep(stub, args)
			case "Delete":
				return Delete(stub, args)
			default:
				chainlogger.Info("Error::: function " + function + " not found!!!")
				return shim.Error("Error::: function not found!!!")

									}
											// return shim.Success(nil)
			}

var consommation [users]string
var list_reputation [users]string

func NewData(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//	 variables
    var sm string
    var id string
    var data2 string
    
    //var timestamp string
	var err error
	var dataObject DataStruct
	var x int
	var today time.Time
			
	x = 1
	timestamp := "0"

	ArgsNb := 1             // number of arguments expected"
	if len(args) < ArgsNb { //verify the number of arguments
		return shim.Error("Expecting more values")				
	}
//	// retrieving the parametres/arguments
	   
	   sm = args[0]
	   id = args[1]
	   //data1 = args[2]
	   if len(args) > 2 {
		data2 = args[2]
							}


	// str to int		             
	tx, err := strconv.Atoi(id)
	if err != nil {
      // handle error
   				  }

    consommation[tx] = data2

//}
	for i:=0; i<users; i++{
		if consommation[i] == "" {
		x = 0
		break
}
}
///////////////////////////////////////////////////////////
	if x == 1{
												// Merkle root DATA
	  	var MR_D []merkletree.Content
		for i:=0; i<users; i++{
			MR_D = append(MR_D, TestContent{x: consommation[i]})

							}

  		//Create a new Merkle Tree from the list of Content
  		t, err := merkletree.NewTree(MR_D)
  		if err != nil {
    		log.Fatal(err)
  						}

  		//Get the Merkle Root of the tree
  		mr1 := t.MerkleRoot()
  		log.Println(mr1)
		 //sha := base64.URLEncoding.EncodeToString(mr)  
  		 //fmt.Println(sha)  
		 mr1_hash := hex.EncodeToString(mr1) 


		// verifying if the id is valid or not
		
	    dataBytes, err := stub.GetState(sm)
	    if err != nil {
			return shim.Error(err.Error())
	    } else if dataBytes != nil {
			chainlogger.Info("Error::: This id already exists: " + id + "!!!")
			return shim.Error("This id " + id + " already exists!!!")

													}
	// ********************************************
	// InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response
	// ********************************************
	// creating the data object???

		today = time.Now()
		timestamp = today.Format("2006.01.02 15:04:05")
		//timestamp := dt. Format("2006.01.02 15:04:05")

		dataObject = DataStruct{sm, timestamp, consommation, mr1_hash} // this is what we store in the blockchain(after the serialisation!!!)
			//	// converting to slice of bytes in order to be stored in the blockchain
		dataJsonBytes, err := json.Marshal(dataObject)
		if err != nil {
			return shim.Error(err.Error())
			}
		err = stub.PutState(sm, dataJsonBytes)
		if err != nil {
			return shim.Error(err.Error())
						}


		for i:=0; i<users; i++{
		consommation[i]= ""
			    }
			             
	}

			return shim.Success(nil)
				}


func GetData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
						// variables
    var sm string
    var err error
	ArgsNb := 1             // number of arguments expected"
	if len(args) < ArgsNb { //verify the number of arguments
		return shim.Error("Expecting more values")
																															}

	//	// finding the object
	sm = args[0]
	dataObjectBytes, err := stub.GetState(sm)
	if err != nil {
		chainlogger.Info("Error::: when retrieving the data with id " + sm + "!!!")
		return shim.Error(err.Error())
	} else if dataObjectBytes == nil {
		chainlogger.Info("invalid id" + sm + "!!! no data is found!!! ")
		return shim.Error("invalid id" + sm + "!!! no data is found!!! ")
												}

		return shim.Success(dataObjectBytes)
										}





const peers = 4
var votes_recus = peers - 1 // le nombre de votes que ce peer attend avant de faire la liste des votes. (-1 pour eliminer le leader car on attend pas un vote de sa part).
var b1_votes [users+1] string


func Votes(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// variables
    var vote_id string // B1_votes
    var sm string // sm = user
    var vote string

	var voteObject VoteStruct

	votes_pos := 0
	votes_neg := 0
	votes_globale := ""

	vote_id = args[0]
	sm = args[1]
	vote = args[2]
	 


	//str to int
	tx, err := strconv.Atoi(sm)
	if err != nil {
      // handle error
   				  }

    b1_votes[tx] = vote
    votes_recus = votes_recus - 1



    
	if votes_recus == 0{   // if received all the votes.


		for i:=0; i<users; i++{
			if b1_votes[i] == "accepter" {
				votes_pos = votes_pos + 1
			}
			if b1_votes[i] == "rejeter" {
				votes_neg = votes_neg + 1
			}
		}


		if votes_pos > 2*votes_neg{
			votes_globale = "valide"

		}else{
			votes_globale = "notvalide"
		}

		b1_votes[users] = votes_globale



		// verifying if the id is valid or not
		//str1 := strconv.Itoa(tx)
	    dataBytes, err := stub.GetState(vote_id)
	    if err != nil {
			return shim.Error(err.Error())
	    } else if dataBytes != nil {
			chainlogger.Info("Error::: This id already exists: " + vote_id + "!!!")
			return shim.Error("This id " + vote_id + " already exists!!!")

													}
	// ********************************************
	// InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response
	// ********************************************
	// creating the data object???
		voteObject = VoteStruct{vote_id, b1_votes} // this is what we store in the blockchain(after the serialisation!!!)
			//	// converting to slice of bytes in order to be stored in the blockchain
		dataJsonBytes, err := json.Marshal(voteObject)
		if err != nil {
			return shim.Error(err.Error())
			}
		err = stub.PutState(vote_id, dataJsonBytes)
		if err != nil {
			return shim.Error(err.Error())
						}
		

		for i:=0; i<=users; i++{
			b1_votes[i] = ""
		}
		votes_recus = peers - 1

}
			return shim.Success(nil)
				}



func GetVotes(stub shim.ChaincodeStubInterface, args []string) pb.Response {
						// variables
    var vote_retrieve string
    var err error
	ArgsNb := 1             // number of arguments expected"
	if len(args) < ArgsNb { //verify the number of arguments
		return shim.Error("Expecting more values")
																															}

	//	// finding the object
	vote_retrieve = args[0]
	dataObjectBytes, err := stub.GetState(vote_retrieve)
	if err != nil {
		chainlogger.Info("Error::: when retrieving the data with id " + vote_retrieve + "!!!")
		return shim.Error(err.Error())
	} else if dataObjectBytes == nil {
		chainlogger.Info("invalid id" + vote_retrieve + "!!! no data is found!!! ")
		return shim.Error("invalid id" + vote_retrieve + "!!! no data is found!!! ")
												}

		return shim.Success(dataObjectBytes)
										}

func NewRep(stub shim.ChaincodeStubInterface, args []string) pb.Response {

//	 variables
     var id string  // id is the id to retrieve the data from the BC with it"in this case id = maj_reputations"
     var hash_rep string
     var err error
     var repObject RepStruct

     ArgsNb := 1             // number of arguments expected"
	if len(args) < ArgsNb { //verify the number of arguments
		return shim.Error("Expecting more values")				
				}


   id = args[0]
   hash_rep = args[1]

												
    // verifying if the id is valid or not
	dataBytes, err := stub.GetState(id)
    if err != nil {
	return shim.Error(err.Error())
   } else if dataBytes != nil {
	chainlogger.Info("Error::: This id already exists: " + id + "!!!")
	return shim.Error("This id " + id + " already exists!!!")
			}
																																																		// ********************************************
//nvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response
//********************************************
// creating the data object???
	repObject = RepStruct{id, hash_rep} // this is what we store in the blockchain(after the serialisation!!!)
//rting to slice of bytes in order to be stored in the blockchain
													dataJsonBytes, err := json.Marshal(repObject)
													if err != nil {
	return shim.Error(err.Error())
			}
													err = stub.PutState(id, dataJsonBytes)
     if err != nil {
	return shim.Error(err.Error())
		}

    return shim.Success(nil)
				}



func GetRep(stub shim.ChaincodeStubInterface, args []string) pb.Response {
						// variables
    var id string
    var err error
	ArgsNb := 1             // number of arguments expected"
	if len(args) < ArgsNb { //verify the number of arguments
		return shim.Error("Expecting more values")
																															}

	//	// finding the object
	id = args[0]
	dataObjectBytes, err := stub.GetState(id)
	if err != nil {
		chainlogger.Info("Error::: when retrieving the data with id " + id + "!!!")
		return shim.Error(err.Error())
	} else if dataObjectBytes == nil {
		chainlogger.Info("invalid id ggggg" + id + "!!! no data is found!!! ")
		return shim.Error("invalid id" + id + "!!! no data is found!!! ")
												}

		return shim.Success(dataObjectBytes)
										}


func Delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}









func (t *DataStruct) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
		return shim.Success(nil)
				}

func main() {
	err := shim.Start(new(DataStruct))
	if err != nil {
		fmt.Printf("Error!!! cannot start chaincode %s: %s", this_chaincode, err)
						}
					}

