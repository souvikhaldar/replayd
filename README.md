# EPO

## Deploying the server to cloud
Enter IP of the target host and username in the inventory file. Then run `./deploy.bash`

## Ping the host
Check the accessibility of the target host by running `./ping.bash <targethostname>`. Eg. `./ping.bash ubuntu`.

## Check logs of the replayd http server
Read the logs of the server by using `./readLogs.bash <username> <target_ip> <type>`. Eg. `./readLogs.bash ubuntu 34.219.22.227 -f`

## Read data from the buffer

`./replay.bash <hostname>`. Eg.- `./replay.bash 34.219.22.227`

## Insert data into the buffer

`./set-replay.bash <hostname> “<playload”>`. Eg- `./set-replay.bash 34.219.22.227 "hello"`

