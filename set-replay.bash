#echo 'Address recieved: ' $1
echo 'Data recieved: ' $1
curl -XPOST http://34.219.22.227 -d "$1"