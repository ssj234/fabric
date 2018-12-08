// 1.在main包下
package main
// 2.引入必要的依赖
import(
   "fmt"
   "github.com/hyperledger/fabric/core/chaincode/shim"
   pb "github.com/hyperledger/fabric/protos/peer"
)
// 3.定义一个结构体
type mychaincode struct {
}
// 4.为结构体绑定init和invoke方法
func (t *mychaincode) Init(stub shim.ChaincodeStubInterface) pb.Response{
    fmt.Println(" << ====[Init] success init it is view in docker ======")
    return shim.Success([]byte("success init"))
}

func (t *mychaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
    fmt.Println(" << ====[Invoke] success init it is view in docker ======")
    return shim.Success([]byte("success init"))
}
// 5.主方法
func main(){
    err := shim.Start(new(mychaincode))
    if err != nil{
        fmt.Println("Error starting Simple chaincode : %s",err)
    }
}
