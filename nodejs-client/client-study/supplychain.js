// import 

var co = require("co")
var path = require("path")
var fs = require("fs")
var util = require("util")
var hfclient = require("fabric-client")
var Peer = require("fabric-client/lib/Peer.js");
var EventHub = require("fabric-client/lib/ChannelEventHub");
var User = require("fabric-client/lib/User.js");
var crypto = require("crypto")
var FabricCAService = require("fabric-ca-client")

// 证书文件的缓存目录
var tempdir = "/home/shisj/mycode/github_wksp/fabric/nodejs-client/client-study/kvs";

// fabric client agent

var client = new hfclient();
var cryptoSuite = hfclient.newCryptoSuite()
cryptoSuite.setCryptoKeyStore(hfclient.newCryptoKeyStore({path:tempdir}))
client.setCryptoSuite(cryptoSuite);
const ORDERER_IP = "106.12.196.74";
const PEER_IP_supplier = "106.75.115.63";
const PEER_IP_core = "106.75.62.194";
const PEER_IP_bank = "106.75.16.132";
// 创建CA客户端 


var channel = client.newChannel("cmbcchannel666");
var order = client.newOrderer("grpc://" + ORDERER_IP + ":7050");
channel.addOrderer(order);

var peerSupplier = client.newPeer("grpc://"+PEER_IP_supplier+":7051");
channel.addPeer(peerSupplier);

var peerCore = client.newPeer("grpc://"+PEER_IP_core+":7051");
channel.addPeer(peerCore);

var peerBank = client.newPeer("grpc://"+PEER_IP_bank+":7051");
channel.addPeer(peerBank);


function readAllFiles(dirPath){
    var files = fs.readdirSync(dirPath);  
    var data = fs.readFileSync(path.join(dirPath,files[0]), 'utf8');
    return data;
}

/**
 * 根据cryptogen模块生成的帐号通过Fabric接口进行相关的操作
 */
function getOrgUser4Local(){
    var keyPath = "/home/shisj/mycode/github_wksp/fabric/nodejs-client/client-study/supplychain/msp-supplier/keystore";
    var keyPEM = Buffer.from(readAllFiles(keyPath)).toString()
    var certPath = "/home/shisj/mycode/github_wksp/fabric/nodejs-client/client-study/supplychain/msp-supplier/signcerts";
    var certPEM = readAllFiles(certPath).toString()

    return hfclient.newDefaultKeyValueStore({path:tempdir})
    .then((store) =>{
        client.setStateStore(store);
        return client.createUser({
            username:"admin",
            mspid:"SupplierMSP",
            cryptoContent:{
                privateKeyPEM:keyPEM,
                signedCertPEM:certPEM
            }
        })
    })
    
}


co((function *(){
    let member = yield getOrgUser4Local()
    // let resultPeerInfo = yield channel.queryInfo(peer)
    console.info("============当前peer加入的某个channel的区块数============");
    // console.info(JSON.stringify(resultPeerInfo))

    console.info("============当前peer加入了那些channel============");
    // var result = yield client.queryChannels(peer)
    // console.info(JSON.stringify(result))

    //console.info("============当前peer加入的某个channel的区块数的区块号============");
    //result = yield client.queryChannels(2,peer,null)
    //console.info(JSON.stringify(result))

    console.info("============当前peer加入chaincode[install]的信息============");
    // result = yield client.queryInstalledChaincodes(peer)
    // var chaincodes = result.chaincodes;
    // chaincodes.forEach(element => {
    //     console.info("chaincode name is %s, version is %s ",element.name,element.version);
    // });
    // console.info(JSON.stringify(result))

   let tx_id = client.newTransactionID();
   var request = {
       
       chaincodeId:"testchaincode",
       fcn:"invoke",
       args:["a","b","1"],
       chainId:"cmbcchannel666",
       txId:tx_id
   };

   let chaincodeInvokeResult = yield channel.sendTransactionProposal(request);
   var proposalResponse = chaincodeInvokeResult[0];
   var proposal = chaincodeInvokeResult[1];
   var header = chaincodeInvokeResult[2];
   var all_good = true;
   for(var i in proposalResponse){
       let one_good = false;
       if(proposalResponse && proposalResponse[0].response &&
        proposalResponse[0].response.status === 200){
            one_good = true;
            console.log("transaction is good");
        }else{
            console.error("transaction is bad");
        }
        all_good = all_good & one_good;
   }
   if(all_good){
       /*console.info(util.format(

        'successfullly: status - %s,message - "%s",metadata - "%s",endorsemenet signature:"%s"',
        proposalResponse[0].response.status,proposalResponse[0].response.message,
        proposalResponse[0].response.payload,proposalResponse[0].response.endorsement.signature
     ));*/


     var request = {
        proposalResponses:proposalResponse,
        proposal:proposal,
        header:header
    }
 
    var transactionID = tx_id.getTransactionID();
    var sendPromise = yield channel.sendTransaction(request);

   }

   

})())