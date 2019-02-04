#!/bin/bash
echo "---reading logs from replayd http cache server---"
ssh $1@$2 "journalctl -u replayd.service $3"

