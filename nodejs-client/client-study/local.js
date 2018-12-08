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
const PEER_IP = "106.12.196.74";
// 创建CA客户端 


var channel = client.newChannel("cmbcchannel666");
var order = client.newOrderer("grpc://" + ORDERER_IP + ":7050");
channel.addOrderer(order);

var peer = client.newPeer("grpc://"+PEER_IP+":7051");
channel.addPeer(peer);


function readAllFiles(dirPath){
    var files = fs.readdirSync(dirPath);  
    var data = fs.readFileSync(path.join(dirPath,files[0]), 'utf8');
    return data;
}

/**
 * 根据cryptogen模块生成的帐号通过Fabric接口进行相关的操作
 */
function getOrgUser4Local(){
    var keyPath = "/home/shisj/mycode/github_wksp/fabric/nodejs-client/client-study/fabric-config/msp/keystore";
    var keyPEM = Buffer.from(readAllFiles(keyPath)).toString()
    var certPath = "/home/shisj/mycode/github_wksp/fabric/nodejs-client/client-study/fabric-config/msp/signcerts";
    var certPEM = readAllFiles(certPath).toString()

    return hfclient.newDefaultKeyValueStore({path:tempdir})
    .then((store) =>{
        client.setStateStore(store);
        return client.createUser({
            username:"user87",
            mspid:"Org1MSP",
            cryptoContent:{
                privateKeyPEM:keyPEM,
                signedCertPEM:certPEM
            }
        })
    })
    
}


co((function *(){
    let member = yield getOrgUser4Local()
    let resultPeerInfo = yield channel.queryInfo(peer)
    console.info("============当前peer加入的某个channel的区块数============");
    console.info(JSON.stringify(resultPeerInfo))

    console.info("============当前peer加入了那些channel============");
    var result = yield client.queryChannels(peer)
    console.info(JSON.stringify(result))

    //console.info("============当前peer加入的某个channel的区块数的区块号============");
    //result = yield client.queryChannels(2,peer,null)
    //console.info(JSON.stringify(result))

    console.info("============当前peer加入chaincode[install]的信息============");
    result = yield client.queryInstalledChaincodes(peer)
    var chaincodes = result.chaincodes;
    chaincodes.forEach(element => {
        console.info("chaincode name is %s, version is %s ",element.name,element.version);
    });
    // console.info(JSON.stringify(result))

   let tx_id = client.newTransactionID();
   var request = {
       targets: peer,
       chaincodeId:"statechaincode",
       fcn:"invoke",
       args:["a","b","1"],
       chainId:"cmbcchannel666",
       txId:tx_id
   };

   let chaincodeInvokeResult = yield channel.sendTransactionProposal(request);


})())