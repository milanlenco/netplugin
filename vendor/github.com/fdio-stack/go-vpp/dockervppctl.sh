#!/usr/bin/env bash

if [ $USER != "root" ] ; then
    echo "Restarting script with sudo..."
    sudo $0 ${*}
    exit
fi

DOCKER1="web1"
DOCKER2="web2"

service vpp restart

#Temporary vppctl commands to test....
vppctl create host-interface name $DOCKER1
vppctl set int ip address host-$DOCKER1 192.199.1.1/24
vppctl set int state host-$DOCKER1 up

vppctl create host-interface name $DOCKER2
vppctl set int ip address host-$DOCKER2 192.199.2.1/24
vppctl set int state host-$DOCKER2 up
