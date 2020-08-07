#!/bin/bash

head -n 21 My-network-origine.json > My-network.json


VAL=`expr $1 - 1`

for i in `seq 0 $VAL`;
do	
	if [ $i == $VAL ]
	then
		echo "                                \"peer$i.org1.dz\": {}" >> My-network.json
	else
		echo "                                \"peer$i.org1.dz\": {}," >> My-network.json
	fi
done
		
awk 'FNR>=24 && FNR<=37' My-network-origine.json >> My-network.json
#head -n 37 My-network-origine | tail -n 13 >> My-network.json
for i in `seq 0 $VAL`;
do	
	if [ $i == $VAL ]
	then
		echo "                                \"peer$i.org1.dz\": {}" >> My-network.json
	else
		echo "                                \"peer$i.org1.dz\": {}," >> My-network.json
	fi
done

awk 'FNR>=40 && FNR<=55' My-network-origine.json >> My-network.json
#head -n 55 My-network-origine | tail -n 15 >> My-network.json

cd /home/Adel/Desktop/PFE/Two_Chain_Network_Template/pfe-project/crypto-config/peerOrganizations/org1.dz/users/Admin@org1.dz/msp/keystore
fileKeyAdminOrderer=$(ls)
cd /home/Adel/hyperledger_explorer/connection-profile
echo "                                \"path\": \"/tmp/crypto/peerOrganizations/org1.dz/users/Admin@org1.dz/msp/keystore/$fileKeyAdminOrderer\"
}," >> My-network.json
	
PEERS=""
for i in `seq 0 $VAL`;
do
	if [ $i == $VAL ]
	then
		PEER="\"peer$i.org1.dz\" "
	else
		PEER="\"peer$i.org1.dz\", "
	fi
	PEERS=$PEERS$PEER
done
echo "                        \"peers\": [$PEERS]," >> My-network.json


awk 'FNR>=59 && FNR<=64' My-network-origine.json >> My-network.json
#head -n 64 My-network-origine | tail -n 5 >> My-network.json
PORT=7051
for i in `seq 0 $VAL`;
do
    x="                \"peer$i.org1.dz\": {
                        \"tlsCACerts\": {
                                \"path\": \"/tmp/crypto/peerOrganizations/org1.dz/peers/peer$i.org1.dz/tls/ca.crt\"
                        },
                        \"url\": \"grpcs://peer$i.org1.dz:$PORT\",
                        \"eventUrl\": \"grpcs://peer$i.org1.dz:`expr $PORT + 2`\",
                        \"grpcOptions\": {
                                \"ssl-target-name-override\": \"peer$i.org1.dz\"
                        }
                },"
				
	if [ $i == $VAL ]
	then
		echo "${x%?}" >> My-network.json
	else
		echo "$x" >> My-network.json
	fi
		PORT=`expr $PORT + 1000`
done

echo "
        }
}" >> My-network.json
		
