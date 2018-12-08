var co = require("co");
var fabricservice = require("./supplychainExplorer.js")
var express = require("express")

var app = express()

var chaincodeid = "supplychain-chaincode";
var channelid = "cmbcchannel666";

app.get("/sendTransaction1",function(req,res){
    co(function * (){
        var blockinfo = yield fabricservice.sendTransaction(chaincodeid,"invoke",["putvalue","tans_id1","1"],channelid)
        res.send("success")
    })
})


app.get("/sendTransaction2",function(req,res){
    co(function * (){
        var blockinfo = yield fabricservice.sendTransaction(chaincodeid,"invoke",["putvalue","tans_id1","2"],channelid)
        res.send("success")
    })
})


app.get("/sendTransaction3",function(req,res){
    co(function * (){
        var blockinfo = yield fabricservice.sendTransaction(chaincodeid,"invoke",["putvalue","tans_id1","100"],channelid)
        res.send("success")
    })
})



app.get("/queryhistory",function(req,res){
    co(function * (){
        var blockinfo = yield fabricservice.sendTransaction(chaincodeid,"invoke",["gethistory","tans_id1","-1"],channelid)
        res.send("success")
    })
})

var server = app.listen(3000,function(){
    var host = server.address().address;
    var port = server.address().port;
    console.log("start at %s,%s",host,port)
})


process.on("unhandledRejection",function(err){
    console.error(err.stack);
})

process.on("unhandledException",function(err){
    console.error(err.stack);
})

