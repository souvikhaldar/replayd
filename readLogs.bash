#!/bin/bash
echo "---reading logs from replayd http cache server---"
ssh ubuntu@34.219.22.227 "journalctl -u replayd.service $1"

