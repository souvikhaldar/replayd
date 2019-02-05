#echo 'Address recieved: ' $1
echo 'Date recieved: ' $2
curl -XPOST http://34.219.22.227 -d "$2"