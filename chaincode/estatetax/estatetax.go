package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	peer "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

// Define the Smart Contract structure
type EstateTaxContract struct {
}

type EstateTax struct {
	TaxID  string `json:"taxid"`  //核税编号
	BookID string `json:"bookid"` //不动产权证书编号
	Taxer  string `json:"taxer"`  //纳税人
	Area   int    `json:"area"`   //房屋面积
	Tax    int    `json:"tax"`    //纳税金额
}

type RecordsInfo struct {
	Size  uint64
	Start string
	End   string
}

var key_recordinfo = "recordeinfo"
var evn_estatetax = "evn_estatetax"

func (s *EstateTaxContract) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *EstateTaxContract) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	function, args := APIstub.GetFunctionAndParameters()
	switch function {
	case "create":
		return s.create(APIstub, args)
	case "queryByTaxID":
		return s.queryByTaxID(APIstub, args)
	case "queryByPara":
		return s.queryByPara(APIstub, args)
	case "queryAll":
		return s.queryAll(APIstub, args)
	default:
		return shim.Error("Invalid Smart Contract function name.")
	}
}

func (s *EstateTaxContract) create(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6. ")
	}
	key := args[0]
	area, err := strconv.Atoi(args[4])
	if err != nil {
		return shim.Error("area value wrong.")
	}
	tax, err := strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("area value wrong.")
	}
	EstateTax := &EstateTax{
		TaxID:  args[1],
		BookID: args[2],
		Taxer:  args[3],
		Area:   area,
		Tax:    tax,
	}
	jsBytes, err := json.Marshal(EstateTax)
	if err != nil {
		return shim.Error("marshal json error:" + err.Error())
	}
	err = APIstub.PutState(key, jsBytes)
	if err != nil {
		return shim.Error("error on putstate:" + err.Error())
	}
	//update recodeinfo
	recordInfo := &RecordsInfo{}
	rebs, err := APIstub.GetState(key_recordinfo)
	if len(rebs) == 0 {
		recordInfo.Size = 1
		recordInfo.Start = key
		recordInfo.End = key
	} else {
		err = json.Unmarshal(rebs, &recordInfo)
		if err != nil {
			return shim.Error("error on unmarsh recorderinfo:" + err.Error())
		}
		recordInfo.Size = recordInfo.Size + 1
		recordInfo.End = key
	}
	rebs, err = json.Marshal(recordInfo)
	if err != nil {
		return shim.Error("error on marsh new recorderinfo:" + err.Error())
	}
	err = APIstub.PutState(key_recordinfo, rebs)
	if err != nil {
		return shim.Error("error on put new recorderinfo:" + err.Error())
	}
	//broadcast event
	err = APIstub.SetEvent(evn_estatetax, []byte("new EstateTax created with key:"+key))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("new EstateTax created with key:" + key))
}

func (s *EstateTaxContract) queryByTaxID(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 ")
	}
	queryString := fmt.Sprintf("{\"selector\":{\"taxid\":\"%s\"}}", args[0])
	qis, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error("queryByTaxID error:" + err.Error())
	}
	defer qis.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for qis.HasNext() {
		queryResponse, err := qis.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	return shim.Success(buffer.Bytes())
}

func (s *EstateTaxContract) queryByPara(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 ")
	}
	queryString := fmt.Sprintf("{\"selector\":{\""+args[0]+"\":\"%s\"}}", args[1])
	qis, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error("queryByPara error:" + err.Error())
	}
	defer qis.Close()
	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	return shim.Success(buffer.Bytes())
}

func (s *EstateTaxContract) queryAll(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	rebs, err := APIstub.GetState(key_recordinfo)
	if err != nil {
		return shim.Error("error on get recorderinfo:" + err.Error())
	}
	recordInfo := &RecordsInfo{}
	if len(rebs) == 0 {
		return shim.Success([]byte{})
	}
	err = json.Unmarshal(rebs, &recordInfo)
	if err != nil {
		return shim.Error("error on unmarsh recorderinfo:" + err.Error())
	}
	resultsIterator, err := APIstub.GetStateByRange(recordInfo.Start, recordInfo.End+"1")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	println("<recordeinfo> size:" + strconv.FormatUint(recordInfo.Size, 10) + "  start:" + recordInfo.Start + " end:" + recordInfo.End)
	return shim.Success(buffer.Bytes())
}

func main() {
	err := shim.Start(new(EstateTaxContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
