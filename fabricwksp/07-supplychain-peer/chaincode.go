package main

import (
   "github.com/hyperledger/fabric/core/chaincode/shim"
        pb "github.com/hyperledger/fabric/protos/peer"
        "fmt"
        "encoding/json"
        "time"
)

var asset_name = "asset_name_a"
type scfinancechaincode struct {
	
}

func (t *scfinancechaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Printf(" init success")
	return shim.Success([]byte("init success"))
}

func (t *scfinancechaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response{

	_,args := stub.GetFunctionAndParameters()
	var opttype = args[0]
	var assetname = args[1]
	var optcontent = args[2]

	fmt.Printf("param is %s %s %s \n",opttype,assetname,optcontent)

	if opttype == "putvalue"{
		stub.PutState(assetname,[]byte(optcontent))
		return shim.Success([]byte("success put" + optcontent))
	}else if opttype == "getlastvalue"{
		var kv []byte
		var err error
		kv,err = stub.GetState(assetname)
		if(err != nil){
			return shim.Error("find error!")
		}
		return shim.Success(kv)
	}else if opttype == "gethistory"{
		keysIter,err := stub.GetHistoryForKey(assetname)
		if err != nil {
			return shim.Error(fmt.Sprintf("GetHistoryForKey failed,Error state: %s",err))
		}
		defer keysIter.Close()
		var keys []string
		for keysIter.HasNext(){
			response,iterErr := keysIter.Next()
			if iterErr != nil {
				return shim.Error(fmt.Sprintf("GetHistoryForKey operation failed,Error state: %s",err))
			}

			txid := response.TxId 
			txvalue := response.Value
			txstate := response.IsDelete
			txtimestamp := response.Timestamp
			tm := time.Unix(txtimestamp.Seconds,0)
			datestr := tm.Format("2006-01-02 03:04:05 PM")

			fmt.Printf(" Tx info - txid: %s value: %s if delete: %t datetime: %s\n",txid,string(txvalue),txstate,datestr)
			keys = append(keys,txid)
		}

		jsonKeys,err := json.Marshal(keys)
		if err != nil {
			return shim.Error(fmt.Sprintf("query operation failed,error is %s",err))
		}
		return shim.Success(jsonKeys)
	}else{
		return  shim.Success([]byte("success invoke and No operation"))
	}
}

func main() {
        err := shim.Start(new(scfinancechaincode))
        if err != nil {
                fmt.Printf("Error starting Simple chaincode: %s", err)
        }
}
