#!/bin/bash

# Configure VPP
while sudo vppctl sh version | grep FileNotFoundError;
do
	echo "Restarting dysfunctional VPP..."
	sudo systemctl disable vpp.service # disable auto-restarts
	#sudo service vpp stop
	#sleep 3
	if pidof vpp; then
		sudo kill `pidof vpp`
		sleep 3
	fi
	sudo rm -f /dev/shm/vpe-api
	sudo ./dpdk.py -u 0000:00:0a.0
	sudo modprobe uio_pci_generic
	sudo ./dpdk.py -b uio_pci_generic 0000:00:0a.0
	sudo vpp "api-trace { on } dpdk { dev 0000:00:0a.0 }"
	sleep 5
done
if [[ $(hostname) == "netplugin-node1" ]]; then
       sudo vppctl set int ip address GigabitEthernet0/a/0 192.168.2.10/24
else
       sudo vppctl set int ip address GigabitEthernet0/a/0 192.168.2.11/24
fi
sudo vppctl set interface state GigabitEthernet0/a/0 up

