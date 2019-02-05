echo 'Address recieved: ' $1
echo 'Data recieved: ' $2
curl -XPOST $1 -d "$2"