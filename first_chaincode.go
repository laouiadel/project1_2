// Look the official documentation documentation, fabric-sample/chaincode, and the fabric source code for deeper informations.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	// "strconv"
	"time"
	// /fabric/vendor/github.com/op/go-logging
)

const this_chaincode = "first_chaincode.go" // This is the name of the file that contains the chaincode, used for error messages
const layout = time.RFC3339                 //or:ANSIC, UnixDate::: this is used with time.parse() to convert time.Time to string
var chainlogger = shim.NewLogger("FirstLogger")

// var chainLogLevel shim.logging.Level
// var chainLogger *shim.ChaincodeLogger

type DataStruct struct { // This is a basic struct
	Id   string    `json:"Timestamp"`

	Data string    `json:"Data"` 
	Mr_data string    `json:"Merkle root data (MR-D)"`
	Rep string    `json:"Liste des reputations"` 
	Mr_rep string    `json:"Merkle root de la liste des reputations (MR-R)"`
	
	//Time time.Time `json:"time"`
}


type TestStruct struct { // This is a basic struct
	Id   string    `json:"id"`
	Rep string    `json:"Liste des reputations"` 
	//Time time.Time `json:"time"`
						}





func (t *DataStruct) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// This function must be implimented, it is executed when instantiating or upgrading the chaincode, i don't think that we need to specify any actions yet (maybe in the future [-_-] )
	// chainLogger = chainlogger.Info(loggername)

	chainlogger.Info("Init function::: first_chaincode.go")

	return shim.Success(nil) // this return value is required
}

func (t *DataStruct) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// This function must be implimented, it is executed when consulting the chaincode or the blockchain(when read/write/update...)
	// I think that we need the function '''response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)'''
	// Look "https://github.com/hyperledger/fabric/blob/release-1.1/examples/chaincode/go/chaincode_example05/chaincode_example05.go#L94"

	chainlogger.Info("Invoke function::: first_chaincode.go")
	function, args := stub.GetFunctionAndParameters() // function: the name of the function to execute

	chainlogger.Info("Executing function " + function + " ...")
	switch function {
	case "NewData":
		return NewData(stub, args)
	case "GetData":
		return GetData(stub, args)
	case "test":
		return test(stub, args)
	default:
		chainlogger.Info("Error::: function " + function + " not found!!!")
		return shim.Error("Error::: function not found!!!")

	}
	// return shim.Success(nil)
}




func test(stub shim.ChaincodeStubInterface, args []string) pb.Response {

		//	// variables
	var id string
	var data string
        //var val2 string
	//var Time time.Time
	var err error
	var dataObject TestStruct
	ArgsNb := 1             // number of arguments expected"
	if len(args) < ArgsNb { //verify the number of arguments
													return shim.Error("Expecting more values")
														}

													//	// retrieving the parametres/arguments
	id = args[0]
	data = args[1]
	//if len(args) > 2 {
	//	val2 = args[2]
	//	if len(args) > 3 {
	//		Time, err = time.Parse(layout, args[3])
	//		if err != nil {
	//			return shim.Error(err.Error())
	//		}
	///	}
	//		}
	// verifying if the id is valid or not
	dataBytes, err := stub.GetState(id)
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
dataObject = TestStruct{id, data} // this is what we store in the blockchain(after the serialisation!!!)
//converting to slice of bytes in order to be stored in the blockchain
dataJsonBytes, err := json.Marshal(dataObject)
if err != nil {
	return shim.Error(err.Error())
			}
	err = stub.PutState(id, dataJsonBytes)
if err != nil {
	return shim.Error(err.Error())
			}
	return shim.Success(nil)
				}







func NewData(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//	// variables
	var time string
	var mesures string
	var mr_mesures string

	var reputations string
	var mr_reputations string

//	var Time time.Time
	var err error
	var dataObject DataStruct
	ArgsNb := 1             // number of arguments expected"
	if len(args) < ArgsNb { //verify the number of arguments
		return shim.Error("Expecting more values")
	}

	//	// retrieving the parametres/arguments
	time = args[0]
	mesures = args[1]
	mr_mesures = args[2]

	reputations = args[3]
	mr_reputations = args[4]

	//	if len(args) > 5 {
	//		Time, err = time.Parse(layout, args[5])
	//		if err != nil {
	//			return shim.Error(err.Error())
	//		}
	//	}
	//

	//	// verifying if the id is valid or not
	dataBytes, err := stub.GetState(time)
	if err != nil {
		return shim.Error(err.Error())
	} else if dataBytes != nil {
		chainlogger.Info("Error::: This id already exists: " + time + "!!!")
		return shim.Error("This id " + time + " already exists!!!")
	}
	// ********************************************
	// InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response
	// ********************************************
	//	// creating the data object???
	dataObject = DataStruct{time, mesures, mr_mesures, reputations, mr_reputations} // this is what we store in the blockchain(after the serialisation!!!)
	//	// converting to slice of bytes in order to be stored in the blockchain
	dataJsonBytes, err := json.Marshal(dataObject)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(string(dataJsonBytes), dataJsonBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func GetData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
		chainlogger.Info("invalid id" + id + "!!! no data is found!!! ")
		return shim.Error("invalid id" + id + "!!! no data is found!!! ")
	}

	return shim.Success(dataObjectBytes)
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
