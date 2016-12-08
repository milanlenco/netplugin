#!/usr/bin/env bash

if [ $USER != "root" ] ; then
    echo "Restarting script with sudo..."
    sudo $0 ${*}
    exit
fi

DOCKER1="web1"
DOCKER2="web2"

echo "Ping from "$DOCKER1" via TAP > HostVPP > Host"
docker exec $DOCKER1 /bin/ping -c3 192.199.2.2

echo "Ping from "$DOCKER2" via TAP > HostVPP > Host"
docker exec $DOCKER2 /bin/ping -c3 192.199.1.2

