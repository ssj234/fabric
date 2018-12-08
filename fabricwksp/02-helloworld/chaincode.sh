export set FABRIC_CFG_PATH=/home/ssj234/fabricwksp/02-helloworld/peer
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_ADDRESS=peer0.org1.cmbc.com:7051
export set CORE_PEER_MSPCONFIGPATH=/home/ssj234/fabricwksp/02-helloworld/fabricconfig/crypto-config/peerOrganizations/org1.cmbc.com/users/Admin@org1.cmbc.com/msp

# 部署chaincode 代码
# /home/ssj234/software/go1.11.1/bin/src/home/ssj234/software/go1.11.1/bin/src/github.com/hyperledger/fabric/examples/chaincode/go/example01

peer chaincode install -n r_test_cc6 -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02

# 实例化chaincode代码

peer chaincode instantiate -o orderer.cmbc.com:7050 -C cmbcchannel666 -n r_test_cc6 -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "OR ('Org1MSP.member','Org2MSP.member')"

# 调用代码

peer chaincode invoke -o orderer.cmbc.com:7050 -C cmbcchannel666 -n r_test_cc6 -c '{"Args":["invoke","a","b","1"]}'


# 查询数据

peer chaincode query -C cmbcchannel666 -n r_test_cc6 -c '{"Args":["query","a"]}'
