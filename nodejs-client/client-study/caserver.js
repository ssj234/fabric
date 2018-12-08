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
const CA_SERVER_IP = "106.12.196.74";
const ORDERER_IP = "106.12.196.74";
const PEER_IP = "106.12.196.74";
// 创建CA客户端 

var caClient = new FabricCAService("http://" + CA_SERVER_IP + ":7054",null,'',cryptoSuite);

var channel = client.newChannel("cmbcchannel666");

var order = client.newOrderer("grpc://" + ORDERER_IP + ":7050");
channel.addOrderer(order);

var peer = client.newPeer("grpc://"+PEER_IP+":7051");
channel.addPeer(peer);




/**
 * 通过CA获取当前用户的证书信息
 * @param {*} username 
 * @param {*} password 
 */
function getOrgUser4FabricCa(username,password){
  var member;
  return hfclient.newDefaultKeyValueStore({path:tempdir})
  .then((store) => {
      client.setStateStore(store);
      client._userContext = null;
      return client.getUserContext(username,true).then((user) => {
          if(user && user.isEnrolled()){
              console.info('success enrolled %s',username)
              client.setUserContext(user)
              return user;
          }else{
              return caClient.enroll({enrollmentID:username,enrollmentSecret:password}).then(
                  (enrollment) => {

                      console.info("Successfully enrolled user %s",username)
                      member = new User(username)
                      member.setCryptoSuite(client.getCryptoSuite())
                      return member.setEnrollment(enrollment.key,enrollment.certificate,"Org1MSP")
                   
              ).then( (user) =>{
                  return client.setUserContext(user)
              }).then((user)=>{
                  return user;
              }).catch((err) => {
                  console.error("enroll admin error" + err.stack)
              })
          }
      })
  })
}


co((function *(){
    let member = yield getOrgUser4FabricCa("admin","adminpw")
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