echo 'Address recieved: ' $1
echo 'Date recieved: ' $2
curl -XPOST $1 -d $2