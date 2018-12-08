#!/bin/bash
if [[ "$1" = "" || "$2" = "" ]]; then
	echo "certCheck.sh  certfile keyfile"  
	exit 0;
else
	#certModuleMd5=`openssl x509 -noout -modulus -in $1 | openssl md5`
	#privateModuleMd5=`openssl rsa -noout -modulus -in $2 | openssl md5`


        #if [  "$certModuleMd5" = "$privateModuleMd5" ] ; then
        #        echo "ok"
	#else
	#	echo "not ok"
        #fi
        value=`openssl x509 -text -noout -in $1  | grep "Public Key Algorithm:" | awk  -F ':'  'BEGIN {}  {print $2} END {}'`

        if [ "$value" = " rsaEncryption" ] ; then
                echo $value

                requestModuleMd5=`openssl x509 -modulus -in $1 | grep Modulus | openssl md5`
                privateModuleMd5=`openssl rsa -noout -modulus -in $2 | openssl md5`

        else
                `openssl ec -in $2 -pubout -out ecpubkey.pem `
                privateModuleMd5=`cat ecpubkey.pem | openssl md5`
                requestModuleMd5=`openssl x509 -in $1  -pubkey  -noout | openssl md5`

        fi
        if [  "$requestModuleMd5" = "$privateModuleMd5" ] ; then
                echo "ok"
        fi


fi
