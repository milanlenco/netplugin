#!/usr/bin/env bash

if [ $USER != "root" ] ; then
    echo "Restarting script with sudo..."
    sudo $0 ${*}
    exit
fi

DOCKER1="web1"
DOCKER2="web2"

exposedockernetns () {
    if [ "$1" == "" ]; then
            echo "usage: $0 <container_name>"
            echo "Exposes the netns of a docker container to the host"
	    exit 1
    fi

    pid=`docker inspect -f '{{.State.Pid}}' $1`
    ln -s /proc/$pid/ns/net /var/run/netns/$1
    
    echo "netns of ${1} exposed as /var/run/netns/${1}"
    
    echo "try: ip netns exec ${1} ip addr list"
    return 0
}

dockerrmf () {
    #Cleanup all containers on the host (dead or alive).
    echo "Cleaning up any docker containers"
    docker stop `docker ps --no-trunc -aq` > /dev/null ; docker rm `docker ps --no-trunc -aq` > /dev/null
    echo "Restarting VPP to wipe config"
    service vpp restart
}

echo "#Remove old netns simlink"
rm -Rf /var/run/netns/*
mkdir /var/run/netns

echo "Killing all containers"
dockerrmf

echo "Creating docker containers hasvppinterface1 & 2"
docker pull ubuntu
docker run -itd --name $DOCKER1 ubuntu:14.04 /bin/bash
docker run -itd --name $DOCKER2 ubuntu:14.04 /bin/bash

echo "Expose our container to the ip netns exec tools"
exposedockernetns $DOCKER1
exposedockernetns $DOCKER2

echo "Create and move veth"
ip link add name veth_$DOCKER1 type veth peer name $DOCKER1
ip link set dev $DOCKER1 up
ip link set dev veth_$DOCKER1 up netns $DOCKER1
ip netns exec $DOCKER1 ip addr add 192.199.1.2/24 dev veth_$DOCKER1
ip netns exec $DOCKER1 ip link set veth_$DOCKER1 up
ip netns exec $DOCKER1 ip route add 192.199.2.0/24 via 192.199.1.1


ip link add name veth_$DOCKER2 type veth peer name $DOCKER2
ip link set dev $DOCKER2 up
ip link set dev veth_$DOCKER2 up netns $DOCKER2
ip netns exec $DOCKER2 ip addr add 192.199.2.2/24 dev veth_$DOCKER2
ip netns exec $DOCKER2 ip link set veth_$DOCKER2 up
ip netns exec $DOCKER2 ip route add 192.199.1.0/24 via 192.199.2.1


#Block ICMP out of the default docker0 container interfaces to prevent false positive results
ip netns exec $DOCKER1 iptables -A OUTPUT -p icmp -o eth0 -j REJECT
ip netns exec $DOCKER2 iptables -A OUTPUT -p icmp -o eth0 -j REJECT


