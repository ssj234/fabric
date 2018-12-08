// 1.在main包下
package main
// 2.引入必要的依赖
import(
   "fmt"
   "strconv"
   "github.com/hyperledger/fabric/core/chaincode/shim"
   pb "github.com/hyperledger/fabric/protos/peer"
)
// 3.定义一个结构体
type mychaincode struct {
}
// 4.为结构体绑定init和invoke方法
func (t *mychaincode) Init(stub shim.ChaincodeStubInterface) pb.Response{
    var args []string
    fmt.Println(" << ====[Init] success init it is view in docker ======")
    _,args = stub.GetFunctionAndParameters()
    var a ,b string
    var aAmt,bAmt int
    var err error
    a = args[0]
    aAmt,err = strconv.Atoi(args[1])
    b = args[2]
    bAmt,err = strconv.Atoi(args[3])
    fmt.Println("aAmt is ",aAmt)
    fmt.Println("bAmt is ",bAmt)
    stub.PutState(a,[]byte(args[1]))
    stub.PutState(b,[]byte(args[3]))
    if err != nil{
       fmt.Println("error")
    } 
    return shim.Success([]byte("success init"))
}

func (t *mychaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
   var args []string 
   var aAmtBytes,bAmtBytes []byte
   fmt.Println(" << ====[Invoke] success init it is view in docker ======")
   _,args = stub.GetFunctionAndParameters()
   var a ,b string
   var aAmt,bAmt int
   var transAmt int
   //var err error
   a = args[0]
   b = args[1]
   transAmt,_ = strconv.Atoi(args[2])
   aAmtBytes,_ = stub.GetState(a)
   bAmtBytes,_ = stub.GetState(b)
   aAmt, _ = strconv.Atoi(string(aAmtBytes))
   bAmt, _ = strconv.Atoi(string(bAmtBytes))
   fmt.Println("aAmt form fabric is ",aAmt)
   fmt.Println("bAmt form fabric is ",bAmt)
   aAmt = aAmt - transAmt
   bAmt = bAmt +  transAmt
   stub.PutState(a, []byte(strconv.Itoa(aAmt)))
   stub.PutState(b, []byte(strconv.Itoa(bAmt)))
    return shim.Success([]byte("success init"))
}
// 5.主方法
func main(){
    err := shim.Start(new(mychaincode))
    if err != nil{
        fmt.Println("Error starting Simple chaincode : %s",err)
    }
}
