#echo 'Address recieved: ' $1
echo 'Data recieved: ' $1
curl -XPOST http://54.189.131.192/ -d "$1"
