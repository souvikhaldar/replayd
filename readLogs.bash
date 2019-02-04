#!/bin/bash
echo "reading error logs from 34.219.22.227 replayd http cache server"
ssh ubuntu@34.219.22.227 "journalctl -u replayd.service $1"

