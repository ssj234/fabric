package main
import(
   "fmt"
   "github.com/hyperledger/fabric/core/chaincode/shim"
   pb "github.com/hyperledger/fabric/protos/peer"
)

type mychaincode struct {
}

func (t *mychaincode) Init(stub shim.ChaincodeStubInterface) pb.Response{
    fmt.Println(" << ====[Init] success init it is view in docker ======")
    return shim.Success([]byte("success init"))
}

func (t *mychaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	fmt.Println(" << ====[Invoke] success init it is view in docker ======")
    return shim.Success([]byte("success init"))
}

func main(){
	err := shim.Start(new(mychaincode))
	if err != nil{
		fmt.Println("Error starting Simple chaincode : %s",err)
	}
}
